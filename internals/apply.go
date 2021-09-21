package internals
/**
 *
 * @todo Read the named plugin and store in memory
 * @body Cache data and automatically use them to install the plugin
 */
import (
	"gopkg.in/yaml.v2"
	"github.com/alefesta/k3ai/log"
	utils "github.com/alefesta/k3ai/shared"
	data "github.com/alefesta/k3ai/config"
)

var dataRes *data.K3ai

func Apply() {
	utils.InitK8s()
}

func InfraDeployment(pluginUrl string, pluginName string) {

log.Info("Preparing to install: " + pluginName)
pluginUrl = "https://raw.githubusercontent.com/k3ai/plugins/main/infra/k3s/k3ai.yaml"
data,_ := getContent(pluginUrl)		
err := yaml.Unmarshal([]byte(data), &dataResults)
if err != nil {
	log.Error(err)
}
log.Info(dataRes.Api)
}

func BundlesDeployment() {

}