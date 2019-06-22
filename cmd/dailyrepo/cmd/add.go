package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "今日の日報を追加",
	Long:  "テンプレートを元に今日の日報の雛形を作成する",
	RunE: func(cmd *cobra.Command, args []string) error {
		fileName, _ := cmd.Flags().GetString("name")
		_ = generateReport(fileName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("name", "n", time.Now().Format("2006-01-02")+".md", "filename")
}

func generateReport(filename string) error {
	fmt.Println(filename)
	return nil
}
