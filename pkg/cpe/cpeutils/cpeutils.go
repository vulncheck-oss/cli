package cpeutils

import (
	"fmt"
	hcversion "github.com/hashicorp/go-version"
	"regexp"
	"strconv"
	"strings"
)

func CompareVersions(a string, b string) (int, error) {
	var c int
	var err error

	// try to handle the e.g., 5.x case
	m, err := regexp.Match(`^\d{1,3}\.x$`, []byte(b))
	if err == nil && m {
		amaj := strings.Split(a, ".")[0]
		amajv, errA := strconv.Atoi(amaj)
		bmaj := strings.Split(b, ".")[0]
		bmajv, errB := strconv.Atoi(bmaj)
		if errA == nil && errB == nil {
			if amajv < bmajv {
				return -1, nil
			} else if amajv == bmajv {
				return 0, nil
			} else {
				return 1, nil
			}
		}
	}

	c, err = CompareVersionsByUint64(a, b)
	if err == nil {
		return c, nil
	}

	hcvA, err := hcversion.NewVersion(a)
	if err != nil {
		return 0, fmt.Errorf("unable to parse version: %s", a)
	}
	hcvB, err := hcversion.NewVersion(b)
	if err != nil {
		return 0, fmt.Errorf("unable to parse version: %s", b)
	}

	if hcvA.LessThan(hcvB) {
		return -1, nil
	} else if hcvA.Equal(hcvB) {
		return 0, nil
	}
	return 1, nil
}

func CompareVersionsByUint64(a, b string) (int, error) {
	au64, err1 := GetUint64FromVersion(a)
	bu64, err2 := GetUint64FromVersion(b)

	if err1 != nil {
		return 0, err1
	}
	if err2 != nil {
		return 0, err2
	}
	if au64 < bu64 {
		return -1, nil
	}
	if au64 == bu64 {
		return 0, nil
	}
	return 1, nil
}

// GetUint64FromVersion attempts to map period separated numeric version
// string to a uint64. Needs testing.
func GetUint64FromVersion(version string) (uint64, error) {
	iv, err := versionToUint64Slice(version)
	if err != nil {
		return 0, err
	}
	return uint64SliceToUint64(iv)
}

func uint64SliceToUint64(isl []uint64) (uint64, error) {
	if len(isl) == 0 {
		return 0, fmt.Errorf("no fields detected in version slice")
	}
	if len(isl) > 8 {
		return 0, fmt.Errorf("too many fields detected in version slice")
	}
	uver := uint64(0)
	mask := uint64(0xff00000000000000)
	for i, donor := range isl {
		uver |= (donor << (56 - (i * 8))) & mask
		mask = mask >> 8
	}
	return uver, nil
}

func versionToUint64Slice(version string) ([]uint64, error) {
	s := strings.Split(version, ".")
	if len(s) > 8 {
		return nil, fmt.Errorf("version %s contains greater than 8 fields", version)
	}

	ivals := make([]uint64, len(s))
	for i, v := range s {
		ival, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("version %s contains unsupported field: %v", version, err)
		}
		ivals[i] = ival
	}
	return ivals, nil
}

func RemoveCVEEntry(b []string, xv string) []string {
	for bi, bv := range b {
		if bv == xv {
			return append(b[:bi], b[bi+1:]...)
		}
	}
	return b
}

func RemoveCVEEntries(b []string, x []string) []string {
	for _, xv := range x {
		b = RemoveCVEEntry(b, xv)
	}
	return b
}

func IsParseableVersion(a string) bool {
	if _, err := GetUint64FromVersion(a); err == nil {
		return true
	}
	_, err := hcversion.NewVersion(a)
	return err == nil
}
