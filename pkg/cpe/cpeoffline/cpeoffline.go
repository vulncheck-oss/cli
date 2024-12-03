package cpeoffline

import (
	"fmt"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpeutils"
	"strings"
)

func Query(cpe *cpeutils.CPE) (string, error) {

	if cpe.Vendor == "*" && cpe.Product == "*" {
		return "", fmt.Errorf("need at least vendor or product specified")
	}

	return buildCPEQuery(cpe, !cpeutils.IsParseableVersion(cpeutils.Unquote(cpe.Version)))
}

func addCondition(field, value string) string {
	if value != "*" {
		return fmt.Sprintf(`(.%s == %q or .%s == "*")`, field, value, field)
	}
	return ""
}

func buildCPEQuery(cpe *cpeutils.CPE, queryVersion bool) (string, error) {
	if cpe.Vendor == "*" && cpe.Product == "*" {
		return "", fmt.Errorf("need at least vendor or product specified")
	}

	var conditions []string

	fields := []struct {
		name  string
		value string
	}{
		{"vendor", cpe.Vendor},
		{"product", cpe.Product},
		{"update", cpe.Update},
		{"edition", cpe.Edition},
		{"language", cpe.Language},
		{"sw_edition", cpe.SoftwareEdition},
		{"target_sw", cpe.TargetSoftware},
		{"target_hw", cpe.TargetHardware},
		{"other", cpe.Other},
	}

	// Add version field only if withVersion is true
	if queryVersion {
		fields = append(fields, struct{ name, value string }{"version", cpeutils.Unquote(cpe.Version)})
	}

	for _, field := range fields {
		if condition := addCondition(field.name, field.value); condition != "" {
			conditions = append(conditions, condition)
		}
	}

	if len(conditions) == 0 {
		return "true", nil
	}

	query := strings.Join(conditions, " and ")
	return query, nil
}
