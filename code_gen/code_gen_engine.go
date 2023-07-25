package code_gen

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
	"text/template"

	"github.com/Mrpye/maui-gen/lib"
	"github.com/iancoleman/strcase"
	"gopkg.in/yaml.v2"
)

type CodeGen struct {
	Output_folder      string
	Templates          []Template
	requited_resources map[string]Template
}

func (m *CodeGen) AddRequiredResource(requires Template) {
	if _, exists := m.requited_resources[requires.TemplateFile]; !exists {
		m.requited_resources[requires.TemplateFile] = requires
	}
}

func (m *CodeGen) AddRequiredResources(requires []Template) {
	for _, o := range requires {
		m.AddRequiredResource(o)
	}
}

func (m *CodeGen) GetTemplate(name string) []Template {
	var tmp_templates []Template
	for _, s := range m.Templates {
		if strings.EqualFold(s.Name, name) {
			tmp_templates = append(tmp_templates, s)
		}
	}
	return tmp_templates
}

func (m *CodeGen) getTemplates(template_names []string) ([]Template, []Template) {
	var tmp_templates []Template
	var tmp_resource []Template
	for _, t := range template_names {
		result := m.GetTemplate(t)
		for _, r := range result {
			if r.SchemaTemplate {
				tmp_templates = append(tmp_templates, r)
			} else {
				tmp_resource = append(tmp_resource, r)
			}
			//********************
			//Add the depends upon
			//********************
			dep_templates, dep_resource := m.getTemplates(r.DependsOn)
			tmp_templates = append(tmp_templates, dep_templates...)
			tmp_resource = append(tmp_resource, dep_resource...)
		}
	}
	return tmp_templates, tmp_resource
}

func (m *CodeGen) GetTemplates(template_names []string) []Template {
	tmp_map_templates := make(map[string]Template)
	dep_templates, dep_resource := m.getTemplates(template_names)
	//*****************
	//Remove duplicates
	//*****************
	for _, t := range dep_templates {
		if _, exists := tmp_map_templates[t.TemplateFile]; !exists {
			tmp_map_templates[t.TemplateFile] = t
		}
	}
	//*****************
	//Add the resources
	//*****************
	m.AddRequiredResources(dep_resource)

	//****************
	//Convert to array
	//****************
	var tmp_templates []Template
	for _, v := range tmp_map_templates {
		tmp_templates = append(tmp_templates, v)
	}
	return tmp_templates
}

// ***************************
// Load the template resources
// ***************************
func (m *CodeGen) LoadTemplates(template_path string) error {
	//****************
	//Load the package
	//****************
	var t []Template

	file, _ := os.ReadFile(template_path)
	err := yaml.Unmarshal(file, &t)
	if err != nil {
		return err
	}
	m.Templates = t
	return nil
}

func Create(output_folder string) (*CodeGen, error) {
	cg := &CodeGen{Output_folder: output_folder}
	cg.requited_resources = make(map[string]Template)
	//******************
	//Load the templates
	//******************
	err := cg.LoadTemplates("templates/templates.yaml")
	if err != nil {
		return nil, err
	}
	return cg, nil
}

