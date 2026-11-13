package main

/*
#cgo LDFLAGS: -framework ApplicationServices
#include <ApplicationServices/ApplicationServices.h>

static int trusted(void) { return AXIsProcessTrusted(); }
static int moveMouse(void) {
    CGEventRef current = CGEventCreate(NULL);
    if (!current) return 0;
    CGPoint p = CGEventGetLocation(current);
    CFRelease(current);
    // Move by one pixel, then return to the original position.
    CGPoint next = CGPointMake(p.x + 1, p.y);
    CGEventRef event = CGEventCreateMouseEvent(NULL, kCGEventMouseMoved, next, kCGMouseButtonLeft);
    if (!event) return 0;
    CGEventPost(kCGHIDEventTap, event);
    CGEventSetLocation(event, p);
    CGEventPost(kCGHIDEventTap, event);
    CFRelease(event);
    return 1;
}
*/
import "C"

import (
	"fmt"
	"os"
	"time"
)

func interval(args []string) (time.Duration, error) {
	if len(args) != 1 {
		return 0, fmt.Errorf("usage: mouse-timer <interval> (example: 30s)")
	}
	d, err := time.ParseDuration(args[0])
	if err != nil || d <= 0 {
		return 0, fmt.Errorf("interval must be a positive duration, such as 30s or 2m")
	}
	return d, nil
}

func main() {
	d, err := interval(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	if C.trusted() == 0 {
		fmt.Fprintln(os.Stderr, "Accessibility permission is required. Enable it in System Settings > Privacy & Security > Accessibility, then retry.")
		os.Exit(1)
	}
	fmt.Printf("Moving mouse every %s. Press Ctrl+C to stop.\n", d)
	ticker := time.NewTicker(d)
	defer ticker.Stop()
	for range ticker.C {
		if C.moveMouse() == 0 {
			fmt.Fprintln(os.Stderr, "could not create mouse event")
			os.Exit(1)
		}
	}
}
