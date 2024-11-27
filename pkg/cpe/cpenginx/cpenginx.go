package cpenginx

import (
	"github.com/vulncheck-oss/cli/pkg/cpe/cpetypes"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpeutils"
	"strings"
)

func Process(cpe cpetypes.CPE, results []cpetypes.NginxAdvisory) ([]string, error) {
	var CVEs []string

	for _, result := range results {
		isVulnVer := false

		for _, vvr := range result.VulnVersions {
			vlh := strings.Split(vvr, "-")
			if len(vlh) != 2 {
				// log
				continue
			}
			cmprGE, err := cpeutils.CompareVersions(cpetypes.Unquote(cpe.Version), vlh[0])
			if err != nil {
				// log
				continue
			}
			cmprLE, err := cpeutils.CompareVersions(cpetypes.Unquote(cpe.Version), vlh[1])
			if err != nil {
				// log
				continue
			}
			if cmprLE <= 0 && cmprGE >= 0 {
				isVulnVer = true
				break
			}
		}

		if !isVulnVer {
			continue
		}
		isNotVulnVer := false
		var nvv string
		for _, nvv = range result.NotVulnVersions {

			if nvv[len(nvv)-1] != '+' {
				// log
				continue
			}

			nvv := nvv[:len(nvv)-1]
			i := strings.LastIndex(nvv, ".")
			if i <= 0 {
				continue
			}
			cmpv := nvv[:i]
			i = strings.LastIndex(cpetypes.Unquote(cpe.Version), ".")
			if i <= 0 {
				break
			}
			cmpvuq := cpetypes.Unquote(cpe.Version)[:i]

			if cmpv != cmpvuq {
				continue
			}
			fcmp, err := cpeutils.CompareVersions(cpetypes.Unquote(cpe.Version), nvv)
			if err != nil {
				continue
			}
			if fcmp >= 0 {
				isNotVulnVer = true
				break
			}
		}
		if isNotVulnVer {
			continue
		}
		CVEs = append(CVEs, result.CVE...)

	}

	return CVEs, nil
}
