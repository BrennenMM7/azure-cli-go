package account

import (
	"fmt"

	"github.com/spf13/cobra"
)

// storageCmd represents the storage command
var AccountCmd = &cobra.Command{
	Use:   "account",
	Short: "Azure Storage Account Commands",
	Long:  `Azure Storage Account Commands are used to manage Azure Storage Accounts.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("account called")
	},
}

func init() {
	AccountCmd.AddCommand(listCmd)
}
