#!env python3

from generate import Component, Property, init_method, Param, Method, Return

if __name__ == "__main__":
    w = Component(
        Type="appkit.GridView",
        super_type='appkit.View',
        description="a container that aligns views in a flexible grid of rows and columns",
        properties=[
            Property(name='numberOfRows', Type='int', readonly=True,
                     description='the number of rows in the grid view'),
            Property(name='numberOfColumns', Type='int', readonly=True,
                     description='the number of columns in the grid view'),
            Property(name='columnSpacing', Type='float64',
                     description='the column spacing for the grid view'),
            Property(name='rowSpacing', Type='float64',
                     description='the row spacing for the grid view'),
            Property(name='rowAlignment', Type='int', go_alias_type='GridRowAlignment',
                     description='the row alignment for the grid view'),
            Property(name='xPlacement', Type='int', go_alias_type='GridCellPlacement',
                     description='The placement of the cell within the grid column'),
            Property(name='yPlacement', Type='int', go_alias_type='GridCellPlacement',
                     description='The placement of the cell within the grid row'),

        ],
        methods=[
            init_method(
                name='initWithFrame',
                Type='appkit.GridView',
                params=[
                    Param(name='frameRect', Type='foundation.Rect'),
                ],
            ),
            # Method(
            #     name='gridViewWithViews',
            #     register_delegate=True,
            #     static=True,
            #     params=[Param(name='views', Type='appkit.View', array=True)],
            #     return_value=Return(Type='appkit.GridView'),
            #     description='creates a newly allocated grid view object with the specified array of arrays of views',
            # ),
            Method(
                name='gridViewWithNumberOfColumns',
                register_delegate=True,
                static=True,
                params=[Param(name='columnCount', Type='int'), Param(
                    name='rowCount', Type='int', objc_param_name='rows')],
                return_value=Return(Type='appkit.GridView'),
                description='creates a newly allocated grid view object with the specified number of columns and rows',
            ),
            Method(
                name='indexOfColumn',
                params=[Param(name='column', Type='appkit.GridColumn')],
                return_value=Return(Type='int'),
                description='returns the index of the specified grid column',
            ),
            Method(
                name='indexOfRow',
                params=[Param(name='row', Type='appkit.GridRow')],
                return_value=Return(Type='int'),
                description='returns the index of the specified grid row',
            ),
            Method(
                name='columnAtIndex',
                params=[Param(name='index', Type='int')],
                return_value=Return(Type='appkit.GridColumn'),
                description='returns the grid column object at the specified index',
            ),
            Method(
                name='rowAtIndex',
                params=[Param(name='index', Type='int')],
                return_value=Return(Type='appkit.GridRow'),
                description='returns the grid row object at the specified index',
            ),
            Method(
                name='addRowWithViews',
                params=[Param(name='views', Type='appkit.View', array=True)],
                return_value=Return('appkit.GridRow'),
                description='adds an array of views to a new row',
            ),
            Method(
                name='insertRowAtIndex',
                params=[
                    Param(name='index', Type='int'),
                    Param(name='views', Type='appkit.View', array=True, objc_param_name='withViews'),
                ],
                return_value=Return('appkit.GridRow'),
                description='inserts the array of view objects into the grid view at the index',
            ),
            Method(
                name='removeRowAtIndex',
                params=[
                    Param(name='index', Type='int'),
                ],
                description='removes the row from the grid view at the index',
            ),
            Method(
                name='moveRowAtIndex',
                params=[
                    Param(name='fromIndex', Type='int'),
                    Param(name='toIndex', Type='int', objc_param_name='toIndex'),
                ],
                description='moves the specified row to the new row location',
            ),
            Method(
                name='addColumnWithViews',
                params=[Param(name='views', Type='appkit.View', array=True)],
                return_value=Return('appkit.GridColumn'),
                description='adds an array of views to a new column',
            ),
            Method(
                name='insertColumnAtIndex',
                params=[
                    Param(name='index', Type='int'),
                    Param(name='views', Type='appkit.View', array=True, objc_param_name='withViews'),
                ],
                return_value=Return('appkit.GridColumn'),
                description='inserts the array of view objects into the grid view at the index',
            ),
            Method(
                name='removeColumnAtIndex',
                params=[
                    Param(name='index', Type='int'),
                ],
                description='removes the column from the grid view at the index',
            ),
            Method(
                name='moveColumnAtIndex',
                params=[
                    Param(name='fromIndex', Type='int'),
                    Param(name='toIndex', Type='int', objc_param_name='toIndex'),
                ],
                description='moves the specified column to the new column location',
            ),
            Method(
                name='cellAtColumnIndex',
                go_func_name='CellAt',
                params=[
                    Param(name='columnIndex', Type='int'),
                    Param(name='rowIndex', Type='int', objc_param_name='rowIndex'),
                ],
                return_value=Return('appkit.GridCell'),
                description='returns the grid cell object at the specified column and row index',
            ),
            Method(
                name='cellForView',
                params=[
                    Param(name='view', Type='appkit.View'),
                ],
                return_value=Return('appkit.GridCell'),
                description='returns the grid cell object that contains the given view or one of its ancestors',
            ),
            Method(
                name='mergeCellsInHorizontalRange',
                params=[
                    Param(name='hRange', Type='foundation.Range'),
                    Param(name='vRange', Type='foundation.Range', objc_param_name='verticalRange'),
                ],
                description='expands the cell at the top-leading corner of the horizontal and vertical range to cover the entire area',
            ),
            
        ]
    )
    w.generate_code()
