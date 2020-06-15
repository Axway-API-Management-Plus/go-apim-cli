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
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/antihax/optional"
	"github.com/skckadiyala/apimanager/apimgr"
	"github.com/skckadiyala/kubecrt-vms/utils"
	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var (
	apiCmd = &cobra.Command{
		Use:   "api",
		Short: "Create an api",
		Long: `Create an api from a swagger file. 

For example:

# Create an api using the data in swagger.json
apimanager create api -n <name> -f swagger.json -o <orgName> `,
		Run: createBackendAPI,
	}

	apiListCmd = &cobra.Command{
		Use:   "apis",
		Short: "List all backend apis",
		Long: `List all backend apis from a swagger file. 

For example:

# List all apis using the data 
apimanager list apis `,
		Run: listBackendAPI,
	}

	apiDelCmd = &cobra.Command{
		Use:   "api",
		Short: "Delete an API",
		Long: `Delete an API from a swagger file. 

For example:ß

# Delete an api using the data 
apimanager delete api -n <name> `,
		Run: deleteAPI,
	}

	apiDescCmd = &cobra.Command{
		Use:   "api",
		Short: "Describe an API",
		Long: `Describe an API from a swagger file. 

For example:ß

# Describe an api using the data 
apimanager describe api -n <name> `,
		Run: describeAPI,
	}
)

func init() {
	createCmd.AddCommand(apiCmd)
	listCmd.AddCommand(apiListCmd)
	deleteCmd.AddCommand(apiDelCmd)
	describeCmd.AddCommand(apiDescCmd)

	apiCmd.Flags().StringVarP(&file, "swagger", "f", "", "The filename of the swagger api to be stored")
	apiCmd.MarkFlagRequired("swagger")
	apiCmd.Flags().StringVarP(&apiName, "name", "n", "", "The name to store API name")
	apiCmd.MarkFlagRequired("name")
	apiCmd.Flags().StringVarP(&orgName, "orgName", "o", "", "The name to store Organization name")
	apiCmd.MarkFlagRequired("orgName")

	apiDelCmd.Flags().StringVarP(&apiName, "name", "n", "", "The name to store API name")
	apiDelCmd.MarkFlagRequired("name")

	apiDescCmd.Flags().StringVarP(&apiName, "name", "n", "", "The name to store API name")
	apiDescCmd.MarkFlagRequired("name")
}

func createBackendAPI(cmd *cobra.Command, args []string) {
	// utils.PrettyPrintInfo("Creating backend API")

	cfg := getConfig()
	orgID := getOrganizationByName(args)

	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	file, err := os.Open(file)
	if err != nil {
		utils.PrettyPrintErr("Error Opening file: %v", err)
	}
	beAPI, _, err := client.APIRepositoryApi.ApirepoImportPost(context.Background(), orgID, apiName, "swagger", file)
	if err != nil {
		utils.PrettyPrintErr("Error Creating Backend API: %v", err)
		return
	}
	utils.PrettyPrintInfo("Backend API %v with ID: %v created", beAPI.Name, beAPI.Id)
	return
}

func listBackendAPI(cmd *cobra.Command, args []string) {
	cfg := getConfig()
	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)
	apiGetOpts := &apimgr.ApirepoGetOpts{}
	stdout := fmtDisplay()

	apis, _, err := client.APIRepositoryApi.ApirepoGet(context.Background(), apiGetOpts)
	if err != nil {
		utils.PrettyPrintErr("Error Creating Backend API: %v", err)
		return
	}

	if len(apis) != 0 {
		fmt.Fprintf(stdout, "ID\tNAME\tORGANIZATION\tBACKEND URL\tVERSION\n")
		for _, api := range apis {
			org, _, _ := client.OrganizationsApi.OrganizationsIdGet(context.Background(), api.OrganizationId)
			fmt.Fprintf(stdout, "%v\t%v\t%v\t%v\t%v\n", api.Id, api.Name, org.Name, api.BasePath+api.ResourcePath, api.Version)
		}
		fmt.Fprint(stdout)
		stdout.Flush()
	} else {
		utils.PrettyPrintInfo("No backend api's found ")
		return
	}
}

func getAPIByName(args []string) string {
	cfg := getConfig()

	// utils.PrettyPrintInfo("Finding API %v ....", apiName)

	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	apiGetOpts := &apimgr.ApirepoGetOpts{}

	apiGetOpts.Field = optional.NewInterface("name")
	apiGetOpts.Op = optional.NewInterface("eq")
	apiGetOpts.Value = optional.NewInterface(apiName)

	apis, _, err := client.APIRepositoryApi.ApirepoGet(context.Background(), apiGetOpts)
	if err != nil {
		utils.PrettyPrintErr("Error finding backend API: %v", err)
		os.Exit(0)
	}
	if len(apis) != 0 {
		// utils.PrettyPrintInfo("Backend API found: %v", apis[0].Name)
		return apis[0].Id
	}
	utils.PrettyPrintInfo("Backend API %v not found ", apiName)
	os.Exit(0)
	return apis[0].Id
}

func deleteAPI(cmd *cobra.Command, args []string) {
	cfg := getConfig()

	apiID := getAPIByName(args)

	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	_, err := client.APIRepositoryApi.ApirepoIdDelete(context.Background(), apiID)
	if err != nil {
		utils.PrettyPrintErr("Unable to delete the backend API: %v", err)
		return
	}
	utils.PrettyPrintInfo("Backend API %v Deleted", apiName)
	return
}

func describeAPI(cmd *cobra.Command, args []string) {
	cfg := getConfig()
	apiID := getAPIByName(args)

	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	api, _, err := client.APIRepositoryApi.ApirepoIdGet(context.Background(), apiID)
	if err != nil {
		utils.PrettyPrintErr("Unable to delete the backend API: %v", err)
		return
	}

	prettyJSON, err := json.MarshalIndent(api, "", "    ")
	if err != nil {
		log.Fatal("Failed to marshal to json", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))
	return
}
