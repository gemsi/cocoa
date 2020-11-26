#!env python3

from generate import Component, Property, init_method, Param, Method, Return, DelegateMethod


if __name__ == "__main__":
    w = Component(
        Type="appkit.TextField",
        super_type='appkit.Control',
        description="text that the user can select or edit and that sends its action message to its target when the user presses the Return key",
        properties=[
            Property(name='bezeled', Type='bool',
                     description='whether the receiver draws a bezeled border around its contents'),
            Property(name='drawsBackground', Type='bool', getter_prefix_is=False, description='whether the receiver’s cell draws its background color behind its text'),
            Property(name='editable', Type='bool', description='whether the user can edit the value in the text field'),
            Property(name='selectable', Type='bool', description='whether the receiver is selectable (but not editable)'),
            Property(name='textColor', Type='appkit.Color', description='the color used to draw the receiver’s text'),
            Property(name='backgroundColor', Type='appkit.Color', description='the color of the background that the receiver’s cell draws behind the text.'),
        ],
        methods=[ 
            init_method(name="initWithFrame", Type="appkit.TextField", params=[Param(name='frame', Type='foundation.Rect')]),
        ],
        delegate_methods=[
            DelegateMethod(
                name='controlTextDidChange',
                params=[Param(name='notification', Type='foundation.Notification')],
                description=''
            ),
            DelegateMethod(
                name='controlTextDidEndEditing',
                params=[Param(name='notification', Type='foundation.Notification')],
                description=''
            ),
            DelegateMethod(
                name='controlTextDidBeginEditing',
                params=[Param(name='notification', Type='foundation.Notification')],
                description=''
            ),
        ]
    )
    w.generate_code()
