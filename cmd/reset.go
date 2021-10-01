package cmd

import (
	"os"
	"time"
	log "github.com/k3ai/log"
	"github.com/spf13/cobra"
	"github.com/briandowns/spinner"
)

// resetCmd represents the version command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset K3ai.",
	Long:  `
Reset will uninstall k3ai from the system.
`,
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: Remove any local cluster before remove binary
		//BODY: In order to clean up the system we should have a script 
		// that manage the logic
		icon := []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
		s := spinner.New(icon, 100*time.Millisecond,spinner.WithColor("fgHiYellow"))
		s.Start()
		time.Sleep(1 * time.Second)
		log.Info("Removing K3ai from the system...")
		time.Sleep(1 * time.Second)
		homeDir,_ := os.UserHomeDir()
		err := os.RemoveAll(homeDir + "/.k3ai")
		time.Sleep(1 * time.Second)
		log.Info("K3ai successfully removed.Have a nice day!")
		if err != nil {
			log.Error(err)
		}
	},
	Example: `
k3ai reset
	`,

}

func init() {
	
	rootCmd.AddCommand(resetCmd)

}