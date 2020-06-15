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
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/antihax/optional"
	"github.com/skckadiyala/apimanager/apimgr"
	"github.com/skckadiyala/kubecrt-vms/utils"
	"github.com/spf13/cobra"
)

// proxyCmd represents the proxy command
var (
	proxyCmd = &cobra.Command{
		Use:   "proxy",
		Short: "Create Proxy",
		Long: `Create Proxy from a file or with params. 

For example:

# Create proxy using the data 
apimanager create proxy -f proxy.json
apimanager create proxy -n <name> -a 'Backend API' -o <orgName> -s passthrough`,
		PreRun: func(cmd *cobra.Command, args []string) {
			if security != "passthrough" {
				cmd.MarkFlagRequired("appName")
			}
		},
		Run: createProxy,
	}
	proxyList = &cobra.Command{
		Use:   "proxies",
		Short: "List all proxies",
		Long: `list all proxies from a file or with params. 

For example:

# list all proxy 
apimanager list proxies `,
		Run: listProxies,
	}
	proxyDelete = &cobra.Command{
		Use:   "proxy",
		Short: "Delete a proxy",
		Long: `Delete a proxy by name. 

For example:

# Delete a proxy 
apimanager delete proxy -n <ProxyName> `,
		Run: deleteProxy,
	}

	proxyDescribe = &cobra.Command{
		Use:   "proxy",
		Short: "Describe a proxy",
		Long: `Describe a proxy by name. 

For example:

# Describe a proxy 
apimanager describe proxy -n <ProxyName> `,
		Run: describeProxy,
	}
)

func init() {
	createCmd.AddCommand(proxyCmd)
	listCmd.AddCommand(proxyList)
	deleteCmd.AddCommand(proxyDelete)
	describeCmd.AddCommand(proxyDescribe)

	proxyDelete.Flags().StringVarP(&name, "name", "n", "", "proxy name")
	proxyDelete.MarkFlagRequired("name")

	proxyDescribe.Flags().StringVarP(&name, "name", "n", "", "proxy name")
	proxyDescribe.MarkFlagRequired("name")

	proxyCmd.Flags().StringVarP(&file, "file", "f", "", "The filename of the swagger api to be stored")

	proxyCmd.Flags().StringVarP(&name, "name", "n", "", "proxy name")
	proxyCmd.MarkFlagRequired("name")
	proxyCmd.Flags().StringVarP(&orgName, "orgName", "o", "", "organization name")
	proxyCmd.MarkFlagRequired("orgName")
	proxyCmd.Flags().StringVarP(&apiName, "apiName", "b", "", "backend API name")
	proxyCmd.MarkFlagRequired("apiName")

	proxyCmd.Flags().StringVarP(&appName, "appName", "a", "", "application name")
	proxyCmd.Flags().StringVarP(&security, "security", "s", "passthrough", "provide the security to use for proxy: \napikey \nhttpbasic \noauth \npassthrough")
	proxyCmd.Flags().StringVarP(&resourcePath, "resourcePath", "r", "", "provide the resource path for the proxy")
	proxyCmd.Flags().StringVarP(&certPath, "certPath", "c", "", "provide the location of the backend api cert")

	proxyCmd.Flags().StringVarP(&proxyVersion, "proxyVersion", "v", "1.0", "provide the proxy version")
	proxyCmd.Flags().StringVarP(&proxyState, "proxyState", "p", "published", "provide the proxy state")

}

func createProxy(cmd *cobra.Command, args []string) {
	cfg := getConfig()
	proxyBody := apimgr.VirtualizedApi{}
	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	if resourcePath == "" {
		fmt.Fprintln(os.Stderr, "Resource path is empty,so adding a random path ")
		resourcePath = "/api/" + getUniqueID(5) + "/v1"
	}
	if certPath != "" {
		certs, _ := importCerts(cfg, certPath)
		proxyBody.CaCerts = certs
	}

	apiID := getAPIByName(args)

	orgID := getOrganizationByName(args)

	pro, err := getProxyByName(args, cfg)
	if err == nil {
		utils.PrettyPrintInfo("Proxy with name %v already exists", pro.Name)
		return
	}

	switch security {
	case "passthrough":
		proxyBody.SecurityProfiles = getSecurityProfilePassThrough()
	case "apikey":
		proxyBody.SecurityProfiles = getSecurityProfileAPIKey()
	case "httpbasic":
		proxyBody.SecurityProfiles = getSecurityProfileHTTPBasic()
	case "oauth":
		proxyBody.SecurityProfiles = getSecurityProfileOAuth()
	default:
		fmt.Fprintln(os.Stderr, "Invalid security data - allowed security name: passthrough,apikey,oauth, httpbasic")
		return
	}

	proxyBody.Path = resourcePath //"/bank/v1"
	proxyBody.ApiId = apiID
	proxyBody.OrganizationId = orgID
	proxyBody.Name = name
	proxyBody.Version = proxyVersion
	proxyBody.State = proxyState

	proxy, _, err := client.APIProxyRegistrationApi.ProxiesPost(context.Background(), proxyBody)
	if err != nil {
		utils.PrettyPrintErr("Error creating proxy :%v", err)
		return
	}
	utils.PrettyPrintInfo("Proxy %v created", proxy.Name)
	if appName != "" {
		appID := getApplicationByName(args)
		reqApplicationAPIAccess(appID, proxy.Id, cfg)
	}
	return
}

