using CommunityToolkit.Mvvm.Messaging.Messages;
namespace {{.FullNS}}
{
    public class {{.Schema.FuncName}}ChangedMessage : ValueChangedMessage<{{.Schema.FuncName}}Model> {
        public {{.Schema.FuncName}}ChangedMessage({{.Schema.FuncName}}Model value) : base(value) {
           
        }
    }
}
