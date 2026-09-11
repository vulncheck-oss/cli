package bill

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/anchore/syft/syft"
	"github.com/anchore/syft/syft/cataloging/pkgcataloging"
	"github.com/anchore/syft/syft/format"
	"github.com/anchore/syft/syft/format/cyclonedxjson"
	"github.com/anchore/syft/syft/sbom"
	"github.com/package-url/packageurl-go"
	"github.com/vulncheck-oss/cli/pkg/cache"
	"github.com/vulncheck-oss/cli/pkg/client"
	"github.com/vulncheck-oss/cli/pkg/cmd/offline/packages"
	"github.com/vulncheck-oss/cli/pkg/cmd/offline/sync"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpeuri"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpeutils"
	"github.com/vulncheck-oss/cli/pkg/db"
	"github.com/vulncheck-oss/cli/pkg/models"
	"github.com/vulncheck-oss/cli/pkg/sdk"
	"github.com/vulncheck-oss/cli/pkg/session"
)

type InputSbomRef struct {
	SbomRef string
	PURL    string
	CPE     string
}

// SBOMOptions carries knobs for GetSBOM. Kept as a struct so future toggles
// can be added without churning every caller.
type SBOMOptions struct {
	// Enrich is a passthrough of syft's `--enrich` scope list. Empty means
	// enrichment is disabled and syft runs with its default (offline-safe)
	// behaviour. Non-empty entries opt in to network-backed metadata lookups
	// (proxy.golang.org, Maven Central, NPM registry, PyPI) — see
	// buildEnrichConfig for which fields each scope flips.
	Enrich []string
}

func GetSBOM(dir string, opts SBOMOptions) (*sbom.SBOM, error) {
	src, err := syft.GetSource(context.Background(), dir, nil)
	if err != nil {
		return nil, err
	}

	sbm, err := syft.CreateSBOM(context.Background(), src, buildSBOMConfig(opts))
	if err != nil {
		return nil, err
	}

	return sbm, nil
}

// Scope alias groups, mirroring syft's own task-name mapping
// (internal/task/package_tasks.go). Every name syft's CLI accepts for a
// language is listed so `--enrich <alias>` behaves the same here as it would
// under syft itself. These are the single source of truth for both
// buildSBOMConfig's dispatch and ValidateEnrich's accepted set, so a scope can
// never be dispatchable but unrecognised (or vice versa).
var (
	scopeGolang     = []string{"golang", "go"}
	scopeJavaScript = []string{"javascript", "node", "npm"}
	scopePython     = []string{"python"}
	scopeJava       = []string{"java", "maven"}

	// scopeMeta are the collective directives syft accepts alongside language
	// names. Their precedence is handled by enrichmentScope.
	scopeMeta = []string{"all", "none"}
)

// vcpkgUnsupported explains why we reject a scope syft itself accepts. Kept as
// a named constant because it is the one rejection that is a deliberate product
// decision rather than a typo, and the reasoning needs to survive contact with
// the next person who wonders why parity is missing.
//
// Must stay a single line: ui.Error renders through lipgloss, which pads a
// multi-line block out to its widest line and leaves trailing whitespace on
// every row.
const vcpkgUnsupported = "--enrich %s is not supported by this CLI (syft accepts it, but it clones git registries named in " +
	"the scanned repository's own vcpkg.json). Packages declared in vcpkg.json are still catalogued without it; " +
	"run syft directly if you need registry resolution. Supported scopes: all, golang, java, javascript, python"

// unsupportedScopes are scopes syft recognises that this CLI deliberately does
// not wire up, mapped to the explanation shown to the user. syft enables vcpkg
// registry cloning under both "vcpkg" and "cpp", so both are rejected.
var unsupportedScopes = map[string]string{
	"vcpkg": vcpkgUnsupported,
	"cpp":   vcpkgUnsupported,
}

