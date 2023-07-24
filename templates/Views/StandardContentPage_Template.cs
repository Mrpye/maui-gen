using System.Windows.Input;
namespace {{.FullNS}} {
    [XamlCompilation(XamlCompilationOptions.Compile)]
    public partial class StandardContentPage : ContentPage {

        public static readonly BindableProperty LeftButtonCommandProperty = BindableProperty.Create(nameof(LeftButtonCommand), typeof(ICommand), typeof(StandardContentPage));

        public System.Windows.Input.ICommand LeftButtonCommand {
            get { return (System.Windows.Input.ICommand)GetValue(LeftButtonCommandProperty); }
            set { SetValue(LeftButtonCommandProperty, value); }
        }
        public static readonly BindableProperty RightButtonCommandProperty = BindableProperty.Create(nameof(RightButtonCommand), typeof(ICommand), typeof(StandardContentPage));

        public System.Windows.Input.ICommand RightButtonCommand {
            get { return (System.Windows.Input.ICommand)GetValue(RightButtonCommandProperty); }
            set { SetValue(RightButtonCommandProperty, value); }
        }

        public static readonly BindableProperty LeftButtonVisibleProperty = BindableProperty.Create(nameof(LeftButtonVisible), typeof(bool), typeof(StandardContentPage), true);
        public static readonly BindableProperty RightButtonVisibleProperty = BindableProperty.Create(nameof(RightButtonVisible), typeof(bool), typeof(StandardContentPage), true);
        
        public static readonly BindableProperty LeftButtonTitleProperty = BindableProperty.Create(nameof(LeftButtonTitle), typeof(string), typeof(StandardContentPage), "Cancel");
        public static readonly BindableProperty RightButtonTitleProperty = BindableProperty.Create(nameof(RightButtonTitle), typeof(string), typeof(StandardContentPage), "OK");


        public static readonly BindableProperty IsRefreshingProperty = BindableProperty.Create(nameof(IsRefreshing), typeof(bool), typeof(StandardContentPage), false);
        private static void IsRefreshingPropertyChanged(BindableObject bindable, object oldvalue, object newvalue) {

           
        }
        public bool IsRefreshing {
            get => (bool)GetValue(IsRefreshingProperty);
            set => SetValue(IsRefreshingProperty, value);
        }

        public static readonly BindableProperty IndicatorMsgProperty = BindableProperty.Create(nameof(IndicatorMsg), typeof(string), typeof(StandardContentPage), "Downloading Data Please Wait!!!");
        public string IndicatorMsg {
            get => (string)GetValue(IndicatorMsgProperty);
            set => SetValue(IndicatorMsgProperty, value);
        }

        public string LeftButtonTitle {
            get => (string)GetValue(LeftButtonTitleProperty);
            set => SetValue(LeftButtonTitleProperty, value);
        }
        public string RightButtonTitle {
            get => (string)GetValue(RightButtonTitleProperty);
            set => SetValue(RightButtonTitleProperty, value);
        }

        public bool LeftButtonVisible {
            get => (bool)GetValue(LeftButtonVisibleProperty);
            set => SetValue(LeftButtonVisibleProperty, value);
        }

        public bool RightButtonVisible {
            get => (bool)GetValue(RightButtonVisibleProperty);
            set => SetValue(RightButtonVisibleProperty, value);
        }
        
    }
}