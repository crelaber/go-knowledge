// Package cmd /*
package cmd

import (
	"github.com/spf13/cobra"
)

// dbCmd represents the db command
var dbCmd = &cobra.Command{
	Use:     "db",
	Short:   "数据库相关的信息",
	Example: "go-knowledge db",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	RootCmd.AddCommand(dbCmd)
}

func run() {

}
