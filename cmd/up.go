package cmd

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"

	internal "github.com/k3ai/internal"
	auth "github.com/k3ai/pkg/auth"
	color "github.com/k3ai/pkg/color"
	db "github.com/k3ai/pkg/db"
	config "github.com/k3ai/pkg/env"
	http "github.com/k3ai/pkg/http"
	loader "github.com/k3ai/pkg/loader"
)

var envConfig *internal.Env
var token string

func upCommand() *cobra.Command {
	homeDir, _ := os.UserHomeDir()
	up := internal.Options{}
	upCmd := &cobra.Command{
		Use:   "up [-h --help] [-q --quiet] [-c fileOrUrl] [-p --pat]",
		Short: "K3ai starting point. Up configure K3ai to work on local environment",
		Run: func(cmd *cobra.Command, args []string) {
			bQuiet, _ := cmd.Flags().GetBool("quiet")
			sConfig, _ := cmd.Flags().GetString("config")
			pat, _ := cmd.Flags().GetString("pat")
			if _, err := os.Stat(homeDir + "/.config"); os.IsNotExist(err) {
				err := os.Mkdir(homeDir+"/.config", 0755)
				if err != nil {
					log.Fatal(err)
				}
			}
			if _, err := os.Stat(homeDir + "/.config/k3ai/.env"); !os.IsNotExist(err) {
				viper.AddConfigPath(homeDir + "/.config/k3ai")
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
				token = envConfig.K3AI_TOKEN
				_, err, _ := auth.GitHub(token)
				if err != nil {
					log.Fatal("GitHub Authentication Token:  NOT OK")
				}
			} else {
				if !bQuiet {
					color.White()
					time.Sleep(700 * time.Millisecond)
					fmt.Println("🎉🎉🎉 Welcome to K3ai 🎉🎉🎉")
					fmt.Println("📢	Give us a second and we will start the process...")
					time.Sleep(1 * time.Second)
					color.Alert()
					// comment from here to below to bypass
					fmt.Println(" ❌	Missing GitHub Authentication Token, please paste it here:")
					if pat != "" {
						token = pat
					} else {
						bytepw, err := term.ReadPassword(int(syscall.Stdin))
						if err != nil {
							os.Exit(1)
						}
						token = string(bytepw)
					}
					//comment above to bypass
					// token ="" //add token to bypass
					time.Sleep(700 * time.Millisecond)
					_, err, _ = auth.GitHub(token)
					if err != nil {
						fmt.Println(" ❌	GitHub Authentication Token:  NOT OK")
						os.Exit(1)
					} else {
						fmt.Println(" ✔️	GitHub Authentication Token: OK")
						time.Sleep(800 * time.Millisecond)
					}
					fmt.Println(" ❔	To avoid asking for the token everytime, are you ok if we save it as .env variable?")
					prompt := promptui.Select{
						Label: "Select[Yes/No]",
						Items: []string{"Yes", "No"},
					}
					_, result, err := prompt.Run()
					if err != nil {
						log.Fatalf("Prompt failed %v\n", err)
					}
					if result == "Yes" {
						err := os.Mkdir(homeDir+"/.config/k3ai", 0755)
						if err != nil {
							log.Fatal(err)
						}
						// os.WriteFile(homeDir + ".config/k3ai/.env",bytepw,0664)
						viper.AddConfigPath(homeDir + "/.config/k3ai")
						viper.Set("K3AI_TOKEN", token)
						err = viper.SafeWriteConfigAs(homeDir + "/.config/k3ai/.env")
						if err != nil {
							log.Fatal(err)
						}
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
				go config.InitConfig(ch, msg, sConfig)
				if !bQuiet {
					loader.StandardLoader(msg)
					color.Done()
					msg = "⏳	Completing configuration..."
					fmt.Printf("\r %v", msg)
				} else {
					color.Disable()
					msg = "Completing configuration..."
					fmt.Printf("\r %v", msg)
				}
				<-ch
				if !bQuiet {
					msg = "✔️	Done...                          "
					fmt.Printf("\r %v", msg)
				} else {
					msg = "Done...                          "
					fmt.Printf("\r %v", msg)
				}

				time.Sleep(900 * time.Millisecond)
				msg = "" //nolint:ineffassign,staticcheck
				msg = "Creating database...          "
				ch = make(chan bool)
				go db.InitDB(ch)
				if !bQuiet {
					loader.StandardLoader(msg)
				} else {
					log.Print(msg)
				}
				<-ch
				time.Sleep(500 * time.Millisecond)
				msg = "" //nolint:ineffassign,staticcheck
				msg = "Retrieving plugin list...     "
				ch = make(chan bool)
				action := "config"
				go http.RetrievePlugins(token, action, ch)
				if !bQuiet {
					loader.StandardLoader(msg)
					color.Done()
					msg = "⏳	Finishing retrieving plugins..."
					fmt.Printf("\r %v", msg)
				} else {
					color.Disable()
					msg = "Finishing Retrieving plugins..."
					fmt.Printf("\r %v", msg)
				}
				<-ch
				if !bQuiet {
					msg = "✔️	Done...                          "
					fmt.Printf("\r %v", msg)
				} else {
					msg = "Done...                          "
					fmt.Printf("\r %v", msg)
				}
				time.Sleep(500 * time.Millisecond)
				if !bQuiet {
					color.Done()
					fmt.Println(" ")
					fmt.Println(" ✔️	K3ai Configuration completed ")
				} else {
					log.Print("K3ai  Configuration completed")
				}

			} else {
				fmt.Println(" ✔️	Proceeding with configuration,please wait...")
				fmt.Printf("\n")
				time.Sleep(1 * time.Second)
				color.Done()
				os.Stdin.Close()
				color.Disable()
				ch := make(chan bool)
				msg := "Updating Plugins..."
				action := "update"
				go http.RetrievePlugins(token, action, ch)
				if !bQuiet {
					loader.StandardLoader(msg)
				}
				color.Done()
				msg = "⏳	Completing configuration..."
				fmt.Printf("\r %v", msg)
				time.Sleep(1 * time.Second)
				msg = "✔️	Done...                          "
				fmt.Printf("\r %v", msg)
				fmt.Println(" ")
				<-ch
			}
		},
	}
	flags := upCmd.Flags()
	flags.BoolVarP(&up.Quiet, "quiet", "q", false, "Suppress output messages. Useful when k3ai is used within scripts.")
	flags.StringVarP(&up.Config, "config", "c", "", "Configure K3ai using a custom config file.[-c /path/tofile] [-c https://urlToFile]")
	flags.StringVarP(&up.PAT, "pat", "p", "", "Send PAT (Personal Access Token) directly by skipping input.")
	return upCmd
}
