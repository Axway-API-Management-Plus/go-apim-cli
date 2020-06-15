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
	"io/ioutil"
	"log"
	"os"
	"reflect"

	"github.com/antihax/optional"
	"github.com/skckadiyala/apimanager/apimgr"
	"github.com/skckadiyala/kubecrt-vms/utils"
	"github.com/spf13/cobra"
)

// appCmd represents the app command
var (
	appCmd = &cobra.Command{
		Use:   "app",
		Short: "Create an application",
		Long: `Create an application by name. 

	For example:
	
	  # Create an application by name
	  apimanager create app -n <appname> -o <orgName>
	`,
		PreRun: func(cmd *cobra.Command, args []string) {
			if file == "" {
				cmd.MarkFlagRequired("name")
			}
		},
		Run: createApplication,
	}
	appDelCmd = &cobra.Command{
		Use:   "app",
		Short: "Delete an application",
		Long: `Delete application from a file.  
	
	For example:
	
	# delete application by name
	apimanager delete app -n <appname> 
	`,
		Run: deleteApplication,
	}

	appDescCmd = &cobra.Command{
		Use:   "app",
		Short: "Describe an application",
		Long: `Describe an application. 
	
	For example:
	
	# Describe application by name
	apimanager desc app -n <appname> 
	`,
		Run: describeApplication,
	}

	appEditCmd = &cobra.Command{
		Use:   "app",
		Short: "Edit an application",
		Long: `Edit an application. 
	
	For example:
	
	# Edit application by name
	apimanager edit app -n <appname> 
	`,
		Run: editApplication,
	}

	appListCmd = &cobra.Command{
		Use:   "apps",
		Short: "List all applications",
		Long: `List all applications. 
	
	For example:
	
	# list all the applications 
	apimanager list apps 
	`,
		Run: listApplications,
	}
)

func init() {
	createCmd.AddCommand(appCmd)
	deleteCmd.AddCommand(appDelCmd)
	listCmd.AddCommand(appListCmd)
	describeCmd.AddCommand(appDescCmd)
	editCmd.AddCommand(appEditCmd)

	appCmd.Flags().StringVarP(&file, "file", "f", "", "filename used to create an application resource")

	appCmd.Flags().StringVarP(&orgName, "orgName", "o", "", "The name to store Organization name")
	appCmd.MarkFlagRequired("orgName")
	appCmd.Flags().StringVarP(&appName, "name", "n", "", "The name to store application name")

	appDelCmd.Flags().StringVarP(&appName, "name", "n", "", "The name to store application name")
	appDelCmd.MarkFlagRequired("name")

	appDescCmd.Flags().StringVarP(&appName, "name", "n", "", "The name to store application name")
	appDescCmd.MarkFlagRequired("name")
	appEditCmd.Flags().StringVarP(&appName, "name", "n", "", "The name to store application name")
	appEditCmd.MarkFlagRequired("name")
}

func createApplication(cmd *cobra.Command, args []string) {
	cfg := getConfig()
	orgID := getOrganizationByName(args)
	app := apimgr.ApplicationRequest{}

	if file != "" {
		appBody, err := ioutil.ReadFile(file) // pass the file name with path
		if err != nil {
			fmt.Print(err)
		}
		err = json.Unmarshal([]byte(appBody), &app)
		if err != nil {
			utils.PrettyPrintErr("Error unmarshaling org json: %v", err)
		}
		if app.Name == "" && appName == "" {
			fmt.Printf("Application name is required\n")
			return
		}
		if app.Name == "" {
			app.Name = appName
		}
		app.OrganizationId = orgID
	} else {
		app.Name = appName
		app.Description = appName + " Application"
		app.Phone = "+1 877-564-7700"
		app.Email = appName + "@apimanager.com"
		app.Apis = []string{}
		app.OrganizationId = orgID
	}

	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	appVars := &apimgr.ApplicationsPostOpts{}
	appVars.Body = optional.NewInterface(app)

	appResp, _, err := client.ApplicationsApi.ApplicationsPost(context.Background(), appVars)
	if err != nil {
		utils.PrettyPrintErr("Error creating application: %v", err)
		return
	}
	utils.PrettyPrintInfo("Application %v Created", appResp.Name)
	return
}

func deleteApplication(cmd *cobra.Command, args []string) {
	cfg := getConfig()
	appID := getApplicationByName(args)

	// utils.PrettyPrintInfo("Deleting application %v ....", appName)

	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	_, err := client.ApplicationsApi.ApplicationsIdDelete(context.Background(), appID)
	if err != nil {
		utils.PrettyPrintErr("Unable to delete the application: %v", err)
		return
	}
	utils.PrettyPrintInfo("application %v deleted", appName)
	return
}

