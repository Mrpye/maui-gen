using SQLite;
namespace {{.FullNS}} {
    public partial class {{.Schema.FuncName}}Model: ObservableValidator, ICloneable, IEntity {

        public {{.Schema.FuncName}}Model() { }

        public {{.Schema.FuncName}}Model({{range $idx, $field := .Schema.GetConstructors}}{{if $idx}},{{end}}{{$field.Type}} {{$field.PrivVarName}}{{end}}) {
            
            {{range $idx, $field := .Schema.GetConstructors}}this.{{$field.PubVarName}} = {{$field.SafeName}};
            {{end}}
        }

        #region Readonly Fields 
        {{range $field := .Schema.ReadOnly}}public {{$field.Type}} {{$field.PubVarName}}{{$field.FormatValue `=>` false}};
        {{end}}#endregion
        
        #region Fields 
        {{range $field := .Schema.Fields}}{{if $field.FlagExists `Model_Template_EXCLUDE`}}{{else}}[ObservableProperty]
        {{if $field.PrimaryKey}}[property: PrimaryKey, AutoIncrement]
        {{end}}{{if $field.FlagExists `RELATIONSHIP`}}[property: OneToOne("{{func_name ($field.GetFlagParam `RELATIONSHIP` 0) }}_model_id", "{{$field.PubVarName}}", CascadeOperations = CascadeOperation.All, ReadOnly = false)]
        {{end}}{{if $field.FlagExists `VALIDATE_Required`}}[Required]
        {{end}}{{if $field.FlagExists `VALIDATE_MinLength`}}[MinLength({{$field.GetFlagParam `VALIDATE_MinLength` 0}})]
        {{end}}{{if $field.FlagExists `VALIDATE_MaxLength`}}[System.ComponentModel.DataAnnotations.MaxLength({{$field.GetFlagParam `VALIDATE_MaxLength` 0}})]
        {{end}}{{range $notify := $field.NotifyProperty}}[NotifyPropertyChangedFor(nameof({{pub_var_name $notify }}))]
        {{end}}{{if $field.NoFlag `RELATIONSHIP`}}private {{$field.Type}} {{$field.PrivVarName}}{{if eq $field.InitValueRaw true}}{{$field.InitValue}};{{else}}{{$field.FormatValue `=` true}};{{end}}
        {{else}}private {{func_name ($field.GetFlagParam `RELATIONSHIP` 0) }}Model {{$field.PrivVarName}}{{if eq $field.InitValueRaw true}}{{$field.InitValue}};{{else}}{{$field.FormatValue `=` true}};{{end}}
        
        [ObservableProperty]
        [property: ForeignKey(typeof({{func_name ($field.GetFlagParam `RELATIONSHIP` 0) }}Model))]
        private int {{priv_var_name ($field.GetFlagParam `RELATIONSHIP` 0) }}_model_id;
        {{end}}
        {{end}}
        {{end}}
        #endregion

        #region Validation
        [ObservableProperty]
        [property: Ignore]
        string errors;

        public Dictionary<string, string> Validate() {
            ValidateAllProperties();
            if (HasErrors)
                Errors = string.Join(Environment.NewLine, GetErrors().Select(e => e.ErrorMessage));
            else
                Errors = String.Empty;

            return GetErrors().ToDictionary(k => k.MemberNames.First(), v => v.ErrorMessage) ?? new Dictionary<string, string?>();

        }
        #endregion  Validation
        
        public void Merge({{.Schema.FuncName}}Model to_merge) {
            {{range $idx, $field := .Schema.Fields}}this.{{$field.PubVarName}} = to_merge.{{$field.PubVarName}};
            {{end}}
        }

        public object Clone() {
            return this.MemberwiseClone();
        }
    }
}
