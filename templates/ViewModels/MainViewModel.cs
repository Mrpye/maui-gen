namespace {{.FullNS}};

public partial class MainViewModel : BaseViewModel {

    {{range $s := .RootSchema.Schemas}}[RelayCommand]
    private async Task Show{{$s.FuncName}}() {
        await Shell.Current.GoToAsync(nameof({{$s.FuncName}}ListPage));
    }
    {{end}}
}
