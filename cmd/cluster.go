package cmd

import (
	"os"
	"fmt"
  	"github.com/spf13/cobra"
	

	color "github.com/k3ai/pkg/color"
	internal "github.com/k3ai/internal"
  
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
			if len(args) >= 0 &&  strType == "" && strConf != ""{
				cmd.Help()
				os.Exit(0)
			}
			if !boolQuiet {

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
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
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
	listFlags.StringVarP(&cluster.Config,"all","a","","Show all possible cluster configurations available.")
	listFlags.StringVarP(&cluster.Name,"name","n","","NAME of cluster to list")
	
	
	return clusterCmd
}