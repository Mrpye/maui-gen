namespace {{.RootNS}};

public partial class App : Application
{
    public static IServiceProvider Services;
    public static IAlertService AlertSvc;
    public App(IServiceProvider provider)
	{
		InitializeComponent();

        Services = provider;
        AlertSvc = Services.GetService<IAlertService>();

        // Uncomment the following as a quick way to test loading resources for different languages
        //CultureInfo.CurrentCulture = CultureInfo.CurrentUICulture = new CultureInfo("es");

        MainPage = new AppShell();

        
	}
}
