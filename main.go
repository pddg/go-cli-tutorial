package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var flgVersion bool
var flgVerbose bool

func main() {
	rootCmd := flag.NewFlagSet("Root", flag.ContinueOnError)
	rootCmd.BoolVar(&flgVersion, "v", false, "Print version")
	rootCmd.BoolVar(&flgVersion, "version", false, "Print version")
	rootCmd.BoolVar(&flgVerbose, "verbose", false, "Output log")

	if err := rootCmd.Parse(os.Args[1:]); err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		}
		log.Fatal(err)
	}
	if flgVersion {
		fmt.Println("dailyrepo v0.0.1")
	}
}
