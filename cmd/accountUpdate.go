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

// accountUpdate represents the accountChangeSettings command
var accountUpdate = &cobra.Command{
	Use:   "update",
	Short: "Update account settings for the currently authenticated user",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		accountUpdated := api.UpdateSettings(cmd)
		if accountUpdated {
			fmt.Println("Successfully updated account")
		} else {
			fmt.Println("Failed to update account")
		}
	},
}

func init() {
	accountCmd.AddCommand(accountUpdate)
	accountUpdate.Flags().StringP("bio", "", "", "Update the biography for your profile.")
	accountUpdate.Flags().StringP("public-images", "", "public", "Set the images to private or public by default.")
	accountUpdate.Flags().StringP("album-privacy", "", "secret", "public | hidden | secret - Sets the default privacy level of albums")
	accountUpdate.Flags().StringP("username", "", "", "A valid Imgur username (between 4 and 63 alphanumeric characters)")
	accountUpdate.Flags().BoolP("messaging", "", false, "Enable private messaging")
	accountUpdate.Flags().BoolP("show-mature", "", false, "Toggle display of mature images in gallery list endpoints")
	accountUpdate.Flags().BoolP("newsletter-subscribed", "", false, "Toggle subscription to email newsletter.")
}