// ValidateEnrich checks --enrich directives before any cataloguing starts.
// Without it an unrecognised scope is silently ignored: enrichmentScope simply
// never matches it, so `--enrich typo` or `--enrich vcpkg` looks accepted and
// quietly enriches nothing. Failing up front is the difference between "your
// flag did nothing" and "your SBOM is missing the metadata you asked for".
//
// Matching is case-sensitive and the +/- prefixes are accepted, both for parity
// with syft, whose own directive comparison is an exact match against lowercase
// task names. Scopes in unsupportedScopes are rejected only in their enabling
// forms; see the negation carve-out below.
func ValidateEnrich(directives []string) error {
	known := make(map[string]struct{})
	for _, group := range [][]string{scopeMeta, scopeGolang, scopeJavaScript, scopePython, scopeJava} {
		for _, name := range group {
			known[name] = struct{}{}
		}
	}

	for _, d := range directives {
		scope, negated := normalizeScope(d)
		if scope == "" {
			return fmt.Errorf("empty --enrich scope in %q; supported scopes: %s", d, supportedScopes())
		}
		if reason, ok := unsupportedScopes[scope]; ok {
			// A negated directive asks for the behaviour we already have, so
			// honour it rather than failing a command whose intent we satisfy —
			// `all,-vcpkg` is what someone hardening a syft invocation writes.
			// Only the enabling forms are an error.
			if negated {
				continue
			}
			return fmt.Errorf(reason, scope)
		}
		if _, ok := known[scope]; !ok {
			return fmt.Errorf("unknown --enrich scope %q; supported scopes: %s", scope, supportedScopes())
		}
	}
	return nil
}

// supportedScopes renders the publicised scope list for error messages. It
// mirrors syft's publicisedEnrichmentOptions rather than every alias, so the
// hint stays short and matches what --help advertises.
func supportedScopes() string {
	return "all, golang, java, javascript, python"
}

// normalizeScope strips the +/- prefix from a directive, returning the bare
// scope name and whether it was negated. Shared by enrichmentScope and
// ValidateEnrich so validation accepts exactly the syntax the dispatcher
// understands.
func normalizeScope(d string) (scope string, negated bool) {
	d = strings.TrimPrefix(d, "+")
	if strings.HasPrefix(d, "-") {
		return d[1:], true
	}
	return d, false
}

// buildSBOMConfig returns nil when no options need setting, preserving syft's
// default behaviour end-to-end. When Enrich is populated it starts from syft's
// defaults and flips the same per-language fields that syft's own CLI flips
// for `--enrich <scope>` — the scope mapping lives in syft's internal options
// package (cmd/syft/internal/options/catalog.go) and cannot be imported, so we
// mirror it here. Scopes match syft's publicised list: all, golang, java,
// javascript, python. The `+scope` / `-scope` / `all` / `none` precedence
// rules are handled by enrichmentScope.
func buildSBOMConfig(opts SBOMOptions) *syft.CreateSBOMConfig {
	if len(opts.Enrich) == 0 {
		return nil
	}

	pkgCfg := pkgcataloging.DefaultConfig()

	// Aliases mirror syft's own task-name mapping (internal/task/package_tasks.go).
	// Advertising only the publicised names in --help keeps parity with syft's help
	// output, but users typing an alias syft accepts should still get the same
	// behaviour they'd get from `syft --enrich <alias>`.
	if enrichmentScope(opts.Enrich, scopeGolang...) {
		pkgCfg.Golang = pkgCfg.Golang.
			WithSearchLocalModCacheLicenses(true).
			WithSearchLocalVendorLicenses(true).
			WithSearchRemoteLicenses(true).
			WithUsePackagesLib(true)
	}
	if enrichmentScope(opts.Enrich, scopeJavaScript...) {
		pkgCfg.JavaScript = pkgCfg.JavaScript.WithSearchRemoteLicenses(true)
	}
	if enrichmentScope(opts.Enrich, scopePython...) {
		pkgCfg.Python = pkgCfg.Python.
			WithSearchRemoteLicenses(true).
			WithGuessUnpinnedRequirements(true)
	}
	if enrichmentScope(opts.Enrich, scopeJava...) {
		pkgCfg.JavaArchive = pkgCfg.JavaArchive.
			WithUseMavenLocalRepository(true).
			WithUseNetwork(true)
	}

	return syft.DefaultCreateSBOMConfig().WithPackagesConfig(pkgCfg)
}

