package scripts

import (
	"fmt"
	"os/exec"
)

func PrintCommandOutput(args ...string) {
	cmd := exec.Command("echo")
	cmd.Args = args
	bytesOut := cmd.CombinedOutput
	fmt.Printf("\n%q\n", bytesOut)
}
