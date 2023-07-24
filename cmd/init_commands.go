/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"github.com/Mrpye/maui-gen/code_gen"
	"github.com/spf13/cobra"
)

// target_createCmd represents the create command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init systems",
	Long:  `init systems`,
}

func initCmd_Templates() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "templates",
		Short: "This command will make a copy of the templates to the home directory and use them as source templates",
		Long:  `This command will make a copy of the templates to the home directory and use them as source templates`,
		RunE: func(cmd *cobra.Command, args []string) error {

			//***********************************
			//Copy the Template to home directory
			//***********************************
			return code_gen.CopyTemplatesToHome()
		},
	}
	return cmd
}

func initCmd_Examples() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "examples",
		Short: "This command will make a copy of the templates to the home directory and use them as source templates",
		Long:  `This command will make a copy of the templates to the home directory and use them as source templates`,
		RunE: func(cmd *cobra.Command, args []string) error {

			//***********************************
			//Copy the Template to home directory
			//***********************************
			return code_gen.CopyExamplesToHome()
		},
	}
	return cmd
}

func initCmd_All() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "all",
		Short: "This command will make a copy of the templates to the home directory and use them as source templates",
		Long:  `This command will make a copy of the templates to the home directory and use them as source templates`,
		RunE: func(cmd *cobra.Command, args []string) error {

			//***********************************
			//Copy the Template to home directory
			//***********************************
			err := code_gen.CopyTemplatesToHome()
			if err != nil {
				return err
			}
			return code_gen.CopyExamplesToHome()

		},
	}
	return cmd
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.AddCommand(initCmd_Templates())
	initCmd.AddCommand(initCmd_Examples())
	initCmd.AddCommand(initCmd_All())

}
