package cmd

import (
	"fmt"
	"runtime"
  	"github.com/spf13/cobra"

  
)
// GitCommit returns the git commit that was compiled. This will be filled in by the compiler.
var GitCommit string

// Version returns the main version number that is being run at the moment.
const Version = "1.0"

// BuildDate returns the date the binary was built
var BuildDate = ""

// GoVersion returns the version of the go runtime used to compile the binary
var GoVersion = runtime.Version()

// OsArch returns the os and arch used to build the binary
var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
func versionCommand() *cobra.Command{
	versionCommand := &cobra.Command{
		Use:   "version",
		Short: "K3ai actual version. Print current binary version and info's.",
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Build Date:", BuildDate)
			fmt.Println("Git Commit:", GitCommit)
			fmt.Println("Version:", Version)
			fmt.Println("Go Version:", GoVersion)
			fmt.Println("OS / Arch:", OsArch)
		},
	  }
	  return versionCommand
}