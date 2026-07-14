package output

import (
	"context"
	"testing"

	"github.com/spf13/cobra"
)

// TestFromContextReturnsInjected confirms the round-trip contract used
// by root.PersistentPreRunE + every command RunE.
func TestFromContextReturnsInjected(t *testing.T) {
	want := New(Options{Mode: ModeJSON})
	ctx := WithContext(context.Background(), want)
	got := FromContext(ctx)
	if got != want {
		t.Fatal("FromContext should return the exact renderer that was injected")
	}
}

// TestFromContextDefaultOnMissing guarantees FromContext never returns
// nil — critical for RunE bodies that unconditionally do r := ...FromCmd
// and immediately call r.IsJSON() etc. A panic here would hit every
// command that runs before PersistentPreRunE has attached a renderer.
func TestFromContextDefaultOnMissing(t *testing.T) {
	got := FromContext(context.Background())
	if got == nil {
		t.Fatal("FromContext should never return nil")
	}
	if got.Mode() != ModeText {
		t.Errorf("default mode = %v, want ModeText", got.Mode())
	}
}

func TestFromContextDefaultOnNilContext(t *testing.T) {
	//nolint:staticcheck // deliberately testing the nil-ctx branch
	got := FromContext(nil)
	if got == nil {
		t.Fatal("FromContext(nil) should never return nil")
	}
}

// TestFromContextWithWrongType covers the case where something else
// stores a value under our key type. The lookup must fail closed to a
// default renderer, not return the wrong-type value.
func TestFromContextWithWrongType(t *testing.T) {
	//nolint:staticcheck // testing a defensive path
	ctx := context.WithValue(context.Background(), ctxKey{}, "not a renderer")
	got := FromContext(ctx)
	if got == nil {
		t.Fatal("FromContext should not panic or return nil on wrong-type value")
	}
	// Should be a default renderer, not something dressed up as one.
	if got.Mode() != ModeText {
		t.Errorf("wrong-type value should fall through to default; got mode %v", got.Mode())
	}
}

// TestFromCmdReadsFromCommandContext exercises the shorthand agents use.
func TestFromCmdReadsFromCommandContext(t *testing.T) {
	want := New(Options{Mode: ModeJSON})
	cmd := &cobra.Command{}
	cmd.SetContext(WithContext(context.Background(), want))

	got := FromCmd(cmd)
	if got != want {
		t.Fatal("FromCmd should return the renderer attached via cmd.Context")
	}
}

func TestFromCmdDefaultOnNil(t *testing.T) {
	got := FromCmd(nil)
	if got == nil {
		t.Fatal("FromCmd(nil) should never return nil")
	}
}
