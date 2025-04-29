package gitops

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func DownloadFile(url, outputPath string) error {
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %v", dir, err)
	}

	cmd := exec.Command("curl", "-sSL", "-o", outputPath, url)
	cmd.Env = os.Environ()

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download file: %v\nstdout: %s\nstderr: %s",
			err, stdout.String(), stderr.String())
	}

	fmt.Println("Downloaded:", outputPath)
	return nil
}
