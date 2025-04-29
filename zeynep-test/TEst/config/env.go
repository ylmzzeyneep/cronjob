package config

import (
		"fmt"
		"os"
		"time"

		"golang.org/x/crypto/ssh"
		"bufio"
		"strings"
)

var (

		SourceFileURLs = getSourceFileURLs("path.txt")

		TargetRepoURL  = os.Getenv("TARGET_REPO_URL")

		TargetBranch   = os.Getenv("TARGET_BRANCH")
		TargetFilePath = os.Getenv("TARGET_FILE_PATH")
		TargetUsername = os.Getenv("TARGET_USERNAME")

		SourcePath = getEnv("SOURCE_PATH", "/tmp/git-source")
		TargetPath = getEnv("TARGET_PATH", "/tmp/git-target")

		PollInterval   = 60 * time.Second

		
		TargetSSHKeyPath = os.Getenv("TARGET_SSH_KEY_PATH")

		TargetAuth = createSSHAuth(TargetSSHKeyPath)
)


func createSSHAuth(keyPath string) ssh.AuthMethod {
	if keyPath == "" {
		panic("SSH key path is empty")
	}

	sshKey, err := os.ReadFile(keyPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to read SSH key from %s: %v", keyPath, err))
	}

	signer, err := ssh.ParsePrivateKey(sshKey)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse SSH key from %s: %v", keyPath, err))
	}

	return ssh.PublicKeys(signer)
}

// path.txt dosyasındaki URL'leri okur
func getSourceFileURLs(filePath string) []string {
	var urls []string

	file, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("Failed to open path file %s: %v", filePath, err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())
		if url != "" {
			urls = append(urls, url)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Failed to read from path file %s: %v", filePath, err))
	}

	return urls
}



func getEnv(key string, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}
/*


export TARGET_REPO_URL=git@github.com:ylmzzeyneep/zeynep-target.git

export TARGET_BRANCH=main
export TARGET_FILE_PATH=targetfile
export TARGET_USERNAME=ylmzzeyneep

export TARGET_SSH_KEY_PATH=/home/ylmzzeyneep/.ssh/id_ed25519

export TARGET_PATH=/mnt/c/Users/zeyne/deneme


*/