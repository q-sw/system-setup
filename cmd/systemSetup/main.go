package systemsetup

import (
	"fmt"
	"os"

	"github.com/q-sw/system-setup/internal/apt"
	"github.com/q-sw/system-setup/internal/install"
	"github.com/q-sw/system-setup/internal/repository"
	"github.com/spf13/viper"
)

func Setup(configPath string) {

	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("error to load config file")
		fmt.Println(err)
		os.Exit(1)
	}

	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("error to decode config file")
		fmt.Println(err)
		os.Exit(1)
	}

	apt.UpdateRepository()
	apt.UpgradeSystem()

	for _, pkg := range config.AptStandardPackages {
		fmt.Println(pkg)
		apt.InstallAptPackage(pkg)
	}

	for _, repo := range config.AptRepositories {
		switch repo.Mode {
		case "gpg":
			apt.AddAptRepo(repo.Name, repo.RepoUrl, repo.GPGUrl, repo.GPGKeyName, repo.ToSign)
		case "ppa":
			apt.AddPPARepo(repo.GPGUrl)
		}
	}

	apt.UpdateRepository()
	for _, pkg := range config.AdditionnalPackages {
		apt.InstallAptPackage(pkg)
	}

	for _, gpkg := range config.PackagesFromGithub {
		install.InstallFromGithub(gpkg)
	}

	for _, repo := range config.PersonalGithubRepo {
		repository.CloneGitRepositories(repo)
	}
}
