package cmd

import (
	"fmt"
  	"github.com/spf13/cobra"

	color "github.com/k3ai/pkg/color"
	"github.com/k3ai/internal"
  
)

func downCommand() *cobra.Command{
	down := internal.Options{}
	downCmd := &cobra.Command{
		Use:   "down [-h --help] [-q --quiet] [-c fileOrUrl]",
		Short: "K3ai ending point. Remove completly K3ai from local environment.",
		Run: func(cmd *cobra.Command, args []string) {
		  color.Alert()
		  fmt.Println("This is and error")
		  color.Disable()
		  fmt.Println("This is not")
		},
	  }
	  flags := downCmd.Flags()
	  flags.BoolVarP(&down.Quiet,"quiet","q",false,"Suppress output messages. Useful when k3ai is used within scripts.")
	  flags.StringVarP(&down.Config,"config","c","","Configure K3ai using a custom config file.[-c /path/tofile] [-c https://urlToFile]")
	  return downCmd
}