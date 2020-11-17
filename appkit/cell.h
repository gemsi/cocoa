#import <stdbool.h>
#import <stdlib.h>

bool Cell_Wraps(void* ptr);
void Cell_SetWraps(void* ptr, bool wraps);
bool Cell_IsScrollable(void* ptr);
void Cell_SetScrollable(void* ptr, bool scrollable);
bool Cell_IsEditable(void* ptr);
void Cell_SetEditable(void* ptr, bool editable);
bool Cell_IsSelectable(void* ptr);
void Cell_SetSelectable(void* ptr, bool selectable);

