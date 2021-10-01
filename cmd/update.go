package cmd

import (
    internals "github.com/k3ai/internals"
	"github.com/spf13/cobra"
)

//initCmd represents the version command
var updateCmd = &cobra.Command{
	Use:   "update",
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
		internals.Update()

	},
	Example: `
	k3ai init		//Will initialize K3ai on the local computer.
	`,

}

func init() {
	rootCmd.AddCommand(updateCmd)
}