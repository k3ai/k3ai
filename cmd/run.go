package cmd

import (
	"os"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"

	"github.com/k3ai/internal"
	color "github.com/k3ai/pkg/color"
	loader "github.com/k3ai/pkg/loader"
	runner "github.com/k3ai/pkg/runner"
)

func runCommand() *cobra.Command{
	run := internal.Options{}
	runCmd := &cobra.Command{
		Use:   "run [-h --help] [-q --quiet] [-s --source] [-t --target] [-b --backend] [-e --extras]",
		Short: "K3ai Run allow a user to execute a piece of code on a give cluster.",
		Run: func(cmd *cobra.Command, args []string) {
			
			bQuiet,_ := cmd.Flags().GetBool("quiet")
			sExtras,_ := cmd.Flags().GetString("extras")
			sTarget,_ := cmd.Flags().GetString("target")
			sBackend,_ := cmd.Flags().GetString("backend")
			sSource,_ := cmd.Flags().GetString("source")
			if len(args) == 0 &&  sTarget == "" && sBackend == "" && sSource == "" && sExtras ==""{
				cmd.Help()
				os.Exit(0)
			}
			
			color.InProgress()
			fmt.Println("🧪	Executing code...")
			time.Sleep(700 * time.Millisecond)
			err := runner.Loader(sSource,sTarget,sBackend,sExtras)
			if err != nil {
				log.Println("An error occurred, please retry and if persist, please open an issue on our GitHub repository.")
			}
			if !bQuiet {
				msg := "Completing execution..."
				loader.SuperLoader(msg)
			}else{
				log.Println("Uploading code")
			}

			time.Sleep(500 * time.Millisecond)
			if !bQuiet{
				color.Done()
				fmt.Println("✔️	Done. You may check your code directly on the backend platform.")
			}
		},
	  }
	  flags := runCmd.Flags()
	  flags.StringVarP(&run.Source,"source","s","","Source code to execute. Can be YAML,Python,Jupyter Notebooks as extension.")
	  flags.StringVarP(&run.Backend,"backend","b","","Backend for the code to be run. Cane be: Kubeflow,Argo,MLFlow.")
	  flags.StringVarP(&run.Target,"target","t","","Where the run need to be executed, it's the name of a registered cluster.")
	  flags.StringVarP(&run.Extras,"extras","e","","This is equivalent of requirements.txt, it is useful to install extra packages.")
	  flags.BoolVarP(&run.Quiet,"quiet","q",false,"Suppress output messages. Useful when k3ai is used within scripts.")
	  return runCmd
}