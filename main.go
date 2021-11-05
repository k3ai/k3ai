package main

import (
	"log"

	"github.com/k3ai/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
