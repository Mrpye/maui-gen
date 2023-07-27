package code_gen

import (
	"fmt"
	"path"

	"github.com/Mrpye/maui-gen/lib"
)

type Template struct {
	Name              string       `json:"name" yaml:"name"`
	Description       string       `json:"description" yaml:"description"`
	TemplateFile      string       `json:"template_file" yaml:"template_file"`
	Output            string       `json:"output" yaml:"output"`
	NameSpace         string       `json:"name_space" yaml:"name_space"`
	ProcessTemplate   bool         `json:"process_template" yaml:"process_template"`
	SchemaTemplate    bool         `json:"schema_template" yaml:"schema_template"`
	DependsOn         []string     `json:"depends_on" yaml:"depends_on"`
	InjectTag         string       `json:"inject_tag" yaml:"inject_tag"`
	Backup            bool         `json:"backup" yaml:"backup"`
	RegisterSingleton string       `json:"register_singleton" yaml:"register_singleton"`
	RegisterRoute     string       `json:"register_route" yaml:"register_route"`
	Flags             []FlagSchema `json:"flags" yaml:"flags"`
}

type FlagSchema struct {
	FieldFlag       bool              `json:"field_flag" yaml:"field_flag"`
	Name            string            `json:"name" yaml:"name"`
	Description     string            `json:"description" yaml:"description"`
	Params          []FlagParamSchema `json:"params" yaml:"params"`
	Example         string            `json:"example" yaml:"example"`
	OptionsRequired bool              `json:"options_required" yaml:"options_required"`
}

type FlagParamSchema struct {
	Name        string                 `json:"name" yaml:"name"`
	Description string                 `json:"description" yaml:"description"`
	DataType    string                 `json:"data_type" yaml:"data_type"`
	ParamsEnums []FlagParamEnumsSchema `json:"param_enums" yaml:"param_enums"`
	Options     string                 `json:"options" yaml:"options"`
}

type FlagParamEnumsSchema struct {
	Value       string `json:"value" yaml:"value"`
	Description string `json:"description" yaml:"description"`
}

// *****************************************
// Function to build documents for templates
// *****************************************
func (m *CodeGen) BuildTemplateDocs(template_path string, template_doc string, template_doc_menu string, output_dir string) error {

	m.LoadTemplates(template_path)
	fmt.Printf("Building Template %s\n", template_doc)

	for _, t := range m.Templates {
		d := &Data{Template: &t, Templates: &m.Templates}

		//*****************
		//Read the template
		//*****************
		template, err := lib.ReadFileToString(template_doc)
		if err != nil {
			return err
		}

		//***************************
		//Parse the template document
		//***************************
		doc, err := m.parseTemplate(template, d)
		if err != nil {
			return err
		}

		//*****************
		//Save the template
		//*****************
		template_file := path.Base(t.TemplateFile)
		//template_file = strings.ReplaceAll(template_file, ".", "_")
		save_path := path.Join(output_dir, t.Name+"_"+template_file+".md")
		lib.SaveStringToFile(save_path, doc)

	}
	fmt.Printf("Building Template %s\n", template_doc_menu)
	//***************************
	//Parse the template menu
	//***************************
	d := &Data{Templates: &m.Templates}
	//*****************
	//Read the template
	//*****************
	template, err := lib.ReadFileToString(template_doc_menu)
	if err != nil {
		return err
	}

	//***************************
	//Parse the template document
	//***************************
	doc, err := m.parseTemplate(template, d)
	if err != nil {
		return err
	}

	//*****************
	//Save the template
	//*****************

	save_path := path.Join(output_dir, "Templates.md")
	lib.SaveStringToFile(save_path, doc)
	return nil
}
