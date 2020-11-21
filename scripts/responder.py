#!env python3

from .generate import Component, Property, Method, Return

if __name__ == "__main__":
    w = Component(
        Type="appkit.Responder",
        super_type='foundation.Object',
        description="the basis of event and command processing in AppKit",
        properties=[
            Property(name='acceptsFirstResponder', Type='bool', getter_prefix_is=False, readonly=True,
                     description='whether the responder accepts first responder status'),
            Property(name='nextResponder', Type='appkit.Responder',
                     description='the next responder after this one, or nil if it has none'),
        ],
        methods=[
            Method(name='becomeFirstResponder', return_value=Return(Type='bool'),
                   description='notifies the receiver that it’s about to become first responder in its NSWindow')

        ]
    )
    w.generate_code()
