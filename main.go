package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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
	var err error
	args := rootCmd.Args()
	if len(args) > 0 {
		switch args[0] {
		case "add":
			_ = addCmd.Parse(args[1:])
			err = handleAddCmd(fileName)
		default:
			fmt.Printf("Unknown command: %v\n", args[1:])
			os.Exit(2)
		}
	}
	if err != nil {
		fmt.Println(err)
	}
}

func handleAddCmd(fileName string) error {
	filePath := fmt.Sprintf("./templates/%s", fileName)
	byteTmpl, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	fmt.Println(string(byteTmpl))
	return nil
}
