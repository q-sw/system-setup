package install

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/q-sw/system-setup/internal/utils"
)

func InstallFromGithub(pkg PackageFromGithub) {
	fmt.Printf("--- Install %s Started ---\n", pkg.Name)

	installPath := filepath.Join(pkg.InstallPath, pkg.Name)
	os.Mkdir(installPath, 0644)

	switch pkg.Mode {
	case "gz":
		tarfile := filepath.Join(installPath, fmt.Sprintf("%s.tar.gz", pkg.Name))
		utils.DownloadFile(pkg.Url, tarfile)
		err := utils.UntarGz(tarfile, installPath)
		if err != nil {
			fmt.Println(err)
		}
	case "xz":
		tarfile := filepath.Join(installPath, fmt.Sprintf("%s.tar.xz", pkg.Name))
		utils.DownloadFile(pkg.Url, tarfile)
		err := utils.UntarXz(tarfile, installPath)
		if err != nil {
			fmt.Println(err)
		}
	}

	if pkg.BinaryPath != "" {
		os.Chmod(pkg.BinaryPath, 0755)
	}
	fmt.Println("--- Install Ended ---")

}
