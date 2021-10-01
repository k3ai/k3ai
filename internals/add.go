package internals

import (
	// "github.com/spf13/viper"
	"strings"
	"gopkg.in/yaml.v2"
	"github.com/k3ai/log"
	data "github.com/k3ai/config"
	utils "github.com/k3ai/shared"
)

var pluginContents = data.Plugin{}
var pluginRoot = "https://raw.githubusercontent.com/k3ai/plugins/main/"
// var dataResults = data.K3ai{}

//AppsDeployment install the selected plugin
func AppsDeployment(pluginUrl string, pluginName string) {
	log.Info("Preparing to install: " + pluginName)

	err := rootPlugin(pluginName, pluginUrl)
	if err != nil {
		log.Error(err)
	}

	
}


//BundlesDeployment install a bundle
func BundlesDeployment() {
	log.Info("bundle installed")
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