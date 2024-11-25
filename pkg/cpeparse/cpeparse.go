package cpeparse

import (
	"fmt"
	"strings"
	"unicode"
)

type CPE struct {
	Part            string `json:"part"`
	Vendor          string `json:"vendor"`
	Product         string `json:"product"`
	Version         string `json:"version"`
	Update          string `json:"update"`
	Edition         string `json:"edition"`
	Language        string `json:"language"`
	SoftwareEdition string `json:"sw_edition"`
	TargetSoftware  string `json:"target_sw"`
	TargetHardware  string `json:"target_hw"`
	Other           string `json:"other"`
}

func Parse(s string) (*CPE, error) {

	var cpe CPE
	var err error

	cpe, err = GetCPEStructFromString(s)

	if err != nil {
		return nil, fmt.Errorf("invalid CPE string: %v", err)
	}

	if cpe.Vendor == "*" || cpe.Product == "*" {
		return nil, fmt.Errorf("CPE Vendor or Product cannot be *")
	}

	return &cpe, nil
}

func GetCPEStructFromString(s string) (CPE, error) {
	if IsCPEFormattedString(s) {
		c, e := UnbindCPEFormattedString(s)
		if e != nil {
			return c, e
		}
		ConvertEmptyToAny(&c)
		return c, nil
	}
	if IsCPEURIString(s) {
		c, e := UnbindCPEURIString(s)
		if e != nil {
			return c, e
		}
		ConvertEmptyToAny(&c)
		return c, nil
	}
	return CPE{}, fmt.Errorf("unrecognized cpe binding form")
}

// IsCPEURIString determines if a string looks like it is a CPE bound as 2.2
// URI
func IsCPEURIString(s string) bool {
	if !isAllASCII(s) {
		return false
	}
	if !strings.HasPrefix(s, "cpe:/a") &&
		!strings.HasPrefix(s, "cpe:/h") &&
		!strings.HasPrefix(s, "cpe:/o") {
		return false
	}
	// meant as a weak check, ignores escapes.
	if strings.Count(s, ":") < 4 {
		return false
	}
	return true
}

// ConvertEmptyToAny is used to fill in empty fields with ANY (i.e., "*"). This
// is primarily used when CPE URI binding string is the input.
func ConvertEmptyToAny(cpe *CPE) {
	if cpe.Part == "" {
		cpe.Part = "*"
	}
	if cpe.Vendor == "" {
		cpe.Vendor = "*"
	}
	if cpe.Product == "" {
		cpe.Product = "*"
	}
	if cpe.Version == "" {
		cpe.Version = "*"
	}
	if cpe.Update == "" {
		cpe.Update = "*"
	}
	if cpe.Edition == "" {
		cpe.Edition = "*"
	}
	if cpe.Language == "" {
		cpe.Language = "*"
	}
	if cpe.SoftwareEdition == "" {
		cpe.SoftwareEdition = "*"
	}
	if cpe.TargetSoftware == "" {
		cpe.TargetSoftware = "*"
	}
	if cpe.TargetHardware == "" {
		cpe.TargetHardware = "*"
	}
	if cpe.Other == "" {
		cpe.Other = "*"
	}
}

// IsCPEFormattedString determines if a string looks like it is a CPE bound as
// 2.3 formatted string.
func IsCPEFormattedString(s string) bool {
	if !isAllASCII(s) {
		return false
	}
	if !strings.HasPrefix(s, "cpe:2.3:a:") &&
		!strings.HasPrefix(s, "cpe:2.3:h:") &&
		!strings.HasPrefix(s, "cpe:2.3:o:") {
		return false
	}
	// meant as a weak check, ignores escapes.
	if strings.Count(s, ":") < 5 {
		return false
	}
	return true
}

func isAllASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

// UnbindCPEFormattedString attempts to unbind a cpe 2.3 formatted string to
// a CPE struct
func UnbindCPEFormattedString(str string) (CPE, error) {

	if !isAllASCII(str) {
		return CPE{}, fmt.Errorf("cpe string contains non-ASCII chars")
	}

	cpe := CPE{}
	for a := 2; a <= 12; a++ {
		v := getCompFS(str, a)
		v, err := unbindValueFS(v)
		if err != nil {
			return CPE{}, err
		}

		switch a {
		case 2:
			cpe.Part = v

		case 3:
			cpe.Vendor = v

		case 4:
			cpe.Product = v

		case 5:
			cpe.Version = v

		case 6:
			cpe.Update = v

		case 7:
			cpe.Edition = v

		case 8:
			cpe.Language = v

		case 9:
			cpe.SoftwareEdition = v

		case 10:
			cpe.TargetSoftware = v

		case 11:
			cpe.TargetHardware = v

		case 12:
			cpe.Other = v
		}
	}
	return cpe, nil
}

