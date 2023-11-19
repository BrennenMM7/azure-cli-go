/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package login

import (
	"github.com/spf13/cobra"

	"github.com/BrennenMM7/azure-cli-go/pkg/auth"
)

// loginCmd represents the login command
var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Azure",
	Long: `Login to Azure using the default browser. 
		You will be prompted to enter your credentials.
		
		Example: az login
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if auth.CheckIfAuthenticationFileExists() {
			auth.DefaultLogin()
		} else {
			auth.FirstTimeLogin()
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
