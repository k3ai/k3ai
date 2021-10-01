package cmd

import (
	"os"
	"strings"
	log "github.com/k3ai/log"
    internals "github.com/k3ai/internals"
	shared "github.com/k3ai/shared"
	"github.com/spf13/cobra"
)

// resetCmd represents the version command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset K3ai.",
	Long:  `
Reset completly uninstall k3ai from the system.
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
		} else {
			internals.Remove()
		}


	},
	Example: `
k3ai reset
	`,

}

func init() {
	
	rootCmd.AddCommand(resetCmd)
	

}