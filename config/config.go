package config

import (
	"os"
	"fmt"
	"time"
	"syscall"

	"golang.org/x/term"
	"github.com/spf13/viper"
)

// Provider defines a set of read-only methods for accessing the application
// configuration params as defined in one of the config files.
type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
}

var defaultConfig *viper.Viper

// Config returns a default config providers
func Config() Provider {
	return defaultConfig
}

// LoadConfigProvider returns a configured viper instance
func LoadConfigProvider(appName string) Provider {
	return readViperConfig(appName)
}

func init() {
	defaultConfig = readViperConfig("K3AI")
}

func readViperConfig(appName string) *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix(appName)
	v.AutomaticEnv()

	// global defaults
	
	v.SetDefault("json_logs", false)
	v.SetDefault("loglevel", "debug")
	

	return v
}


func InitEnv() {
	//set the default configurations 
	fmt.Println("Missing GitHub authentication token, please paste it here: ")
	bytePassword, _:= term.ReadPassword(int(syscall.Stdin))
	token := string(bytePassword)

	viper.SetDefault("GITHUB_AUTH_TOKEN",token)
	viper.SetDefault("K3AI_REPO","https://github.com/k3ai/plugins")
	viper.SetDefault("COMMUNITY",false)
	viper.SetDefault("SYNC", false)

	//set the path and name
	homeDir,_ := os.UserHomeDir()
	viper.SetConfigName(".env") // name of config file (without extension)
	viper.SetConfigType("env") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(homeDir +  "/.k3ai/" )   // path to look for the config file in
	viper.AddConfigPath("$HOME/.k3ai")  // call multiple times to add many search paths

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.SafeWriteConfigAs(homeDir + "/.k3ai/.env")
		}
	}
}

