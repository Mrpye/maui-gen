package code_gen

type Template struct {
	Name              string   `json:"name" yaml:"name"`
	TemplateFile      string   `json:"template_file" yaml:"template_file"`
	Output            string   `json:"output" yaml:"output"`
	NameSpace         string   `json:"name_space" yaml:"name_space"`
	ProcessTemplate   bool     `json:"process_template" yaml:"process_template"`
	SchemaTemplate    bool     `json:"schema_template" yaml:"schema_template"`
	DependsOn         []string `json:"depends_on" yaml:"depends_on"`
	InjectTag         string   `json:"inject_tag" yaml:"inject_tag"`
	Backup            bool     `json:"backup" yaml:"backup"`
	RegisterSingleton string   `json:"register_singleton" yaml:"register_singleton"`
	RegisterRoute     string   `json:"register_route" yaml:"register_route"`
}