// enrichmentScope reports whether enrichment should be enabled for a language
// under the user's directives. Multiple aliases can be passed for languages
// syft's CLI accepts under more than one name (e.g. golang/go, javascript/node/npm,
// java/maven); an explicit `+alias` / bare `alias` / `-alias` on any of them
// wins over the `all` / `none` fallback, matching syft's own precedence.
func enrichmentScope(directives []string, aliases ...string) bool {
	lookup := func(name string) (found, enable bool) {
		for _, d := range directives {
			scope, neg := normalizeScope(d)
			if scope == name {
				return true, !neg
			}
		}
		return false, false
	}
	for _, alias := range aliases {
		if found, en := lookup(alias); found {
			return en
		}
	}
	if _, disableAll := lookup("none"); disableAll {
		return false
	}
	if _, enableAll := lookup("all"); enableAll {
		return true
	}
	return false
}

func SaveSBOM(sbm *sbom.SBOM, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("unable to create file %s: %w", file, err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			_ = err
		}
	}()
	encoder, err := cyclonedxjson.NewFormatEncoderWithConfig(cyclonedxjson.DefaultEncoderConfig())
	if err != nil {
		return err
	}

	data, err := format.Encode(*sbm, encoder)
	if err != nil {
		return fmt.Errorf("unable to encode SBOM: %w", err)
	}

	_, err = f.Write(data)
	if err != nil {
		return fmt.Errorf("unable to write to file %s: %w", file, err)
	}

	return nil
}

// maxSBOMBytes caps the in-memory read of a user-supplied SBOM file. 1 GiB
// is well above any realistic real-world SBOM but bounded enough that a
// pathological (or malicious) input can't OOM-kill the scan.
const maxSBOMBytes = 1 << 30 // 1 GiB

func LoadSBOM(inputFile string) (*sbom.SBOM, []InputSbomRef, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to open SBOM file %s: %w", inputFile, err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			_ = err
		}
	}()

	// Bounded read so a runaway or hostile SBOM (e.g. /dev/zero, a
	// symlinked infinite stream, or a genuine multi-GB file) can't
	// exhaust memory. Anything beyond maxSBOMBytes gets a clear error.
	content, err := io.ReadAll(io.LimitReader(file, maxSBOMBytes+1))
	if err != nil {
		return nil, nil, fmt.Errorf("unable to read SBOM file %s: %w", inputFile, err)
	}
	if int64(len(content)) > maxSBOMBytes {
		return nil, nil, fmt.Errorf("SBOM file %s exceeds %d-byte size cap; refusing to load", inputFile, maxSBOMBytes)
	}

	// Parse JSON to extract bom-ref and purl
	var rawSBOM map[string]interface{}
	err = json.Unmarshal(content, &rawSBOM)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to parse SBOM JSON from file %s: %w", inputFile, err)
	}

	var inputSbomRefs []InputSbomRef

	// Extract bom-ref, purl and cpe from every CycloneDX object that can carry
	// them. We read these straight from the raw JSON because Syft only surfaces
	// what its catalogers recognise as packages: CycloneDX components of type
	// "file" (and others), and metadata.component (the object the SBOM
	// describes), carry purls/cpes that never make it into
	// sbm.Artifacts.Packages, so relying on the decoded SBOM alone drops them.
	appendRef := func(component map[string]interface{}) {
		bomRef, _ := component["bom-ref"].(string)
		purl, _ := component["purl"].(string)
		cpe, _ := component["cpe"].(string)
		if purl != "" || cpe != "" {
			inputSbomRefs = append(inputSbomRefs, InputSbomRef{
				SbomRef: bomRef,
				PURL:    purl,
				CPE:     cpe,
			})
		}
	}

	if components, ok := rawSBOM["components"].([]interface{}); ok {
		for _, comp := range components {
			if component, ok := comp.(map[string]interface{}); ok {
				appendRef(component)
			}
		}
	}
	if metadata, ok := rawSBOM["metadata"].(map[string]interface{}); ok {
		if component, ok := metadata["component"].(map[string]interface{}); ok {
			appendRef(component)
		}
	}

	// Reset file pointer to the beginning for Syft to read
	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to reset file pointer: %w", err)
	}

	// Decode SBOM using Syft
	sbm, _, _, err := format.Decode(file)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to decode SBOM from file %s: %w", inputFile, err)
	}

	return sbm, inputSbomRefs, nil
}

