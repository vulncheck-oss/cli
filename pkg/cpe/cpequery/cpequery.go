package cpequery

import (
	"fmt"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpetypes"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpeutils"
	"strings"
)

func Query(cpe cpetypes.CPE) (string, error) {
	if cpe.Product == "" {
		return "true", nil
	}
	if cpe.IsMozilla() {
		return fmt.Sprintf(".products | any(. == %q)", cpe.ProductUcFirst()), nil
	}

	if cpe.Vendor == "*" && cpe.Product == "*" {
		return "", fmt.Errorf("need at least vendor or product specified")
	}

	// if IsParseableVersion - do NOT query for the version in the JSON
	return BuildCPEQuery(cpe, !cpeutils.IsParseableVersion(cpetypes.Unquote(cpe.Version)))
}

func addCondition(field, value string) string {
	if value != "*" {
		return fmt.Sprintf(`(.%s == %q or .%s == "*")`, field, value, field)
	}
	return ""
}

func BuildCPEQuery(cpe cpetypes.CPE, queryVersion bool) (string, error) {
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
		fields = append(fields, struct{ name, value string }{"version", cpetypes.Unquote(cpe.Version)})
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
