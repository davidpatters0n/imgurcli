/*
Copyright © 2021 DJ Patterson

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
	"fmt"

	"github.com/davidPatters0n/imgurcli/api"
	"github.com/spf13/cobra"
)

// sendVerifyCmd represents the sendVerifyCmd command
var sendVerifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Sends a verification email to the currently authenticated user",
	Run: func(cmd *cobra.Command, args []string) {
		sentEmail := api.SendVerificationEmail()

		if sentEmail {
			fmt.Println("Email verification sent")
		}
	},
}

// accountVerifyCmd represents the accountVerify command
var accountVerifyCmd = &cobra.Command{
	Use:   "verified",
	Short: "Checks to see if the currently authenticated user has verified their email address",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(api.AccountVerified())
	},
}

func init() {
	accountCmd.AddCommand(sendVerifyCmd)
	accountCmd.AddCommand(accountVerifyCmd)
}
