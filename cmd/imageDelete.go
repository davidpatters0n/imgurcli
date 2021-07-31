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
	"log"
	"os"

	"github.com/davidPatters0n/imgurcli/api"
	"github.com/spf13/cobra"
)

var imageDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an image given an image ID",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		deleted, err := api.Delete(args[0])

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(deleted)
	},
}

func init() {
	imageCmd.AddCommand(imageDeleteCmd)
}
