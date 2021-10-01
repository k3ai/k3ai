package internals


import (
	"github.com/spf13/viper"
	"strings"
	"gopkg.in/yaml.v2"
	"github.com/k3ai/log"
	// data "github.com/k3ai/config"
	utils "github.com/k3ai/shared"
)


func Cluster(pluginUrl string, pluginName string, pluginType string) error {
 
	if pluginName == "k3s" {
		viper.Set("KUBECONFIG", "/etc/rancher/k3s/k3s.yaml")
	}
	
	data,_ := getContent(pluginUrl)		
	err := yaml.Unmarshal([]byte(data), &dataResults)
	if err != nil {
		log.Error(err)
	}
	pluginUrl = strings.TrimSuffix(pluginUrl,"k3ai.yaml")
	log.Info("Starting to install " + pluginName)
	 for i:=0; i < len(dataResults.Resources); i++ {
		 if strings.HasPrefix(dataResults.Resources[i], "../../") {
			pluginUrl = strings.TrimSuffix(pluginUrl,pluginName + "/k3ai.yaml")
			outer := strings.TrimPrefix(dataResults.Resources[i],"../../")
			log.Info("Outerfolder is: " + pluginUrl + outer)
		 } else if strings.HasPrefix(dataResults.Resources[i], "http://") || strings.HasPrefix(dataResults.Resources[i], "https://") {
			log.Info("url is: " + dataResults.Resources[i] )
		 } else  {	 	
			url := pluginUrl + dataResults.Resources[i] + "plugin.yaml"
			data,_ := getContent(url)		
			err := yaml.Unmarshal([]byte(data), &pluginContents)
			if err != nil {
				log.Error(err)
				}
			for i := range pluginContents.Resources {
				pluginEx := string(pluginContents.Resources[i].Path)
				pluginArgs := string(pluginContents.Resources[i].Args)
				pluginKube := string(pluginContents.Resources[i].Kubecfg)
				pluginType := string(pluginContents.Resources[i].PluginType)
				pluginWait := pluginContents.Resources[i].Wait
				err := utils.InitExec(pluginName,pluginEx,pluginArgs,pluginKube,pluginType,pluginWait)
				if err != nil {
					log.Error(err)
					return err
				}
			}
	
			}
		 
		 }
		 if pluginName == "k3s" {
			log.Warn("Do not forget to add K3s config file to your KUBECONFIG variable...")
			log.Warn("Please copy and paste the following line...")
			log.Warn("export KUBECONFIG=/etc/rancher/k3s/k3s.yaml")
			log.Info("Cluster is up and running enjoy K3ai...")
		 }
		 return nil
	}