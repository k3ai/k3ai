package cmd

import (
	// "os"
	// "fmt"

	"github.com/spf13/cobra"

)

const cliName = "k3ai"
var version bool

var (

	rootCmd = &cobra.Command{

	Use:   cliName + "[options]",
	Short: "K3ai is a very fast tool to run AI Infrastructure stacks",
	// By default (no Run/RunE in parent command) for typos in subcommands, cobra displays the help of parent command but exit(0) !
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Help()
		}
		if version {
			return versionCommand().Execute()
		}
		_ = cmd.Help()
		return nil
	},

	}
)

func Execute() error {
	return rootCmd.Execute()
	}

func init() {
	cobra.OnInitialize()
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Flags().SortFlags = false
	rootCmd.DisableFlagsInUseLine = true
	rootCmd.PersistentFlags().BoolP("help", "h", false , "Help usage")
	rootCmd.PersistentFlags().Lookup("help").Hidden = true
	rootCmd.AddCommand(
		upCommand(),
		downCommand(),
		clusterCommand(),
		pluginCommand(),
		versionCommand(),
	)
}