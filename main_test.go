package main

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestMyOut(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	MyOut("This")
	_ = w.Close()
	result, _ := io.ReadAll(r)
	os.Stdout = stdOut
	fmt.Println(result)
}
