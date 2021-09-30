package cmd

import (
	"fmt"
	internals "github.com/k3ai/internals"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of generated code example",
	Long:  `All software has versions. This is generated code example`,
	Run: func(cmd *cobra.Command, args []string) {
				
		fmt.Println("Build Date:", internals.BuildDate)
		fmt.Println("Git Commit:", internals.GitCommit)
		fmt.Println("Version:", internals.Version)
		fmt.Println("Go Version:", internals.GoVersion)
		fmt.Println("OS / Arch:", internals.OsArch)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
