package code_gen

import (
	"strings"
)

type Data struct {
	RootNS            string
	Templates         *[]Template
	Template          *Template
	RootSchema        *RootSchema
	Schema            *Schema
	NameSpaces        map[string]string
	RegisterSingleton map[string]string
	RegisterRoute     map[string]string
	Data              Field
}

func (m *Data) GetTemplates(template_name string) []Template {
	result := []Template{}
	for _, t := range *m.Templates {
		if strings.EqualFold(t.Name, template_name) {
			result = append(result, t)
		}
	}
	return result
}

func (m *Data) AddNamespace(value string) {
	if m.NameSpaces == nil {
		m.NameSpaces = make(map[string]string)
	}
	if value == "" {
		return
	}
	if _, exists := m.NameSpaces[value]; !exists {
		m.NameSpaces[value] = value
	}
}
func (m *Data) AddSingleton(value string) {
	if m.RegisterSingleton == nil {
		m.RegisterSingleton = make(map[string]string)
	}
	if value == "" {
		return
	}
	if _, exists := m.RegisterSingleton[value]; !exists {
		m.RegisterSingleton[value] = value
	}
}
func (m *Data) AddRoute(value string) {
	if m.RegisterRoute == nil {
		m.RegisterRoute = make(map[string]string)
	}
	if value == "" {
		return
	}
	if _, exists := m.RegisterRoute[value]; !exists {
		m.RegisterRoute[value] = value
	}
}
func (m *Data) FullNS() string {
	parts := strings.Split(m.RootNS, ".")
	parts = append(parts, m.Template.NameSpace)
	return strings.Join(parts, ".")
}
