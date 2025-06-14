package cpeuri

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"github.com/vulncheck-oss/cli/pkg/cpe/cpeutils"
)

func ToStruct(s string) (*cpeutils.CPE, error) {

	var cpe cpeutils.CPE
	var err error

	cpe, err = GetCPEStructFromString(s)

	if err != nil {
		return nil, fmt.Errorf("invalid CPE string: %v", err)
	}

	if cpe.Vendor == "*" && cpe.Product == "*" {
		return nil, fmt.Errorf("CPE Vendor and Product cannot be *")
	}

	return &cpe, nil
}

func GetCPEStructFromString(s string) (cpeutils.CPE, error) {
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
	return cpeutils.CPE{}, fmt.Errorf("unrecognized cpe binding form")
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
func ConvertEmptyToAny(cpe *cpeutils.CPE) {
	v := reflect.ValueOf(cpe).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.String && field.String() == "" {
			field.SetString("*")
		}
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

func UnbindCPEFormattedString(str string) (cpeutils.CPE, error) {
	str = strings.ToLower(str)
	if !isAllASCII(str) {
		return cpeutils.CPE{}, fmt.Errorf("cpe string contains non-ASCII chars")
	}

	cpe := cpeutils.CPE{}
	for a := 2; a <= 12; a++ {
		v := getCompFS(str, a)
		v, err := unbindValueFS(v)
		if err != nil {
			return cpeutils.CPE{}, err
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

	if fcount == i {
		return str[sidx:]
	}

	return ""
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
func UnbindCPEURIString(uri string) (cpeutils.CPE, error) {
	cpe := cpeutils.CPE{}

	if !isAllASCII(uri) {
		return cpeutils.CPE{}, fmt.Errorf("cpe string contains non-ASCII chars")
	}

	var err error
	for i := 1; i <= 7; i++ {
		v := getCompURI(uri, i)
		switch i {
		case 1:
			cpe.Part, err = decodeCompURI(v)
			if err != nil {
				return cpeutils.CPE{}, err
			}

		case 2:
			cpe.Vendor, err = decodeCompURI(v)
			if err != nil {
				return cpeutils.CPE{}, err
			}

		case 3:
			cpe.Product, err = decodeCompURI(v)
			if err != nil {
				return cpeutils.CPE{}, err
			}

		case 4:
			cpe.Version, err = decodeCompURI(v)
			if err != nil {
				return cpeutils.CPE{}, err
			}

		case 5:
			cpe.Update, err = decodeCompURI(v)
			if err != nil {
				return cpeutils.CPE{}, err
			}

		case 6:
			if v == "" || v == "-" || v[0] != '~' {
				cpe.Edition, err = decodeCompURI(v)
				if err != nil {
					return cpeutils.CPE{}, err
				}

			} else {
				err = unpackCompURI(v, &cpe)
				if err != nil {
					return cpeutils.CPE{}, err
				}
			}

		case 7:
			cpe.Language, err = decodeCompURI(v)
			if err != nil {
				return cpeutils.CPE{}, err
			}
		}
	}

	return cpe, nil
}

func decodeCompURI(s string) (string, error) {
	switch s {
	case "":
		return "*", nil
	case "-":
		return "-", nil
	}

	s = strings.ToLower(s)
	result := strings.Builder{}
	percentEncodings := map[string]string{
		"%21": "!", "%22": "\"", "%23": "#", "%24": "$", "%25": "%", "%26": "&",
		"%27": "'", "%28": "(", "%29": ")", "%2a": "*", "%2b": "+", "%2c": ",",
		"%2f": "/", "%3a": ":", "%3b": ";", "%3c": "<", "%3d": "=", "%3e": ">",
		"%3f": "?", "%40": "@", "%5b": "[", "%5c": "\\", "%5d": "]", "%5e": "^",
		"%60": "`", "%7b": "{", "%7c": "|", "%7d": "}", "%7e": "~",
	}

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '%':
			if i+2 >= len(s) {
				return "", fmt.Errorf("invalid percent-encoding")
			}
			encoded := s[i : i+3]
			switch encoded {
			case "%01":
				if (i == 0 || i == len(s)-3) ||
					(i >= 3 && s[i-3:i] == "%01") ||
					(i+6 <= len(s) && s[i+3:i+6] == "%01") {
					result.WriteRune('?')
					i += 2
				} else {
					return "", fmt.Errorf("invalid %%01 encoding")
				}
			case "%02":
				if i == 0 || i == len(s)-3 {
					result.WriteRune('*')
					i += 2
				} else {
					return "", fmt.Errorf("invalid %%02 encoding")
				}
			default:
				if decoded, ok := percentEncodings[encoded]; ok {
					result.WriteString("\\" + decoded)
					i += 2
				} else {
					return "", fmt.Errorf("unrecognized percent-encoding: %s", encoded)
				}
			}
		case '.', '-', '~':
			result.WriteString("\\" + string(s[i]))
		default:
			result.WriteRune(rune(s[i]))
		}
	}

	return result.String(), nil
}

func unpackCompURI(s string, cpe *cpeutils.CPE) error {
	if len(s) <= 1 {
		return fmt.Errorf("invalid packed URI: too short")
	}

	fields := []string{"Edition", "SoftwareEdition", "TargetSoftware", "TargetHardware", "Other"}
	components := strings.Split(s[1:], "~")

	if len(components) < len(fields) {
		return fmt.Errorf("invalid packed URI: not enough components")
	}

	for i, field := range fields {
		value := "*"
		if i < len(components) && components[i] != "" {
			value = components[i]
		}

		reflect.ValueOf(cpe).Elem().FieldByName(field).SetString(value)
	}

	return nil
}

func getCompURI(uri string, i int) string {
	if i < 1 || i > 7 {
		return ""
	}

	parts := strings.SplitN(uri, ":", 8)
	if len(parts) <= i {
		return ""
	}

	if i == 1 {
		// Handle the special case for the part component
		return strings.TrimPrefix(parts[i], "/")
	}

	return parts[i]
}

// RemoveSlashes - remove backslashes from a string
func RemoveSlashes(s string) string {
	return strings.ReplaceAll(s, "\\", "")
}
