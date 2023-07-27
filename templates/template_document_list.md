# **Templates**

|Template Group Name|Template File|Description|
|--------------|-----|-----|{{range  $t := .Templates}}
|{{$t.Name}}|[{{path_base $t.TemplateFile}}]({{$t.Name}}_{{path_base $t.TemplateFile}}.md)|{{$t.Description}}|{{end}}

