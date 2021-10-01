package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "k3ai [flags]",
	Short: "A simple way to learn and use Artificial Intelligence.",
	Long: `
K3ai allow anyone to start with Artificial Intelligence.
Like a cooking recipe, you select the ingredients we take care of everything else.

The logic is super simple, check out our documentation at: https://k3ai.github.io/docs

Do not forget to add a star to the project: https://github.com/k3ai`,
	Example:`
k3ai create cluster --type k3s --name mycluster
k3ai delete cluster --name mycluster

k3ai add <plugin name> --to mycluster
k3ai add <plugin name>  --to <cluster group name or config file>

k3ai remove <plugin name> --from mycluster (or the cluster group or config file)`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//      Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

}
