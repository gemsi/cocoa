#!env python3

from generate import Component, Property, InitMethod, Method, Param, Return

if __name__ == "__main__":
    w = Component(
        Type="appkit.ProgressIndicator",
        super_type='appkit.View',
        description="an interface that provides visual feedback to the user about the status of an ongoing task",
        init_method=InitMethod(name="initWithFrame", params=[Param(name='frame', Type='foundation.Rect')]),
        properties=[
            Property(name="usesThreadedAnimation", Type='bool', getter_prefix_is=False,
                     description='whether the progress indicator implements animation in a separate thread'),
            Property(name="doubleValue", Type='float64',
                     description='the value that indicates the current extent of the progress indicator'),
            Property(name="minValue", Type='float64', description='the minimum value for the progress indicator'),
            Property(name="maxValue", Type='float64', description='the maximum value for the progress indicator'),
            Property(name="indeterminate", Type='bool', description='whether the progress indicator is indeterminate'),
            Property(name="bezeled", Type='bool',
                     description='whether the progress indicator’s frame has a three-dimensional bezel'),
            Property(name="displayedWhenStopped", Type='bool',
                     description='whether the progress indicator hides itself when it isn’t animating'),
        ],
        methods=[
            Method(name="startAnimation",
                   params=[Param(name='sender', Type='foundation.Object')],
                   return_value=Return(''),
                   description='starts the animation of an indeterminate progress indicator',
                   ),
            Method(name="stopAnimation",
                   params=[Param(name='sender', Type='foundation.Object')],
                   return_value=Return(''),
                   description='stops the animation of an indeterminate progress indicator',
                   ),
            Method(name="incrementBy",
                   params=[Param(name='delta', Type='float64')],
                   return_value=Return(''),
                   description='advances the progress bar of a determinate progress indicator by the specified amount',
                   ),
        ],

    )
    w.generate_code()
