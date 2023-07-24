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
	"fmt"
	"path"

	"github.com/Mrpye/maui-gen/lib"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// target_createCmd represents the create command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set setting values",
	Long:  `Set setting values`,
}

func setCmd_Edit_Config() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "edit",
		Short: "This command will open the config",
		Long:  `This command will open the config`,
		RunE: func(cmd *cobra.Command, args []string) error {

			fmt.Println(viper.ConfigFileUsed())
			open.RunWith(viper.ConfigFileUsed(), "Notepad.exe")
			/*err := open.Run(viper.ConfigFileUsed())
			if err != nil {

			}*/
			return nil
		},
	}
	return cmd
}

func setCmd_Template_Path() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "template [path]",
		Short: "This command will set the default template path",
		Long:  `This command will set the default template path`,
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) > 0 {
				if lib.DirExists(args[0]) && lib.FileExists(path.Join(args[0], "templates.yaml")) {
					viper.Set("templates", args[0])
					viper.WriteConfig()
				} else {
					return fmt.Errorf("the path is not a valid template directory, make sure you are pointing to a directory where templates exist")
				}
			} else {
				return fmt.Errorf("no path is specified. Please specify a valid path to the templates folder")
			}

			return nil
		},
	}
	return cmd
}

func setCmd_Target_Path() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "output [path]",
		Short: "This command will set the default output path",
		Long:  `This command will set the default output path`,
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) > 0 {
				viper.Set("output", args[0])
				viper.WriteConfig()
			} else {
				return fmt.Errorf("no path is specified. Please specify a valid path to the templates folder")
			}

			return nil
		},
	}
	return cmd
}

func setCmd_Schema_Path() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "schema [file]",
		Short: "This command will set the default schema file",
		Long:  `This command will set the default schema file`,
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) > 0 {
				if lib.FileExists(args[0]) {
					viper.Set("schema", args[0])
					viper.WriteConfig()
				} else {
					return fmt.Errorf("the path does not exist exist")
				}
			} else {
				return fmt.Errorf("no path is specified. Please specify a valid path to the schema file")
			}

			return nil
		},
	}
	return cmd
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.AddCommand(setCmd_Template_Path())
	setCmd.AddCommand(setCmd_Target_Path())
	setCmd.AddCommand(setCmd_Schema_Path())
	setCmd.AddCommand(setCmd_Edit_Config())

}
