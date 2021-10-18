package cmd

import (
	"fmt"
	"log"
	"os"
	"time"
	// "syscall"

	// "golang.org/x/term"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	internal "github.com/k3ai/internal"
	auth "github.com/k3ai/pkg/auth"
	color "github.com/k3ai/pkg/color"
	db "github.com/k3ai/pkg/db"
	config "github.com/k3ai/pkg/env"
	loader "github.com/k3ai/pkg/loader"
	http "github.com/k3ai/pkg/http"
)
var envConfig *internal.Env
var token string
func upCommand() *cobra.Command{
	homeDir,_ := os.UserHomeDir()
	up := internal.Options{}
	upCmd := &cobra.Command{
		Use:   "up [-h --help] [-q --quiet] [-f --force] [-c fileOrUrl]",
		Short: "K3ai starting point. Up configure K3ai to work on local environment",
		Run: func(cmd *cobra.Command, args []string) {
			bQuiet,_ :=cmd.Flags().GetBool("quiet")
			bForce, _ := cmd.Flags().GetBool("force")
			if _, err := os.Stat(homeDir + "/.k3ai/.env"); !os.IsNotExist(err) {
			
			    viper.AddConfigPath(homeDir + "/.k3ai")
				viper.SetConfigName(".env")
				viper.SetConfigType("env")
				viper.AutomaticEnv()
				err = viper.ReadInConfig()
				if err != nil {
					log.Fatal(err)
				}
				err = viper.Unmarshal(&envConfig)
				if err != nil {
					log.Fatal(err)
				}
				token = envConfig.GH_AUTH_TOKEN
				_,err,_ := auth.GitHub(token)
				if err != nil {
					log.Fatal("GitHub Authentication Token:  NOT OK")
				} 
				} else {
				if !bQuiet {
					color.Alert()
					// fmt.Println(" ❌	Missing GitHub Authentication Token, please paste it here:")
					// bytepw, err := term.ReadPassword(int(syscall.Stdin))
					// if err != nil {
					// 	os.Exit(1)
					// } 
					time.Sleep(700 * time.Millisecond)
					// token = string(bytepw)
					token ="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"
					
					_,err,_ = auth.GitHub(token)
					if err != nil {
						fmt.Println(" ❌	GitHub Authentication Token:  NOT OK")
						os.Exit(1)
					}else {
					fmt.Println(" ✔️	GitHub Authentication Token: OK")
					time.Sleep(800 * time.Millisecond)
					}
					fmt.Println(" ✔️	Proceeding with configuration,please wait...")
					fmt.Printf("\n")
					time.Sleep(1 * time.Second)
					color.Done()
					os.Stdin.Close()
					color.Disable()
				} else {
					log.Println("Please create a .env file and place it under K3ai folder. See docs https://docs.k3ai.in")
				}
			}

			if _, err := os.Stat(homeDir + "/.k3ai/k3ai.db"); os.IsNotExist(err) { 
			msg := "Loading K3ai configuration... "
			ch := make(chan bool)
			go config.InitConfig(ch,msg,bForce)
			if !bQuiet {
				loader.StandardLoader(msg)	
			} 
			color.Done()
			msg = "⏳	Completing configuration..."
			fmt.Printf("\r %v", msg)
			<-ch
			msg = "✔️	Done...                          "
			fmt.Printf("\r %v", msg)
			time.Sleep(900 * time.Millisecond)
			msg = ""
			msg = "Creating database...          "
			ch = make(chan bool)
			go db.InitDB(ch)
			if !bQuiet {
				loader.StandardLoader(msg)	
			}else{
				log.Print(msg)
			}
			time.Sleep(500 * time.Millisecond)
			msg = ""
			msg = "Retrieving plugin list...     "
			action:="config"
			http.RetrievePlugins(token,action)
			if !bQuiet {
				loader.StandardLoader(msg)	
			}else{
				log.Print(msg)
			}
			if !bQuiet {
				color.Done()
				fmt.Println(" ✔️	K3ai Configuration completed ")	
			} else {
				log.Print("K3ai  Configuration completed")
			}
		}else {
			fmt.Println(" ✔️	Proceeding with configuration,please wait...")
			fmt.Printf("\n")
			time.Sleep(1 * time.Second)
			color.Done()
			os.Stdin.Close()
			color.Disable()
			// ch := make(chan bool)
			msg := "Updating Plugins..."
			action:="update"
			go 	http.RetrievePlugins(token,action)
			if !bQuiet {
				loader.StandardLoader(msg)	
			} 
			color.Done()
			msg = "⏳	Completing configuration..."
			fmt.Printf("\r %v", msg)
			// <-ch
			// ch <- true
			time.Sleep(1 * time.Second)
			msg = "✔️	Done...                          "
			fmt.Printf("\r %v", msg)
			fmt.Println(" ")
		}
		},
	  }
	  flags := upCmd.Flags()
	  flags.BoolVarP(&up.Quiet,"quiet","q",false,"Suppress output messages. Useful when k3ai is used within scripts.")
	  flags.BoolVarP(&up.Force,"force","f",false,"Force re-configuration of K3ai. Will overwrite existing configuration.")
	  flags.StringVarP(&up.Config,"config","c","","Configure K3ai using a custom config file.[-c /path/tofile] [-c https://urlToFile]")
	  return upCmd
}

