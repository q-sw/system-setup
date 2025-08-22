package systemsetup

import (
	"github.com/q-sw/system-setup/internal/apt"
	"github.com/q-sw/system-setup/internal/install"
	"github.com/q-sw/system-setup/internal/repository"
)

type Config struct {
	AptStandardPackages []string                    `mapstructure:"aptStandardPackages"`
	AptRepositories     []apt.AptRepository         `mapstructure:"aptRepositories"`
	AdditionnalPackages []string                    `mapstructure:"additionnalPackages"`
	PackagesFromGithub  []install.PackageFromGithub `mapstructure:"packagesFromGithub"`
	PersonalGithubRepo  []repository.GitRepo        `mapstructure:"personalGithubRepo"`
}
