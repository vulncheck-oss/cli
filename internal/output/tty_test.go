package output

import (
	"os"
	"testing"
)

func TestColorEnabledHonoursNoColor(t *testing.T) {
	t.Setenv("NO_COLOR", "1")
	if ColorEnabled() {
		t.Fatal("NO_COLOR set: ColorEnabled must be false")
	}
}

func TestColorEnabledHonoursDumbTerm(t *testing.T) {
	t.Setenv("NO_COLOR", "")
	t.Setenv("TERM", "dumb")
	if ColorEnabled() {
		t.Fatal("TERM=dumb: ColorEnabled must be false")
	}
}

func TestIsCIDetectsCommonVars(t *testing.T) {
	cases := []string{"CI", "BUILD_NUMBER", "RUN_ID"}
	for _, v := range cases {
		t.Run(v, func(t *testing.T) {
			// Clear all CI vars first so the test doesn't get a false positive
			// from a different one set in the host environment.
			for _, other := range cases {
				t.Setenv(other, "")
			}
			t.Setenv(v, "1")
			if !IsCI() {
				t.Fatalf("%s=1 should be detected as CI", v)
			}
		})
	}
}

func TestInteractiveFalseUnderCI(t *testing.T) {
	t.Setenv("CI", "1")
	if Interactive() {
		t.Fatal("CI=1: Interactive must be false")
	}
}

func TestIsTTYNilFile(t *testing.T) {
	if IsTTY(nil) {
		t.Fatal("IsTTY(nil) must be false")
	}
}

func TestIsTTYDevNull(t *testing.T) {
	f, err := os.Open(os.DevNull)
	if err != nil {
		t.Fatalf("open /dev/null: %v", err)
	}
	defer f.Close()
	if IsTTY(f) {
		t.Fatal("/dev/null is not a TTY")
	}
}
