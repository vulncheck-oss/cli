package cpeparse

import (
	"fmt"
	"github.com/facebookincubator/nvdtools/wfn"
	"strings"
)

func Parse(s string) (*wfn.Attributes, error) {

	var attr *wfn.Attributes
	var err error

	// Check if the CPE string is in 2.2 format
	if strings.HasPrefix(s, "cpe:/") {
		// Convert 2.2 to 2.3 format
		attr, err = wfn.UnbindURI(s)
	} else {
		// Assume 2.3 format
		attr, err = wfn.UnbindFmtString(s)
	}

	if err != nil {
		return nil, fmt.Errorf("invalid CPE string: %v", err)
	}

	// if the CPE Vendor or Product is *, return an error
	if attr.Vendor == "*" || attr.Product == "*" {
		return nil, fmt.Errorf("CPE Vendor or Product cannot be *")
	}

	return attr, nil
}
