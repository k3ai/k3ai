package cmd

import (
	"os"
	"strings"
	"github.com/spf13/cobra"
	internals "github.com/k3ai/internals"
	log "github.com/k3ai/log"
	shared "github.com/k3ai/shared"
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
		if len(args) <= 0 {
			log.Warn("No plugin has been indicated. The right form is: k3ai apply <plugin-name>")
			os.Exit(0)
		}
		
		pluginType, pluginUrl := shared.SelectPlugin(strings.ToLower(args[0]))

		if pluginType == "Infra" {
			log.Info("Cluster up, waiting for remaining services to start...")
			internals.InfraDeployment(pluginUrl,args[0])
		} else if pluginType == "Bundle" {
			internals.BundlesDeployment()
		} else {
			internals.AppsDeployment(pluginUrl,args[0])
			log.Info("plugin installed...")
		}


	},
	Example: `
	k3ai apply		//shows all the possible combination of the command apply
	k3ai apply k3s	//apply the plugin with name "k3s"
	`,

}

func init() {
	rootCmd.AddCommand(applyCmd)
	// initCmd.AddCommand(deleteCmd,updateCmd)
}