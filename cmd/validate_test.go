package cmd

import "testing"

func TestDetectVersion(t *testing.T) {
	version, err := detectVersion("../test/non-oauth-scopes.json")
	if err != nil {
		t.Fatalf("failed test: %v", err)
	}
	if version != V31 {
		t.Fatalf("failed test: %v", version)
	}
}
