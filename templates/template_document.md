# **{{.Template.Name}}** Template

### Main Menu
[Template List](Templates.md)


## Parameters

|Field Name|Yaml field|Value|
|--------------|-----|-----|
|Name|name|{{na .Template.Name}}|
|Description|description|{{na .Template.Description}}|
|TemplateFile|template_file|{{na .Template.TemplateFile}}|
|Output|output|{{na .Template.Output}}|
|NameSpace|name_space|{{na .Template.NameSpace}}|
|ProcessTemplate|description|{{.Template.ProcessTemplate}}|
|SchemaTemplate|schema_template|{{.Template.SchemaTemplate}}|
|InjectTag|inject_tag|{{na .Template.InjectTag}}|
|RegisterSingleton|register_singleton|{{na .Template.RegisterSingleton}}|
|RegisterRoute|register_route|{{na .Template.RegisterRoute}}|
|Backup|backup|{{.Template.Backup}}|
{{if gt (len .Template.DependsOn) 0}}|DependsOn|depends_on|[DependsOn](#depends-on)|
{{end}}{{if gt (len .Template.Flags) 0}}|Flags|flags|[Flags](#flags)|
{{end}}

{{if gt (len .Template.DependsOn) 0}}## Depends On
{{if gt (len .Template.DependsOn) 0}}
|Template Group Name|Template File|
|--------------|-----|{{range  $field := .Template.DependsOn}}{{range  $dep := (get_templates $field)}}
|{{$field}}|[{{path_base $dep.TemplateFile}}]({{$field}}_{{path_base $dep.TemplateFile}}.md)|{{end}}{{end}}
{{end}}{{end}}

{{if gt (len .Template.Flags) 0}}
## Flags
{{if gt (len .Template.Flags) 0}}
{{range $idx, $field := .Template.Flags}}

<details>
<summary>{{$field.Name}}</summary>

|Template Field Name|Yaml field|Value|
|--------------|-----|-----|
|Name|name|{{$field.Name}}|
|Description|description|{{$field.Description}}|
|FieldFlag|field_flag|{{$field.FieldFlag}}|
|OptionsRequired|options_required|{{$field.OptionsRequired}}|
{{if gt (len $field.Params) 0}}|Params|params|[params](#{{lc $field.Name}}-parameters)|

---

### **{{$field.Name}}** Parameters: 

{{range $idx, $p := $field.Params}}
#### Parameter: **{{$p.Name}}**

|Template Field Name|Yaml field|Value|
|--------------|-----|-----|
|Name|name|{{$p.Name}}|
|Description|description|{{$p.Description}}|
|DataType|data_type|{{$p.DataType}}|
---

{{end}}{{end}}
### Example:
```yaml
{{$field.Example}}
```
</details>
{{end}}
{{end}}{{end}}

### Main Menu
[Template List](Templates.md)