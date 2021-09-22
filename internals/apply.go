package internals

import (
	"strings"
	"gopkg.in/yaml.v2"
	"github.com/alefesta/k3ai/log"
	data "github.com/alefesta/k3ai/config"
	utils "github.com/alefesta/k3ai/shared"
	
)
var pluginContents = data.Plugin{}
func Apply() {
	var kubeconfig = ""
	var pluginName = ""
	utils.InitK8s(kubeconfig, pluginName)
}

func InfraDeployment(pluginUrl string, pluginName string) {
 
log.Info("Preparing to install: " + pluginName)
data,_ := getContent(pluginUrl)		
err := yaml.Unmarshal([]byte(data), &dataResults)
if err != nil {
	log.Error(err)
}
pluginUrl = strings.TrimSuffix(pluginUrl,"k3ai.yaml")
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
		pluginEx := string(pluginContents.Resources[0].Path)
		pluginArgs := string(pluginContents.Resources[0].Args)
		pluginKube := string(pluginContents.Resources[0].Kubecfg)
		pluginType := string(pluginContents.Resources[0].PluginType)
		pluginWait := pluginContents.Resources[0].Wait
		utils.InitExec(pluginName,pluginEx,pluginArgs,pluginKube,pluginType,pluginWait)
		}
	 
 	}
}

func BundlesDeployment() {

}