package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var flgVersion bool
var flgVerbose bool
var fileName string

func main() {
	rootCmd := flag.NewFlagSet("Root", flag.ContinueOnError)
	rootCmd.BoolVar(&flgVersion, "v", false, "Print version")
	rootCmd.BoolVar(&flgVersion, "version", false, "Print version")
	rootCmd.BoolVar(&flgVerbose, "verbose", false, "Output log")

	addCmd := flag.NewFlagSet("Add", flag.ContinueOnError)
	addCmd.StringVar(&fileName, "name", time.Now().Format("2006-01-02")+".md", "File name")

	if err := rootCmd.Parse(os.Args[1:]); err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		}
		log.Fatal(err)
	}
	if flgVersion {
		fmt.Println("dailyrepo v0.0.1")
	}

	// Handle sub commands
	args := rootCmd.Args()
	if len(args) > 0 {
		switch args[0] {
		case "add":
			_ = addCmd.Parse(args[1:])
			fmt.Println(fileName)
		default:
			fmt.Printf("Unknown command: %v\n", args[1:])
			os.Exit(2)
		}
	}
}
