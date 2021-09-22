package internals
/**
 *
 * @to-do Read the named plugin and store in memory
 * @body Cache data and automatically use them to install the plugin
 */
import (
	"gopkg.in/yaml.v2"
	"github.com/alefesta/k3ai/log"
	utils "github.com/alefesta/k3ai/shared"
	
)

func Apply() {
	utils.InitK8s()
}

func InfraDeployment(pluginUrl string, pluginName string) {

log.Info("Preparing to install: " + pluginName)
data,_ := getContent(pluginUrl)		
err := yaml.Unmarshal([]byte(data), &dataResults)
if err != nil {
	log.Error(err)
}
log.Info(dataResults.Resources)
}

func BundlesDeployment() {

}