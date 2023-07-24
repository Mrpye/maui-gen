namespace {{.FullNS}};

public partial class {{.Schema.FuncName}}EditPage : StandardContentPage {

    {{.Schema.FuncName}}EditViewModel ViewModel;

    public {{.Schema.FuncName}}EditPage({{.Schema.FuncName}}EditViewModel viewModel)
	{
	    InitializeComponent();
        BindingContext = ViewModel= viewModel;
    }

     protected override async void OnNavigatedTo(NavigatedToEventArgs args) {
        base.OnNavigatedTo(args);
       // {{if gt (len (.Schema.GetFieldsByFlag `RELATIONSHIP`)) 0}}await ViewModel.LoadData();{{end}}
    }

    protected override bool OnBackButtonPressed() {
        return true;
    }
    
}