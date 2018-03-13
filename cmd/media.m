#import <AppKit/AppKit.h>
#import <IOKit/hidsystem/ev_keymap.h>
#import <media.h>

void StartApp(){
  [NSAutoreleasePool new];
  [MediaKeysApplication sharedApplication];
  [NSApp setActivationPolicy: NSApplicationActivationPolicyAccessory];
  [NSApp run];
}

@implementation MediaKeysApplication

- (void)sendEvent: (NSEvent*)event
{
	if( [event type] == NSEventTypeSystemDefined && [event subtype] == 8 )
	{
		int keyCode = (([event data1] & 0xFFFF0000) >> 16);
		int keyFlags = ([event data1] & 0x0000FFFF);
		int keyState = (((keyFlags & 0xFF00) >> 8)) == 0xA;
		int keyRepeat = (keyFlags & 0x1);

    if (keyCode == NX_KEYTYPE_PLAY && keyState == 0)
      play_pause();
	}

	[super sendEvent: event];
}
@end
