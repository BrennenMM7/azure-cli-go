/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package storage

import (
	"fmt"

	"github.com/BrennenMM7/azure-cli-go/cmd/storage/account"
	"github.com/spf13/cobra"
)

// storageCmd represents the storage command
var StorageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Azure Storage Commands",
	Long: `Azure Storage Commands are used to manage Azure Storage Accounts and Blob Containers.
		These commands are used to create, delete, and list storage accounts and blob containers.
		They are also used to upload and download blobs to and from blob containers.
		`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("storage called")
	},
}

func init() {
	StorageCmd.AddCommand(account.AccountCmd)
}
