# CapsLockX - Mac/Windows

A keyboard utility that turns **Space** into a powerful modifier key, giving you mouse control, cursor navigation, and desktop management — all without leaving the home row.

Hold **Space** down to activate CLX mode. Release it to type a normal space (if no other key was pressed).

## Key Bindings

### Mouse Control (Space + ...)

| Keys | Action |
|------|--------|
| `Space + W` | Move mouse up |
| `Space + A` | Move mouse left |
| `Space + S` | Move mouse down |
| `Space + D` | Move mouse right |
| `Space + E` | Left click (hold to drag) |
| `Space + Q` | Right click (hold to drag) |
| `Space + R` | Scroll up |
| `Space + F` | Scroll down |
| `Space + Shift + R/F` | Horizontal scroll |

Mouse movement uses a **physics model** with acceleration and friction — the longer you hold, the faster it moves, and it glides to a smooth stop.

### Cursor / Edit Control (Space + ...)

| Keys | Action |
|------|--------|
| `Space + H` | ← Left arrow |
| `Space + J` | ↓ Down arrow |
| `Space + K` | ↑ Up arrow |
| `Space + L` | → Right arrow |
| `Space + Y` | Home |
| `Space + O` | End |
| `Space + U` | Page Down |
| `Space + I` | Page Up |
| `Space + T` | Backspace/Delete |
| `Space + G` | Enter |
| `Space + N` | Tab |
| `Space + P` | Shift+Tab |

Hold any cursor key for turbo repeat — it accelerates the longer you hold.

Modifier keys (Shift, Ctrl, Option/Alt, Cmd/Win) are passed through, so `Space + Shift + L` sends Shift+Right (extend selection).

### Desktop / Window Control (Space + ...)

| Keys | Action |
|------|--------|
| `Space + 1` | Switch to previous desktop |
| `Space + 2` | Switch to next desktop |
| `Space + X` | Close window/tab (Cmd+W) |
| `Space + /` or `Space + \` | Quit CLX |

## How It Works

The program registers global hotkeys using [`golang.design/x/hotkey`](https://pkg.go.dev/golang.design/x/hotkey). When Space is held:

1. All CLX sub-modules activate and register their own hotkeys.
2. On Space release, all sub-hotkeys are unregistered.
3. If no CLX key was pressed during the hold, a regular space keystroke is sent.

Mouse and scroll movement run a **physics simulation** in a background goroutine: applying acceleration, capping velocity, and decaying with friction each tick (every 10ms). This gives smooth, natural-feeling cursor control from the keyboard.

## Platform Support

- **macOS** — primary target. Uses AppleScript (`mack`) for desktop switching.
- **Windows** — supported via `mods_windows.go` (uses Win key as Cmd equivalent).

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.19+
- macOS: grant **Accessibility** permission to the terminal app (System Preferences → Privacy & Security → Accessibility)

### Build & Run

```bash
git clone https://github.com/snomiao/capslockx-mac
cd capslockx-mac

go build -o clx
chmod +x clx
./clx
```

### Watch Mode (auto-rebuild on save)

```bash
# Requires Node.js
npm install -g nodemon

bash watch.sh
```

`watch.sh` builds and relaunches CLX automatically whenever a `.go` file changes.

### Windows Cross-Compile (from Linux/Mac)

```bash
bash x-build-win-from-linux.sh
```

## Dependencies

| Package | Purpose |
|---------|---------|
| [`golang.design/x/hotkey`](https://pkg.go.dev/golang.design/x/hotkey) | Global hotkey registration |
| [`github.com/go-vgo/robotgo`](https://github.com/go-vgo/robotgo) | Mouse/keyboard automation |
| [`github.com/andybrewer/mack`](https://github.com/andybrewer/mack) | AppleScript bridge (macOS desktop switching) |
| [`golang.design/x/mainthread`](https://pkg.go.dev/golang.design/x/mainthread) | Run on OS main thread (required by macOS) |

## Roadmap

- [x] Space as modifier key
- [x] Mouse control with physics (WASD)
- [x] Mouse buttons (Q/E)
- [x] Scroll (R/F)
- [x] Cursor control (HJKL)
- [x] Page/Home/End navigation
- [x] Desktop switching (1/2)
- [ ] CapsLock as alternative modifier
- [ ] Jump to specific desktop (1–8)
- [ ] Horizontal scroll (Shift+R/F)
- [ ] Smooth physics for cursor movement

## References

- [Make the Caps Lock key useful (Reddit)](https://www.reddit.com/r/MacOS/comments/wj1cyt/make_the_caps_lock_key_useful/)
