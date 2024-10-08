package task

import (
	"bytes"
	"context"
	"path/filepath"
	"testing"

	"github.com/go-task/task/v3"
	"github.com/go-task/task/v3/taskfile/ast"
)

func TestTarget(t *testing.T) {
	var buf bytes.Buffer
	taskFilePath, err := filepath.Abs("testdata/Taskfile.yml")
	if err != nil {
		t.Fatal(err)
	}

	e := task.Executor{
		Dir:        "testdata",
		Entrypoint: taskFilePath,
		Stdout: &buf,
		Stderr: &buf,
	}

	err = e.Setup()
	if err != nil {
		t.Fatalf("error calling Setup: %v", err)
	}

	err = e.Run(context.Background(), &ast.Call{Task: "greet"})
	if err != nil {
		t.Fatalf("error calling Run: %v", err)
	}

	t.Log(buf.String())
}
