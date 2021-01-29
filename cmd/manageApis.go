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
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Axway-API-Management-Plus/go-apim-cli/apimgr"
	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// createCmd represents the create command
var (
	loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Stores the login info of API Manager",
		Long: `Stores the login info of API Manager. Stores
	Hostname, 
	Port, 
	Username 
	Password
`,
		PreRun: func(cmd *cobra.Command, args []string) {
			if name != "" {
				// cmd.MarkFlagRequired("name")
				cmd.MarkFlagRequired("port")
				cmd.MarkFlagRequired("user")
				cmd.MarkFlagRequired("password")
			}
		},

		Run: login,
	}

	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create an API Manager resource from a file",
		Long: `Create an API Manager resource from a file. JSON format is accepted. 

For example:

  # Create an organization using the data in org.json
  apimctl create org -f ./org.json

  # Create an application  using the data in org.json
  apimctl create app -f ./app.json
	`,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Println("create called")
		// },
	}

	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update an API Manager resource",
		Long: `Update an API Manager resource. 

For example:

  # Create an organization using the data in org.json

  # Create an application  using the data in org.json
  apimctl update app -f ./app.json
	`,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Println("create called")
		// },
	}

	listCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"get"},
		Short:   "list all API Manager resource",
		Long: `list all API Manager resource by name. 
	
	For example:
	
	  # list all organizations
	  apimctl list org 
	
	  # Create all applications  
	  apimctl get app 
		`,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Println("list called")
		// },
	}
	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "delete an API Manager resource",
		Long: `delete an API Manager resource by name. 
	
	For example:
	
	  # Delete an organization using the name
	  apimctl delete org -n Avengers
	
	  # Dlete an application  using the name
	  apimctl delete app -n Asgard
		`,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Println("delete called")
		// },
	}
	describeCmd = &cobra.Command{
		Use:     "describe",
		Aliases: []string{"desc"},
		Short:   "describe an API Manager resource",
		Long: `describe an API Manager resource by name. 
	
	For example:
	
	  # Delete an organization using the name
	  apimctl describe org -n Avengers
	
	  # Dlete an application  using the name
	  apimctl describe app -n Asgard
		`,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Println("delete called")
		// },
	}
	editCmd = &cobra.Command{
		Use:     "edit",
		Aliases: []string{"edit"},
		Short:   "edit an API Manager resource",
		Long: `edit an API Manager resource by name. 
	
	For example:
	
	  # Edit an organization using the name
	  apimanager describe org -n Avengers
	
	  # Edit an application  using the name
	  apimanager describe app -n Asgard
		`,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Println("delete called")
		// },
	}
)

func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(describeCmd)
	rootCmd.AddCommand(editCmd)

	loginCmd.Flags().StringVarP(&name, "hostname", "H", "", "The name to store API Manager Hostname")
	loginCmd.Flags().StringVarP(&keyID, "port", "P", "", "The name to store API Manager Port")
	loginCmd.Flags().StringVarP(&userName, "user", "u", "", "The name to store API Manager Username")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "The name to store API Manager password")
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "host", Description: "Store API Manager Host"},
		{Text: "port", Description: "Store API Manager Port"},
		{Text: "username", Description: "Store the API Manager username"},
		{Text: "password", Description: "Store the API Manager password"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func login(cmd *cobra.Command, args []string) {

	if name == "" {
		fmt.Print("\nAPI Manager Hostname")
		name = prompt.Input(": ", completer)
		fmt.Print("API Manager Port")
		keyID = prompt.Input(": ", completer)
		fmt.Print("Username")
		userName = prompt.Input(": ", completer)
		fmt.Print("Password")
		password = prompt.Input(": ", completer)
	}

	conf := configAPI{}

	data := []byte(userName + ":" + password)
	basicAuth := base64.StdEncoding.EncodeToString(data)

	getSysInfo(basicAuth, name, keyID)

	conf.APIManagerHost = name
	conf.APIManagerPort = keyID
	conf.Authorization = basicAuth

	out, err := yaml.Marshal(conf)
	if err != nil {
		fmt.Println("Error to marshal config yaml", err)
		return
	}
	home := os.Getenv("HOME")
	err = ioutil.WriteFile(home+"/.apimanager.yaml", out, 0644)
	if err != nil {
		fmt.Println("Error to write config yaml file", err)
	}
}

func getSysInfo(basicAuth, apiHost, apiPort string) {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	cfg := apimgr.NewConfiguration()
	cfg.Host = apiHost + ":" + apiPort
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Basic "+basicAuth)
	cfg.HTTPClient = &http.Client{Transport: transCfg}

	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	_, _, err := client.APIManagerServicesApi.SysconfigGet(context.Background())
	if err != nil {
		fmt.Println("Error logging in to APIManager: ", err)
		os.Exit(0)
	}
	fmt.Println("Login Succeeded")
}
