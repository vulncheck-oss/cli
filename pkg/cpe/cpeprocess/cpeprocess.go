package cpeprocess

import (
	"encoding/json"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpemozilla"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpetypes"
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
	return []string{}, nil
}
