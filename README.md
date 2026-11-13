<<<<<<< HEAD
# mouse-timer-macos
This is a terminal based app for corporate employees with Apple Mac that cannot set sleep timeout. Inspired by Jiggler
=======
# Mouse Timer

A small macOS command-line tool that moves the mouse one pixel back and forth at a fixed interval to provide periodic mouse activity.

Requires macOS 26.6 or later, Go, Xcode Command Line Tools, and Accessibility permission for the process running the program. Your terminal or the built binary may need permission in System Settings → Privacy & Security → Accessibility.

```sh
go build -o mouse-timer .
./mouse-timer 30s
```

The required argument is a positive Go duration, such as `5s` or `2m`. The first movement occurs after the interval. Press Ctrl+C to stop. Run tests with `go test -race ./...`.
>>>>>>> 4f10e42 (Create mouse timer MVP)
