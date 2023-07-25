package code_gen

import (
	"fmt"
	"os"
	"path"

	"github.com/Mrpye/maui-gen/examples"
	"github.com/Mrpye/maui-gen/lib"
	"github.com/Mrpye/maui-gen/templates"
	"github.com/spf13/viper"
)

func CopyTemplatesToHome() error {
	//**************
	//Build the path
	//**************
	home_dir := lib.UserHomeDir()
	maui_path := path.Join(home_dir, ".maui-gen")
	maui_template_path := path.Join(maui_path, "templates")

	//******************************
	//Create the templates directory
	//******************************
	if !lib.DirExists(maui_template_path) {
		err := os.MkdirAll(maui_template_path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating path %s", maui_path)
		}
	}

	//*********************************************
	//Check the the templates exists and copy files
	//*********************************************
	err := lib.SaveEmbeddedFile(&templates.FS, ".", maui_template_path)
	if err != nil {
		return err
	}

	//****************************************************
	//Update the config file with the path to the template
	//****************************************************
	viper.Set("templates", maui_template_path)
	viper.WriteConfig()

	//***********************
	//Show the templates path
	//***********************
	fmt.Printf("Templates directory: %s\n", maui_template_path)

	return nil
}

func CopyExamplesToHome() error {
	//**************
	//Build the path
	//**************
	home_dir := lib.UserHomeDir()
	maui_path := path.Join(home_dir, ".maui-gen")
	maui_examples_path := path.Join(maui_path, "examples")

	//****************************
	//Create the example directory
	//****************************
	if !lib.DirExists(maui_examples_path) {
		err := os.MkdirAll(maui_examples_path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating path %s", maui_path)
		}
	}

	//********************************************
	//Check the the examples exists and copy files
	//********************************************
	err := lib.SaveEmbeddedFile(&examples.FS, ".", maui_examples_path)
	if err != nil {
		return err
	}

	//*********************
	//Show the example path
	//*********************
	fmt.Printf("Examples directory: %s\n", maui_examples_path)

	return nil
}
