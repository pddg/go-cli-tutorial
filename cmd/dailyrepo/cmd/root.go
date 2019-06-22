package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const version = "v0.0.3"

var flgVerbose bool

var rootCmd = &cobra.Command{
	Use:   "dailyrepo",
	Short: "日報作成ツール",
	Long:  "テンプレートから日報の雛形を作成します",
	RunE: func(cmd *cobra.Command, args []string) error {
		v, _ := cmd.Flags().GetBool("version")
		if v {
			printVersion()
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		if flgVerbose {
			fmt.Println("Verbose output is enabled")
		}
	})
	rootCmd.Flags().BoolP("version", "v", false, "Print version")
	rootCmd.PersistentFlags().BoolVar(&flgVerbose, "verbose", false, "Print log")
}

func printVersion() {
	fmt.Printf("dailyrepo %s\n", version)
}
