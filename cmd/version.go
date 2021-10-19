package cmd

import (
  	"github.com/spf13/cobra"

  
)

func versionCommand() *cobra.Command{
	versionCommand := &cobra.Command{
		Use:   "version",
		Short: "K3ai actual version. Print current binary version and info's.",
		Run: func(cmd *cobra.Command, args []string) {

		},
	  }
	  return versionCommand
}