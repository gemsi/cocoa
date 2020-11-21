#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

bool Responder_AcceptsFirstResponder(void* ptr);
void* Responder_NextResponder(void* ptr);
void Responder_SetNextResponder(void* ptr, void* nextResponder);

bool Responder_BecomeFirstResponder(void* ptr);
