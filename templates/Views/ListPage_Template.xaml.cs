
namespace {{.FullNS}} {
    public partial class {{.Schema.FuncName}}ListPage : StandardContentPage {
        {{.Schema.FuncName}}ListViewModel ViewModel;
        public {{.Schema.FuncName}}ListPage({{.Schema.FuncName}}ListViewModel viewModel) {
            InitializeComponent();
            BindingContext = ViewModel = viewModel;
        }
        protected override async void OnNavigatedTo(NavigatedToEventArgs args) {
            base.OnNavigatedTo(args);

            await ViewModel.LoadDataAsync();
        }
    }
}