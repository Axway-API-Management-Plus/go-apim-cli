/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"io/ioutil"

	"github.com/skckadiyala/apimanager/apimgr"
	"github.com/skckadiyala/kubecrt-vms/utils"
	"github.com/spf13/cobra"
)

// oauthCmd represents the oauth command
var (
	oauthCmd = &cobra.Command{
		Use:   "oauth",
		Short: "Create an Oauth",
		Long: `Create an Oauth for an application.

For example:

  # Create an oauth in an application 
  apimanager create oauth -a <appname>
`,
		Run: createOAuth,
	}
	oauthListCmd = &cobra.Command{
		Use:   "oauths",
		Short: "List all Oauth Keys",
		Long: `List all Oauth Keys from an application. 

For example:

# list all the applications 
apimanager list oauths -a <appName> 
`,
		Run: listOAuthKeys,
	}
	oauthDelCmd = &cobra.Command{
		Use:   "oauth",
		Short: "Delete an Oauth",
		Long: `Delete an Oauth from an application. 
	
	For example:
	
	# Delete an Oauth from the application 
	apimanager delete Oauth -a <appName> -k <oauthID>
	`,
		Run: deleteOAuthKey,
	}
)

func init() {
	createCmd.AddCommand(oauthCmd)
	listCmd.AddCommand(oauthListCmd)
	deleteCmd.AddCommand(oauthDelCmd)

	oauthCmd.Flags().StringVarP(&appName, "appName", "a", "", "The name to store application name")
	oauthCmd.MarkFlagRequired("appName")

	oauthCmd.Flags().StringVarP(&certPath, "certPath", "c", "", "provide the location of the oauth cert file")

	oauthListCmd.Flags().StringVarP(&appName, "appName", "a", "", "The name to store application name")
	oauthListCmd.MarkFlagRequired("appName")

	oauthDelCmd.Flags().StringVarP(&appName, "appName", "a", "", "The name to store application name")
	oauthDelCmd.MarkFlagRequired("appName")
	oauthDelCmd.Flags().StringVarP(&oauthID, "oauthID", "k", "", "The keyID to delete")
	oauthDelCmd.MarkFlagRequired("oauthID")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// oauthCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// oauthCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createOAuth(cmd *cobra.Command, args []string) {
	cfg := getConfig()
	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	oauthBody := apimgr.OAuthClient{}
	certContent, err := ioutil.ReadFile(certPath)
	if err != nil {
		utils.PrettyPrintErr("Error reading a file %v", err)
	}
	appID := getApplicationByName(args)
	// oauthBody.Id = "oauth_" + app.Name + "_0"
	oauthBody.ApplicationId = appID
	// oauthBody.Secret = "secret_" + app.Name + "_0"
	oauthBody.RedirectUrls = []string{"https://localhost/oauth_callback"}
	oauthBody.Type = "public"
	oauthBody.Cert = string(certContent)

	oauthClient, _, err := client.ApplicationsApi.ApplicationsIdOauthPost(context.Background(), appID, oauthBody)
	if err != nil {
		utils.PrettyPrintErr("Error creating Oauth %v", err)
	}
	utils.PrettyPrintInfo("oauth Id %v  with Secret %v created", oauthClient.Id, oauthClient.Secret)
}

func listOAuthKeys(cmd *cobra.Command, args []string) {
	cfg := getConfig()
	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)
	stdout := fmtDisplay()

	appID := getApplicationByName(args)

	oauths, _, err := client.ApplicationsApi.ApplicationsIdOauthGet(context.Background(), appID)
	if err != nil {
		utils.PrettyPrintErr("Error listing the oauth: %v", err)
		return
	}
	if len(oauths) != 0 {
		fmt.Fprintf(stdout, "OAUTH\tSECRET\n")
		for _, oauth := range oauths {
			fmt.Fprintf(stdout, "%v\t%v\n", oauth.Id, oauth.Secret)
		}
		fmt.Fprint(stdout)
		stdout.Flush()
	} else {
		utils.PrettyPrintInfo("No OAuth Keys found in the application %v", appName)
		return
	}
}

func deleteOAuthKey(cmd *cobra.Command, args []string) {
	cfg := getConfig()
	appID := getApplicationByName(args)

	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	_, err := client.ApplicationsApi.ApplicationsIdOauthOauthIdDelete(context.Background(), appID, keyID)
	if err != nil {
		utils.PrettyPrintErr("Unable to delete the apikey: %v", err)
		return
	}
	utils.PrettyPrintInfo("OAuth key %v deleted from the applocation %v ....", oauthID, appName)
	return
}