func getCompFS(str string, i int) string {
	fcount := 0
	sidx := 0

	if i < 0 || i > 12 {
		return ""
	}

	for idx, v := range str {
		if v == ':' && (idx == 0 || str[idx-1] != '\\') {
			if i == fcount {
				return str[sidx:idx]
			}
			fcount++
			sidx = idx + 1
		}

	}
	return str[sidx:]
}

func unbindValueFS(str string) (string, error) {
	switch str {
	case "*":
		return "*", nil

	case "":
		return "*", nil

	case "-":
		return "-", nil

	}
	return addQuoting(str)
}

func addQuoting(str string) (string, error) {
	result := ""
	idx := 0
	embedded := false

	for idx < len(str) {
		c := str[idx]

		// alphanum or _
		if (c >= 'A' && c <= 'Z') ||
			(c >= 'a' && c <= 'z') ||
			(c >= '0' && c <= '9') ||
			c == '_' {
			result = fmt.Sprintf("%s%c", result, c)
			idx++
			embedded = true
			continue
		}

		// handle escaping
		if c == '\\' {
			// the 2.3 specification does not do this check, but should
			if len(str)-idx > 1 {
				result = fmt.Sprintf("%s%c%c", result, c, str[idx+1])
				idx += 2
				embedded = true
				continue
			} else {
				return "", fmt.Errorf("escaping character length failure")
			}
		}

		// wildcard must be at start or end
		if c == '*' {
			if idx == 0 || idx == len(str)-1 {
				result = fmt.Sprintf("%s%c", result, c)
				idx++
				embedded = true
				continue
			} else {
				return "",
					fmt.Errorf("unquoted asterisk not at start or end of string")
			}
		}

		// handle ? modifier
		if c == '?' {
			if ((idx == 0) || (idx == len(str)-1)) ||
				(!embedded && str[idx-1] == '?') ||
				(embedded && str[idx+1] == '?') {
				result = fmt.Sprintf("%s%c", result, c)
				idx++
				embedded = false
				continue
			} else {
				return "",
					fmt.Errorf("unquoted ? must be at start or end of string")
			}
		}

		// all others must be escaped
		result = fmt.Sprintf("%s\\%c", result, c)
		idx++
		embedded = true
	}
	return result, nil
}

// UnbindCPEURIString is based on the pseudocode algorithm found in 6.1.3.2 of
// NISTIR 7695
func UnbindCPEURIString(uri string) (CPE, error) {
	cpe := CPE{}

	if !isAllASCII(uri) {
		return CPE{}, fmt.Errorf("cpe string contains non-ASCII chars")
	}

	var err error
	for i := 1; i <= 7; i++ {
		v := getCompURI(uri, i)
		switch i {
		case 1:
			cpe.Part, err = decodeCompURI(v)
			if err != nil {
				return CPE{}, err
			}

		case 2:
			cpe.Vendor, err = decodeCompURI(v)
			if err != nil {
				return CPE{}, err
			}

		case 3:
			cpe.Product, err = decodeCompURI(v)
			if err != nil {
				return CPE{}, err
			}

		case 4:
			cpe.Version, err = decodeCompURI(v)
			if err != nil {
				return CPE{}, err
			}

		case 5:
			cpe.Update, err = decodeCompURI(v)
			if err != nil {
				return CPE{}, err
			}

		case 6:
			if v == "" || v == "-" || v[0] != '~' {
				cpe.Edition, err = decodeCompURI(v)
				if err != nil {
					return CPE{}, err
				}

			} else {
				err = unpackCompURI(v, &cpe)
				if err != nil {
					return CPE{}, err
				}
			}

		case 7:
			cpe.Language, err = decodeCompURI(v)
			if err != nil {
				return CPE{}, err
			}
		}
	}

	return cpe, nil
}

