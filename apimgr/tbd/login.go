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

// loginCmd represents the login command
// var loginCmd = &cobra.Command{
// 	Use:   "login",
// 	Short: "Login to API Manager",
// 	Long: `login to API Manager by providing
// 	hostname, post, username and password

// 	apimanager login
// 	`,
// 	Run: login,
// }

// func init() {
// 	rootCmd.AddCommand(loginCmd)
// }

// func completer(d prompt.Document) []prompt.Suggest {
// 	s := []prompt.Suggest{
// 		{Text: "host", Description: "Store API Manager Host"},
// 		{Text: "port", Description: "Store API Manager Port"},
// 		{Text: "username", Description: "Store the API Manager username"},
// 		{Text: "password", Description: "Store the API Manager password"},
// 	}
// 	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
// }

// func login(cmd *cobra.Command, args []string) {

// 	fmt.Print("\nAPI Manager Hostname")
// 	host := prompt.Input(": ", completer)
// 	fmt.Print("API Manager Port")
// 	port := prompt.Input(": ", completer)
// 	fmt.Print("Username")
// 	username := prompt.Input(": ", completer)
// 	fmt.Print("Password")
// 	password := prompt.Input(": ", completer)

// 	conf := configAPI{}

// 	data := []byte(username + ":" + password)
// 	basicAuth := base64.StdEncoding.EncodeToString(data)

// 	conf.APIManagerHost = host
// 	conf.APIManagerPort = port
// 	conf.Authorization = basicAuth

// 	out, err := yaml.Marshal(conf)
// 	if err != nil {
// 		fmt.Println("Error to marshal config yaml", err)
// 		return
// 	}
// 	home := os.Getenv("HOME")
// 	err = ioutil.WriteFile(home+"/.apimanager.yaml", out, 0644)
// 	if err != nil {
// 		fmt.Println("Error to write config yaml file", err)

// 	}
// }
