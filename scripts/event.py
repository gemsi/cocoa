#!env python3

from generate import Component, Property, InitMethod

if __name__ == "__main__":
    w = Component(
        Type="appkit.Event",
        super_type='foundation.Object',
        description="an object that contains information about an input action, such as a mouse click or a key press",
        properties=[
            Property(name='locationInWindow', Type='foundation.Point', readonly=True,
                     description='the receiver’s location in the base coordinate system of the associated window'),
            Property(name='window', Type='appkit.Window', readonly=True,
                     description='the window object associated with the event'),
            Property(name='windowNumber', Type='int64', readonly=True,
                     description='the identifier for the window device associated with the event'),
            Property(name='timestamp', Type='float64', readonly=True,
                     description='the time when the event occurred in seconds since system startup.'),
        ],
    )
    w.generate_code()
