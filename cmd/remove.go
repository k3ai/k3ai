package cmd

import (
	"os"
	"strings"
	"github.com/alefesta/k3ai/log"
    initialize "github.com/alefesta/k3ai/internals"
	utils "github.com/alefesta/k3ai/shared"
	"github.com/spf13/cobra"
)

// applyCmd represents the version command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a K3ai plugin.",
	Long:  `
Remove is meant to uninstall a specific kind of plugin: infrastructure,application or bundle.
Through the remove command a user may have a certain plugin removed from the target device.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Warn("No plugin has been indicated. The right form is: k3ai apply <plugin-name>")
			os.Exit(0)
		}
		
		pluginType, pluginUrl := utils.SelectPlugin(strings.ToLower(args[0]))

		if pluginType == "Infra" {

			initialize.InfraRemoval(pluginUrl,args[0])
		} else if pluginType == "Bundle" {

			initialize.BundlesRemoval()
		} else {
			initialize.Remove()
		}


	},
	Example: `
	k3ai remove		//shows all the possible combination of the command apply
	k3ai remove k3s	//apply the plugin with name "k3s"
	`,

}

// // deleteCmd represents the version command
// var deleteCmd = &cobra.Command{
// 	Use:   "delete",
// 	Short: "Delete k3ai from the local computer.",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		log.Error("Currently Delete function has not been yet implemented...")

// 	},
// }

// // updateCmd represents the version command
// var updateCmd = &cobra.Command{
// 	Use:   "update",
// 	Short: "Update k3ai on the local computer.",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		initialize.Update()

// 	},
// }

func init() {
	
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().Bool("list",false,"test")
	rootCmd.AddCommand(removeCmd)
	// initCmd.AddCommand(deleteCmd,updateCmd)
}