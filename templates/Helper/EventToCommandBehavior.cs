using System.Reflection;
using System.Windows.Input;

namespace {{.FullNS}}
{
    public class EventToCommandBehavior : BehaviorBase<VisualElement>
    {
        private Delegate _eventHandler;
        public static readonly BindableProperty EventNameProperty = BindableProperty.Create("EventName", typeof(string), typeof(EventToCommandBehavior), null, propertyChanged: OnEventNameChanged);
        public static readonly BindableProperty CommandProperty = BindableProperty.Create("Command", typeof(ICommand), typeof(EventToCommandBehavior));

        public string EventName
        {
            get => (string)GetValue(EventNameProperty);
            set => SetValue(EventNameProperty, value);
        }
        public ICommand Command
        {
            get => (ICommand)GetValue(CommandProperty);
            set => SetValue(CommandProperty, value);
        }

        protected override void OnAttachedTo(VisualElement bindable)
        {
            base.OnAttachedTo(bindable);
            RegisterEvent(EventName);
        }

        protected override void OnDetachingFrom(VisualElement bindable)
        {
            DeregisterEvent(EventName);
            base.OnDetachingFrom(bindable);
        }

        private static void OnEventNameChanged(BindableObject bindable, object oldValue, object newValue)
        {
            var behavior = (EventToCommandBehavior)bindable;

            if (behavior.AssociatedObject == null) return;

            var oldEventName = (string)oldValue;
            var newEventName = (string)newValue;

            behavior.DeregisterEvent(oldEventName);
            behavior.RegisterEvent(newEventName);
        }

        private void RegisterEvent(string name)
        {
            if (string.IsNullOrWhiteSpace(name)) return;

            var eventInfo = AssociatedObject.GetType().GetRuntimeEvent(name);

            if (eventInfo == null)
                throw new ArgumentException($"EventToCommandBehavior: Can't register the '{EventName}' event.");

            var methodInfo = typeof(EventToCommandBehavior).GetTypeInfo().GetDeclaredMethod("OnEvent");
            _eventHandler = methodInfo.CreateDelegate(eventInfo.EventHandlerType, this);
            eventInfo.AddEventHandler(AssociatedObject, _eventHandler);
        }

        private void DeregisterEvent(string name)
        {
            if (string.IsNullOrWhiteSpace(name) || _eventHandler == null)
                return;

            var eventInfo = AssociatedObject.GetType().GetRuntimeEvent(name);

            if (eventInfo == null)
                throw new ArgumentException($"EventToCommandBehavior: Can't de-register the '{EventName}' event.");

            eventInfo.RemoveEventHandler(AssociatedObject, _eventHandler);
            _eventHandler = null;
        }

        private void OnEvent(object sender, object eventArgs)
        {
            if (Command == null) return;

            if (Command.CanExecute(eventArgs))
                Command.Execute(eventArgs);
        }
    }
}
