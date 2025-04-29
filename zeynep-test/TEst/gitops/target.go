package gitops

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"zeynep-test/config"
)

func PushDownloadedFile(fileName string) error {
	targetFilePath := filepath.Join(config.TargetPath, fileName)

	err := os.Rename(config.SourcePath, targetFilePath)
	if err != nil {
		return fmt.Errorf("failed to move downloaded file to target repo: %v", err)
	}

	cmd := exec.Command("git", "-C", config.TargetPath, "add", fileName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git add failed: %v", err)
	}

	cmd = exec.Command("git", "-C", config.TargetPath, "commit", "-m", "automated update")
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
