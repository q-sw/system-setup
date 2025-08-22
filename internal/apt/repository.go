package apt

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/q-sw/system-setup/internal/utils"
)

func AddAptRepo(repoName, repoUrl, keyUrl, keyName string, toSign bool) {
	fmt.Println("Downloading file...")

	var keyPath string
	if toSign {
		keyPath = fmt.Sprintf("/tmp/%s", repoName)
		err := utils.DownloadFile(keyUrl, keyPath)
		if err != nil {
			panic(err)
		}
		gpgInstall(keyPath, repoName)
		keySignPath := fmt.Sprintf("/usr/share/keyrings/%s", repoName)
		addRepoList(repoName, keySignPath, repoUrl)
	} else {
		keyPath = fmt.Sprintf("/etc/apt/keyrings/%s", keyName)
		err := utils.DownloadFile(keyUrl, keyPath)
		if err != nil {
			panic(err)
		}
		addRepoList(repoName, keyPath, repoUrl)
	}
}

func AddPPARepo(ppaURL string) {
	fmt.Println("--- Install PPA Repo  ---")

	cmd := exec.Command("add-apt-repository", "-y", ppaURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Erreur lors de l'installation du repo ppa %v\n", err)
	}

	fmt.Println("--- PPA repo install Successfully ---")

}

func gpgInstall(gpgKey, RepoName string) {
	fmt.Println("--- Install GPG key Started  ---")

	o := fmt.Sprintf("/usr/share/keyrings/%s", RepoName)

	cmd := exec.Command("gpg", "--dearmor", "-o", o, gpgKey)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("error to install gpg key")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("--- Install GPG key Ended  ---")
}

func addRepoList(repoName, gpgKey, repoUrl string) {
	repoStr := fmt.Sprintf("deb [arch=amd64 signed-by=%s] %s\n", gpgKey, repoUrl)
	filePath := fmt.Sprintf("/etc/apt/sources.list.d/%s.list", repoName)

	os.WriteFile(filePath, []byte(repoStr), 0644)
}
