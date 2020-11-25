#!env python3

from generate import Component, Property, InitMethod, Param

if __name__ == "__main__":
    w = Component(
        Type="appkit.StackView",
        super_type='appkit.View',
        description="a view that arranges an array of views horizontally or vertically and updates their placement and sizing when the window size changes",
        init_method=InitMethod(name="stackViewWithViews", is_factory=True,
                               params=[Param(name='views', Type='appkit.View', array=True)]),
        properties=[
            Property(name='views', Type='appkit.View', readonly=True, array=True,
                     description='the array of views owned by the stack view'),
        ],
    )
    w.generate_code()