func GetCPEDetail(sbm *sbom.SBOM, inputRefs []InputSbomRef) []string {
	var cpes []string
	seen := make(map[string]struct{})

	add := func(cpeStr string) {
		cpeStr = strings.TrimSpace(cpeStr)
		if cpeStr == "" || strings.Contains(cpeStr, ".github/workflows") {
			return
		}
		norm := cpeutils.NormalizeCPEString(cpeStr)
		if _, exists := seen[norm]; exists {
			return
		}
		seen[norm] = struct{}{}
		cpes = append(cpes, cpeStr)
	}

	if sbm != nil {
		if sbm.Artifacts.LinuxDistribution != nil {
			add(sbm.Artifacts.LinuxDistribution.CPEName)
		}

		for p := range sbm.Artifacts.Packages.Enumerate() {
			for _, cpe := range p.CPEs {
				add(cpe.Attributes.BindToFmtString())
			}
		}
	}

	// CycloneDX components of type "file" (and others) carry CPEs that Syft does
	// not surface as packages, so pull them straight from the parsed SBOM
	for _, ref := range inputRefs {
		add(ref.CPE)
	}

	return cpes
}

func GetPURLDetail(sbm *sbom.SBOM, inputRefs []InputSbomRef) []models.PurlDetail {
	var purls []models.PurlDetail
	seen := make(map[string]struct{})

	add := func(purl models.PurlDetail) {
		if purl.Purl == "" || strings.HasPrefix(purl.Purl, "pkg:github") {
			return
		}
		if _, exists := seen[purl.Purl]; exists {
			return
		}
		seen[purl.Purl] = struct{}{}
		purls = append(purls, purl)
	}

	if sbm != nil {
		for p := range sbm.Artifacts.Packages.Enumerate() {
			locations := make([]string, len(p.Locations.ToSlice()))
			for i, l := range p.Locations.ToSlice() {
				locations[i] = l.RealPath
			}

			purlDetail := models.PurlDetail{
				Purl:        p.PURL,
				PackageType: string(p.Type),
				Cataloger:   p.FoundBy,
				Locations:   locations,
			}

			for _, ref := range inputRefs {
				if ref.PURL == p.PURL {
					purlDetail.SbomRef = ref.SbomRef
					break
				}
			}

			add(purlDetail)
		}
	}

	// CycloneDX components of type "file" (and others) carry PURLs that Syft
	// does not surface as packages, so pull them straight from the parsed SBOM
	// — mirrors what GetCPEDetail does for the CPE side.
	for _, ref := range inputRefs {
		add(models.PurlDetail{
			Purl:    ref.PURL,
			SbomRef: ref.SbomRef,
		})
	}

	return purls
}

