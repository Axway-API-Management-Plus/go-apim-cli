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
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

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
		Run: login,
	}

	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create an API Manager resource from a file",
		Long: `Create an API Manager resource from a file. JSON format is accepted. 

For example:

  # Create an organization using the data in org.json
  apimanager create org -f ./org.json

  # Create an application  using the data in org.json
  apimanager create app -f ./app.json
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
	  apimanager list org 
	
	  # Create all applications  
	  apimanager get app 
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
	  apimanager delete org -n Avengers
	
	  # Dlete an application  using the name
	  apimanager delete app -n Asgard
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
	  apimanager describe org -n Avengers
	
	  # Dlete an application  using the name
	  apimanager describe app -n Asgard
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
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(describeCmd)
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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

	fmt.Print("\nAPI Manager Hostname")
	host := prompt.Input(": ", completer)
	fmt.Print("API Manager Port")
	port := prompt.Input(": ", completer)
	fmt.Print("Username")
	username := prompt.Input(": ", completer)
	fmt.Print("Password")
	password := prompt.Input(": ", completer)

	conf := configAPI{}

	data := []byte(username + ":" + password)
	basicAuth := base64.StdEncoding.EncodeToString(data)

	conf.APIManagerHost = host
	conf.APIManagerPort = port
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