func describeApplication(cmd *cobra.Command, args []string) {

	app, err := descApplication(cmd, args)
	prettyJSON, err := json.MarshalIndent(app, "", "    ")
	if err != nil {
		log.Fatal("Failed to marshal to json", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))
	return
}

func editApplication(cmd *cobra.Command, args []string) {
	cfg := getConfig()

	editedApp := apimgr.Application{}

	fname := getUniqueID(5)
	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	app, err := descApplication(cmd, args)
	if err != nil {
		return
	}
	prettyJSON, err := json.MarshalIndent(app, "", "    ")
	if err != nil {
		log.Fatal("Failed to marshal org to json", err)
		return
	}

	err = createTempFile(fname, prettyJSON)
	if err != nil {
		return
	}

	editedObj, _ := ioutil.ReadFile("/tmp/" + fname)
	err = json.Unmarshal(editedObj, &editedApp)
	if err != nil {
		fmt.Println("Invalid json object, no changes applied")
		return
	}
	appVars := &apimgr.ApplicationsIdPutOpts{}
	appVars.Body = optional.NewInterface(editedApp)

	if editedApp.Id != app.Id || editedApp.CreatedOn != app.CreatedOn {
		fmt.Println("Application can't be updated with the changed values")
		return
	}
	// fmt.Println(app)
	// fmt.Println(editedApp)
	if reflect.DeepEqual(app, editedApp) {
		utils.PrettyPrintInfo("Application %v has no changes to update", editedApp.Name)
		return
	}
	updatedApp, _, err := client.ApplicationsApi.ApplicationsIdPut(context.Background(), editedApp.Id, appVars)
	if err != nil {
		utils.PrettyPrintErr("Error updating Application")
		return
	}
	utils.PrettyPrintInfo("Application %v updated with the valid changes", updatedApp.Name)
	deleteTempFile(fname)
	return

}

func getApplicationByName(args []string) string {
	cfg := getConfig()

	// utils.PrettyPrintInfo("Finding application %v ....", appName)

	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	getAppVars := &apimgr.ApplicationsGetOpts{}

	getAppVars.Field = optional.NewInterface("name")
	getAppVars.Op = optional.NewInterface("eq")
	getAppVars.Value = optional.NewInterface(appName)

	apps, _, err := client.ApplicationsApi.ApplicationsGet(context.Background(), getAppVars)
	if err != nil {
		utils.PrettyPrintErr("Error finding the application: %v", err)
		os.Exit(0)
	}
	if len(apps) != 0 {
		// utils.PrettyPrintInfo("application found: %v", apps[0].Name)
		return apps[0].Id
	}
	utils.PrettyPrintInfo("application %v not found ", appName)
	os.Exit(0)
	return apps[0].Id
}

func listApplications(cmd *cobra.Command, args []string) {
	cfg := getConfig()
	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	getAppVars := &apimgr.ApplicationsGetOpts{}

	stdout := fmtDisplay()

	apps, _, err := client.ApplicationsApi.ApplicationsGet(context.Background(), getAppVars)
	if err != nil {
		utils.PrettyPrintErr("Error listing the applications: %v", err)
		return
	}
	if len(apps) != 0 {
		// fmt.Printf("Name \t\t ID \n")
		fmt.Fprintf(stdout, "ID\tNAME\tDESCRIPTION\tORGANIZATION\n")
		for _, app := range apps {
			// fmt.Printf("%v \t %v \n", app.Name, app.Id)
			org, _, _ := client.OrganizationsApi.OrganizationsIdGet(context.Background(), app.OrganizationId)
			fmt.Fprintf(stdout, "%v\t%v\t%vv\t%v\n", app.Id, app.Name, app.Description, org.Name)
		}
		fmt.Fprint(stdout)
		stdout.Flush()
	} else {
		utils.PrettyPrintInfo("No application found ")
		return
	}
}

func reqApplicationAPIAccess(appID, apiID string, cfg *apimgr.Configuration) {

	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	reqBody := apimgr.ApiAccess{}
	reqBody.ApiId = apiID
	reqBody.Enabled = true

	optVars := &apimgr.ApplicationsIdApisPostOpts{}
	optVars.Body = optional.NewInterface(reqBody)

	apiAccess, _, err := client.ApplicationsApi.ApplicationsIdApisPost(context.Background(), appID, optVars)
	if err != nil {
		utils.PrettyPrintErr("Error creating access to apis %v", err)
	}

	utils.PrettyPrintInfo(apiAccess.ApiId)
}

func descApplication(cmd *cobra.Command, args []string) (apimgr.Application, error) {
	cfg := getConfig()
	appID := getApplicationByName(args)

	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	app, _, err := client.ApplicationsApi.ApplicationsIdGet(context.Background(), appID)
	if err != nil {
		utils.PrettyPrintErr("Unable to get the application: %v", err)
		return apimgr.Application{}, err
	}

	return app, nil
}
