package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "k3ai [flags]",
	Short: "A simple way to learn and use Artificial Intelligence.",
	Long: `
	What is K3ai?

	K3ai allow anyone to start with Artificial Intelligence.
	You focus on your needs we take care of everything you need.
	
	How it works?
	
	Like a cooking recipe, you select the ingredients we take care of everything else.
	
	- Infrastructure deployment
	- PreRequisites and PosRequisites
	- AI Tools deployment
	
	https://github.com/k3ai`,
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

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.DisableFlagParsing = false

}
