package cpetypes

import "fmt"

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

func (c CPE) IsMozilla() bool {
	return Unquote(c.Vendor) == "mozilla"
}

// ProductUcFirst - returns the product name with the first letter capitalized.
func (c CPE) ProductUcFirst() string {
	if len(c.Product) > 0 {
		return string(c.Product[0]-32) + c.Product[1:]
	}
	return ""
}

func Unquote(b string) string {
	return bindValueFS(b)
}

func bindValueFS(v string) string {
	if v == "*" || v == "-" {
		return v
	}
	if v == "" {
		return "*"
	}
	return processQuotedChars(v)
}

func processQuotedChars(s string) string {
	result := ""
	idx := 0

	for idx < len(s) {
		c := s[idx]

		if c != '\\' {
			result = fmt.Sprintf("%s%c", result, c)

		} else {

			nextchr := s[idx+1]
			switch nextchr {
			case '.',
				'-',
				'_':
				result = fmt.Sprintf("%s%c", result, nextchr)
				idx += 2

			default:
				result = fmt.Sprintf("%s\\%c", result, nextchr)
				idx += 2
			}
			continue
		}
		idx++
	}
	return result
}
