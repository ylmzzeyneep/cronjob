package gitops

import (
		"bytes"
		"fmt"
		"os/exec"
)

func DownloadFile(url string, path string) error {
	cmd := exec.Command("curl", "-o", path, url)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = nil
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to download file: %v", err)
	}
	return nil
}
