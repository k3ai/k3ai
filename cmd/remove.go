package cmd

import (
	"os"
	"strings"
	log "github.com/k3ai/log"
    internals "github.com/k3ai/internals"
	shared "github.com/k3ai/shared"
	"github.com/spf13/cobra"
)

// removeCmd represents the version command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a K3ai plugin.",
	Long:  `
Remove is meant to uninstall a specific kind of plugin: application or bundle.
Through the remove command a user may have a certain plugin removed from the target device.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Warn("No plugin has been indicated. The right form is: k3ai apply <plugin-name>")
			os.Exit(0)
		}
		
		pluginType, pluginUrl,err := shared.SelectPlugin(strings.ToLower(args[0]))
		_ = log.CheckErrors(err)

		if pluginType == "Infra" {

			internals.InfraRemoval(pluginUrl,args[0])
		} else if pluginType == "Bundle" {

			internals.BundlesRemoval()
		} else {
			internals.Remove()
		}


	},
	Example: `
k3ai remove	<plugin name> --from <cluster name>	
k3ai remove	<plugin name> --from <cluster group name>		
k3ai remove	<plugin name> --from <config file>	
	`,

}

func init() {
	
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().String("from","","Indicate from where the plugin need to be removed: if a single cluster, a group of clusters or reading a config file [local or remote]")
}