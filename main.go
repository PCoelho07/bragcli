package main

import (
	"brag/cmd"
	"log"
)

func main() {
	rootCmd := cmd.CreateRootCmd()
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