// GetBatchVulns looks up every purl and also returns the components the API
// could not assess. Those never appear in the findings, so a caller that
// ignores the second return cannot tell a partial scan from a clean one.
func GetBatchVulns(ctx context.Context, purls []models.PurlDetail, iterator func(cur int, total int)) ([]models.ScanResultVulnerabilities, []models.UnprocessedComponent, error) {
	const batchSize = 100

	var vulns []models.ScanResultVulnerabilities
	var unprocessed []models.UnprocessedComponent

	purlStrings := make([]string, 0, len(purls))
	for _, purl := range purls {
		purlStrings = append(purlStrings, purl.Purl)
	}

	total := len(purlStrings)

	for start := 0; start < total; start += batchSize {
		end := min(start+batchSize, total)

		batch := purlStrings[start:end]

		response, err := session.ConnectWithContext(ctx, config.Token()).GetPurls(batch)
		if err != nil {
			return nil, nil, fmt.Errorf("error fetching purls %v: %w", batch, err)
		}

		for _, purlResponse := range response.PurlData {
			for _, vuln := range purlResponse.Vulnerabilities {
				vulns = append(vulns, models.ScanResultVulnerabilities{
					Name:          purlResponse.PurlMeta.Name,
					Version:       purlResponse.PurlMeta.Version,
					CVE:           vuln.Detection,
					FixedVersions: vuln.FixedVersion,
				})
			}
		}

		// Empty against an API predating partial results, where an unusable purl
		// failed the whole batch rather than being reported.
		for _, item := range response.Meta.Unprocessed {
			unprocessed = append(unprocessed, models.UnprocessedComponent{
				Purl:   item.Purl,
				Reason: item.Reason,
			})
		}

		iterator(start, total)
	}

	return vulns, unprocessed, nil
}

func GetVulns(ctx context.Context, purls []models.PurlDetail, iterator func(cur int, total int)) ([]models.ScanResultVulnerabilities, error) {
	var vulns []models.ScanResultVulnerabilities

	i := 0
	for _, purl := range purls {
		i++
		response, err := session.ConnectWithContext(ctx, config.Token()).GetPurl(purl.Purl)
		if err != nil {
			return nil, fmt.Errorf("error fetching purl %s: %v", purl.Purl, err)
		}
		if len(response.Data.Vulnerabilities) > 0 {
			for _, vuln := range response.Data.Vulnerabilities {
				vulns = append(vulns, models.ScanResultVulnerabilities{
					Name:          response.PurlMeta().Name,
					Version:       response.PurlMeta().Version,
					CVE:           vuln.Detection,
					FixedVersions: vuln.FixedVersion,
					PurlDetail:    purl,
				})
			}
		}
		iterator(i, len(purls))
	}

	return vulns, nil
}

func GetOfflineCpeVulns(indices cache.InfoFile, cpes []string, iterator func(cur int, total int), warnOnly bool) ([]models.ScanResultVulnerabilities, error) {
	var vulns []models.ScanResultVulnerabilities
	i := 0
	seen := make(map[string]struct{})

	indexAvailable, err := sync.EnsureIndexSync(indices, "cpecve", true)
	if err != nil {
		if warnOnly {
			fmt.Printf("[WARNING]: %s\n", err.Error())
			return nil, nil
		} else {
			return nil, err
		}
	}

	if !indexAvailable {
		if warnOnly {
			fmt.Printf("[WARNING]: index cpecve is required to proceed\n")
			return nil, nil
		} else {
			return nil, fmt.Errorf("index cpecve is required to proceed")
		}
	}

	for _, cpestring := range cpes {
		i++
		cpe, err := cpeuri.ToStruct(cpestring)
		if err != nil {
			return nil, err
		}

		results, _, err := db.CPESearch("cpecve", *cpe)
		if err != nil {
			return nil, err
		}

		cves, err := cpeutils.Process(cpe, results)
		if err != nil {
			return nil, err
		}

		for _, cve := range cves {
			key := cpestring + "|" + cve
			if _, exists := seen[key]; !exists {
				vulns = append(vulns, models.ScanResultVulnerabilities{
					Name:    cpeuri.RemoveSlashes(cpe.Product),
					Version: cpeuri.RemoveSlashes(cpe.Version),
					CVE:     cve,
					CPE:     cpestring,
				})
				seen[key] = struct{}{}
			}
		}
		iterator(i, len(cpes))
	}

	return vulns, nil
}

