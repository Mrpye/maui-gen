namespace {{.FullNS}};

public partial class {{.Schema.FuncName}}DetailsPage : StandardContentPage {
    public {{.Schema.FuncName}}DetailsPage({{.Schema.FuncName}}EditViewModel viewModel)
	{
	    InitializeComponent();
        BindingContext =  viewModel;
    }

    protected override bool OnBackButtonPressed() {
        return true;
    }
    
}