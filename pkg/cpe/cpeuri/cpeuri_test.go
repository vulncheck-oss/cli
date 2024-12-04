package cpeuri

import (
	"github.com/vulncheck-oss/cli/pkg/cpe/cpeutils"
	"reflect"
	"testing"
)

func TestToStruct(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *cpeutils.CPE
		wantErr bool
	}{
		{
			name:  "Valid CPE URI",
			input: "cpe:/a:vendor:product:version:update",
			want: &cpeutils.CPE{
				Part:            "a",
				Vendor:          "vendor",
				Product:         "product",
				Version:         "version",
				Update:          "update",
				Edition:         "*",
				Language:        "*",
				SoftwareEdition: "*",
				TargetSoftware:  "*",
				TargetHardware:  "*",
				Other:           "*",
			},
			wantErr: false,
		},
		{
			name:    "Invalid CPE string",
			input:   "invalid:cpe:string",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "CPE with * as vendor",
			input:   "cpe:/a:*:product:1.0",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToStruct(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsCPEURIString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"Valid CPE URI", "cpe:/a:vendor:product:version", true},
		{"Invalid prefix", "invalid:/a:vendor:product:version", false},
		{"Not enough components", "cpe:/a:vendor:product", false},
		{"Non-ASCII characters", "cpe:/a:vend□r:product:version", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCPEURIString(tt.input); got != tt.want {
				t.Errorf("IsCPEURIString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsCPEFormattedString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"Valid CPE 2.3", "cpe:2.3:a:vendor:product:version:update:edition:lang:sw_edition:target_sw:target_hw:other", true},
		{"Invalid prefix", "invalid:2.3:a:vendor:product:version:update:edition:lang:sw_edition:target_sw:target_hw:other", false},
		{"Not enough components", "cpe:2.3:a:vendor:product", false},
		{"Non-ASCII characters", "cpe:2.3:a:vend□r:product:version:update:edition:lang:sw_edition:target_sw:target_hw:other", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCPEFormattedString(tt.input); got != tt.want {
				t.Errorf("IsCPEFormattedString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnbindCPEFormattedString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    cpeutils.CPE
		wantErr bool
	}{
		{
			name:  "Valid CPE 2.3",
			input: "cpe:2.3:a:vendor:product:version:update:edition:lang:sw_edition:target_sw:target_hw:other",
			want: cpeutils.CPE{
				Part:            "a",
				Vendor:          "vendor",
				Product:         "product",
				Version:         "version",
				Update:          "update",
				Edition:         "edition",
				Language:        "lang",
				SoftwareEdition: "sw_edition",
				TargetSoftware:  "target_sw",
				TargetHardware:  "target_hw",
				Other:           "other",
			},
			wantErr: false,
		},
		{
			name:    "Invalid CPE 2.3 (not enough components)",
			input:   "cpe:2.3:a:vendor:product",
			want:    cpeutils.CPE{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnbindCPEFormattedString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnbindCPEFormattedString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnbindCPEFormattedString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnbindCPEURIString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    cpeutils.CPE
		wantErr bool
	}{
		{
			name:  "Valid CPE URI",
			input: "cpe:/a:vendor:product:version:update",
			want: cpeutils.CPE{
				Part:            "a",
				Vendor:          "vendor",
				Product:         "product",
				Version:         "version",
				Update:          "update",
				Edition:         "*",
				Language:        "*",
				SoftwareEdition: "",
				TargetSoftware:  "",
				TargetHardware:  "",
				Other:           "",
			},
			wantErr: false,
		},
		{
			name:    "Invalid CPE URI (non-ASCII characters)",
			input:   "cpe:/a:vend□r:product:version",
			want:    cpeutils.CPE{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnbindCPEURIString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnbindCPEURIString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnbindCPEURIString() = %v, want %v", got, tt.want)
			}
		})
	}
}
