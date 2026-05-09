//go:build !windows

package mondtest

import (
	"os"
	"os/exec"
)

// setOSNodeCmdOptions sets platform-specific options needed to run mond.
func setOSNodeCmdOptions(n *nodeConfig, cmd *exec.Cmd) {
	cmd.ExtraFiles = []*os.File{
		n.pipeTX.w,
		n.pipeRX.r,
	}
}

// appendOSNodeArgs appends platform-specific arguments needed to run mond.
func appendOSNodeArgs(n *nodeConfig, args []string) []string {
	args = append(args, "--pipetx=3")
	args = append(args, "--piperx=4")
	return args
}
