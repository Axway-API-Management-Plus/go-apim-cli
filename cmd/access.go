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

	"github.com/antihax/optional"
	"github.com/skckadiyala/apimanager/apimgr"
	"github.com/skckadiyala/kubecrt-vms/utils"
	"github.com/spf13/cobra"
)

// accessCmd represents the access command
var (
	accessCmd = &cobra.Command{
		Use:   "access",
		Short: "request for an application to access an API",
		Long: `Request for an application to access an API. For example:

		apimctl create access -a Avengers -p ProxyName`,
		Run: reqAppAPIAccess,
	}
	accessListCmd = &cobra.Command{
		Use:   "access",
		Short: "Get the list of APIs that the application can access",
		Long: `Get the list of APIs that the application can access. For example:

		apimctl list access -a Avengers`,
		Run: listAppAPIAccess,
	}

	accessDelCmd = &cobra.Command{
		Use:   "access",
		Short: "removes API access of proxy from application",
		Long: `Removes API access of proxy from application. For example:

		apimctl delete access -a Avengers -n ProxyName`,
		Run: deleteAppAPIAccess,
	}
)

func init() {
	createCmd.AddCommand(accessCmd)
	listCmd.AddCommand(accessListCmd)
	deleteCmd.AddCommand(accessDelCmd)

	accessCmd.Flags().StringVarP(&appName, "appName", "a", "", "The name to store application name")
	accessCmd.MarkFlagRequired("appName")
	accessCmd.Flags().StringVarP(&name, "name", "n", "", "The name to store proxy name")
	accessCmd.MarkFlagRequired("name")

	accessListCmd.Flags().StringVarP(&appName, "appName", "a", "", "The name to store application name")
	accessCmd.MarkFlagRequired("appName")

	accessDelCmd.Flags().StringVarP(&appName, "appName", "a", "", "The name to store application name")
	accessDelCmd.MarkFlagRequired("appName")
	accessDelCmd.Flags().StringVarP(&name, "name", "n", "", "The name to store proxy name")
	accessDelCmd.MarkFlagRequired("name")
}

func reqAppAPIAccess(cmd *cobra.Command, args []string) {
	cfg := getConfig()
	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	proxy, err := getProxyByName(args, cfg)
	if err != nil {
		utils.PrettyPrintErr("unable to find the proxy : %v", err)
		return
	}
	appID, err := getApplicationByName(args)
	if err != nil {
		utils.PrettyPrintErr("unable to find the application : %v", err)
		return
	}

	reqBody := apimgr.ApiAccess{}
	reqBody.ApiId = proxy.Id
	reqBody.Enabled = true

	optVars := &apimgr.ApplicationsIdApisPostOpts{}
	optVars.Body = optional.NewInterface(reqBody)

	apiAccess, _, err := client.ApplicationsApi.ApplicationsIdApisPost(context.Background(), appID, optVars)
	if err != nil {
		utils.PrettyPrintErr("Error creating access to apis %v", err)
	}
	utils.PrettyPrintInfo("API Access grandted to Proxy %v with ID %v", proxy.Name, apiAccess.Id)
}

func listAppAPIAccess(cmd *cobra.Command, args []string) {
	cfg := getConfig()
	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)
	stdout := fmtDisplay()

	appID, err := getApplicationByName(args)
	if err != nil {
		utils.PrettyPrintErr("unable to find the application : %v", err)
		return
	}
	apiAccess, _, err := client.ApplicationsApi.ApplicationsIdApisGet(context.Background(), appID)
	if err != nil {
		utils.PrettyPrintErr("Error creating access to apis %v", err)
	}

	if len(apiAccess) != 0 {
		fmt.Fprintf(stdout, "ID\tNAME\tENABLED\tSTATE\n")
		for _, access := range apiAccess {
			proxy, _, _ := client.APIProxyRegistrationApi.ProxiesIdGet(context.Background(), access.ApiId)
			fmt.Fprintf(stdout, "%v\t%v\t%v\t%v\n", access.Id, proxy.Name, access.Enabled, access.State)
		}
		fmt.Fprint(stdout)
		stdout.Flush()
	} else {
		utils.PrettyPrintInfo("No API Access found ")
		return
	}
}

func deleteAppAPIAccess(cmd *cobra.Command, args []string) {
	cfg := getConfig()
	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	appID, err := getApplicationByName(args)
	if err != nil {
		utils.PrettyPrintErr("unable to find the application : %v", err)
		return
	}

	apiAccess, _, err := client.ApplicationsApi.ApplicationsIdApisGet(context.Background(), appID)
	if err != nil {
		utils.PrettyPrintErr("Error creating access to apis %v", err)
	}
	if len(apiAccess) != 0 {
		for _, access := range apiAccess {
			proxy, _, _ := client.APIProxyRegistrationApi.ProxiesIdGet(context.Background(), access.ApiId)
			if proxy.Name == name {
				_, _, err = client.ApplicationsApi.ApplicationsIdApisApiAccessIdDelete(context.Background(), appID, access.Id)
				if err != nil {
					utils.PrettyPrintErr("Error delete access to apis %v", err)
				}
				utils.PrettyPrintInfo("access for  proxy %v  removed from application %v", proxy.Name, appName)
				return
			}
		}
	} else {
		utils.PrettyPrintInfo("No API Access found ")
		return
	}

}
