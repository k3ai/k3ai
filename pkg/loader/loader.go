package loader

import (
	"fmt"
	"log"
	"time"

	"github.com/schollz/progressbar/v3"
)

func StandardLoader(msg string) {
	tasks := 100
	bar := progressbar.NewOptions(tasks,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetDescription("[green] 🚀	"+msg+"[yellow]"),
		progressbar.OptionOnCompletion(func() {
			fmt.Printf("\n")

		}))
	for i := 0; i < tasks; i++ {
		err := bar.Add(1)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(40 * time.Millisecond)
	}
}

func SuperLoader(msg string) {
	tasks := 100
	bar := progressbar.NewOptions(tasks,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetDescription("[green]🧪	"+msg+"[yellow]"),
		progressbar.OptionOnCompletion(func() {
			fmt.Printf("\n")

		}))
	for i := 0; i < tasks; i++ {
		err := bar.Add(1)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(40 * time.Millisecond)
	}
}

func Working(msg string) {
	tasks := -1
	bar := progressbar.NewOptions(tasks,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetDescription("[green]🧪	"+msg+"[yellow]"),
		progressbar.OptionOnCompletion(func() {
			fmt.Printf("\n")

		}))
	for i := 0; i < tasks; i++ {
		err := bar.Add(1)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(40 * time.Millisecond)
	}
}
