package task

import (
	"bytes"
	"context"
	"testing"

	"github.com/go-task/task/v3"
	"github.com/go-task/task/v3/taskfile/ast"
)

func TestRunTargetFromTaskfile(t *testing.T) {
	var buf bytes.Buffer

	e := task.Executor{
		Dir:    "testdata",
		Stdout: &buf,
		Stderr: &buf,
		Silent: true,
	}

	err := e.Setup()
	if err != nil {
		t.Fatalf("error calling Setup: %v", err)
	}

	err = e.Run(context.Background(), &task.Call{
		Task: "greet",
	})
	if err != nil {
		t.Fatalf("error calling Run: %v", err)
	}

	want := "Hello Taskfile\n"
	got := buf.String()
	if got != want {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func RunTargetFromAST(t *testing.T) {
	var buf bytes.Buffer

	taskfile := ast.Taskfile{}

	e := task.Executor{
		Taskfile: &taskfile,
		Stdout:   &buf,
		Stderr:   &buf,
		Silent:   true,
	}

	err := e.Setup()
	if err != nil {
		t.Fatalf("error calling Setup: %v", err)
	}

	err = e.Run(context.Background(), &task.Call{
		Task: "greet",
	})
	if err != nil {
		t.Fatalf("error calling Run: %v", err)
	}

	want := "Hello Taskfile\n"
	got := buf.String()

	if got != want {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
