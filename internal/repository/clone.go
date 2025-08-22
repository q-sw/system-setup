package repository

import (
	"fmt"
	"os"
	"regexp"

	git "github.com/go-git/go-git/v5"
)

func CloneGitRepositories(repo GitRepo) {
	fmt.Printf("--- Git Clone %s ---\n", repo.Name)
	targetDir := checkPath(repo.Path)

	fmt.Printf("Clone %s to %s \n", repo.Name, targetDir)
	_, err := git.PlainClone(targetDir, false, &git.CloneOptions{
		URL:      repo.Url,
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Printf("error to clone  %s\n", repo.Url)
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("--- Git Clone Ended ---")
}

func checkPath(path string) string {
	check := regexp.MustCompile(`^HOME`)
	if check.Match([]byte(path)) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("error to find home directory \n")
			fmt.Println(err)
			os.Exit(1)
		}

		return string(check.ReplaceAll([]byte(path), []byte(homeDir)))
	}
	return path
}
