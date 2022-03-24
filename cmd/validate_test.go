package cmd

import (
	"testing"
)

func TestDetectVersion(t *testing.T) {
	target, err := readFromFile("../test/non-oauth-scopes.json")
	if err != nil {
		t.Fatalf("failed test: %v", err)
	}
	version, err := detectVersion(target)
	if err != nil {
		t.Fatalf("failed test: %v", err)
	}
	if version != V31 {
		t.Fatalf("failed test: %v", version)
	}
}
