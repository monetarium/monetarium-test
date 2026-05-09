package mondtest

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// buildMond builds a monetarium node binary in a temp file and returns the path to the
// binary. This requires the Go toolchain to be installed and available in
// the machine. The version of the node package built depends on the currently
// required version of the github.com/monetarium/monetarium-node module, which may be defined
// by either this go mod, a parent go mod (when this package is included as a
// library in a project) or the current workspace.
func buildMond() (string, error) {
	const mondMainPkg = "github.com/monetarium/monetarium-node"
	outDir, err := os.MkdirTemp("", "mondtestmonnnode")
	if err != nil {
		return "", err
	}

	mondBin := "mond"
	if runtime.GOOS == "windows" {
		mondBin += ".exe"
	}

	mondPath := filepath.Join(outDir, mondBin)
	cmd := exec.Command("go", "build", "-o", mondPath, mondMainPkg)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to build mond: %v", err)
	}

	return mondPath, nil
}
