global using CommunityToolkit.Maui;
global using CommunityToolkit.Mvvm.ComponentModel;
global using  CommunityToolkit.Maui.Views;
global using CommunityToolkit.Mvvm.Input;
global using System.Collections.ObjectModel;
global using System.ComponentModel;
global using System.Diagnostics;
global using System.Globalization;
global using System.Runtime.CompilerServices;
global using System.Text.Json;
global using System.Threading.Tasks;
global using SQLiteNetExtensions.Attributes;
global using SQLiteNetExtensionsAsync.Extensions;
global using System.ComponentModel.DataAnnotations;
{{range $key, $value := .NameSpaces}}global using {{$.RootNS}}.{{$value}};
{{end}}





