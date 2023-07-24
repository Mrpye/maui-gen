package code_gen

import (
	"strings"
)

type Data struct {
	RootNS            string
	Template          *Template
	RootSchema        *RootSchema
	Schema            *Schema
	NameSpaces        map[string]string
	RegisterSingleton map[string]string
	RegisterRoute     map[string]string
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
