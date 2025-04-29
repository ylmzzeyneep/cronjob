package main

import (
	"log"
	"os"
	"os/exec"
	"time"
	"strings"

	"zeynep-test/config"
	"zeynep-test/gitops"
)

func main() {
	for {
		log.Println("Preparing target directory...")
		os.RemoveAll(config.TargetPath)
		err := os.MkdirAll(config.TargetPath, 0755)
		if err != nil {
			log.Fatalf("Failed to create target path: %v", err)
		}

		log.Println("Cloning target repository...")
		cmd := exec.Command("git", "clone", config.TargetRepoURL, config.TargetPath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("Git clone failed: %v", err)
		}

		for _, url := range config.SourceFileURLs {
			log.Printf("Downloading file from source URL: %s", url)

			_, fileName := getFileNameFromURL(url)

			err := gitops.DownloadFile(url, config.SourcePath)
			if err != nil {
				log.Printf("Error downloading file from %s: %v", url, err)
				continue 
			}
			log.Println("File downloaded successfully.")

			cmd = exec.Command("git", "add", fileName)
			cmd.Dir = config.TargetPath 
			if err := cmd.Run(); err != nil {
				log.Printf("Error adding file %s to git: %v", fileName, err)
				continue
			}

			cmd = exec.Command("git", "status", "--porcelain")
			cmd.Dir = config.TargetPath 
			output, err := cmd.CombinedOutput()
			if err != nil {
				log.Printf("Error running git status: %v", err)
				continue
			}

			if len(output) == 0 {
				log.Printf("No changes detected for %s, push not made.", fileName) 
				continue
			}

			err = gitops.PushDownloadedFile(fileName) 
			if err != nil {
				log.Printf("Error pushing changes for file %s: %v", fileName, err)
				continue 
			}

			log.Println("Changes pushed successfully.")
		}

		log.Println("Waiting for next iteration...")
		time.Sleep(config.PollInterval)
	}
}

func getFileNameFromURL(url string) (string, string) {
	parts := strings.Split(url, "/")
	fileName := parts[len(parts)-1]
	return parts[len(parts)-2], fileName
}
