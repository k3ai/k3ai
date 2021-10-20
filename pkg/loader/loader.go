package loader

import (
	"fmt"
	"time"

	"github.com/schollz/progressbar/v3"
)

func StandardLoader(msg string) {
	tasks := 100
	bar := progressbar.NewOptions(tasks,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetDescription("[green] 🚀	"+ msg + "[yellow]"),
		progressbar.OptionOnCompletion(func() {
			fmt.Printf("\n")

		}))
		for i := 0; i < tasks; i++ {
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)
		}
}

func SuperLoader(msg string) {
	tasks := 100
	bar := progressbar.NewOptions(tasks,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetDescription("[green]🧪	"+ msg + "[yellow]"),
		progressbar.OptionOnCompletion(func() {
			fmt.Printf("\n")

		}))
		for i := 0; i < tasks; i++ {
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)
		}
}

