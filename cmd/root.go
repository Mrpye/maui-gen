/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Mrpye/maui-gen/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"
)

func GenerateDoc() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "gen_docs",
		Short: "This command will build the documents for the cli",
		Long: `
Description:
This command will build the documents for the cli

Example Command:
- maui-gen gen_docs
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if !lib.DirExists("./documents") {
				os.MkdirAll("./documents", os.ModePerm)
			}
			err := doc.GenMarkdownTree(rootCmd, "./documents")
			if err != nil {
				log.Fatal(err)
			}
			lib.ActionLogOK("Documents Generated", '-')
			//lib.PrintlnOK("Documents Generated")
			return nil
		},
	}
	return cmd
}

func rootCmd_About() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "about",
		Short: "This commands with give information about the maui-gen",
		Long:  `This commands with give information about the maui-gen`,
		RunE: func(cmd *cobra.Command, args []string) error {

			lib.ActionLog("About", '-')
			fmt.Println("Author: Andrew Pye")
			fmt.Println("Version: 0.1.1")
			fmt.Println("License: Apache")
			return nil
		},
	}
	return cmd
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "maui-gen",
	Short: "Maui-Gen is a CLI tool to help build .Net Maui data forms from a schema",
	Long: `Maui-Gen is a CLI tool to help build .Net Maui data forms from a schema.
	Using template and a data schema you can configure required fields and Maui-Gen will
	generate the code for you, saving time and effort.
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		lib.PrintlnFail(err.Error())
		lib.ActionLogFail("Failed", '*')
		return
	}
	lib.ActionLogOK("Completed", '*')
}

func init() {
	SetupConfig()
	rootCmd.AddCommand(rootCmd_About())
	rootCmd.AddCommand(GenerateDoc())

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.maui-gen.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func SetupConfig() {
	//viper.SetConfigFile("config")
	/*home, err := os.UserHomeDir()
	if err == nil {
		viper.AddConfigPath(home)
	}*/
	viper.SetConfigName("config")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AddConfigPath(".")
	viper.SetConfigFile("./config.json")

	viper.SetDefault("output", "./output")
	viper.SetDefault("templates", "./templates")
	viper.SetDefault("schema", "")
	viper.ReadInConfig()

}
