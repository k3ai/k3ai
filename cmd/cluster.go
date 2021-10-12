package cmd

import (
	"fmt"
	"os"
	"log"


	"github.com/spf13/cobra"

	internal "github.com/k3ai/internal"
	names "github.com/k3ai/internal/names"
	color "github.com/k3ai/pkg/color"
	db "github.com/k3ai/pkg/db"
	clusterOperation "github.com/k3ai/pkg/io/clusters"

	tables "github.com/k3ai/pkg/tables"
)


func clusterCommand() *cobra.Command{
	cluster := internal.Options{}
	clusterCmd := &cobra.Command{
		Use:   "cluster",
		Short: "K3ai cluster management. Create/Delete a cluster environment.",
		
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		  color.Alert()
		  fmt.Println("This is and error")
		  color.Disable()
		  fmt.Println("This is not")
		},
	  }
	deployCmd := &cobra.Command{
		Use:"deploy [-t TYPE] [-n NAME] [-q] [-c]",
		Short: "Deploy a given cluster based on TYPE",
		Run: func(cmd *cobra.Command, args []string) {
			strType,_ := cmd.Flags().GetString("type")
			strConf,_ := cmd.Flags().GetString("config")
			boolQuiet, _ := cmd.Flags().GetBool("quiet")
			strName,_ := cmd.Flags().GetString("name")
			if len(args) >= 0 &&  strType == "" && strConf != "" && strName == ""{
				cmd.Help()
				os.Exit(0)
			}
			if !boolQuiet  && strName == ""{
				strName = names.GeneratedName(0)
				res := db.CheckClusterName(strName)
				if res != "" {
					strName = names.GeneratedName(1)
				}


				statusOk,_ := clusterOperation.Deployment(strName, strType)
				if statusOk {			
					clusterConfig := []string{strName,strType,"","Installed"}
					err := db.InsertCluster(clusterConfig)
					if err != nil {
						log.Fatal(err)
					}
					color.Done()
					fmt.Println(" ✔️ Installation Done.")
					// clusterOperation.Client(strName,strType)
				}
			} else if !boolQuiet && strName != "" {
				res := db.CheckClusterName(strName)
				if res != "" {
					strName = names.GeneratedName(1)
				}
				statusOk,_ := clusterOperation.Deployment(strName, strType)
				if statusOk {			
					clusterConfig := []string{strName,strType,"","Installed"}
					fmt.Printf("here")
					err := db.InsertCluster(clusterConfig)
					if err != nil {
						log.Fatal(err)
					}
					color.Done()
					fmt.Println(" ✔️ Installation Done.")
				}
			}
		},
	}
	removeCmd := &cobra.Command{
		Use:"remove [-n NAME] [-q] [-c]",
		Short: "Remove a given cluster based on NAME",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
	}
	listCmd := &cobra.Command{
		Use:"list",
		Short: "List installable cluster types or configuration of a given cluster.",
		Run: func(cmd *cobra.Command, args []string) {
			allList,_ := cmd.Flags().GetBool("all")
			nameList,_ := cmd.Flags().GetString("name")
			deployedList, _ := cmd.Flags().GetBool("deployed")

			if len(args) == 0 && !allList && nameList == "" && !deployedList{
				cmd.Help()
				os.Exit(0)
			} else {
				if allList && nameList != "" && ! deployedList{
					appsResults,infraResults,bundlesResults,commsResults := db.ListPlugins()
					tables.List("infra",appsResults,infraResults,bundlesResults,commsResults)
				} else if allList && !deployedList{
					appsResults,infraResults,bundlesResults,commsResults := db.ListPlugins()
					tables.List("infra",appsResults,infraResults,bundlesResults,commsResults)
				} else if !deployedList {
					results := db.ListPluginsByName(nameList)
					tables.ListByName(results)
				}
				if deployedList {
					results := db.ListClustersByName()
					tables.ListClusters(results)
				}
			}
		},
	}

	deployFlags := deployCmd.Flags()
	listFlags := listCmd.Flags()
	removeFlags := removeCmd.Flags()
	clusterCmd.DisableFlagsInUseLine = true
	clusterCmd.Flags().SortFlags = false

	//cluster subcommands
	clusterCmd.AddCommand(deployCmd,removeCmd,listCmd)

	//cluster deploy flags
	deployFlags.StringVarP(&cluster.Type,"type","t","","Select cluster type to be created/deleted")
	deployFlags.StringVarP(&cluster.Name,"name","n","","NAME of cluster to be created/deleted")
	deployFlags.BoolVarP(&cluster.Quiet,"quiet","q",false,"Suppress output messages. Useful when k3ai is used within scripts.")
	deployFlags.StringVarP(&cluster.Config,"config","c","","Configure K3ai using a custom config file.[-c /path/tofile] [-c https://urlToFile]")
	
	//cluster deploy flags
	removeFlags.StringVarP(&cluster.Name,"name","n","","NAME of cluster to be created/deleted")
	removeFlags.BoolVarP(&cluster.Quiet,"quiet","q",false,"Suppress output messages. Useful when k3ai is used within scripts.")
	removeFlags.StringVarP(&cluster.Config,"config","c","","Configure K3ai using a custom config file.[-c /path/tofile] [-c https://urlToFile]")
	
	//list listFlags available
	listFlags.BoolVarP(&cluster.All,"all","a",false,"Show all possible cluster configurations available.")
	listFlags.StringVarP(&cluster.Name,"name","n","","List configurations by CLUSTER NAME")
	listFlags.BoolVarP(&cluster.Deployed,"deployed","d",false,"List deployed clusters")
	
	
	return clusterCmd
}