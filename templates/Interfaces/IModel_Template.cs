using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace {{.FullNS}} {
    public interface I{{.Schema.FuncName}} {
        #region Fields 
        {{range $field := .Schema.Fields}}{{$field.Type}} {{$field.PubVarName}} { get; set; }
        {{end}}
        #endregion Fields

        #region Methods
        public Merge({{.Schema.FuncName}}Model to_merge)
        public object Clone()
        #endregion Methods
    }

}
