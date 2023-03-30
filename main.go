package main

import (
	"os"
	"github.com/christian-elsee/kafka/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
