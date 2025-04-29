package gitops

import (
	"fmt"
	"os"
	"os/exec"
	"zeynep-test/config"
)

func PushDownloadedFile(fileName string) error {
	cmd := exec.Command("git", "-C", config.TargetPath, "commit", "-am", "automated update")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git commit failed: %v", err)
	}

	cmd = exec.Command("git", "-C", config.TargetPath, "push", "origin", config.TargetBranch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git push failed: %v", err)
	}

	return nil
}
