namespace {{.FullNS}};

public partial class NumInput : ContentView {
    public NumInput() {
        InitializeComponent();
        // BindingContext = this;
    }
    public static readonly BindableProperty MaxValueProperty = BindableProperty.Create(nameof(MaxValue), typeof(int), typeof(NumInput), 100, defaultBindingMode: BindingMode.OneWay);
    public int MaxValue {
        get => (int)GetValue(NumInput.MaxValueProperty);
        set => SetValue(NumInput.MaxValueProperty, value);
    }

    public static readonly BindableProperty MinValueProperty = BindableProperty.Create(nameof(MinValue), typeof(int), typeof(NumInput), 0, defaultBindingMode: BindingMode.OneWay);
    public int MinValue {
        get => (int)GetValue(NumInput.MinValueProperty);
        set => SetValue(NumInput.MinValueProperty, value);
    }

    public static readonly BindableProperty ValueProperty = BindableProperty.Create(nameof(Value), typeof(int), typeof(NumInput), 0, defaultBindingMode: BindingMode.TwoWay);
    public int Value {
        get => (int)GetValue(NumInput.ValueProperty);
        set => SetValue(NumInput.ValueProperty, value);
    }

    public static readonly BindableProperty SliderVisibleProperty = BindableProperty.Create(nameof(SliderVisible), typeof(bool), typeof(NumInput), true, defaultBindingMode: BindingMode.OneWay);
    public bool SliderVisible {
        get => (bool)GetValue(NumInput.SliderVisibleProperty);
        set => SetValue(NumInput.SliderVisibleProperty, value);
    }

    void OnSliderValueChanged(object sender, ValueChangedEventArgs args) {
        double value = args.NewValue;
        Value = Convert.ToInt32(value);
    }

    private void Button_Clicked(object sender, EventArgs e) {
        Value++;
        if (Value > MaxValue) {
            Value = MaxValue;
        }

    }

    private void Button_Clicked_1(object sender, EventArgs e) {
        Value--;
        if (Value < MinValue) {
            Value = MinValue;
        }
    }
}