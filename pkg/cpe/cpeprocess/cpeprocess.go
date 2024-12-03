package cpeprocess

import (
	"encoding/json"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpemozilla"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpenginx"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpetypes"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpeutils"
)

func Process(cpe cpetypes.CPE, entries []interface{}) ([]string, error) {
	if cpe.IsMozilla() {
		mozillaAdvisories := make(cpetypes.MozillaAdvisories, len(entries))
		for i, entry := range entries {
			jsonData, err := json.Marshal(entry)
			if err != nil {
				return nil, err
			}

			var advisory cpetypes.MozillaAdvisory
			if err := json.Unmarshal(jsonData, &advisory); err != nil {
				return nil, err
			}

			mozillaAdvisories[i] = advisory
		}
		return cpemozilla.Process(cpe, mozillaAdvisories)
	}

	if cpe.IsNginx() {
		nginxAdvisories := make([]cpetypes.NginxAdvisory, len(entries))
		for i, entry := range entries {
			jsonData, err := json.Marshal(entry)
			if err != nil {
				return nil, err
			}

			var advisory cpetypes.NginxAdvisory
			if err := json.Unmarshal(jsonData, &advisory); err != nil {
				return nil, err
			}

			nginxAdvisories[i] = advisory
		}
		return cpenginx.Process(cpe, nginxAdvisories)
	}

	cveFindings := make([]string, 0)

	if cpeutils.IsParseableVersion(cpetypes.Unquote(cpe.Version)) {
		for _, entry := range entries {
			jsonData, err := json.Marshal(entry)
			if err != nil {
				return nil, err
			}

			var vuln cpetypes.CPEVulnerabilities
			if err := json.Unmarshal(jsonData, &vuln); err != nil {
				return nil, err
			}

			cmpr, err := cpeutils.CompareVersions(cpetypes.Unquote(cpe.Version), cpetypes.Unquote(vuln.Version))
			if err == nil {
				if cmpr == 0 {
					cveFindings = append(cveFindings, vuln.Cves...)
				}
			}
		}
	} else {

		for _, entry := range entries {
			jsonData, err := json.Marshal(entry)
			if err != nil {
				return nil, err
			}

			var vuln cpetypes.CPEVulnerabilities
			if err := json.Unmarshal(jsonData, &vuln); err != nil {
				return nil, err
			}
			cveFindings = append(cveFindings, vuln.Cves...)

		}
	}

	cveFindings = RemoveDuplicatesUnordered(cveFindings)

	return cveFindings, nil
}
func RemoveDuplicatesUnordered(elements []string) []string {
	encountered := map[string]bool{}

	func() {
		for v := range elements {
			encountered[elements[v]] = true
		}
	}()

	result := []string{}
	for key := range encountered {
		result = append(result, key)
	}
	return result
}
