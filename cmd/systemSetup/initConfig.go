package systemsetup

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates
var TemplateFile embed.FS

func InitConfiguration() {
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Println("error to get the current path")
		fmt.Println(err)
		os.Exit(1)
	}

	tmpl, err := template.ParseFS(TemplateFile, "templates/config.yaml.tmpl")
	if err != nil {
		fmt.Println("error to find the config file template")
		fmt.Println(err)
		os.Exit(1)
	}

	var configFile *os.File
	configFile, err = os.Create(filepath.Join(currentPath, "config.yaml"))
	if err != nil {
		fmt.Println("error to create the config file")
		fmt.Println(err)
		os.Exit(1)
	}

	err = tmpl.Execute(configFile, "")
	if err != nil {
		fmt.Println("error to generate the config file")
		fmt.Println(err)
		os.Exit(1)
	}

}
