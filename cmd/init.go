package cmd

import (
	"github.com/alefesta/k3ai/log"
    initialize "github.com/alefesta/k3ai/internals"
	"github.com/spf13/cobra"
)

// initCmd represents the version command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize k3ai on the local computer.",
	Long:  `
Init initialize K3ai on the computer where it has been invoked.
We'll create a folder named .k3ai under user home path where we will store
a light database to store the plugins list and their status.
Init has also two sub-commands:

- delete : to remove k3ai configurations
- update : to update current list of plugins
`,
	Run: func(cmd *cobra.Command, args []string) {
		initialize.Init()

	},
	Example: `
	k3ai init		//initialize a local copy of all possible options
	k3ai init update	//update current list of options
	k3ai init delete	//delete current configuration of k3ai
	`,

}

// deleteCmd represents the version command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete k3ai from the local computer.",
	Run: func(cmd *cobra.Command, args []string) {
		log.Error("Currently Delete function has not been yet implemented...")

	},
}

// updateCmd represents the version command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update k3ai on the local computer.",
	Run: func(cmd *cobra.Command, args []string) {
		initialize.Update()

	},
}

func init() {
	
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().Bool("list",false,"test")
	rootCmd.AddCommand(initCmd)
	initCmd.AddCommand(deleteCmd,updateCmd)
}