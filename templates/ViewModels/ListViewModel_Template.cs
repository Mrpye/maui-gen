using CommunityToolkit.Mvvm.Messaging;
namespace {{.FullNS}};

public partial class {{.Schema.FuncName}}ListViewModel : BaseViewModel
{
    readonly {{.Schema.FuncName}}ModelDB dataService;

    [ObservableProperty]
    bool isRefreshing;

    [ObservableProperty]
    ObservableCollection<{{.Schema.FuncName}}Model> items;

    public {{.Schema.FuncName}}ListViewModel({{.Schema.FuncName}}ModelDB service)
    {
        dataService = service;
         WeakReferenceMessenger.Default.Register<PersonChangedMessage>(this, (r, m) => {
            //****************************************
            //Lets just get it to load the match model
            //****************************************
            //IsRefreshing = true;

        });
    }

    [RelayCommand]
    private async Task OnRefreshing()
    {
        IsRefreshing = true;

        try
        {
            await LoadDataAsync();
        }
        finally
        {
            IsRefreshing = false;
        }
    }

    

    [RelayCommand]
    public async Task LoadDataAsync()
    {
        IsRefreshing = true;
        {{if gt (len (.Schema.GetFieldsByFlag `RELATIONSHIP`)) 0}}
        Items = new ObservableCollection<{{.Schema.FuncName}}Model>(await dataService.GetItemsWithChildrenAsync());
        {{else}}
        Items = new ObservableCollection<{{.Schema.FuncName}}Model>(await dataService.GetItemsAsync());
        {{end}}    
        IsRefreshing = false;
    }

    [RelayCommand]
    private async Task GoToDetails({{.Schema.FuncName}}Model item)
    {
        await Shell.Current.GoToAsync(nameof({{.Schema.FuncName}}DetailsPage), true, new Dictionary<string, object>
        {
            { "Item", item }
        });
    }

    [RelayCommand]
    private async Task Add{{.Schema.FuncName}}() {
       var new_item= new {{.Schema.FuncName}}Model();
        await Shell.Current.GoToAsync(nameof({{.Schema.FuncName}}EditPage), true, new Dictionary<string, object>
        {
            { "Item", new_item }
        });
    }

    //*************
    //Delete Player
    //*************
    [RelayCommand]
    private async Task Delete{{.Schema.FuncName}}({{.Schema.FuncName}}Model line) {

        //******************
        //Confirm the delete
        //******************
        var answer = await App.AlertSvc.ShowConfirmationAsync("Delete", "Are you sure you want to delete item", "Yes", "No");
        if (answer == true) {
            
            var p = ({{.Schema.FuncName}}Model)line;
            await dataService.DeleteItemAsync(p);
            Items.Remove(p);
            
            WeakReferenceMessenger.Default.Send(new {{.Schema.FuncName}}ChangedMessage(p));
        }
       
    }


    [RelayCommand]
    private async Task Edit{{.Schema.FuncName}}({{.Schema.FuncName}}Model item) {
        await Shell.Current.GoToAsync(nameof({{.Schema.FuncName}}EditPage), true, new Dictionary<string, object>{
            { "Item",item}
        });
    }

    [RelayCommand]
    private async Task {{.Schema.FuncName}}Details({{.Schema.FuncName}}Model item) {
        await Shell.Current.GoToAsync(nameof({{.Schema.FuncName}}DetailsPage), true, new Dictionary<string, object>{
            { "Item",item}
        });
    }

    [RelayCommand]
    private async Task Ok() {
        await Shell.Current.GoToAsync("..");
    }

    [RelayCommand]
    private async Task Cancel() {
        await Shell.Current.GoToAsync("..");
    }
}
