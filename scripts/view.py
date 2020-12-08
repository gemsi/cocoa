#!env python3

from generate import Component, Property, Method, Return, Param


if __name__ == "__main__":
    w = Component(
        Type="appkit.View",
        super_type='appkit.Responder',
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

            Property(name='bottomAnchor', Type='appkit.LayoutYAxisAnchor', readonly=True,
                     description='a layout anchor representing the bottom edge of the view’s frame'),
            Property(name='centerXAnchor', Type='appkit.LayoutXAxisAnchor',readonly=True,
                     description='a layout anchor representing the horizontal center of the view’s frame'),
            Property(name='centerYAnchor', Type='appkit.LayoutYAxisAnchor',readonly=True,
                     description='a layout anchor representing the vertical center of the view’s frame'),
            Property(name='firstBaselineAnchor', Type='appkit.LayoutYAxisAnchor', readonly=True, description=''),
            Property(name='heightAnchor', Type='appkit.LayoutDimension', readonly=True, description=''),
            Property(name='lastBaselineAnchor', Type='appkit.LayoutYAxisAnchor', readonly=True, description=''),
            Property(name='leadingAnchor', Type='appkit.LayoutXAxisAnchor', readonly=True, description=''),
            Property(name='leftAnchor', Type='appkit.LayoutXAxisAnchor', readonly=True, description=''),
            Property(name='rightAnchor', Type='appkit.LayoutXAxisAnchor', readonly=True, description=''),
            Property(name='topAnchor', Type='appkit.LayoutYAxisAnchor', readonly=True, description=''),
            Property(name='trailingAnchor', Type='appkit.LayoutXAxisAnchor', readonly=True, description=''),
            Property(name='widthAnchor', Type='appkit.LayoutDimension', readonly=True, description=''),

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
