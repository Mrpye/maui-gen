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

	"github.com/Mrpye/maui-gen/code_gen"
	"github.com/Mrpye/maui-gen/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func buildCmd_Build() *cobra.Command {
	var template_path string
	var output_path string
	var schema_path string
	var name_space string = ""
	var cmd = &cobra.Command{
		Use:   "build",
		Short: "This command will build the output",
		Long:  `This command will build the output`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if template_path == "" {
				template_path = viper.GetString("templates")
				if template_path == "" {
					template_path = "./templates"
				}
			}
			if !lib.DirExists(template_path) || !lib.FileExists(path.Join(template_path, "templates.yaml")) {
				return fmt.Errorf("the path %s is not a valid template directory, make sure you are pointing to a directory where templates exist", template_path)
			}

			if output_path == "" {
				output_path = viper.GetString("output")
				if output_path == "" {
					return fmt.Errorf("no output directory set or passed as an flag -o --output")
				}
			}

			if schema_path == "" {
				schema_path = viper.GetString("schema")
				if schema_path == "" {
					return fmt.Errorf("no schema file set or passed as an flag -s --schema")
				}
			}

			if !lib.FileExists(schema_path) {
				return fmt.Errorf("the schema %s does not exist", schema_path)
			}

			//***********************************
			//create out instance of the code gen
			//***********************************
			maui_gen, err := code_gen.Create(output_path)
			if err != nil {
				return err
			}
			//*****************
			//Generate the code
			//*****************
			err = maui_gen.Build(template_path, schema_path, name_space)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&template_path, "templates", "t", "", "template path")
	cmd.Flags().StringVarP(&output_path, "output", "o", "", "template path")
	cmd.Flags().StringVarP(&schema_path, "schema", "s", "", "schema file")
	cmd.Flags().StringVarP(&name_space, "namespace", "n", "", "name space")
	return cmd
}

func init() {
	rootCmd.AddCommand(buildCmd_Build())

}
