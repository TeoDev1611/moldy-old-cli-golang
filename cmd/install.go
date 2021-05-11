package cmd

import (
	pyGo "github.com/DataDog/go-python3"
	"github.com/fatih/color"
	GitDetect "github.com/openshift/library-go/pkg/git"
	"github.com/spf13/cobra"
)

func tellInstall(module string) {
	color.Green("Installing the moldy module **" + module + "**")
	gitInstalled := GitDetect.IsGitInstalled()
	if gitInstalled == true {
		pyGo.PyRun_AnyFile("gitPython.py")
	} else {
		errorGit := "Error git is not installed"
		color.Red(errorGit)
	}
}

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install www.github.com/User/Repository.git",
	Short: "Install a package of Moldy",
	Long: `Install the package use.
	This clone the package in the current folder.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			color.Red("Error: Not found the module")
		} else {
			modulo := args[0]
			tellInstall(modulo)
		}

	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
