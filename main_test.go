package main

import (
	"testing"
)

func TestEcho(t *testing.T) {
	if Echo("my string") != "my string" {
		t.Error("expected my string")
	}
}
