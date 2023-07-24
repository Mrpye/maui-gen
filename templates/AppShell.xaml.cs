//<%REGISTER_ROUTE%>
//{%Injected Values Start%}
{{range $key, $value := .RegisterRoute}}Routing.RegisterRoute(nameof({{$value}}), typeof({{$value}}));
{{end}}
//{%Injected Values End%}
