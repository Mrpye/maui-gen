namespace {{.FullNS}};

[QueryProperty(nameof(Item), "Item")]
public partial class {{.Schema.FuncName}}DetailsViewModel : BaseViewModel
{
    [ObservableProperty]
    {{.Schema.FuncName}}Model item;

    [RelayCommand]
    private async Task Cancel() {
        await Shell.Current.GoToAsync("..");
    }
}
