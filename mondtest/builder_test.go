package mondtest

import (
	"os/exec"
	"testing"
)

// TestBuilder tests that we can build a new mond node.
func TestBuilder(t *testing.T) {
	path, err := buildMond()
	if err != nil {
		t.Fatalf("Unable to build mond: %v", err)
	}

	t.Logf("Built mond at %s", path)

	cmd := exec.Command(path, "--version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Unable to fetch mond output: %v", err)
	}

	t.Logf("mond version: %s", string(output))
}
