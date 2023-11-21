package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	_ "github.com/Azure/azure-sdk-for-go/sdk/azidentity/cache"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
)

// DefaultLogin logs in using the authentication.json file and the default browser.
func DefaultLogin() error {
	//Read authentication.json
	file, err := os.Open("authentication.json")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	identity := azidentity.AuthenticationRecord{}
	err = decoder.Decode(&identity)
	if err != nil {
		fmt.Println(err)
	}
	identity.TenantID = ""
	credentials, err := azidentity.NewInteractiveBrowserCredential(&azidentity.InteractiveBrowserCredentialOptions{
		TokenCachePersistenceOptions: &azidentity.TokenCachePersistenceOptions{},
		AuthenticationRecord:         identity,
	})
	if err != nil {
		fmt.Println(err)
	}

	clientFactory, err := armsubscriptions.NewClient(credentials, nil)
	if err != nil {
		fmt.Println(err)
	}
	pager, err := clientFactory.Get(context.Background(), "", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pager.DisplayName)

	return nil
}

// FirstTimeLogin logs in for the first time and saves the authentication.json file
func FirstTimeLogin() error {
	credentials, err := azidentity.NewInteractiveBrowserCredential(&azidentity.InteractiveBrowserCredentialOptions{
		TokenCachePersistenceOptions: &azidentity.TokenCachePersistenceOptions{},
	})
	if err != nil {
		fmt.Println(err)
	}

	identity, err := credentials.Authenticate(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
	}

	jsonData, err := json.MarshalIndent(identity, "", "    ")
	if err != nil {
		return err
	}

	// Create or open the JSON file for writing
	outputFile, err := os.Create("authentication.json")
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Write the JSON data to the file
	_, err = outputFile.Write(jsonData)
	if err != nil {
		return err
	}

	fmt.Println("Initial login successful. Authentication file created.")

	return nil
}

// CheckIfAuthenticationFileExists checks if authentication.json exists
func CheckIfAuthenticationFileExists() bool {
	if _, err := os.Stat("authentication.json"); os.IsNotExist(err) {
		return false // File does not exist
	} else if err != nil {
		return false
	} else {
		return true // File exists
	}
}
