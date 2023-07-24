using CommunityToolkit.Mvvm.Messaging;
namespace {{.FullNS}};

[QueryProperty(nameof(Item), "Item")]
public partial class {{.Schema.FuncName}}EditViewModel : BaseViewModel
{
    #region constructor
    public {{.Schema.FuncName}}EditViewModel({{.Schema.FuncName}}ModelDB {{.Schema.PrivVarName}}Service{{range $field := .Schema.GetFieldsByFlag `RELATIONSHIP`}},{{func_name ($field.GetFlagParam `RELATIONSHIP` 0) }}ModelDB {{priv_var_name ($field.GetFlagParam `RELATIONSHIP` 0) }}Service{{end}}) {
        {{.Schema.PrivVarName}}DataService={{.Schema.PrivVarName}}Service;{{range $field := .Schema.GetFieldsByFlag `RELATIONSHIP`}}
        {{priv_var_name ($field.GetFlagParam `RELATIONSHIP` 0) }}DataService = {{priv_var_name ($field.GetFlagParam `RELATIONSHIP` 0) }}Service;
        {{end}}


        {{range $field := .Schema.GetFieldsByFlag `RELATIONSHIP`}}
        WeakReferenceMessenger.Default.Register<{{func_name ($field.GetFlagParam `RELATIONSHIP` 0)}}ChangedMessage>(this, (r, m) => {
            //****************************************
            //Lets just get it to load the match model
            //****************************************
            data_loaded = false;
            Task.Run(() => LoadData());
        });
        {{end}}

    }
    #endregion constructor

    readonly {{.Schema.FuncName}}ModelDB {{.Schema.PrivVarName}}DataService;{{range $idx, $field := .Schema.GetFieldsByFlag `RELATIONSHIP`}}
    readonly {{func_name ($field.GetFlagParam `RELATIONSHIP` 0)}}ModelDB  {{priv_var_name ($field.GetFlagParam `RELATIONSHIP` 0)}}DataService;
    {{end}}private bool data_loaded = false;

    [ObservableProperty]
    {{.Schema.FuncName}}Model edit_item;

    {{$fields:=.Schema.GetFieldsByFlag `RELATIONSHIP`}}

    [ObservableProperty]
    {{.Schema.FuncName}}Model item;
    partial void OnItemChanged({{.Schema.FuncName}}Model oldValue, {{.Schema.FuncName}}Model newValue) { 
        {{if gt (len $fields) 0}}
        data_loaded = false;
        
        if (newValue != null) {
            Task.Run(() => LoadData());
        }
        {{else}}
        if (Item != null && Edit_item == null) {
            Edit_item = ({{.Schema.FuncName}}Model)Item.Clone();
        }
        {{end}}
    }
    
    #region Validation msg
    {{range $field := .Schema.Fields}}[ObservableProperty]
    public string {{$field.PrivVarName}}_error="";

    {{end}}
    #endregion Validation msg

    {{range $field := $fields}}[ObservableProperty]
    List<{{func_name ($field.GetFlagParam `RELATIONSHIP` 0)}}Model> {{priv_var_name ($field.GetFlagParam `RELATIONSHIP` 0)}}_items;
    {{end}}

    {{if gt (len $fields) 0}}public async Task LoadData() {
        if (data_loaded) { return; }
        if (Item == null) { return; }

        //************************
        //Load the data for Model
        //************************
        {{range $field := $fields}}{{pub_var_name ($field.GetFlagParam `RELATIONSHIP` 0)}}_items = await {{priv_var_name ($field.GetFlagParam `RELATIONSHIP` 0)}}DataService.GetItemsWithChildrenAsync();
        {{end}}
        
        if (Item != null && Edit_item == null) {
            Edit_item = ({{.Schema.FuncName}}Model)Item.Clone();
        }

        //**************************
        //Restore the data for Model
        //So the SelectedItem Works
        //**************************if (Edit_item != null) {
        {{range $field := $fields}}if (Edit_item != null) {
            var t = {{pub_var_name ($field.GetFlagParam `RELATIONSHIP` 0)}}_items.Where(n => n.Id == Edit_item.{{pub_var_name ($field.GetFlagParam `RELATIONSHIP` 0)}}.Id).FirstOrDefault();
            if (t != null) {
                Edit_item.{{$field.PubVarName}} = t;
            }
        }
        {{end}}data_loaded = true;
    }{{end}}

    {{range $field := $fields}}[RelayCommand]
    private async Task Add{{func_name ($field.GetFlagParam `RELATIONSHIP` 0)}}() {
        var new_item = new {{func_name ($field.GetFlagParam `RELATIONSHIP` 0)}}Model();
        await Shell.Current.GoToAsync(nameof({{func_name ($field.GetFlagParam `RELATIONSHIP` 0)}}EditPage), true, new Dictionary<string, object>
        {
            { "Item", new_item }
        });
    }
    {{end}}


    [RelayCommand]
    private async Task Save() {
        var errors = Edit_item.Validate();
        if (Edit_item.HasErrors) {
            await App.AlertSvc.ShowAlertAsync("Errors", "There are Validation Errors you Need to Fix", "OK");
            string error = "";
            {{range $field := .Schema.Fields}}errors.TryGetValue("{{$field.PubVarName}}", out error);
            {{$field.PubVarName}}_error = error;
            {{end}}
            return;
        }

        var answer = await App.AlertSvc.ShowConfirmationAsync("Save", "Are you Sure You Want To Save Changes", "Yes", "No");
        if (answer == true) {
            Item.Merge(this.Edit_item);
            {{if gt (len (.Schema.GetFieldsByFlag `RELATIONSHIP`)) 0}}
            await {{.Schema.PrivVarName}}DataService.SaveItemWithChildrenAsync(Item);
            {{else}}
            await {{.Schema.PrivVarName}}DataService.SaveItemAsync(Item);
            {{end}}
            WeakReferenceMessenger.Default.Send(new {{.Schema.FuncName}}ChangedMessage(Item));
            Edit_item=null;
            await Shell.Current.GoToAsync("..");
        }
    }

    [RelayCommand]
    private async Task Cancel() {
        var answer = await App.AlertSvc.ShowConfirmationAsync("Save", "Are you Sure You Want To Cancel, Changes Wont Be Saved", "Yes", "No");
        if (answer == true) {
            Edit_item=null;
            await Shell.Current.GoToAsync("..");
        }
    }
}
