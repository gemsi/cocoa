#!env python3

from generate import Component, Property, init_method, Param, Method, Return, ActionMethod


if __name__ == "__main__":
    w = Component(
        Type="appkit.Button",
        super_type='appkit.Control',
        description="a control that defines an area on the screen that can be used to trigger actions",

        properties=[
            Property(name='title', Type='string',
                     description='the title displayed on the button when it’s in an off state'),
            Property(name='bezelStyle', Type='uint', go_alias_type='BezelStyle',
                     description='the appearance of the button’s border'),
            Property(name='state', Type='int', go_alias_type='ControlStateValue', description='the button’s state'),
        ],
        methods=[
            init_method(name="initWithFrame", Type='appkit.Button', params=[Param(name='frame', Type='foundation.Rect')]),
            Method(
                name='setButtonType',
                params=[Param(name='buttonType', Type='uint', go_alias='ButtonType')],
                description='sets the button’s type, which affects its user interface and behavior when clicked',
            ),
        ],
        action_methods=[
            ActionMethod(name='action')
        ]
    )
    w.generate_code()
