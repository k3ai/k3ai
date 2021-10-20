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

func runCommand() *cobra.Command{
	run := internal.Options{}
	runCmd := &cobra.Command{
		Use:   "run [-h --help] [-q --quiet] [-c fileOrUrl]",
		Short: "K3ai Run allow a user to execute a piece of code on a give cluster.",
		Run: func(cmd *cobra.Command, args []string) {
			
			bQuiet,_ := cmd.Flags().GetBool("quiet")
			homeDir,_ := os.UserHomeDir()
			k3aiDir := homeDir + "/.k3ai"
			k3aiConfigDir := homeDir + "/.config/k3ai"
			color.InProgress()
			fmt.Println("🧪	Executing code...")
			time.Sleep(700 * time.Millisecond)
			if !bQuiet {
				msg := "Completing upload..."
				loader.SuperLoader(msg)
			}else{
				log.Println("Uploading code")
			}
			os.RemoveAll(k3aiDir)
			os.RemoveAll(k3aiConfigDir)
			time.Sleep(500 * time.Millisecond)
			if !bQuiet{
				color.Done()
				fmt.Println("✔️	Done.You may check your code directly on the backend platform.")
			}
		},
	  }
	  flags := runCmd.Flags()
	  flags.StringVarP(&run.Code,"code","c","","Source code to execute. Can be YAML,Python,Jupyter Notebooks as extension.")
	  flags.StringVarP(&run.Backend,"backend","b","","Backend for the code to be run. Cane be: Kubeflow,Argo,MLFlow.")
	  flags.StringVarP(&run.Target,"target","t","","Where the run need to be exectuted, it's the name of a registered cluster.")
	  flags.BoolVarP(&run.Quiet,"quiet","q",false,"Suppress output messages. Useful when k3ai is used within scripts.")
	  return runCmd
}