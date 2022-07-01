#import <AppKit/NSPasteboard.h>

NSPasteboard* clipboard = NULL;

void clipboard_init() {
    clipboard = [NSPasteboard generalPasteboard];
    [clipboard declareTypes:[NSArray arrayWithObject:NSPasteboardTypeString] owner:nil];
}

bool set_clipboard(const char* data) {
    NSString* str = [NSString stringWithUTF8String:data];
    return [clipboard setString:str forType:NSPasteboardTypeString];
}

const char* get_clipboard() {
    NSString* str = [clipboard stringForType:NSPasteboardTypeString];
    return [str UTF8String];
}