func GetOfflineVulns(indices cache.InfoFile, purls []models.PurlDetail, iterator func(cur int, total int), warnOnly bool) ([]models.ScanResultVulnerabilities, error) {
	var vulns []models.ScanResultVulnerabilities

	i := 0
	for _, purl := range purls {
		i++
		instance, err := packageurl.FromString(purl.Purl)
		if err != nil {
			return nil, err
		}

		/*
			if packages.IsOS(instance) {
				return nil, fmt.Errorf("operating system package support coming soon")
			}
		*/

		indexName := packages.IndexFromInstance(instance)

		indexAvailable, err := sync.EnsureIndexSync(indices, indexName, true)
		if err != nil {
			if warnOnly {
				fmt.Printf("[WARNING]: %s\n", err.Error())
				continue
			} else {
				return nil, err
			}
		}

		if !indexAvailable {
			if warnOnly {
				fmt.Printf("[WARNING]: index %s is required to PURL %s \n", indexName, purl.Purl)
				continue
			} else {
				return nil, fmt.Errorf("index %s is required to proceed", instance.Type)
			}
		}

		index := indices.GetIndex(indexName)

		results, _, err := db.PURLSearch(index.Name, instance)
		if err != nil {
			return nil, err
		}

		// loop through results and add to vulns
		for _, purlEntry := range results {
			for _, vuln := range purlEntry.Vulnerabilities {
				vulns = append(vulns, models.ScanResultVulnerabilities{
					Name:          purlEntry.Name,
					Version:       purlEntry.Version,
					CVE:           vuln.Detection,
					FixedVersions: vuln.FixedVersion,
					PurlDetail:    purl,
				})
			}
		}

		iterator(i, len(purls))
	}

	return vulns, nil
}

func GetMeta(ctx context.Context, vulns []models.ScanResultVulnerabilities) ([]models.ScanResultVulnerabilities, error) {
	for i, vuln := range vulns {
		// Honour SIGINT between iterations — each iteration is a fresh
		// HTTP call so the user can cancel a long scan promptly.
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		nvd2Response, err := session.ConnectWithContext(ctx, config.Token()).GetIndexVulncheckNvd2(sdk.IndexQueryParameters{Cve: vuln.CVE})
		if err != nil {
			return nil, err
		}

		vulns[i].InKEV = nvd2Response.Data[0].VulncheckKEVExploitAdd != nil
		vulns[i].Published = *nvd2Response.Data[0].Published
		vulns[i].CVSSBaseScore = baseScore(nvd2Response.Data[0])
		vulns[i].CVSSTemporalScore = temporalScore(nvd2Response.Data[0])
		vulns[i].Metrics = nvd2Response.Data[0].Metrics
		vulns[i].Weaknesses = nvd2Response.Data[0].Weaknesses

	}
	return vulns, nil
}

