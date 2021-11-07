package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"github.com/k3ai/internal"
	color "github.com/k3ai/pkg/color"
	db "github.com/k3ai/pkg/db"
	"github.com/k3ai/pkg/http"
	pluginOperations "github.com/k3ai/pkg/io/execution"
	tables "github.com/k3ai/pkg/tables"
)

func deployPlugin(strName string, strTarget string, extraArray string) {
	statusOk, _ := pluginOperations.Deployment("plugin", strName, strTarget, extraArray)
	if statusOk {
		_, clusterType := db.CheckClusterName(strTarget)
		out := pluginOperations.Client(strTarget, clusterType)
		os.Remove(out)
		fmt.Println(" ")
		strIP := http.GetIP()
		switch strName {
		case "mlflow":
			fmt.Println("We tried to publish MLFLow at:http://" + strIP + ":30500")
		case "kf-pa":
			fmt.Println("We tried to publish Kubeflow Pipelines at:http://" + strIP + ":30900")
		case "kf-katib":
			fmt.Println("We tried to publish Kubeflow Katib at:http://" + strIP + ":30600")
		case "airflow":
			fmt.Println("We tried to publish Apache Airflow at:http://" + strIP + ":30800")
		case "argo-workflow":
			fmt.Println("We tried to publish Argo Workflows at:http://" + strIP + ":32746")
		}

		color.Done()
		fmt.Println(" ✔️ Installation Done.")
	} else {
		log.Println("Error occurred while deploying plugin.")
	}
}

func pluginCommand() *cobra.Command {
	plugin := internal.Options{}
	pluginCmd := &cobra.Command{
		Use:   "plugin",
		Short: "K3ai plugin management. Create/Delete a plugin environment.",

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				err := cmd.Help()
				if err != nil {
					log.Fatal(err)
				}
				os.Exit(0)
			}
		},
	}
	deployCmd := &cobra.Command{
		Use:   "deploy [-n NAME] [other flags]",
		Short: "Deploy a given plugin based on TYPE",
		Run: func(cmd *cobra.Command, args []string) {
			strTarget, _ := cmd.Flags().GetString("target")
			strConf, _ := cmd.Flags().GetString("config")
			boolQuiet, _ := cmd.Flags().GetBool("quiet")
			strName, _ := cmd.Flags().GetString("name")
			strName = strings.ToLower(strName)
			if len(args) == 0 && strTarget == "" && strConf == "" && strName == "" {
				err := cmd.Help()
				if err != nil {
					log.Fatal(err)
				}
				os.Exit(0)
			} else if len(args) >= 0 && strTarget == "" && strConf == "" && strName == "" {
				err := cmd.Help()
				if err != nil {
					log.Fatal(err)
				}
				os.Exit(0)
			}
			if strConf != "" {
				if strings.TrimSpace(strConf)[:4] == "http" {
					data, _ := http.Download(strConf)
					err := yaml.Unmarshal([]byte(data), &k3aiConfig)
					if err != nil {
						log.Fatal(err)
					}
					strName = k3aiConfig.Plugin.Name
					strTarget = k3aiConfig.Plugin.Target
					deployPlugin(strName, strTarget, extraArray)
				} else {
					data, err := ioutil.ReadFile(strConf)
					if err != nil {
						log.Fatal(err)
					}
					err = yaml.Unmarshal(data, &k3aiConfig)
					if err != nil {
						log.Fatal(err)
					}
					strName = k3aiConfig.Plugin.Name
					strTarget = k3aiConfig.Plugin.Target
					deployPlugin(strName, strTarget, extraArray)
				}
			} else if !boolQuiet && strName == "" {
				statusOk, _ := pluginOperations.Deployment("plugin", strName, strTarget, extraArray)
				if statusOk {
					clusterConfig := []string{strName, strTarget, "", "Installed"}
					err := db.InsertCluster(clusterConfig)
					if err != nil {
						log.Fatal(err)
					}

					color.Done()
					fmt.Println(" ✔️ Installation Done.")
					// pluginOperations.Client(strName,strTarget)
				}
			} else if !boolQuiet && strName != "" {
				deployPlugin(strName, strTarget, extraArray)
			}
		},
	}
	removeCmd := &cobra.Command{
		Use:   "remove [-n NAME] [other flags]",
		Short: "Remove a given plugin based on NAME",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				err := cmd.Help()
				if err != nil {
					log.Fatal(err)
				}
				os.Exit(0)
			}
		},
	}
	listCmd := &cobra.Command{
		Use:   "list [-a --all] [-n NAME]",
		Short: "List installable plugin types or configuration of a given plugin.",
		Run: func(cmd *cobra.Command, args []string) {
			allList, _ := cmd.Flags().GetBool("all")
			nameList, _ := cmd.Flags().GetString("name")

			if len(args) == 0 && !allList && nameList == "" {
				err := cmd.Help()
				if err != nil {
					log.Fatal(err)
				}
				os.Exit(0)
			} else {
				if allList && nameList != "" {
					appsResults, infraResults, bundlesResults, commsResults := db.ListPlugins()
					tables.List("plugin", appsResults, infraResults, bundlesResults, commsResults)
				} else if allList {
					appsResults, infraResults, bundlesResults, commsResults := db.ListPlugins()
					tables.List("plugin", appsResults, infraResults, bundlesResults, commsResults)
				} else {
					results := db.ListPluginsByName(nameList)
					tables.ListByName(results)
				}
			}
		},
	}

	deployFlags := deployCmd.Flags()
	listFlags := listCmd.Flags()
	removeFlags := removeCmd.Flags()
	pluginCmd.DisableFlagsInUseLine = true
	pluginCmd.Flags().SortFlags = false
	deployCmd.DisableFlagsInUseLine = true
	deployCmd.Flags().SortFlags = false
	removeCmd.DisableFlagsInUseLine = true
	removeCmd.Flags().SortFlags = false
	listCmd.DisableFlagsInUseLine = true
	listCmd.Flags().SortFlags = false

	//plugin subcommands
	pluginCmd.AddCommand(deployCmd, removeCmd, listCmd)

	//plugin deploy flags
	deployFlags.StringVarP(&plugin.Target, "target", "t", "", "Target where to install plugin.")
	deployFlags.StringVarP(&plugin.Name, "name", "n", "", "NAME of plugin to be created/deleted")
	deployFlags.BoolVarP(&plugin.Quiet, "quiet", "q", false, "Suppress output messages. Useful when k3ai is used within scripts.")
	deployFlags.StringVarP(&plugin.Config, "config", "c", "", "Configure K3ai using a custom config file.[-c /path/tofile] [-c https://urlToFile]")

	//plugin deploy flags
	removeFlags.StringVarP(&plugin.Name, "name", "n", "", "NAME of plugin to be created/deleted")
	removeFlags.StringVarP(&plugin.Target, "target", "t", "", "Target from where to remove plugin.")
	removeFlags.BoolVarP(&plugin.Quiet, "quiet", "q", false, "Suppress output messages. Useful when k3ai is used within scripts.")
	removeFlags.StringVarP(&plugin.Config, "config", "c", "", "Configure K3ai using a custom config file.[-c /path/tofile] [-c https://urlToFile]")

	//list listFlags available
	listFlags.BoolVarP(&plugin.All, "all", "a", false, "Show all possible plugin configurations available.")
	listFlags.StringVarP(&plugin.Name, "name", "n", "", "List plugins by CLUSTER NAME")

	return pluginCmd
}
