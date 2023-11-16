package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "azure",
	Version: "0.0.1",
	Short:   "A fast and flexible CLI for Azure",
	Long:    `A fast and flexible CLI for Azure built with love by, and for, the Azure community in Go.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
