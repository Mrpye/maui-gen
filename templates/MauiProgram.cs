//<%REGISTER_SINGLETON%>
//{%Injected Values Start%}
{{range $key, $value := .RegisterSingleton}}builder.Services.AddSingleton<{{$value}}>();
{{end}}
//{%Injected Values End%}