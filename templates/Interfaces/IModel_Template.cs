using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace {{.FullNS}} {
    public interface I{{.Schema.FuncName}} {
        #region Fields 
        {{range $field := .Schema.Fields}}{{if $field.FlagExists `Model_Template_EXCLUDE`}}{{else}}{{$field.Type}} {{$field.PubVarName}} { get; set; }
        {{end}}{{end}}
        #endregion Fields

        #region Methods
        public Merge({{.Schema.FuncName}}Model to_merge)
        public object Clone()
        #endregion Methods
    }

}
