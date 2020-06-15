/*
Copyright © 2020 Axway, Inc. <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"context"
	"fmt"

	"github.com/antihax/optional"
	"github.com/skckadiyala/apimanager/apimgr"
	"github.com/skckadiyala/kubecrt-vms/utils"
	"github.com/spf13/cobra"
)

// apiKeyCmd represents the apiKey command
var (
	keyCmd = &cobra.Command{
		Use:     "key",
		Aliases: []string{"apikeys"},
		Short:   "Create an APIKey",
		Long: `Create an APIKey for an application.

	For example:

	  # Create an apiKey in an application
	  apimanager create apiKey -o <appname>

	  apimanager create key -a <appname>
	`,
		Run: createAPIKey,
	}

	keyListCmd = &cobra.Command{
		Use:     "keys",
		Aliases: []string{"apikeys"},
		Short:   "List all apiKeys",
		Long: `List all apiKeys from an application. 
	
	For example:
	
	# list all the applications 
	apimanager list keys -a <appName> 
	`,
		Run: listAPIKeys,
	}

	keyDelCmd = &cobra.Command{
		Use:     "key",
		Aliases: []string{"apikeys"},
		Short:   "Delete an apiKey",
		Long: `Delete an apiKey from an application. 
	
	For example:
	
	# Delete an apikey from the application 
	apimanager delete key -a <appName> -k <keyID>
	`,
		Run: deleteAPIKey,
	}
)

func init() {

	createCmd.AddCommand(keyCmd)
	listCmd.AddCommand(keyListCmd)
	deleteCmd.AddCommand(keyDelCmd)

	keyCmd.Flags().StringVarP(&appName, "appName", "a", "", "The name to store application name")
	keyCmd.MarkFlagRequired("appName")

	keyListCmd.Flags().StringVarP(&appName, "appName", "a", "", "The name to store application name")
	keyListCmd.MarkFlagRequired("appName")

	keyDelCmd.Flags().StringVarP(&appName, "appName", "a", "", "The name to store application name")
	keyDelCmd.MarkFlagRequired("appName")
	keyDelCmd.Flags().StringVarP(&keyID, "keyID", "k", "", "The keyID to delete")
	keyDelCmd.MarkFlagRequired("keyID")
}

func createAPIKey(cmd *cobra.Command, args []string) {
	cfg := getConfig()
	name = appName
	appID := getApplicationByName(args)

	apikey := apimgr.ApiKey{}

	apikey.ApplicationId = appID

	apikeyPost := &apimgr.ApplicationsIdApikeysPostOpts{}
	apikeyPost.ApiKey = optional.NewInterface(apikey)

	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	apikey, _, err := client.ApplicationsApi.ApplicationsIdApikeysPost(context.Background(), appID, apikeyPost)
	if err != nil {
		utils.PrettyPrintErr("Error creating apikey: %v ", err)
		return
	}
	utils.PrettyPrintInfo("APIKey %v and Secret %v", apikey.Id, apikey.Secret)
	return
}

func listAPIKeys(cmd *cobra.Command, args []string) {

	cfg := getConfig()
	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)
	stdout := fmtDisplay()

	appID := getApplicationByName(args)

	keys, _, err := client.ApplicationsApi.ApplicationsIdApikeysGet(context.Background(), appID)

	if err != nil {
		utils.PrettyPrintErr("Error listing the apiKeys: %v", err)
		return
	}
	if len(keys) != 0 {
		fmt.Fprintf(stdout, "APIKEY\tSECRET\n")
		for _, key := range keys {
			fmt.Fprintf(stdout, "%v\t%v\n", key.Id, key.Secret)
		}
		fmt.Fprint(stdout)
		stdout.Flush()
	} else {
		utils.PrettyPrintInfo("No apikeys found in the application %v", appName)
		return
	}
}

func deleteAPIKey(cmd *cobra.Command, args []string) {
	cfg := getConfig()
	appID := getApplicationByName(args)

	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	_, err := client.ApplicationsApi.ApplicationsIdApikeysKeyIdDelete(context.Background(), appID, keyID)
	if err != nil {
		utils.PrettyPrintErr("Unable to delete the apikey: %v", err)
		return
	}
	utils.PrettyPrintInfo("APIkey %v deleted from the applocation %v ....", keyID, appName)

	return
}
