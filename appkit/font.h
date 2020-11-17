#import <stdlib.h>
#import <stdbool.h>

void* Font_FontWithName(const char* name, double size);

void* Font_UserFontOfSize(double size);

void* Font_UserFixedPitchFontOfSize(double size);

void* Font_SystemFontOfSize(double size);
void* Font_SystemFontOfSize2(double size, double weight);
void* Font_BoldSystemFontOfSize(double size);
void* Font_LabelFontOfSize(double size);
void* Font_MessageFontOfSize(double size);
void* Font_MenuBarFontOfSize(double size);
void* Font_MenuFontOfSize(double size);
void* Font_ControlContentFontOfSize(double size);
void* Font_TitleBarFontOfSize(double size);
void* Font_PaletteFontOfSize(double size);
void* Font_ToolTipsFontOfSize(double size);
void* Font_MonospacedSystemFontOfSize(double size, double weight);
void* Font_MonospacedDigitSystemFontOfSize(double size, double weight);

double Font_SystemFontSize();
double Font_SmallSystemFontSize();
double Font_LabelFontSize();

bool Font_FixedPitch(void* ptr);
const char* Font_DisplayName(void* ptr);
const char* Font_FamilyName(void* ptr);
const char* Font_FontName(void* ptr);
double Font_PointSize(void* ptr);
bool Font_Vertical(void* ptr);

double Font_FontWeightUltraLight();
double Font_FontWeightThin();
double Font_FontWeightLight();
double Font_FontWeightRegular();
double Font_FontWeightMedium();
double Font_FontWeightSemibold();
double Font_FontWeightBold();
double Font_FontWeightHeavy();
double Font_FontWeightBlack();