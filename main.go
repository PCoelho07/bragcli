package main

import (
	"brag/cmd"
	"log"
	"os"
)

func main() { 
    rootCmd := cmd.CreateRootCmd()
    if err := rootCmd.Execute(); err != nil { 
        log.Fatal(err)
        os.Exit(1)
    }
}