func decodeCompURI(s string) (string, error) {
	// ANY case
	if s == "" {
		return "*", nil
	}

	// NA case
	if s == "-" {
		return "-", nil
	}

	s = strings.ToLower(s)
	result := ""
	idx := 0
	embedded := false

	for idx < len(s) {
		c := s[idx]

		if c == '.' || c == '-' || c == '~' {
			result = fmt.Sprintf("%s\\%c", result, c)
			idx++
			embedded = true
			continue
		}

		if c != '%' {
			result = fmt.Sprintf("%s%c", result, c)
			idx++
			embedded = true
			continue
		}
		if len(s)-idx <= 2 {
			return "", fmt.Errorf("pct encoded length error")
		}

		form := s[idx : idx+3]
		switch form {
		case "%01":
			if ((idx == 0) || (idx == (len(s) - 3))) ||
				(!embedded && s[idx-3:idx] == "%01") ||
				(embedded && len(s) >= idx+6 && s[idx+3:idx+6] == "%01") {
				result = fmt.Sprintf("%s?", result)
				idx += 3
				continue
			} else {
				return "", fmt.Errorf("%%01 error encoding")
			}

		case "%02":
			if idx == 0 || idx == (len(s)-3) {
				result = fmt.Sprintf("%s*", result)
			} else {
				return "", fmt.Errorf("%%02 encoding length error")
			}

		case "%21":
			result = fmt.Sprintf("%s\\!", result)
		case "%22":
			result = fmt.Sprintf("%s\\\"", result)
		case "%23":
			result = fmt.Sprintf("%s\\#", result)
		case "%24":
			result = fmt.Sprintf("%s\\$", result)
		case "%25":
			result = fmt.Sprintf("%s\\%%", result)
		case "%26":
			result = fmt.Sprintf("%s\\&", result)
		case "%27":
			result = fmt.Sprintf("%s\\'", result)
		case "%28":
			result = fmt.Sprintf("%s\\(", result)
		case "%29":
			result = fmt.Sprintf("%s\\)", result)
		case "%2a":
			result = fmt.Sprintf("%s\\*", result)
		case "%2b":
			result = fmt.Sprintf("%s\\+", result)
		case "%2c":
			result = fmt.Sprintf("%s\\,", result)
		case "%2f":
			result = fmt.Sprintf("%s\\/", result)
		case "%3a":
			result = fmt.Sprintf("%s\\:", result)
		case "%3b":
			result = fmt.Sprintf("%s\\;", result)
		case "%3c":
			result = fmt.Sprintf("%s\\<", result)
		case "%3d":
			result = fmt.Sprintf("%s\\=", result)
		case "%3e":
			result = fmt.Sprintf("%s\\>", result)
		case "%3f":
			result = fmt.Sprintf("%s\\?", result)
		case "%40":
			result = fmt.Sprintf("%s\\@", result)
		case "%5b":
			result = fmt.Sprintf("%s\\[", result)
		case "%5c":
			result = fmt.Sprintf("%s\\\\", result)
		case "%5d":
			result = fmt.Sprintf("%s\\]", result)
		case "%5e":
			result = fmt.Sprintf("%s\\^", result)
		case "%60":
			result = fmt.Sprintf("%s\\`", result)
		case "%7b":
			result = fmt.Sprintf("%s\\{", result)
		case "%7c":
			result = fmt.Sprintf("%s\\|", result)
		case "%7d":
			result = fmt.Sprintf("%s\\}", result)
		case "%7e":
			result = fmt.Sprintf("%s\\~", result)

		default:
			return "", fmt.Errorf("unrecognized pct encoded")
		}
		idx += 3
		embedded = true
	}
	return result, nil
}

func unpackCompURI(s string, cpe *CPE) error {

	if len(s) <= 1 {
		return fmt.Errorf("len(s) <= 1 at start")
	}

	// Is 1 ok here?
	start := 1
	end := offsetIndexOf(s, '~', start)
	if end == -1 {
		return fmt.Errorf("unpackCompURI: failed to find Edition ~")
	}
	if start == end {
		cpe.Edition = "*"
	} else {
		cpe.Edition = s[start:end]
	}
	start = end + 1

	end = offsetIndexOf(s, '~', start)
	if end == -1 {
		return fmt.Errorf("unpackCompURI: failed to find SoftwareEdition ~")
	}
	if start == end {
		cpe.SoftwareEdition = "*"
	} else {
		cpe.SoftwareEdition = s[start:end]
	}
	start = end + 1

	end = offsetIndexOf(s, '~', start)
	if end == -1 {
		return fmt.Errorf("unpackCompURI: failed to find TargetSoftware ~")
	}
	if start == end {
		cpe.TargetSoftware = "*"
	} else {
		cpe.TargetSoftware = s[start:end]
	}
	start = end + 1

	end = offsetIndexOf(s, '~', start)
	if end == -1 {
		return fmt.Errorf("unpackCompURI: failed to find TargetHardware ~")
	}
	if start == end {
		cpe.TargetHardware = "*"
	} else {
		cpe.TargetHardware = s[start:end]
	}
	start = end + 1

	if start >= len(s) {
		cpe.Other = "*"
	} else {
		cpe.Other = s[start:]
	}
	return nil
}

func offsetIndexOf(s string, c byte, off int) int {
	if off >= len(s) {
		return -1
	}

	i := strings.IndexByte(s[off:], c)
	if i == -1 {
		return -1
	}

	return off + i
}

func getCompURI(uri string, i int) string {
	fcount := 0
	sidx := 0

	if i < 0 || i > 7 {
		return ""
	}
	for idx, v := range uri {
		if v == ':' {
			if i == fcount {
				if i == 1 {
					// XXX: relies on verified "cpe:/" prefix
					return uri[sidx+1 : idx]
				}
				return uri[sidx:idx]
			}
			fcount++
			sidx = idx + 1
		}
	}
	if fcount < i {
		return ""
	}
	return uri[sidx:]
}
