package infra

import (
	"os"
	"fmt"
	"log"
	"syscall"

	"golang.org/x/term"
	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
	"k8s.io/client-go/util/homedir"

	internal "github.com/k3ai/internal"
)
var envConfig *internal.Env

func InitCivo () (civoKey string, err error) {
	homeDir := homedir.HomeDir()
	viper.AddConfigPath(homeDir + "/.config/k3ai")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&envConfig)

	if envConfig.CIVO_TOKEN == "" {
		// log.Fatal(err)
		fmt.Println(" ❔	It seems we do not have any Civo API Key stored, to avoid asking for it everytime, are you ok if we save it as .env variable?")
		prompt := promptui.Select{
			Label: "Select[Yes/No]",
			Items: []string{"Yes", "No"},
		}
		_, result, err := prompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
		}
		if result == "Yes" {
			bytepw, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				os.Exit(1)
			}
			token := string(bytepw)
			viper.AddConfigPath(homeDir + "/.config/k3ai")
			viper.Set("CIVO_TOKEN", token)
			err = viper.WriteConfigAs(homeDir + "/.config/k3ai/.env")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	civoKey = envConfig.CIVO_TOKEN
	return civoKey,err
}

func ConfigCivo(clusterName string, clusterRegion string) error {
	homeDir := homedir.HomeDir()
	civoPath := homeDir + "/.k3ai/" + clusterName
	if _, err := os.Stat(civoPath); os.IsNotExist(err) {
		err := os.Mkdir(civoPath, 0755)
		if err != nil {
			log.Fatal(err)
		}

	}
	data :=  []byte(clusterName + " " + clusterRegion)
	err := os.WriteFile(civoPath + "/" + clusterName + "." + clusterRegion,data,0755)
	return err
}