func importCerts(cfg *apimgr.Configuration, certpath string) ([]apimgr.CaCert, error) {
	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	cerVar := apimgr.CertinfoPostOpts{}

	cerVar.Inbound = optional.NewBool(false)
	cerVar.Outbound = optional.NewBool(true)

	file, err := os.Open(certpath)
	if err != nil {
		utils.PrettyPrintErr("Unable to open the file: %v", err)
	}

	cerVar.File = optional.NewInterface(file)

	certs, _, err := client.APIManagerServicesApi.CertinfoPost(context.Background(), &cerVar)
	if err != nil {
		utils.PrettyPrintErr("Error creating the cert: %v", err)
	}
	return certs, nil
}

func getProxyByName(args []string, cfg *apimgr.Configuration) (apimgr.VirtualizedApi, error) {
	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	proxy := apimgr.VirtualizedApi{}
	getProxyVars := &apimgr.ProxiesGetOpts{}

	getProxyVars.Field = optional.NewInterface("name")
	getProxyVars.Op = optional.NewInterface("eq")
	getProxyVars.Value = optional.NewInterface(name)

	proxies, _, err := client.APIProxyRegistrationApi.ProxiesGet(context.Background(), getProxyVars)
	if err != nil {
		utils.PrettyPrintErr("Error getting the Proxy: %v", err)
		return proxy, err
	}
	if len(proxies) != 0 {
		// utils.PrettyPrintInfo("Proxy found: %v", proxies[0].Name)
		return proxies[0], nil
	}
	// utils.PrettyPrintInfo("Proxy %v not found ", name)
	return proxy, errors.New("Proxy not found")
}

func listProxies(cmd *cobra.Command, args []string) {
	cfg := getConfig()
	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)
	getProxyVars := &apimgr.ProxiesGetOpts{}

	stdout := fmtDisplay()
	proxies, _, err := client.APIProxyRegistrationApi.ProxiesGet(context.Background(), getProxyVars)
	if err != nil {
		utils.PrettyPrintErr("Error listing the proxies: %v", err)
		return
	}
	if len(proxies) != 0 {
		fmt.Fprintf(stdout, "ID\tNAME\tORGANIZATION\tPATH\tSTATE\tVERSION\n")
		for _, proxy := range proxies {
			org, _, _ := client.OrganizationsApi.OrganizationsIdGet(context.Background(), proxy.OrganizationId)
			fmt.Fprintf(stdout, "%v\t%v\t%v\t%v\t%v\t%v\n", proxy.Id, proxy.Name, org.Name, proxy.Path, proxy.State, proxy.Version)
		}
		fmt.Fprint(stdout)
		stdout.Flush()
		return
	}
	utils.PrettyPrintInfo("No Proxy found ")
	return
}

func deleteProxy(cmd *cobra.Command, args []string) {
	cfg := getConfig()

	proxy, err := getProxyByName(args, cfg)
	if err != nil {
		utils.PrettyPrintErr("unable to find the proxy : %v", err)
		return
	}
	if proxy.State == "published" {
		fmt.Printf("Unable to Delete, Proxy %v is in published state \n", proxy.Name)
		return
	}

	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	_, err = client.APIProxyRegistrationApi.ProxiesIdDelete(context.Background(), proxy.Id)
	if err != nil {
		utils.PrettyPrintErr("Unable to delete the Proxy: %v", err)
		return
	}
	utils.PrettyPrintInfo("Proxy %v Deleted", name)
	return
}

func describeProxy(cmd *cobra.Command, args []string) {
	cfg := getConfig()

	proxy, err := getProxyByName(args, cfg)
	if err != nil {
		utils.PrettyPrintErr("unable to find the proxy : %v", err)
		return
	}
	// if proxy.State == "published" {
	// 	fmt.Printf("Unable to Delete, Proxy %v is in published state \n", proxy.Name)
	// 	return
	// }

	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	proxy, _, err = client.APIProxyRegistrationApi.ProxiesIdGet(context.Background(), proxy.Id)
	if err != nil {
		utils.PrettyPrintErr("Unable to get the Proxy: %v", err)
		return
	}
	prettyJSON, err := json.MarshalIndent(proxy, "", "    ")
	if err != nil {
		log.Fatal("Failed to marshal to json", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))
	return
}