func (m *CodeGen) Build(template_path string, schema_path string, ns string) error {
	//***************
	//Load the schema
	//***************
	s, err := LoadScheme(schema_path)
	if err != nil {
		return err
	}

	//***************************
	//override the root namespace
	//***************************
	if ns != "" {
		s.ProjectNameSpace = ns
	}

	current_data := &Data{
		RootNS:     s.ProjectNameSpace,
		RootSchema: s,
	}

	//*******************
	//Build the templates
	//*******************
	for _, o := range s.Schemas {
		//***********************************
		//Get the template for the data model
		//***********************************
		template_results := m.GetTemplates(o.Templates)
		for _, t := range template_results {

			fmt.Printf("Processing template %s for schema %s\n", t.TemplateFile, o.Name)
			current_data.AddNamespace(t.NameSpace)
			//*****************
			//Load the template
			//*****************
			template_path := path.Join(template_path, t.TemplateFile)
			tpl, err := lib.ReadFileToString(template_path)
			if err != nil {
				return err
			}
			//*************************************
			//Build the data model for the template
			//*************************************
			current_data.Template = &t
			current_data.Schema = &o

			//*************************
			//Parse the register params
			//*************************
			result_singleton_register, err := parseTemplate(t.RegisterSingleton, current_data)
			if err != nil {
				return err
			}
			current_data.AddSingleton(result_singleton_register)

			result_route_register, err := parseTemplate(t.RegisterRoute, current_data)
			if err != nil {
				return err
			}
			current_data.AddRoute(result_route_register)

			//******************
			//Parse the template
			//******************
			result, err := parseTemplate(tpl, current_data)
			if err != nil {
				return err
			}

			//*********************
			//Parse the output path
			//*********************
			new_path, err := parseTemplate(t.Output, current_data)
			if err != nil {
				return err
			}

			//*************
			//Save the file
			//*************
			output_path := path.Join(m.Output_folder, new_path)
			lib.MakeDir(output_path)
			err = lib.SaveStringToFile(output_path, result)
			if err != nil {
				return err
			}
			lib.PrintOK(fmt.Sprintf("Created file %s from template %s\n", output_path, t.TemplateFile))
		}
	}

	//***********************************
	//Lets add in any requested resources
	//***********************************
	_, resources := m.getTemplates(s.Resources)
	m.AddRequiredResources(resources)

	var res_list []Template
	for _, r := range m.requited_resources {
		res_list = append(res_list, r)
	}
	sort.SliceStable(res_list, func(i, j int) bool {
		return res_list[i].InjectTag == "" && res_list[j].InjectTag != ""
	})

	for _, r := range res_list {
		current_data.AddNamespace(r.NameSpace)
	}
	//*********************
	//Process the Resources
	//*********************
	for _, r := range res_list {
		template_path := path.Join("templates", r.TemplateFile)

		fmt.Printf("Processing template %s \n", r.TemplateFile)

		if r.ProcessTemplate {
			tpl, err := lib.ReadFileToString(template_path)
			if err != nil {
				return err
			}
			current_data.RootNS = s.ProjectNameSpace
			current_data.Template = &r
			current_data.RootSchema = s
			current_data.Schema = nil

			//*************************
			//Parse the register params
			//*************************
			result_singleton_register, err := parseTemplate(r.RegisterSingleton, current_data)
			if err != nil {
				return err
			}
			current_data.AddSingleton(result_singleton_register)

			result_route_register, err := parseTemplate(r.RegisterRoute, current_data)
			if err != nil {
				return err
			}
			current_data.AddRoute(result_route_register)

			//******************
			//Parse the template
			//******************
			result, err := parseTemplate(tpl, current_data)
			if err != nil {
				return err
			}

			//*********************
			//Parse the output path
			//*********************
			new_path, err := parseTemplate(r.Output, current_data)
			if err != nil {
				return err
			}

			//*************
			//Save the file
			//*************
			output_path := path.Join(m.Output_folder, new_path)
			lib.MakeDir(output_path)
			if r.InjectTag == "" {
				result = strings.ReplaceAll(result, "\ufeff", "")
				err = lib.SaveStringToFile(output_path, result)
				if err != nil {
					return err
				}
				lib.PrintOK(fmt.Sprintf("Created file %s from template %s\n", output_path, r.TemplateFile))
			} else {
				if lib.FileExists(output_path) {
					file_data, err := lib.ReadFileToString(output_path)
					if err != nil {
						return err
					}

					//remove value between {%Injected Values Start%} and {%Injected Values End%}
					res := ""
					start := "//{%Injected Values Start%}"
					stop := "//{%Injected Values End%}"
					startIndex := strings.Index(file_data, start)
					stopIndex := strings.Index(file_data, stop)

					if startIndex != stopIndex {
						stopIndex = stopIndex + len(stop)
						res = file_data[:startIndex] + file_data[stopIndex:]
						file_data = strings.ReplaceAll(res, "\n\n", "\n")
						//file_data = file_data[:startIndex] + r.InjectTag + file_data[startIndex:]
					}
					if (startIndex > -1 && stopIndex == -1) || (startIndex == -1 && stopIndex > -1) {
						return fmt.Errorf("cannot find 'Injected Values Start' or 'Injected Values End' in file %s", output_path)
					}
					result = strings.ReplaceAll(result, "\ufeff", "")
					file_data = strings.ReplaceAll(file_data, r.InjectTag, result)

					err = lib.SaveStringToFile(output_path, file_data)
					if err != nil {
						return err
					}
					lib.PrintOK(fmt.Sprintf("Injected data into file %s from template %s\n", output_path, r.TemplateFile))
				}
			}

		} else {
			//*************
			//Copy the file
			//*************
			output_path := path.Join(m.Output_folder, r.Output)
			lib.MakeDir(output_path)
			_, err = lib.CopyFile(template_path, output_path)
			if err != nil {
				return err
			}
			lib.PrintOK(fmt.Sprintf("Copied file %s to location %s\n", template_path, output_path))
		}
	}
	return nil
}

func parseTemplate(tpl string, data *Data) (string, error) {
	//*********************
	//Create a function map
	//*********************
	funcMap := template.FuncMap{
		"base64enc":     lib.Base64EncString,
		"base64dec":     lib.Base64DecString,
		"gzip_base64":   lib.GzipBase64,
		"lc":            strings.ToLower,
		"uc":            strings.ToUpper,
		"domain":        lib.GetDomainOrIP,
		"port_string":   lib.GetPortString,
		"port_int":      lib.GetPortInt,
		"clean":         lib.Clean,
		"concat":        lib.Concat,
		"replace":       strings.ReplaceAll,
		"contains":      lib.StringContainsStringListItem,
		"not":           lib.NOT,
		"or":            lib.OR,
		"and":           lib.AND,
		"plus":          lib.Plus,
		"minus":         lib.Minus,
		"multiply":      lib.Multiply,
		"divide":        lib.Divide,
		"camel":         strcase.ToCamel,
		"pub_var_name":  lib.PubVarName,
		"priv_var_name": lib.PrivVarName,

		"func_name": lib.FuncName,
		"safe":      lib.SafeName,
		"display":   lib.DisplayName,
	}

	//*****************
	//Pase the template
	//*****************
	tmpl, err := template.New("CodeRun").Funcs(funcMap).Parse(tpl)
	if err != nil {
		return "", err
	}

	//**************************************
	//Run the template to verify the output.
	//**************************************
	var tpl_buffer bytes.Buffer
	err = tmpl.Execute(&tpl_buffer, data)
	if err != nil {
		return "", err
	}

	return tpl_buffer.String(), nil
}
