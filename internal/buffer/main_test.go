package buffer

import (
	"bytes"
	"testing"
)

func TestBufferWriteRead(t *testing.T) {
	var buf bytes.Buffer
	testData := "hello\n"

	buf.WriteString(testData)

	want := testData
	got := buf.String()

	if want != got {
		t.Fatalf("wanted %v got %v", want, got)
	}

	t.Logf("buffer contents: %v", got)
}
