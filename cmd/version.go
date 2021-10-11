package cmd

import (
	"fmt"
  	"github.com/spf13/cobra"

	color "github.com/k3ai/pkg/color"
  
)

func versionCommand() *cobra.Command{
	versionCommand := &cobra.Command{
		Use:   "version",
		Short: "K3ai actual version. Print current binary version and info's.",
		Run: func(cmd *cobra.Command, args []string) {
		  color.Alert()
		  fmt.Println("This is and error")
		  color.Disable()
		  fmt.Println("This is not")
		},
	  }
	  return versionCommand
}