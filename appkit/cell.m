#import <Appkit/NSCell.h>
# import "cell.h"
# include "_cgo_export.h"


bool Cell_Wraps(void* ptr) {
	NSCell* cell = (NSCell*)ptr;
	return [cell wraps];
}

void Cell_SetWraps(void* ptr, bool wraps) {
	NSCell* cell = (NSCell*)ptr;
	[cell setWraps:wraps];
}

bool Cell_IsScrollable(void* ptr) {
	NSCell* cell = (NSCell*)ptr;
	return [cell isScrollable];
}

void Cell_SetScrollable(void* ptr, bool scrollable) {
	NSCell* cell = (NSCell*)ptr;
	[cell setScrollable:scrollable];
}

bool Cell_IsEditable(void* ptr) {
	NSCell* cell = (NSCell*)ptr;
	return [cell isEditable];
}

void Cell_SetEditable(void* ptr, bool editable) {
	NSCell* cell = (NSCell*)ptr;
	[cell setEditable:editable];
}

bool Cell_IsSelectable(void* ptr) {
	NSCell* cell = (NSCell*)ptr;
	return [cell isSelectable];
}

void Cell_SetSelectable(void* ptr, bool selectable) {
	NSCell* cell = (NSCell*)ptr;
	[cell setSelectable:selectable];
}
