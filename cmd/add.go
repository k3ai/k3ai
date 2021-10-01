package cmd

import (
	"os"
	"strings"
	"github.com/spf13/cobra"
	internals "github.com/k3ai/internals"
	log "github.com/k3ai/log"
	shared "github.com/k3ai/shared"
)

//addCmd represents the version command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a K3ai plugin.",
	Long:  `
Add is meant to deploy a specific plugin: application or bundle.
Through the apply command a user may have a certain plugin deployed on the target device.
`,	
	Example:`
k3ai add <plugin name> --to <cluster name>
k3ai add <plugin name> --to <cluster group name>
k3ai add <plugin name> --to <config file>	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Warn("No plugin has been indicated. The right form is: k3ai apply <plugin-name>")
			os.Exit(0)
		}
		
		pluginType, pluginUrl,err := shared.SelectPlugin(strings.ToLower(args[0]))
		_ = log.CheckErrors(err)

		 if pluginType == "Bundle" {
			internals.BundlesDeployment()
		} else {
			internals.AppsDeployment(pluginUrl,args[0])
			log.Info("plugin installed...")
		}


	},

}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().String("to","","Where to install the plugin: a single cluster, a group or using a config file [local or remote]")
}