// Package tasks is a shim over fumeapp/taskin for headless execution.
//
// Rationale: taskin's DisableUI flag only strips ANSI escapes — it still
// writes task-status lines to os.Stdout, which contaminates --json
// output. RunHeadless walks the same task tree without invoking the
// bubbletea runner so stdout stays clean.
package tasks

import (
	"fmt"
	"io"

	"github.com/fumeapp/taskin"
)

// RunHeadless executes every task in the tree in order, in the caller's
// goroutine, without the bubbletea runner. Task.Progress calls still
// update internal state and Task.Title mutations still work; the
// difference is that nothing reaches stdout.
//
// When progressOut is non-nil, one line per completed task is written
// there (typically the renderer's stderr) so users get a coarse audit
// trail even without the TUI.
func RunHeadless(tasks taskin.Tasks, progressOut io.Writer) error {
	for i := range tasks {
		task := &tasks[i]
		if task.Task != nil {
			if err := task.Task(task); err != nil {
				return err
			}
			if progressOut != nil {
				fmt.Fprintln(progressOut, task.Title)
			}
		}
		if len(task.Tasks) > 0 {
			if err := RunHeadless(task.Tasks, progressOut); err != nil {
				return err
			}
		}
	}
	return nil
}
