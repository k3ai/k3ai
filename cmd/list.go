package cmd

import (
	"github.com/spf13/cobra"
	internals "github.com/k3ai/internals"
)

// listCmd represents the version command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all possible option supported in K3ai",
	Long:  `
List command  always shows ALL options but may be also used along with --tag followed by of the options:

- infra
- apps
- bundles

NOTE: Filter options are mutally exclusive so cannot be used together`,
	Run: func(cmd *cobra.Command, args []string) {
		res,_ := cmd.Flags().GetString("type")
		internals.List(res)

	},
	Example: `
	k3ai list 	          //shows all the possible options
	k3ai list --type infra 	  //shows all the Infrastructure options
	k3ai list --type apps	  //shows all the Applications options
	k3ai list --type bundles   //shows all the Bundles options
	`,

}

func init() {
	
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().Bool("list",false,"test")
	rootCmd.AddCommand(listCmd)
	// listCmd.AddCommand(tagCmd)
	listCmd.Flags().String("type","","Filter list of supported options by value. Possible values are: infra,apps,bundles")
	// tagCmd.Flags().String("infra","","Filter list of supported infra options by value.")
	// tagCmd.AddCommand(infraCmd)
	// tagCmd.AddCommand(appsCmd)
	// tagCmd.AddCommand(bundlesCmd)
}