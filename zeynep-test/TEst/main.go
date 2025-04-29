package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"zeynep-test/config"
	"zeynep-test/gitops"
)

func main() {
	for {
		log.Println("Preparing target directory...")

		if _, err := os.Stat(config.TargetPath); os.IsNotExist(err) {
			log.Println("Creating target path...")
			if err := os.MkdirAll(config.TargetPath, 0755); err != nil {
				log.Fatalf("Failed to create target path: %v", err)
			}
		}

		if _, err := os.Stat(filepath.Join(config.TargetPath, ".git")); os.IsNotExist(err) {
			log.Println("Cloning target repo...")
			cmd := exec.Command("git", "clone", config.TargetRepoURL, config.TargetPath)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				log.Fatalf("Git clone failed: %v", err)
			}
		}

		for _, url := range config.SourceFileURLs {
			fileName := filepath.Base(url)
			targetPath := filepath.Join(config.TargetPath, fileName)

			log.Printf("Downloading: %s", fileName)
			if err := gitops.DownloadFile(url, targetPath); err != nil {
				log.Printf("Download failed for %s: %v", url, err)
				continue
			}

			cmd := exec.Command("git", "-C", config.TargetPath, "add", fileName)
			if err := cmd.Run(); err != nil {
				log.Printf("Git add failed for %s: %v", fileName, err)
				continue
			}

			cmd = exec.Command("git", "-C", config.TargetPath, "status", "--porcelain")
			output, err := cmd.Output()
			if err != nil {
				log.Printf("Git status failed: %v", err)
				continue
			}
			if len(output) == 0 {
				log.Printf("No changes for %s", fileName)
				continue
			}

			if err := gitops.PushDownloadedFile(fileName); err != nil {
				log.Printf("Push failed for %s: %v", fileName, err)
				continue
			}

			log.Println("Push successful.")
		}

		log.Println("Sleeping for next cycle...")
		time.Sleep(config.PollInterval)
	}
}
