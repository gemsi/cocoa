

## Header Files
/Library/Developer/CommandLineTools/SDKs/MacOSX.sdk//usr/include/

/Library/Developer/CommandLineTools/SDKs/MacOSX.sdk/System/Library/Frameworks/Cocoa.framework/Headers
/Library/Developer/CommandLineTools/SDKs/MacOSX.sdk/System/Library/Frameworks/Foundation.framework/Headers
/Library/Developer/CommandLineTools/SDKs/MacOSX.sdk/System/Library/Frameworks/AppKit.framework/Headers

## CGO Types
CGFloat -- float64/float32
NSInteger -- int64/int32

char -->  C.char -->  byte
signed char -->  C.schar -->  int8
unsigned char -->  C.uchar -->  uint8
short int -->  C.short -->  int16
short unsigned int -->  C.ushort -->  uint16
int -->  C.int -->  int
unsigned int -->  C.uint -->  uint32
long int -->  C.long -->  int32 or int64
long unsigned int -->  C.ulong -->  uint32 or uint64
long long int -->  C.longlong -->  int64
long long unsigned int -->  C.ulonglong -->  uint64
float -->  C.float -->  float32
double -->  C.double -->  float64
wchar_t -->  C.wchar_t  -->  
void * -> unsafe.Pointer 