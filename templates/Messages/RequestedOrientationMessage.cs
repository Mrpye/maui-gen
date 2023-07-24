using CommunityToolkit.Mvvm.Messaging.Messages;
/*
    private void X_Appearing(object sender, EventArgs e) {
        WeakReferenceMessenger.Default.Send(new RequestedOrientationMessage(ScreenOrientation.Landscape));
    }

    private void X_Disappearing(object sender, EventArgs e) {
        WeakReferenceMessenger.Default.Send(new RequestedOrientationMessage(ScreenOrientation.Unspecified));
    }

Platforms/Android/MainActivity
 WeakReferenceMessenger.Default.Register<RequestedOrientationMessage>(this, (r, m) => {
	RequestedOrientation = (ScreenOrientation)m.Value;
});
*/
namespace {{.FullNS}} {
    public class RequestedOrientationMessage : ValueChangedMessage<ScreenOrientation> {
        public RequestedOrientationMessage(ScreenOrientation value) : base(value) {

        }
    }
}
