#!env python3

from generate import Component, Property, InitMethod, Method, Return, Param


if __name__ == "__main__":
    w = Component(
        Type="appkit.View",
        super_type='appkit.Responder',
        additional_objc_imports=["Appkit/NSLayoutConstraint"],
        description="The infrastructure for drawing, printing, and handling events in an app.",
        properties=[
            Property(name='frame', Type='foundation.Rect', readonly=False,
                     description='the view’s frame rectangle, which defines its position and size in its superview’s coordinate system'),
            Property(name='autoresizingMask', Type='uint', go_alias_type='AutoresizingMaskOptions', readonly=False,
                     description='the view’s frame rectangle, which defines its position and size in its superview’s coordinate system'),
            Property(name='needsDisplay', Type='bool', getter_prefix_is=False,
                     description='whether the view needs to be redrawn before being displayed'),
            # Property(name='requiresConstraintBasedLayout', Type='bool', readonly=True, getter_prefix_is=False,
                    #  description='whether the view depends on the constraint-based layout system'),
            Property(name='translatesAutoresizingMaskIntoConstraints', Type='bool', getter_prefix_is=False,
                     description='whether the view’s autoresizing mask is translated into Auto Layout constraints'),
        ],
        methods=[
            Method(
                name='addSubview',
                params=[Param(name='subView', Type='appkit.View')],
                return_value=Return(''),
                description='adds a view to the view’s subviews so it’s displayed above its siblings.'
            ),
        ]
    )
    w.generate_code()
