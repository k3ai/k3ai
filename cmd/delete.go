package cmd

import (
	"os"
	"strings"
	log "github.com/k3ai/log"
    internals "github.com/k3ai/internals"
	shared "github.com/k3ai/shared"
	"github.com/spf13/cobra"
)

// deleteCmd represents the version command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a K3ai cluster.",
	Long:  `
Delete is meant to uninstall a specific kind of plugin: application or bundle.
Through the delete command a user may have a certain plugin created from the target device.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Warn("No plugin has been indicated. The right form is: k3ai apply <plugin-name>")
			os.Exit(0)
		}
		
		pluginType, pluginUrl := shared.SelectPlugin(strings.ToLower(args[0]))

		if pluginType == "Infra" {

			internals.InfraRemoval(pluginUrl,args[0])
		} else if pluginType == "Bundle" {

			internals.BundlesRemoval()
		} //else {
			// internals.create()
		//}


	},
	Example: `
k3ai create	<plugin name> --type <cluster type> --name <cluster name>
	`,

}

// clusterCmd represents the version command
var cluster_deleteCmd = &cobra.Command{
	Use:   "cluster",
	Short: "create a K3ai plugin.",
	Long:  `
create is meant to uninstall a specific kind of plugin: application or bundle.
Through the create command a user may have a certain plugin created from the target device.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Warn("No plugin has been indicated. The right form is: k3ai apply <plugin-name>")
			os.Exit(0)
		}
		
		pluginType, pluginUrl := shared.SelectPlugin(strings.ToLower(args[0]))

		if pluginType == "Infra" {

			internals.InfraRemoval(pluginUrl,args[0])
		} else if pluginType == "Bundle" {

			internals.BundlesRemoval()
		} //else {
			// internals.create()
		//}


	},
	Example: `
k3ai create	<plugin name> --type <cluster type> --name <cluster name>
	`,

}
func init() {
	
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(cluster_deleteCmd)
	cluster_deleteCmd.Flags().String("name","","The name of the cluster. This is the name you may see with k3ai list clusters.")
}