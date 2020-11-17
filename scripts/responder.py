#!env python3

from .generate import Component, Property, InitMethod


if __name__ == "__main__":
    w = Component(
        Type = "appkit.Responder", 
        super_type = 'foundation.Object',
        description = "the basis of event and command processing in AppKit",
        properties = [
            Property(name='acceptsFirstResponder', Type = 'bool', getter_prefix_is=False,readonly=True, description='whether the responder accepts first responder status'),
            Property(name='nextResponder', Type = 'appkit.Responder', description='the next responder after this one, or nil if it has none'),
        ],
    )
    w.generate_code()