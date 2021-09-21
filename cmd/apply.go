package cmd

import (
	// "github.com/alefesta/k3ai/log"clear
    initialize "github.com/alefesta/k3ai/internals"
	utils "github.com/alefesta/k3ai/shared"
	"github.com/spf13/cobra"
)

// applyCmd represents the version command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a K3ai plugin.",
	Long:  `
Apply is meant to deploy a specific kind of plugin: infrastructure,application or bundle.
Through the apply command a user may have a certain plugin deployed on the target device.
`,
	Run: func(cmd *cobra.Command, args []string) {
		
		pluginType, pluginUrl := utils.SelectPlugin(args[0])

		if pluginType == "Infra" {

			initialize.InfraDeployment(pluginUrl,args[0])
		} else if pluginType == "Bundle" {

			initialize.BundlesDeployment()
		} else {
			initialize.Apply()
		}


	},
	Example: `
	k3ai apply		//shows all the possible combination of the command apply
	k3ai apply k3s	//apply the plugin with name "k3s"
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
	rootCmd.AddCommand(applyCmd)
	// initCmd.AddCommand(deleteCmd,updateCmd)
}