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

		// Login to Azure
		if auth.CheckIfAuthenticationFileExists() {
			auth.DefaultLogin()
		} else {
			auth.FirstTimeLogin()
		}

		// Obtain Azure Profile Metadata
		
	},
}

func init() {

}
