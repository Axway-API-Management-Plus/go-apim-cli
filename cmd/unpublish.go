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

	"github.com/skckadiyala/apimanager/apimgr"
	"github.com/skckadiyala/kubecrt-vms/utils"
	"github.com/spf13/cobra"
)

// unpublishCmd represents the unpublish command
var (
	unpublishCmd = &cobra.Command{
		Use:     "unpublish",
		Aliases: []string{"unpub"},
		Short:   "unpublish a proxy",
		Long: `unpublish the proxy For example:

apimanager unpublish -p <proxy Name>`,
		Run: unpublishProxy,
	}
	publishCmd = &cobra.Command{
		Use:     "publish",
		Aliases: []string{"pub"},
		Short:   "publish a proxy",
		Long: `publish the proxy For example:

apimanager publish -p <proxy Name>`,
		Run: publishProxy,
	}
)

func init() {
	rootCmd.AddCommand(unpublishCmd)
	rootCmd.AddCommand(publishCmd)

	unpublishCmd.Flags().StringVarP(&name, "name", "n", "", "proxy name")
	unpublishCmd.MarkFlagRequired("name")

	publishCmd.Flags().StringVarP(&name, "name", "n", "", "proxy name")
	publishCmd.MarkFlagRequired("name")
}

// (proxyID string, cfg *apimgr.Configuration)
func unpublishProxy(cmd *cobra.Command, args []string) {
	cfg := getConfig()
	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	proxy, err := getProxyByName(args, cfg)
	if err != nil {
		utils.PrettyPrintErr("Proxy %v not found %v", name, err)
	}
	_, _, err = client.APIProxyRegistrationApi.ProxiesIdUnpublishPost(context.Background(), proxy.Id)
	if err != nil {
		utils.PrettyPrintErr("Error Updating the Proxy: %v", err)
		return
	}
	fmt.Printf("Proxy %v unpublished \n", proxy.Name)
	return
}

func publishProxy(cmd *cobra.Command, args []string) {

	cfg := getConfig()
	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	proxy, err := getProxyByName(args, cfg)
	if err != nil {
		utils.PrettyPrintErr("Proxy %v not found %v", name, err)
	}
	_, _, err = client.APIProxyRegistrationApi.ProxiesIdPublishPost(context.Background(), proxy.Id, proxy.Name, "")
	if err != nil {
		utils.PrettyPrintErr("Error Updating the Proxy: %v", err)
		return
	}
	fmt.Printf("Proxy %v published \n", proxy.Name)
	return
}
