#import <AppKit/NSCursor.h>
#import <AppKit/NSImage.h>
#import <AppKit/NSColor.h>

void get_cursor_size(int *width, int *height) {
    NSCursor *cursor = [NSCursor currentSystemCursor];
    NSImage *image = [cursor image];
    NSSize size = [image size];
    *width = size.width;
    *height = size.height;
}

void cursor_copy(unsigned char* pixels, int width, int height) {
    NSCursor *cursor = [NSCursor currentSystemCursor];
    NSImage *image = [cursor image];
    NSSize size = [image size];
    CGImageRef CGImage = [image CGImageForProposedRect:nil context:nil hints:nil];
    NSBitmapImageRep *bitmap = [[[NSBitmapImageRep alloc] initWithCGImage:CGImage] autorelease];
    for (int y = 0; y < height; y++) {
        if (y > size.height)
            break;
        for (int x = 0; x < width; x++) {
            if (x > size.width)
                break;
            NSColor *color = [bitmap colorAtX:x y:y];
            pixels[y * width * 4 + x * 4 + 0] = [color redComponent] * 255;
            pixels[y * width * 4 + x * 4 + 1] = [color greenComponent] * 255;
            pixels[y * width * 4 + x * 4 + 2] = [color blueComponent] * 255;
            pixels[y * width * 4 + x * 4 + 3] = [color alphaComponent] * 255;
        }
    }
}

