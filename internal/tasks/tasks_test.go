package tasks

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/fumeapp/taskin"
)

// TestRunHeadlessExecutesInOrder locks in the contract that RunHeadless
// runs task closures sequentially and in the caller's slice order.
func TestRunHeadlessExecutesInOrder(t *testing.T) {
	var got []string
	list := taskin.Tasks{
		{Title: "a", Task: func(*taskin.Task) error { got = append(got, "a"); return nil }},
		{Title: "b", Task: func(*taskin.Task) error { got = append(got, "b"); return nil }},
		{Title: "c", Task: func(*taskin.Task) error { got = append(got, "c"); return nil }},
	}
	if err := RunHeadless(list, nil); err != nil {
		t.Fatalf("RunHeadless: %v", err)
	}
	if strings.Join(got, ",") != "a,b,c" {
		t.Fatalf("execution order = %v, want a,b,c", got)
	}
}

// TestRunHeadlessRecursesIntoChildren confirms nested Tasks fields run
// after the parent's Task closure — matching taskin's own semantics.
func TestRunHeadlessRecursesIntoChildren(t *testing.T) {
	var got []string
	list := taskin.Tasks{
		{
			Title: "parent",
			Task:  func(*taskin.Task) error { got = append(got, "parent"); return nil },
			Tasks: taskin.Tasks{
				{Title: "child-1", Task: func(*taskin.Task) error { got = append(got, "child-1"); return nil }},
				{Title: "child-2", Task: func(*taskin.Task) error { got = append(got, "child-2"); return nil }},
			},
		},
	}
	if err := RunHeadless(list, nil); err != nil {
		t.Fatal(err)
	}
	if strings.Join(got, ",") != "parent,child-1,child-2" {
		t.Fatalf("execution order = %v", got)
	}
}

// TestRunHeadlessAbortsOnError stops on the first failing task and
// returns its error unwrapped, so callers can errors.Is against it.
func TestRunHeadlessAbortsOnError(t *testing.T) {
	sentinel := errors.New("boom")
	var got []string
	list := taskin.Tasks{
		{Title: "ok", Task: func(*taskin.Task) error { got = append(got, "ok"); return nil }},
		{Title: "bad", Task: func(*taskin.Task) error { got = append(got, "bad"); return sentinel }},
		{Title: "skipped", Task: func(*taskin.Task) error { got = append(got, "skipped"); return nil }},
	}
	err := RunHeadless(list, nil)
	if !errors.Is(err, sentinel) {
		t.Fatalf("err = %v, want %v", err, sentinel)
	}
	if strings.Contains(strings.Join(got, ","), "skipped") {
		t.Fatalf("task after failure should not have run: %v", got)
	}
}

// TestRunHeadlessWritesProgressOut confirms progressOut receives the
// final Task.Title after each task's closure completes — including the
// task's own mutation of its title.
func TestRunHeadlessWritesProgressOut(t *testing.T) {
	buf := &bytes.Buffer{}
	list := taskin.Tasks{
		{
			Title: "before",
			Task: func(t *taskin.Task) error {
				t.Title = "after"
				return nil
			},
		},
	}
	if err := RunHeadless(list, buf); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "after") {
		t.Fatalf("progressOut should contain post-mutation title, got: %q", buf.String())
	}
	if strings.Contains(buf.String(), "before") {
		t.Fatalf("progressOut should not contain initial title, got: %q", buf.String())
	}
}

// TestRunHeadlessSkipsNilTaskClosure handles Tasks with no Task func —
// they exist purely to group children (as in cache.IndicesSync's parent
// wrapper) and should be silently walked without panic.
func TestRunHeadlessSkipsNilTaskClosure(t *testing.T) {
	var ran bool
	list := taskin.Tasks{
		{
			Title: "grouping-only",
			Task:  nil,
			Tasks: taskin.Tasks{
				{Title: "child", Task: func(*taskin.Task) error { ran = true; return nil }},
			},
		},
	}
	if err := RunHeadless(list, nil); err != nil {
		t.Fatal(err)
	}
	if !ran {
		t.Fatal("child task should have run under a nil-Task parent")
	}
}

// TestRunHeadlessOnEmpty is a nil-safety check.
func TestRunHeadlessOnEmpty(t *testing.T) {
	if err := RunHeadless(nil, nil); err != nil {
		t.Fatalf("empty input should not error: %v", err)
	}
	if err := RunHeadless(taskin.Tasks{}, nil); err != nil {
		t.Fatalf("empty input should not error: %v", err)
	}
}
