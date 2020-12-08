#!env python3

from generate import Component, Property, Param, Method, Return, DelegateMethod, init_method


if __name__ == "__main__":
    w = Component(
        Type="appkit.Application",
        super_type='appkit.Responder',
        description="An object that manages an app’s main event loop and resources used by all of that app’s objects.",
        properties=[
            Property(name='currentEvent', Type='appkit.Event', readonly=True, description='the last event object that the app retrieved from the event queue'),
            Property(name='running', Type='bool', readonly=True, description='whether the main event loop is running'),
            Property(name='active', Type='bool', readonly=True, description='whether this is the active app'),
            Property(name='mainMenu', Type='appkit.Menu', description='the app’s main menu bar'),
            Property(name='windowsMenu', Type='appkit.Menu', description='the Window menu of the app'),
            Property(name='servicesMenu', Type='appkit.Menu', description='the app’s Services menu'),
            Property(name='mainWindow', Type='appkit.Window', readonly=True, description='the app’s main window'),
            Property(name='windows', Type='appkit.Window', readonly=True, array=True, description='an array of the app’s window objects'),
            Property(name='hidden', Type='bool', readonly=True, description='whether the app is hidden'),
            
        ],
        methods=[
            Method(name='run', description='starts the main event loop'),
            Method(name='finishLaunching', description='activates the app, opens any files specified by the NSOpen user default, and unhighlights the app’s icon'),
            Method(name='stop', params=[Param(name='sender', Type='foundation.Object')], description='stops the main event loop'),
            Method(name='terminate', params=[Param(name='sender', Type='foundation.Object')], description='terminates the receiver'),
            Method(name='sendEvent', params=[Param(name='event', Type='appkit.Event')], description='dispatches an event to other objects'),
            Method(name='activateIgnoringOtherApps', params=[Param(name='flag', Type='bool')], description='makes the receiver the active app'),
            Method(name='deactivate', description='deactivates the receiver'),
            Method(name='sharedApplication', static=True, return_value=Return('appkit.Application'), description='returns the application instance, creating it if it doesn’t exist yet'),
            Method(name='setActivationPolicy', params=[Param(name='activationPolicy', Type='int', go_alias='ApplicationActivationPolicy')], description='attempts to modify the app’s activation policy'),
            Method(name='windowWithWindowNumber', params=[Param(name='windowNum', Type='int')], description='returns the window corresponding to the specified window number'),
            Method(name='miniaturizeAll', params=[Param(name='sender', Type='foundation.Object')], description='miniaturizes all the receiver’s windows'),
            Method(name='hide', params=[Param(name='sender', Type='foundation.Object')], description='hides all the receiver’s windows, and the next app in line is activated'),
            Method(name='unhide', params=[Param(name='sender', Type='foundation.Object')], description='restores hidden windows to the screen and makes the receiver active'),
            Method(name='unhideWithoutActivation', description='restores hidden windows without activating their owner (the receiver)'),
            Method(name='updateWindows', description='sends an update message to each onscreen window'),
            Method(name='setWindowsNeedUpdate', params=[Param(name='needUpdate', Type='bool')]),
            Method(name='arrangeInFront', params=[Param(name='sender', Type='foundation.Object')]),
            Method(name='preventWindowOrdering'),
            
        ],
        delegate_methods=[
            DelegateMethod(
                name='applicationWillFinishLaunching',
                params=[Param(name='notification', Type='foundation.Notification')],
                description='sent by the default notification center immediately before the application object is initialized'
            ),
            DelegateMethod(
                name='applicationDidFinishLaunching',
                params=[Param(name='notification', Type='foundation.Notification')],
                description='sent by the default notification center after the application has been launched and initialized but before it has received its first event'
            ),
            DelegateMethod(
                name='applicationShouldTerminateAfterLastWindowClosed',
                params=[Param(name='theApplication', Type='appkit.Application')],
                return_value=Return('bool'),
                default_return_value='true',
                description='invoked when the user closes the last window the application has open.return NO if the application should not be terminated when its last window is closed; otherwise, YES to terminate the application'
            ),
        ]
    )
    w.generate_code()
