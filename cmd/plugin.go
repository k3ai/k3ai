package cmd

import (
	"os"
	"fmt"
  	"github.com/spf13/cobra"

	color "github.com/k3ai/pkg/color"
	"github.com/k3ai/internal"
  
)

func pluginCommand() *cobra.Command{
	plugin := internal.Options{}
	pluginCmd := &cobra.Command{
		Use:   "plugin",
		Short: "K3ai plugin management. Create/Delete a plugin environment.",
		
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
		Use:"deploy [-n NAME] [other flags]",
		Short: "Deploy a given plugin based on TYPE",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
	}
	removeCmd := &cobra.Command{
		Use:"remove [-n NAME] [other flags]",
		Short: "Remove a given plugin based on NAME",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
	}
	listCmd := &cobra.Command{
		Use:"list [-a --all] [-n NAME]",
		Short: "List installable plugin types or configuration of a given plugin.",
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
	pluginCmd.DisableFlagsInUseLine = true
	pluginCmd.Flags().SortFlags = false
	deployCmd.DisableFlagsInUseLine = true
	deployCmd.Flags().SortFlags = false
	removeCmd.DisableFlagsInUseLine = true
	removeCmd.Flags().SortFlags = false
	listCmd.DisableFlagsInUseLine = true
	listCmd.Flags().SortFlags = false

	//plugin subcommands
	pluginCmd.AddCommand(deployCmd,removeCmd,listCmd)

	//plugin deploy flags
	deployFlags.StringVarP(&plugin.Target,"target","t","","Target where to install plugin.")
	deployFlags.StringVarP(&plugin.Name,"name","n","","NAME of plugin to be created/deleted")
	deployFlags.BoolVarP(&plugin.Quiet,"quiet","q",false,"Suppress output messages. Useful when k3ai is used within scripts.")
	deployFlags.StringVarP(&plugin.Config,"config","c","","Configure K3ai using a custom config file.[-c /path/tofile] [-c https://urlToFile]")
	
	//plugin deploy flags
	removeFlags.StringVarP(&plugin.Name,"name","n","","NAME of plugin to be created/deleted")
	removeFlags.StringVarP(&plugin.Target,"target","t","","Target from where to remove plugin.")
	removeFlags.BoolVarP(&plugin.Quiet,"quiet","q",false,"Suppress output messages. Useful when k3ai is used within scripts.")
	removeFlags.StringVarP(&plugin.Config,"config","c","","Configure K3ai using a custom config file.[-c /path/tofile] [-c https://urlToFile]")
	
	//list listFlags available
	listFlags.StringVarP(&plugin.Config,"all","a","","Show all possible plugin configurations available.")
	listFlags.StringVarP(&plugin.Name,"name","n","","NAME of plugin to list")
		
	return pluginCmd
}