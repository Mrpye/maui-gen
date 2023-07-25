package code_gen

import (
	"fmt"
	"os/exec"
)

func InstallNugetPackage(project_folder string, package_name string) error {

	out, err := exec.Command("dotnet", "add", project_folder, "package", package_name).Output()

	if err != nil {
		return err
	}

	fmt.Println(string(out))
	return nil
}
