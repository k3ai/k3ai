package internals
import (
	"github.com/alefesta/k3ai/log"
	utils "github.com/alefesta/k3ai/shared"
)

func Apply() {
	utils.InitK8s()
}

func InfraDeployment(pluginUrl string, pluginName string) {
log.Info("Preparing to install: " + pluginName)

}

func BundlesDeployment() {

}