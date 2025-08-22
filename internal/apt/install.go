package apt

import (
	"fmt"
	"os"
	"os/exec"
)

func InstallAptPackage(aptPackage string) {
	fmt.Printf("--- Install %s Started ---\n", aptPackage)
	cmd := exec.Command("apt", "install", "-y", aptPackage)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("error to install %s package \n", aptPackage)
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("--- Install %s Ended ---\n", aptPackage)
}

func UpdateRepository() {
	fmt.Println("--- Update Repositories Started ---")
	cmd := exec.Command("apt", "update")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("error to update apt repositories")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("--- Update Repositories Ended  ---")
}

func UpgradeSystem() {
	fmt.Println("--- Upgrade system packages Started  ---")
	cmd := exec.Command("apt", "upgrade", "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("error to upgrade apt packages")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("--- Upgrade system packages Ended  ---")
}
