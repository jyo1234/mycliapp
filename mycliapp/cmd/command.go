package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "commit",
	Long:  "commit",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("commit changed")

	},
}
func init() {
	fmt.Println("ting tong")
	rootCmd.AddCommand(commitCmd)
}
