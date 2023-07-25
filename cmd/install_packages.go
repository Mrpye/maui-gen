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

	"github.com/Mrpye/maui-gen/code_gen"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// target_createCmd represents the create command
var nugetCmd = &cobra.Command{
	Use:   "nuget",
	Short: "Set setting values",
	Long:  `Set setting values`,
}

func nugetCmd_Install_Packages() *cobra.Command {
	var output_path string
	var cmd = &cobra.Command{
		Use:   "install",
		Short: "This command will open the config",
		Long:  `This command will open the config`,
		RunE: func(cmd *cobra.Command, args []string) error {

			if output_path == "" {
				output_path = viper.GetString("output")
				if output_path == "" {
					return fmt.Errorf("no output directory set or passed as an flag -o --output")
				}
			}

			err := code_gen.InstallNugetPackage(output_path, "SQLitePCLRaw.bundle_green")
			if err != nil {
				return err
			}

			err = code_gen.InstallNugetPackage(output_path, "SQLitePCLRaw.provider.sqlite3")
			if err != nil {
				return err
			}

			err = code_gen.InstallNugetPackage(output_path, "SQLitePCLRaw.provider.dynamic_cdecl")
			if err != nil {
				return err
			}

			err = code_gen.InstallNugetPackage(output_path, "SQLiteNetExtensions")
			if err != nil {
				return err
			}

			err = code_gen.InstallNugetPackage(output_path, "SQLiteNetExtensions.Async")
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&output_path, "output", "o", "", "template path")
	return cmd
}

func init() {
	rootCmd.AddCommand(nugetCmd)
	nugetCmd.AddCommand(nugetCmd_Install_Packages())
}
