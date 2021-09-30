package internals

import (
	"strings"
	"gopkg.in/yaml.v2"
	"github.com/k3ai/log"
	// data "github.com/k3ai/config"
	utils "github.com/k3ai/shared"
	
)
// var pluginContents = data.Plugin{}
func Remove() {
	// var kubeconfig = ""
	// var pluginName = ""
	// utils.InitK8s(kubeconfig, pluginName)
}

func InfraRemoval(pluginUrl string, pluginName string) {
 
log.Info("Preparing to remove: " + pluginName)
data,_ := getContent(pluginUrl)		
err := yaml.Unmarshal([]byte(data), &dataResults)
if err != nil {
	log.Error(err)
}
pluginUrl = strings.TrimSuffix(pluginUrl,"k3ai.yaml")
log.Info("Starting to remove " + pluginName)
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
			pluginRemove := pluginContents.Resources[i].Remove
			utils.InitRemove(pluginName,pluginEx,pluginArgs,pluginKube,pluginType,pluginWait,pluginRemove)
		}

		}
	 
 	}
log.Info(pluginName + " correctly removed.")
}

func BundlesRemoval() {

}