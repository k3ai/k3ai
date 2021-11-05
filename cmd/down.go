package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/k3ai/internal"
	color "github.com/k3ai/pkg/color"
	loader "github.com/k3ai/pkg/loader"
)

func downCommand() *cobra.Command {
	down := internal.Options{}
	downCmd := &cobra.Command{
		Use:   "down [-h --help] [-q --quiet] [-c fileOrUrl]",
		Short: "K3ai ending point. Remove completly K3ai from local environment.",
		Run: func(cmd *cobra.Command, args []string) {

			bQuiet, _ := cmd.Flags().GetBool("quiet")
			homeDir, _ := os.UserHomeDir()
			k3aiDir := homeDir + "/.k3ai"
			k3aiConfigDir := homeDir + "/.config/k3ai"
			color.Alert()
			fmt.Println("❗️	Proceeding with K3ai uninstall...")
			time.Sleep(700 * time.Millisecond)
			if !bQuiet {
				msg := "Working..."
				loader.StandardLoader(msg)
			} else {
				log.Println("Removing k3ai....")
			}
			os.RemoveAll(k3aiDir)
			os.RemoveAll(k3aiConfigDir)
			time.Sleep(500 * time.Millisecond)
			if !bQuiet {
				color.Done()
				fmt.Println("✔️	Done.Thanks for using K3ai.")
			}
		},
	}
	flags := downCmd.Flags()
	flags.BoolVarP(&down.Quiet, "quiet", "q", false, "Suppress output messages. Useful when k3ai is used within scripts.")
	return downCmd
}
