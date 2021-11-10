package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"github.com/k3ai/internal"
	color "github.com/k3ai/pkg/color"
	"github.com/k3ai/pkg/http"
	loader "github.com/k3ai/pkg/loader"
	runner "github.com/k3ai/pkg/runner"
)

func runCommand() *cobra.Command {
	run := internal.Options{}
	runCmd := &cobra.Command{
		Use:   "run [-h --help] [-q --quiet] [-s --source] [-t --target] [-b --backend] [-e --extras] [-c --config]",
		Short: "K3ai Run allow a user to execute a piece of code on a give cluster.",
		Run: func(cmd *cobra.Command, args []string) {

			bQuiet, _ := cmd.Flags().GetBool("quiet")
			sExtras, _ := cmd.Flags().GetString("extras")
			sConf, _ := cmd.Flags().GetString("config")
			sTarget, _ := cmd.Flags().GetString("target")
			sBackend, _ := cmd.Flags().GetString("backend")
			sSource, _ := cmd.Flags().GetString("source")
			sEntry, _ := cmd.Flags().GetString("entrypoint")
			if len(args) == 0 && sConf == "" && sTarget == "" && sBackend == "" && sSource == "" && sExtras == "" {
				err := cmd.Help()
				if err != nil {
					log.Fatal(err)
				}
				os.Exit(0)
			}

			if sConf != "" {
				if strings.TrimSpace(sConf)[:4] == "http" {
					data, _ := http.Download(sConf)
					err := yaml.Unmarshal([]byte(data), &k3aiConfig)
					if err != nil {
						log.Fatal(err)
					}
					sSource = k3aiConfig.Run.Source
					sTarget = k3aiConfig.Run.Target
					sBackend = k3aiConfig.Run.Backend
				} else {
					data, err := ioutil.ReadFile(sConf)
					if err != nil {
						log.Fatal(err)
					}
					err = yaml.Unmarshal(data, &k3aiConfig)
					if err != nil {
						log.Fatal(err)
					}
					sSource = k3aiConfig.Run.Source
					sTarget = k3aiConfig.Run.Target
					sBackend = k3aiConfig.Run.Backend
				}
			}

			color.InProgress()
			fmt.Println("🧪	Initializing code...")
			time.Sleep(700 * time.Millisecond)
			err := runner.Loader(sSource, sTarget, sBackend, sExtras, sEntry)
			if err != nil {
				log.Println("An error occurred, please retry and if persist, please open an issue on our GitHub repository.")
			}
			if !bQuiet {
				msg := "Completing execution..."
				loader.SuperLoader(msg)
			} else {
				log.Println("Uploading code")
			}

			time.Sleep(500 * time.Millisecond)
			if !bQuiet {
				color.Done()
				fmt.Println("✔️	Done. You may check your code directly on the backend platform.")
			}
		},
	}
	flags := runCmd.Flags()
	flags.StringVarP(&run.Source, "source", "s", "", "Source code to execute. Can be YAML,Python,Jupyter Notebooks as extension.")
	flags.StringVarP(&run.Backend, "backend", "b", "", "Backend for the code to be run. Can be: Kubeflow,Argo,MLFlow.")
	flags.StringVarP(&run.Target, "target", "t", "", "Where the run need to be executed, it's the name of a registered cluster.")
	flags.StringVarP(&run.Extras, "extras", "x", "", "This is equivalent of requirements.txt, it is useful to install extra packages.")
	flags.StringVarP(&run.Entrypoint, "entrypoint", "e", "", "This is the entrypoint for KFP pipelines [ -e pipeline.py].")
	flags.BoolVarP(&run.Quiet, "quiet", "q", false, "Suppress output messages. Useful when k3ai is used within scripts.")
	flags.StringVarP(&run.Config, "config", "c", "", "Configure K3ai using a custom config file.[-c /path/tofile] [-c https://urlToFile]")
	return runCmd
}
