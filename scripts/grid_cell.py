#!env python3

from generate import Component, Property, init_method, Param, Method, Return

if __name__ == "__main__":
    w = Component(
        Type="appkit.GridCell",
        objc_imports=['Appkit/NSGridView'],
        super_type='foundation.Object',
        description="an individual content area within a grid view, typically at the intersection of a row and a column",
        properties=[
            Property(name='column', Type='appkit.GridColumn', readonly=True, description=''),
            Property(name='row', Type='appkit.GridRow', readonly=True, description=''),
            Property(name='contentView', Type='appkit.View', readonly=True, description=''),
            # Property(name='emptyContentView', Type='appkit.View', readonly=True, description=''),
            Property(name='customPlacementConstraints', Type='appkit.LayoutConstraint', array=True, description=''),
            Property(name='rowAlignment', Type='int', go_alias_type='GridRowAlignment', description=''),
            Property(name='xPlacement', Type='int', go_alias_type='GridCellPlacement', description=''),
            Property(name='yPlacement', Type='int', go_alias_type='GridCellPlacement', description=''),
            
        ],
    )
    w.generate_code()
