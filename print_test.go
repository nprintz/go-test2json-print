package main

import (
	"testing"
)

// TestPrint is run with go test, but using -json or test2json omits the completion status
func TestPrint(t *testing.T) {
	print()
}

// TestPrintln is run with got, and also works with -json or test2json
func TestPrintln(t *testing.T) {
	println()
}
