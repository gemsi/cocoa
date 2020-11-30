#!env python3

from generate import Component, Property, Method, Return, Param, init_method


if __name__ == "__main__":
    w = Component(
        Type="appkit.LayoutGuide",
        super_type='foundation.Object',
        description="a rectangular area that can interact with Auto Layout",
        properties=[
            Property(name='frame', Type='foundation.Rect', readonly=True,
                     description='the layout guide’s frame in its owning view’s coordinate system'),

            Property(name='bottomAnchor', Type='appkit.LayoutYAxisAnchor', readonly=True,
                     description='a layout anchor representing the bottom edge of the view’s frame'),
            Property(name='centerXAnchor', Type='appkit.LayoutXAxisAnchor',readonly=True,
                     description='a layout anchor representing the horizontal center of the view’s frame'),
            Property(name='centerYAnchor', Type='appkit.LayoutYAxisAnchor',readonly=True,
                     description='a layout anchor representing the vertical center of the view’s frame'),
            Property(name='heightAnchor', Type='appkit.LayoutDimension', readonly=True, description=''),
            Property(name='leadingAnchor', Type='appkit.LayoutXAxisAnchor', readonly=True, description=''),
            Property(name='leftAnchor', Type='appkit.LayoutXAxisAnchor', readonly=True, description=''),
            Property(name='rightAnchor', Type='appkit.LayoutXAxisAnchor', readonly=True, description=''),
            Property(name='topAnchor', Type='appkit.LayoutYAxisAnchor', readonly=True, description=''),
            Property(name='trailingAnchor', Type='appkit.LayoutXAxisAnchor', readonly=True, description=''),
            Property(name='widthAnchor', Type='appkit.LayoutDimension', readonly=True, description=''),
            Property(name='owningView', Type='appkit.View', readonly=True, description=''),
        ],
        methods=[
            init_method(name="init", Type="appkit.LayoutGuide"),
        ]
    )
    w.generate_code()
