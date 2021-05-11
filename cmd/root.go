package cmd

import (
	"fmt"
	"os"

	//Community packages
	"github.com/spf13/cobra"
)

var asciiMoldy string = `
 _     ____  _    _______  _
/ \__//  _ \/ \  /  _ \  \//
| |\/|| / \|| |  | | \|\  / 
| |  || \_/|| |_/\ |_/|/ /  
\_/  \\____/\____|____/_/   
						   
`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "moldy",
	Short: "Moldy the best project starter of the world",
	Long: asciiMoldy + `
The best tool for initial his project :

Moldy is a Project Starter writed in Go that help for 
start his project and is 100% OPEN SOURCE.

Author: Moldy Community
Contact mail: moldycommunity@gmail.com
Repository: www.github.com/Moldy/Cli
Web Page: www.moldy.github.io
-----------------------------------------------------
Made with love in Ecuador and Colombia`,
}

// Excute the the commands in main go
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
