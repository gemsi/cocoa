#!env python3

from generate import Component, Property, init_method, Param, Method, Return

if __name__ == "__main__":
    w = Component(
        Type="appkit.GridColumn",
        super_type='foundation.Object',
        objc_imports=['Appkit/NSGridView'],
        description="a row within a grid view",
        properties=[
            Property(name='numberOfCells', Type='int', readonly=True, description=''),
            Property(name='hidden', Type='bool', description=''),
            Property(name='leadingPadding', Type='float64', description=''),
            Property(name='trailingPadding', Type='float64', description=''),
            Property(name='width', Type='float64', description=''),

            Property(name='xPlacement', Type='int', go_alias_type='GridCellPlacement', description=''),

            Property(name='gridView', Type='appkit.GridView', readonly=True, description=''),
        ],
        methods=[
            Method(
                name='cellAtIndex',
                params=[Param(name='index', Type='int')],
                return_value=Return('appkit.GridCell'),
                description='',
            ),
            Method(
                name='mergeCellsInRange',
                params=[Param(name='r', Type='foundation.Range')],
                description='',
            ),
        ]
    )
    w.generate_code()
