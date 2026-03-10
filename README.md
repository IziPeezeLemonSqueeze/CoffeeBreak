# CoffeeBreak

Keep your Windows session awake with a tiny Go utility that checks user idle time and sends a synthetic `Shift` press when inactivity gets too high.

![Platform](https://img.shields.io/badge/platform-Windows-0078D6?style=flat-square)
![Language](https://img.shields.io/badge/language-Go-00ADD8?style=flat-square)
![License](https://img.shields.io/badge/license-See%20License-blue?style=flat-square)

## Overview

CoffeeBreak is a lightweight background utility for Windows.  
It reads system idle time using native Win32 APIs and prevents lock/sleep side effects triggered by inactivity policies.

## Features

- Native Windows API integration (`user32.dll`, `kernel32.dll`)
- Idle detection based on real input inactivity
- Automatic keep-awake input (`Shift` key down/up)
- Debug mode with live idle output (`--d`)
- Single static executable build with Go

## How It Works

1. Reads idle time from `GetLastInputInfo` + `GetTickCount64`
2. Loops every `5s`
3. If idle is `>= 290s`, sends `Shift` via `SendInput`
4. In debug mode, prints current idle state continuously

## Requirements

- Windows
- Go `1.25.6` or newer

## Quick Start

Run directly:

```bash
go run .
```

Run with debug logs:

```bash
go run . --d
```

## Build

Build executable:

```bash
go build -o dist/CoffeeBreak.exe .
```

Run executable:

```bash
.\dist\CoffeeBreak.exe
```

Run executable in debug mode:

```bash
.\dist\CoffeeBreak.exe --d
```

## Icon In Executable

This repository already includes `rsrc.syso`, so `go build` embeds the app icon automatically.

If you change the icon (`asset/caffe.ico`), regenerate resources before building:

```bash
rsrc -ico asset/caffe.ico -o rsrc.syso
go build -o dist/CoffeeBreak.exe .
```

Install `rsrc` (one time):

```bash
go install github.com/akavel/rsrc@latest
```

## Runtime Configuration (Current Defaults)

| Setting | Value | Source |
| --- | --- | --- |
| Poll interval | `5s` | `time.Sleep(5 * time.Second)` |
| Keep-awake trigger | `290s` idle | `if idleSeconds >= 290` |
| Debug idle label threshold | `300s` idle | `if idleSeconds >= 300` |
| Simulated key | `Shift` | `VK_SHIFT` |

## Project Layout

- `main.go` - app logic and Windows interop
- `go.mod` - Go module definition
- `asset/caffe.ico` - source icon
- `rsrc.syso` - Windows resource blob included in build

## Notes

- Works only on Windows.
- Sends synthetic keyboard input locally; no network calls are made.
- If Explorer still shows an old icon, rename the EXE or restart Explorer to refresh icon cache.

## License

See [License](./License).

## Author

Stefano Pastore  
p.stefano92@gmail.com
