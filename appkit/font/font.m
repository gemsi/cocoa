#import <AppKit/NSFont.h>
#import "font.h"

void* Font_FontWithName(const char* name, double size) {
    return [NSFont fontWithName:[NSString stringWithUTF8String:name] size:size];
}

void* Font_UserFontOfSize(double size) {
    return [NSFont userFontOfSize:size];
}


void* Font_UserFixedPitchFontOfSize(double size) {
    return [NSFont userFixedPitchFontOfSize:size];
}


void* Font_SystemFontOfSize(double size) {
    return [NSFont systemFontOfSize:size];
}

void* Font_SystemFontOfSize2(double size, double weight) {
    return [NSFont systemFontOfSize:size weight:weight];
}

void* Font_BoldSystemFontOfSize(double size) {
    return [NSFont boldSystemFontOfSize:size];
}

void* Font_LabelFontOfSize(double size) {
    return [NSFont labelFontOfSize:size];
}

void* Font_MessageFontOfSize(double size) {
    return [NSFont messageFontOfSize:size];
}

void* Font_MenuBarFontOfSize(double size) {
    return [NSFont menuBarFontOfSize:size];
}

void* Font_MenuFontOfSize(double size) {
    return [NSFont menuFontOfSize:size];
}

void* Font_ControlContentFontOfSize(double size) {
    return [NSFont controlContentFontOfSize:size];
}

void* Font_TitleBarFontOfSize(double size) {
    return [NSFont titleBarFontOfSize:size];
}

void* Font_PaletteFontOfSize(double size) {
    return [NSFont paletteFontOfSize:size];
}

void* Font_ToolTipsFontOfSize(double size) {
    return [NSFont toolTipsFontOfSize:size];
}

void* Font_MonospacedSystemFontOfSize(double size, double weight) {
    return [NSFont monospacedSystemFontOfSize:size weight:weight];
}
void* Font_MonospacedDigitSystemFontOfSize(double size, double weight) {
    return [NSFont monospacedDigitSystemFontOfSize:size weight:weight];
}

double Font_SystemFontSize() {
    return [NSFont systemFontSize];
}
double Font_SmallSystemFontSize() {
    return [NSFont smallSystemFontSize];
}
double Font_LabelFontSize() {
    return [NSFont labelFontSize];
}

bool Font_FixedPitch(void* ptr) {
    NSFont* font = (NSFont*)ptr;
    return font.fixedPitch;
}

double Font_PointSize(void* ptr) {
    NSFont* font = (NSFont*)ptr;
    return font.pointSize;
}

const char* Font_DisplayName(void* ptr) {
    NSFont* font = (NSFont*)ptr;
    return [font.displayName UTF8String];
}

const char* Font_FamilyName(void* ptr) {
    NSFont* font = (NSFont*)ptr;
    return [font.familyName UTF8String];
}

const char* Font_FontName(void* ptr) {
    NSFont* font = (NSFont*)ptr;
    return [font.fontName UTF8String];
}

bool Font_Vertical(void* ptr) {
    NSFont* font = (NSFont*)ptr;
    return font.vertical;
}

double Font_FontWeightUltraLight() {
    return NSFontWeightUltraLight;
}
double Font_FontWeightThin() {
    return NSFontWeightThin;
}
double Font_FontWeightLight() {
    return NSFontWeightLight;
}
double Font_FontWeightRegular() {
    return NSFontWeightRegular;
}
double Font_FontWeightMedium() {
    return NSFontWeightMedium;
}
double Font_FontWeightSemibold() {
    return NSFontWeightSemibold;
}
double Font_FontWeightBold() {
    return NSFontWeightBold;
}
double Font_FontWeightHeavy() {
    return NSFontWeightHeavy;
}
double Font_FontWeightBlack() {
    return NSFontWeightBlack;
}