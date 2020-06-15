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
	"encoding/base64"
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

// userCmd represents the user command
var (
	userCmd = &cobra.Command{
		Use:   "user",
		Short: "Create User",
		Long: `Create user from a file. JSON format is accepted. 
	
	For example:
	
	# Create an organization using the data in org.json
	apimanager create user -f user.json
	`,
		PreRun: func(cmd *cobra.Command, args []string) {
			if file == "" {
				cmd.MarkFlagRequired("name")
				cmd.MarkFlagRequired("loginName")
				cmd.MarkFlagRequired("role")
			}
		},
		Run: createUser,
	}

	userDelCmd = &cobra.Command{
		Use:   "user",
		Short: "Delete User",
		Long: `Delete user from a file. JSON format is accepted. 
	
	For example:
	
	# delete user by name
	apimanager delete user -n username
	`,
		Run: deleteUser,
	}

	userDescCmd = &cobra.Command{
		Use:   "user",
		Short: "Describe User",
		Long: `Describe user from a file. JSON format is accepted. 
	
	For example:
	
	# Describe user by name
	apimanager describe user -n username
	`,
		Run: describeUser,
	}

	userEditCmd = &cobra.Command{
		Use:   "user",
		Short: "Edit User",
		Long: `Edit user from a file. JSON format is accepted. 
	
	For example:
	
	# Edit user by name
	apimanager edit user -n username
	`,
		Run: editUser,
	}

	userListCmd = &cobra.Command{
		Use:   "users",
		Short: "List all users",
		Long: `List all users. 
	
	For example:
	
	# list all the users
	apimanager list users 
	`,
		Run: listUsers,
	}
)

func init() {
	createCmd.AddCommand(userCmd)
	deleteCmd.AddCommand(userDelCmd)
	listCmd.AddCommand(userListCmd)
	describeCmd.AddCommand(userDescCmd)
	editCmd.AddCommand(userEditCmd)

	userCmd.Flags().StringVarP(&file, "file", "f", "", "The filename of the raw data to be stored")
	// userCmd.MarkFlagRequired("file")

	userCmd.Flags().StringVarP(&orgName, "orgName", "o", "", "The name to store Organization name")
	userCmd.MarkFlagRequired("orgName")
	userCmd.Flags().StringVarP(&password, "password", "p", "", "change password for the user")
	// userCmd.MarkFlagRequired("password")

	userCmd.Flags().StringVarP(&userName, "name", "n", "", "The name of the username")
	userCmd.Flags().StringVarP(&loginName, "loginName", "l", "", "login name for the user")
	userCmd.Flags().StringVarP(&userRole, "role", "r", "", "role for the user")

	userCmd.Flags().StringVarP(&image, "image", "i", "", "filename of the image to be used")

	userDelCmd.Flags().StringVarP(&userName, "name", "n", "", "The name of the username")
	userDelCmd.MarkFlagRequired("name")

	userDescCmd.Flags().StringVarP(&userName, "name", "n", "", "The name of the username")
	userDescCmd.MarkFlagRequired("name")

	userEditCmd.Flags().StringVarP(&userName, "name", "n", "", "The name of the username")
	userEditCmd.MarkFlagRequired("name")
}

func createUser(cmd *cobra.Command, args []string) {

	cfg := getConfig()
	orgID := getOrganizationByName(args)
	newUser := apimgr.User{}

	if file != "" {
		userBody, err := ioutil.ReadFile(file) // pass the file name with path
		if err != nil {
			fmt.Print(err)
		}
		err = json.Unmarshal([]byte(userBody), &newUser)
		if err != nil {
			utils.PrettyPrintErr("Error unmarshaling user json: %v", err)
		}
		if newUser.Name == "" && userName == "" {
			fmt.Printf("User name is required\n")
			return
		}
		if newUser.Name == "" {
			newUser.Name = userName
		}
		newUser.OrganizationId = orgID
	} else {
		if image != "" {
			bImage, err := ioutil.ReadFile(image) // pass the file name with path
			if err != nil {
				fmt.Print(err)
			}
			newUser.Image = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(bImage)
		}
		newUser.Name = userName
		newUser.Description = userName + " is a " + userRole
		newUser.Phone = "+1 877-564-7700"
		newUser.Email = loginName + "@apimanager.com"
		newUser.LoginName = loginName
		newUser.Role = userRole
		newUser.Enabled = true
		newUser.OrganizationId = orgID
	}

	// utils.PrettyPrintInfo("Creating a new user %v, %v %v ....", newUser.Name, newUser.LoginName, newUser.Email)

	if verifyIfUserExists(newUser.Email, newUser.LoginName, args) {
		utils.PrettyPrintInfo("User %v with login: %v or email: %v already exists....", newUser.Name, newUser.LoginName, newUser.Email)
		return
	}
	if password != "" && len(password) < 6 {
		utils.PrettyPrintInfo("Minimum password length is 6 chars")
	}
	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	userVars := &apimgr.UsersPostOpts{}
	userVars.Body = optional.NewInterface(newUser)

	user, _, err := client.UsersApi.UsersPost(context.Background(), userVars)
	if err != nil {
		utils.PrettyPrintErr("Error creating user:%v", err)
	}
	utils.PrettyPrintInfo("New user %v created", user.Name)
	if password != "" {
		changeUserPassword(user.Id, password, cfg)
	}

	return
}