// GetOfflineMeta enriches each vuln with CVSS / KEV / description data from
// the vulncheck-nvd2 index. The second return value reports whether that index
// was actually available; callers use it to surface a hint to the user when
// the table would otherwise show empty score columns. Without warnOnly the
// missing-index condition is a hard error; with warnOnly the input vulns are
// passed through unchanged so the caller can still render what it found.
func GetOfflineMeta(indices cache.InfoFile, vulns []models.ScanResultVulnerabilities, warnOnly bool) ([]models.ScanResultVulnerabilities, bool, error) {
	indexAvailable, err := sync.EnsureIndexSync(indices, "vulncheck-nvd2", true)
	if err != nil {
		if warnOnly {
			fmt.Printf("[WARNING]: %s\n", err.Error())
			return vulns, false, nil
		}
		return nil, false, err
	}

	if !indexAvailable {
		if warnOnly {
			fmt.Printf("[WARNING]: index vulncheck-nvd2 is required to proceed\n")
			return vulns, false, nil
		}
		return nil, false, fmt.Errorf("index vulncheck-nvd2 is required to proceed")
	}

	for i, vuln := range vulns {
		nvd2Response, err := db.MetaByCVE(vuln.CVE)
		if err != nil {
			continue
		}

		if len(nvd2Response.Data) > 0 {
			vulns[i].InKEV = nvd2Response.Data[0].VulncheckKEVExploitAdd != nil
			vulns[i].Published = *nvd2Response.Data[0].Published
			vulns[i].CVSSBaseScore = baseScore(nvd2Response.Data[0])
			vulns[i].CVSSTemporalScore = temporalScore(nvd2Response.Data[0])
			vulns[i].Metrics = nvd2Response.Data[0].Metrics
			vulns[i].Weaknesses = nvd2Response.Data[0].Weaknesses
			vulns[i].Description = nvd2Response.Description
		}
	}
	return vulns, true, nil
}

func baseScore(item client.ApiNVD20CVEExtended) string {
	if item.Metrics == nil {
		return "n/a"
	}
	var score *float32
	if (item.Metrics.CvssMetricV31 != nil) && (len(*item.Metrics.CvssMetricV31) > 0) {
		score = (*item.Metrics.CvssMetricV31)[0].CvssData.BaseScore
	}

	if score == nil && (item.Metrics.CvssMetricV30 != nil) && (len(*item.Metrics.CvssMetricV30) > 0) {
		score = (*item.Metrics.CvssMetricV30)[0].CvssData.BaseScore
	}

	if score == nil && (item.Metrics.CvssMetricV2 != nil) && (len(*item.Metrics.CvssMetricV2) > 0) {
		score = (*item.Metrics.CvssMetricV2)[0].CvssData.BaseScore
	}

	if score == nil {
		return "n/a"
	}

	return formatSingleDecimal(score)
}

func temporalScore(item client.ApiNVD20CVEExtended) string {
	if item.Metrics == nil {
		return "n/a"
	}
	var score *float32

	if item.Metrics.TemporalCVSSV31 != nil {
		score = item.Metrics.TemporalCVSSV31.TemporalScore
	}

	if score == nil && item.Metrics.TemporalCVSSV31Secondary != nil && len(*item.Metrics.TemporalCVSSV31Secondary) > 0 {
		score = (*item.Metrics.TemporalCVSSV31Secondary)[0].TemporalScore
	}

	if score == nil && item.Metrics.CvssMetricV30 != nil && len(*item.Metrics.CvssMetricV30) > 0 {
		score = (*item.Metrics.CvssMetricV30)[0].CvssData.TemporalScore
	}

	if score == nil && item.Metrics.TemporalCVSSV30Secondary != nil && len(*item.Metrics.TemporalCVSSV30Secondary) > 0 {
		score = (*item.Metrics.TemporalCVSSV30Secondary)[0].TemporalScore
	}

	if score == nil && item.Metrics.TemporalCVSSV30 != nil {
		score = item.Metrics.TemporalCVSSV30.TemporalScore
	}

	if score == nil && item.Metrics.CvssMetricV2 != nil && len(*item.Metrics.CvssMetricV2) > 0 {
		score = (*item.Metrics.CvssMetricV2)[0].CvssData.TemporalScore
	}

	if score == nil && item.Metrics.TemporalCVSSV2 != nil {
		score = item.Metrics.TemporalCVSSV2.TemporalScore
	}

	if score == nil && item.Metrics.TemporalCVSSV2Secondary != nil && len(*item.Metrics.TemporalCVSSV2Secondary) > 0 {
		score = (*item.Metrics.TemporalCVSSV2Secondary)[0].TemporalScore
	}

	if score == nil {
		return "n/a"
	}

	return formatSingleDecimal(score)
}

func formatSingleDecimal(value *float32) string {
	return fmt.Sprintf("%.1f", *value)
}
