package cmd

import (
	"time"
	"strings"
	"github.com/briandowns/spinner"
	log "github.com/k3ai/log"
    internals "github.com/k3ai/internals"
	shared "github.com/k3ai/shared"
	"github.com/spf13/cobra"
)



// createCmd represents the version command
var createCmd = &cobra.Command{
	Use:  "create",
	Short: "create a K3ai plugin.",
	Long:  `
create is meant to uninstall a specific kind of plugin: application or bundle.
Through the create command a user may have a certain plugin created from the target device.
`,
	Example: `
k3ai create	<plugin name> --type <cluster type> --name <cluster name>
	`,

}

// clusterCmd represents the version command
var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "create a K3ai plugin.",
	Long:  `
create is meant to create a cluster based on the cluster name type (i.e: k3s or eksa)
`,
	Run: func(cmd *cobra.Command, args []string) {
		clusterType,_ := cmd.Flags().GetString("type")
		clusterName,_ := cmd.Flags().GetString("name")
		if clusterName == "" {
			//TODO: generate cluster name
		}

		icon := []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
		s := spinner.New(icon, 100*time.Millisecond,spinner.WithColor("fgHiYellow"))
		s.Start()
		time.Sleep(500 * time.Millisecond)
		log.Info("Starting cluster installation...")
		_, clusterUrl, err := shared.SelectPlugin(strings.ToLower(clusterType))
		_ = log.CheckErrors(err)
		time.Sleep(500 * time.Millisecond)
		err = internals.Cluster(clusterUrl,clusterName,clusterType)
		_ = log.CheckErrors(err)
	},
	Example: `
k3ai create	<plugin name> --type <cluster type> --name <cluster name>
	`,

}
func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(clusterCmd)
	clusterCmd.Flags().String("type","","The type of cluster to create as listed through k3ai list --type infra")
	clusterCmd.Flags().String("name","","The name of the cluster. This is the name you will refer to not necessarly the real cluster name. If omitted a generated name will be used.")
	clusterCmd.MarkFlagRequired("type")
}