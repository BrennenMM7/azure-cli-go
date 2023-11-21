package account

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Azure Storage Account Commands",
	Long:  `Azure Storage Account Commands are used to manage Azure Storage Accounts.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
}
