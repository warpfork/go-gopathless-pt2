package gopathless2

import (
	"bytes"
	"testing"
)

func TestFuncWithDependencies(t *testing.T) {
	var buf bytes.Buffer
	FuncWithDependencies(&buf)
	if buf.String() == "heji, mundus!\n" {
		t.Log("test passed!")
	} else {
		t.Fatalf("test failed :(")
	}
}
