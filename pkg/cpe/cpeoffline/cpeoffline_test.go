package cpeoffline

import (
	"github.com/vulncheck-oss/cli/pkg/cpe/cpeutils"
	"testing"
)

func TestQuery(t *testing.T) {
	tests := []struct {
		name    string
		cpe     *cpeutils.CPE
		want    string
		wantErr bool
	}{
		{
			name: "Valid CPE",
			cpe: &cpeutils.CPE{
				Vendor:  "vendor",
				Product: "product",
				Version: "1.0",
			},
			want:    `(.vendor == "vendor" or .vendor == "*") and (.product == "product" or .product == "*") and (.update == "" or .update == "*") and (.edition == "" or .edition == "*") and (.language == "" or .language == "*") and (.sw_edition == "" or .sw_edition == "*") and (.target_sw == "" or .target_sw == "*") and (.target_hw == "" or .target_hw == "*") and (.other == "" or .other == "*")`,
			wantErr: false,
		},
		{
			name: "CPE with all fields",
			cpe: &cpeutils.CPE{
				Vendor:          "vendor",
				Product:         "product",
				Version:         "1.0",
				Update:          "update",
				Edition:         "edition",
				Language:        "language",
				SoftwareEdition: "sw_edition",
				TargetSoftware:  "target_sw",
				TargetHardware:  "target_hw",
				Other:           "other",
			},
			want: `(.vendor == "vendor" or .vendor == "*") and (.product == "product" or .product == "*") and ` +
				`(.update == "update" or .update == "*") and (.edition == "edition" or .edition == "*") and ` +
				`(.language == "language" or .language == "*") and (.sw_edition == "sw_edition" or .sw_edition == "*") and ` +
				`(.target_sw == "target_sw" or .target_sw == "*") and (.target_hw == "target_hw" or .target_hw == "*") and ` +
				`(.other == "other" or .other == "*")`,
			wantErr: false,
		},
		{
			name: "CPE with wildcard version",
			cpe: &cpeutils.CPE{
				Vendor:  "vendor",
				Product: "product",
				Version: "x.y",
			},
			want:    `(.vendor == "vendor" or .vendor == "*") and (.product == "product" or .product == "*") and (.update == "" or .update == "*") and (.edition == "" or .edition == "*") and (.language == "" or .language == "*") and (.sw_edition == "" or .sw_edition == "*") and (.target_sw == "" or .target_sw == "*") and (.target_hw == "" or .target_hw == "*") and (.other == "" or .other == "*") and (.version == "x.y" or .version == "*")`,
			wantErr: false,
		},
		{
			name: "Invalid CPE (both vendor and product are wildcards)",
			cpe: &cpeutils.CPE{
				Vendor:  "*",
				Product: "*",
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Query(tt.cpe)
			if (err != nil) != tt.wantErr {
				t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddCondition(t *testing.T) {
	tests := []struct {
		name  string
		field string
		value string
		want  string
	}{
		{
			name:  "Non-wildcard value",
			field: "vendor",
			value: "test",
			want:  `(.vendor == "test" or .vendor == "*")`,
		},
		{
			name:  "Wildcard value",
			field: "product",
			value: "*",
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addCondition(tt.field, tt.value); got != tt.want {
				t.Errorf("addCondition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildCPEQuery(t *testing.T) {
	tests := []struct {
		name         string
		cpe          *cpeutils.CPE
		queryVersion bool
		want         string
		wantErr      bool
	}{
		{
			name: "Valid CPE without version",
			cpe: &cpeutils.CPE{
				Vendor:  "vendor",
				Product: "product",
			},
			queryVersion: false,
			want:         `(.vendor == "vendor" or .vendor == "*") and (.product == "product" or .product == "*") and (.update == "" or .update == "*") and (.edition == "" or .edition == "*") and (.language == "" or .language == "*") and (.sw_edition == "" or .sw_edition == "*") and (.target_sw == "" or .target_sw == "*") and (.target_hw == "" or .target_hw == "*") and (.other == "" or .other == "*")`,
			wantErr:      false,
		},
		{
			name: "Valid CPE with version",
			cpe: &cpeutils.CPE{
				Vendor:  "vendor",
				Product: "product",
				Version: "1.0",
			},
			queryVersion: true,
			want:         `(.vendor == "vendor" or .vendor == "*") and (.product == "product" or .product == "*") and (.update == "" or .update == "*") and (.edition == "" or .edition == "*") and (.language == "" or .language == "*") and (.sw_edition == "" or .sw_edition == "*") and (.target_sw == "" or .target_sw == "*") and (.target_hw == "" or .target_hw == "*") and (.other == "" or .other == "*") and (.version == "1.0" or .version == "*")`,
			wantErr:      false,
		},
		{
			name: "Invalid CPE (both vendor and product are wildcards)",
			cpe: &cpeutils.CPE{
				Vendor:  "*",
				Product: "*",
			},
			queryVersion: false,
			want:         "",
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildCPEQuery(tt.cpe, tt.queryVersion)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildCPEQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("buildCPEQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
