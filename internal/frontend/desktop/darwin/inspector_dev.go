//go:build darwin && (dev || debug || devtools)

package darwin

// 我们在这里使用了私有API，确保此代码仅包含在开发/调试构建中，而不包含在生产构建中。
// 否则，将应用推送到AppStore时，可能会被应用审查团队拒绝。

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework Cocoa -framework WebKit
#import <Foundation/Foundation.h>
#import "WailsContext.h"

extern void processMessage(const char *message);

@interface _WKInspector : NSObject
- (void)show;
- (void)detach;
@end

@interface WKWebView ()
- (_WKInspector *)_inspector;
@end

void showInspector(void *inctx) {
#if MAC_OS_X_VERSION_MAX_ALLOWED >= 120000
    ON_MAIN_THREAD(
		if (@available(macOS 12.0, *)) {
			WailsContext *ctx = (__bridge WailsContext*) inctx;

			@try {
				[ctx.webview._inspector show];
			} @catch (NSException *exception) {
				NSLog(@"Opening the inspector failed: %@", exception.reason);
				return;
			}

			dispatch_time_t popTime = dispatch_time(DISPATCH_TIME_NOW, 1 * NSEC_PER_SEC);
			dispatch_after(popTime, dispatch_get_main_queue(), ^(void){
				// Detach 必须稍后延时调用，并且在 show 操作后直接忽略。
				@try {
					[ctx.webview._inspector detach];
				} @catch (NSException *exception) {
					NSLog(@"Detaching the inspector failed: %@", exception.reason);
				}
			});
		} else {
			NSLog(@"Opening the inspector needs at least MacOS 12");
		}
    );
#endif
}

void setupF12hotkey() {
	[NSEvent addLocalMonitorForEventsMatchingMask:NSEventMaskKeyDown handler:^NSEvent * _Nullable(NSEvent * _Nonnull event) {
		if (event.keyCode == 111 &&
				event.modifierFlags & NSEventModifierFlagFunction &&
				event.modifierFlags & NSEventModifierFlagCommand &&
				event.modifierFlags & NSEventModifierFlagShift) {
			processMessage("wails:openInspector");
			return nil;
		}
		return event;
	}];
}
*/
import "C"
import (
	"unsafe"
)

func init() {
	C.setupF12hotkey()
}

func showInspector(context unsafe.Pointer) {
	C.showInspector(context)
}
