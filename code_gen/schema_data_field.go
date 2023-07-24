package code_gen

import (
	"fmt"
	"strings"

	"github.com/Mrpye/maui-gen/lib"
)

type Flag struct {
	Name    string      `json:"name" yaml:"name"`
	Options interface{} `json:"options" yaml:"options"`
}

func (m *Flag) GetFlag() string {
	//"DetailsView_EXCLUDE:1,5,3"
	parts := strings.Split(m.Name, ":")
	return parts[0]
}

type Field struct {
	PrimaryKey     bool        `json:"primary_key" yaml:"primary_key"`
	Constructor    bool        `json:"constructor" yaml:"constructor"`
	Name           string      `json:"name" yaml:"name"`
	Type           string      `json:"type" yaml:"type"`
	Flags          []Flag      `json:"flags" yaml:"flags"`
	NotifyProperty []string    `json:"notify_property" yaml:"notify_property"`
	InitValue      interface{} `json:"init_value" yaml:"init_value"`
	InitValueRaw   bool        `json:"init_value_raw" yaml:"init_value_raw"`
}

func (m *Field) FlagExists(flag string) bool {
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

func (m *Field) NoFlag(flag string) bool {
	return !m.FlagExists(flag)
}

func (m *Field) OptExists(flag string) bool {
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

func (m *Field) GetOpt(flag string) interface{} {
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

func (m *Field) GetFlagParam(flag string, param_index int) string {
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

func (m *Field) FormatValue(prefix string, quote_value bool) string {
	if m.InitValue == nil {
		return ""
	}
	if quote_value {
		return fmt.Sprintf("%s\"%v\"", prefix, m.InitValue)
	}
	return fmt.Sprintf("%s%v", prefix, m.InitValue)
}

// special chars replaced with _
func (m *Field) SafeName() string {
	return lib.SafeName(m.Name)
}

func (m *Field) PrivVarName() string {
	return lib.PrivVarName(m.Name)
}

func (m *Field) PubVarName() string {
	return lib.PubVarName(m.Name)
}

func (m *Field) FuncName() string {
	return lib.FuncName(m.Name)
}

func (m *Field) DisplayName() string {
	return lib.DisplayName(m.Name)
}
