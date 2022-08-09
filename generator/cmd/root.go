package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var RootCmd = &cobra.Command{
	Use: "genarator",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
}
