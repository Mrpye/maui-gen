package code_gen

import (
	"os"
	"strings"

	"github.com/Mrpye/maui-gen/lib"
	"gopkg.in/yaml.v2"
)

type RootSchema struct {
	ProjectNameSpace string   `json:"name_space" yaml:"name_space"`
	Resources        []string `json:"resources" yaml:"resources"`
	Schemas          []Schema `json:"schemas" yaml:"schemas"`
}

type Schema struct {
	Name      string   `json:"name" yaml:"name"`
	Templates []string `json:"templates" yaml:"templates"`
	Flags     []Flag   `json:"flags" yaml:"flags"`
	Fields    []Field  `json:"fields" yaml:"fields"`
	ReadOnly  []Field  `json:"read_only" yaml:"read_only"`
}

func (m *Schema) GetFieldsByFlag(flag string) []Field {
	var tmp_field []Field
	for _, s := range m.Fields {
		for _, f := range s.Flags {
			if strings.EqualFold(f.GetFlag(), flag) {
				tmp_field = append(tmp_field, s)
			}
		}
	}
	return tmp_field
}

func (m *Schema) FlagExists(flag string) bool {
	for _, f := range m.Flags {
		if f.Name != "" {
			parts := strings.Split(f.Name, ":")
			if len(parts) > 0 {
				if strings.EqualFold(parts[0], flag) {
					return true
				}
			}
		}
	}
	return false
}

func (m *Schema) NoFlag(flag string) bool {
	return !m.FlagExists(flag)
}

func (m *Schema) OptExists(flag string) bool {
	for _, f := range m.Flags {
		parts := strings.Split(f.Name, ":")
		if len(parts) > 0 {
			if strings.EqualFold(parts[0], flag) {
				if f.Options != nil {
					return true
				}
			}
		}
	}
	return false
}

func (m *Schema) GetOpt(flag string) interface{} {
	for _, f := range m.Flags {
		parts := strings.Split(f.Name, ":")
		if len(parts) > 0 {
			if strings.EqualFold(parts[0], flag) {
				return f.Options
			}
		}
	}
	return ""
}

func (m *Schema) GetFlagParam(flag string, param_index int) string {
	for _, f := range m.Flags {
		if f.Name != "" {
			//"DetailsView_EXCLUDE:1,5,3"
			parts := strings.Split(f.Name, ":")
			if strings.EqualFold(parts[0], flag) {
				if len(parts) > 1 {
					params := strings.Split(parts[1], ",")
					if param_index > len(params) {
						return ""
					}
					return params[param_index]
				}
			}
		}
	}
	return ""
}

func (m *Schema) GetConstructors() []Field {
	var tmp_field []Field
	for _, s := range m.Fields {
		if s.Constructor {
			tmp_field = append(tmp_field, s)
		}
	}
	return tmp_field
}
func (m *Schema) GetPrimaryKeys() []Field {
	var tmp_field []Field
	for _, s := range m.Fields {
		if s.PrimaryKey {
			tmp_field = append(tmp_field, s)
		}
	}
	return tmp_field
}

// special chars replaced with _
func (m *Schema) SafeName() string {
	return lib.SafeName(m.Name)
}

func (m *Schema) PrivVarName() string {
	return lib.PrivVarName(m.Name)
}

func (m *Schema) PubVarName() string {
	return lib.PubVarName(m.Name)
}

func (m *Schema) FuncName() string {
	return lib.FuncName(m.Name)
}

func (m *Schema) DisplayName() string {
	return lib.DisplayName(m.Name)
}

func LoadScheme(scheme_path string) (*RootSchema, error) {
	//****************
	//Load the package
	//****************
	var m RootSchema
	file, _ := os.ReadFile(scheme_path)
	err := yaml.Unmarshal(file, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
