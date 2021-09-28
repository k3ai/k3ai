package internals

import (
	"github.com/spf13/viper"
	"strings"
	"gopkg.in/yaml.v2"
	"github.com/alefesta/k3ai/log"
	data "github.com/alefesta/k3ai/config"
	utils "github.com/alefesta/k3ai/shared"

	
)
var pluginContents = data.Plugin{}
var pluginRoot = "https://raw.githubusercontent.com/k3ai/plugins/main/"
// var dataResults = data.K3ai{}

func AppsDeployment(pluginUrl string, pluginName string) {
	log.Info("Preparing to install: " + pluginName)

	err := rootPlugin(pluginName, pluginUrl)
	if err != nil {
		log.Error(err)
	}

	
}

func InfraDeployment(pluginUrl string, pluginName string) {
 
log.Info("Preparing to install: " + pluginName)
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
}

func BundlesDeployment() {

}

func rootPlugin(pluginName string,pluginUrl string) error {
	data,_ := getContent(pluginUrl)		
	err := yaml.Unmarshal([]byte(data), &dataResults)
	if err != nil {
		log.Error(err)
	}
	pluginUrl = strings.TrimSuffix(pluginUrl,"k3ai.yaml")
	log.Info("Starting to install " + pluginName)
	 for i:=0; i < len(dataResults.Resources); i++ {
		 if strings.HasPrefix(dataResults.Resources[i], "../../") {

			outer := strings.TrimPrefix(dataResults.Resources[i],"../../")
			urlRoot := pluginRoot + outer + "/k3ai.yaml"
			data,_ := getContent(urlRoot)		
			err := yaml.Unmarshal([]byte(data), &dataResults)

			if err != nil {
				log.Error(err)
				}
			urlRoot = strings.TrimSuffix(urlRoot,"k3ai.yaml")
			err = innerPlugin(pluginName, urlRoot)
			if err != nil {
				log.Error(err)
			}

		 }
		 if strings.HasPrefix(dataResults.Resources[i], "http://") || strings.HasPrefix(dataResults.Resources[i], "https://") {
			log.Info("url is: " + dataResults.Resources[i] )
		 }

		err = innerPlugin(pluginName, pluginUrl)
		if err != nil {
			log.Error(err)
		}
	}
 return err
}

func innerPlugin(pluginName string,urlRoot string) error  {
	for s:=0; s < len(dataResults.Resources); s++ {
		url := urlRoot + dataResults.Resources[s] + "plugin.yaml"
		data,_ := getContent(url)		
		err := yaml.Unmarshal([]byte(data), &pluginContents)
		if err != nil {
			log.Error(err)
			}
		for d:=0; d < len(pluginContents.Resources); d++ {
			pluginEx := string(pluginContents.Resources[d].Path)
			pluginArgs := string(pluginContents.Resources[d].Args)
			pluginKube := string(pluginContents.Resources[d].Kubecfg)
			pluginType := string(pluginContents.Resources[d].PluginType)
			pluginWait := pluginContents.Resources[d].Wait
			err = utils.InitExec(pluginName,pluginEx,pluginArgs,pluginKube,pluginType,pluginWait)
			if err != nil {
				log.Error(err)
			}
			
		}
	}
return nil
}