func getUserByName(args []string) string {

	cfg := getConfig()

	// utils.PrettyPrintInfo("Finding User %v ....", userName)

	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	getUserVars := &apimgr.UsersGetOpts{}

	getUserVars.Field = optional.NewInterface("name")
	getUserVars.Op = optional.NewInterface("eq")
	getUserVars.Value = optional.NewInterface(userName)

	users, _, err := client.UsersApi.UsersGet(context.Background(), getUserVars)
	if err != nil {
		utils.PrettyPrintErr("Error finding the user: %v", err)
		os.Exit(0)
	}
	if len(users) != 0 {
		// utils.PrettyPrintInfo("User found: %v", users[0].Name)
		return users[0].Id
	}
	utils.PrettyPrintInfo("User %v not found ", userName)
	return "User Not Found"

}

func listUsers(cmd *cobra.Command, args []string) {

	cfg := getConfig()
	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	stdout := fmtDisplay()

	getUserVars := &apimgr.UsersGetOpts{}

	users, _, err := client.UsersApi.UsersGet(context.Background(), getUserVars)
	if err != nil {
		utils.PrettyPrintErr("Error listing the users: %v", err)
		return
	}
	if len(users) != 0 {
		fmt.Fprintf(stdout, "ID\tNAME\tLOGIN\tORGANIZATION\tEMAIL\tROLE\n")
		for _, user := range users {
			org, _, _ := client.OrganizationsApi.OrganizationsIdGet(context.Background(), user.OrganizationId)
			fmt.Fprintf(stdout, "%v\t%v\t%v\t%v\t%v\t%v\n", user.Id, user.Name, user.LoginName, org.Name, user.Email, user.Role)
		}
		fmt.Fprint(stdout)
		stdout.Flush()
	} else {
		utils.PrettyPrintInfo("No users found ")
		return
	}
}

func deleteUser(cmd *cobra.Command, args []string) {
	cfg := getConfig()
	userID := getUserByName(args)

	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	_, err := client.UsersApi.UsersIdDelete(context.Background(), userID)
	if err != nil {
		utils.PrettyPrintErr("Unable to delete the user: %v", err)
		return
	}
	utils.PrettyPrintInfo("User %v Deleted", userName)
	return
}

func describeUser(cmd *cobra.Command, args []string) {
	user, err := descUser(cmd, args)
	prettyJSON, err := json.MarshalIndent(user, "", "    ")
	if err != nil {
		log.Fatal("Failed to marshal to json", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))
	return
}

func editUser(cmd *cobra.Command, args []string) {
	cfg := getConfig()

	editedUser := apimgr.User{}

	fname := getUniqueID(5)
	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	user, err := descUser(cmd, args)
	if err != nil {
		return
	}
	prettyJSON, err := json.MarshalIndent(user, "", "    ")
	if err != nil {
		log.Fatal("Failed to marshal org to json", err)
		return
	}

	err = createTempFile(fname, prettyJSON)
	if err != nil {
		return
	}

	editedObj, _ := ioutil.ReadFile("/tmp/" + fname)
	err = json.Unmarshal(editedObj, &editedUser)
	if err != nil {
		fmt.Println("Invalid json object, no changes applied")
		return
	}
	userVars := &apimgr.UsersIdPutOpts{}
	userVars.Body = optional.NewInterface(editedUser)

	if editedUser.Id != user.Id || editedUser.CreatedOn != user.CreatedOn {
		fmt.Println("User can't be updated with the changed values")
		return
	}
	if reflect.DeepEqual(user, editedUser) {
		utils.PrettyPrintInfo("User %v has no changes to update", editedUser.Name)
		return
	}
	updatedUser, _, err := client.UsersApi.UsersIdPut(context.Background(), editedUser.Id, userVars)
	if err != nil {
		utils.PrettyPrintErr("Error updating User")
		return
	}
	utils.PrettyPrintInfo("User %v updated with the valid changes", updatedUser.Name)
	deleteTempFile(fname)
	return

}

func verifyIfUserExists(uEmail, lName string, args []string) bool {

	cfg := getConfig()

	// utils.PrettyPrintInfo("Verify User %v ....", lName)

	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	getUserVars := &apimgr.UsersGetOpts{}

	getUserVars.Field = optional.NewInterface("loginName")
	getUserVars.Op = optional.NewInterface("eq")
	getUserVars.Value = optional.NewInterface(lName)

	getUserVars.Field = optional.NewInterface("email")
	getUserVars.Op = optional.NewInterface("eq")
	getUserVars.Value = optional.NewInterface(uEmail)

	users, _, err := client.UsersApi.UsersGet(context.Background(), getUserVars)
	if err != nil {
		utils.PrettyPrintErr("Error finding the user: %v", err)
		os.Exit(0)
	}
	if len(users) != 0 {
		// utils.PrettyPrintInfo("User found: %v", users[0].Name)
		return true
	}
	// utils.PrettyPrintInfo("User %v not found ", userName)
	return false

}

func changeUserPassword(userID, newPassword string, cfg *apimgr.Configuration) {
	// utils.PrettyPrintInfo("Change password for the UserId :%v", userID)
	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	_, err := client.UsersApi.UsersIdChangepasswordPost(context.Background(), userID, newPassword)
	if err != nil {
		utils.PrettyPrintErr("Error updating password :%v", err)
		return
	}
	utils.PrettyPrintInfo("Password updated")
	return
}

func descUser(cmd *cobra.Command, args []string) (apimgr.User, error) {
	cfg := getConfig()
	userID := getUserByName(args)

	client := &apimgr.APIClient{}
	client = apimgr.NewAPIClient(cfg)

	user, _, err := client.UsersApi.UsersIdGet(context.Background(), userID)
	if err != nil {
		utils.PrettyPrintErr("Unable to get the user: %v", err)
		return apimgr.User{}, err
	}
	return user, nil
}
