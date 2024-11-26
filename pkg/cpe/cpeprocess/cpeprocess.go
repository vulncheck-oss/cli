package cpeprocess

import (
	"github.com/vulncheck-oss/cli/pkg/cpe/cpemozilla"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpetypes"
	"github.com/vulncheck-oss/cli/pkg/search"
)

func Process(cpe cpetypes.CPE, entries search.AdvisoryEntries) ([]string, error) {
	if cpe.IsMozilla() {
		return cpemozilla.Process(cpe, entries)
	}
	return entries.CVES(), nil
}
