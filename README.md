# ☕ CoffeeBreak

Keep your Windows session awake with a tiny Go utility that checks user idle time and sends a synthetic `Shift` press when inactivity gets too high.

![Platform](https://img.shields.io/badge/platform-Windows-0078D6?style=flat-square)
![Language](https://img.shields.io/badge/language-Go-00ADD8?style=flat-square)
[![License](https://img.shields.io/badge/license-See%20License-blue?style=flat-square)](https://github.com/IziPeezeLemonSqueeze/CoffeeBreak/blob/main/License)

## 🚀 Overview

CoffeeBreak is a lightweight background utility for Windows.  
It reads system idle time using native Win32 APIs and prevents lock/sleep side effects triggered by inactivity policies.

## ✨ Features

- Native Windows API integration (`user32.dll`, `kernel32.dll`)
- Idle detection based on real input inactivity
- Automatic keep-awake input (`Shift` key down/up)
- Debug mode with live idle output (`--d`)
- Single static executable build with Go

## ✅ Requirements

- Windows
- Go `1.25.6` or newer

## 🧪 Quick Start

Run directly:

```bash
go run .
```

Run with debug logs:

```bash
go run . --d
```

## 📄 License

See [License](./License).

## 👤 Author

Stefano Pastore  
p.stefano92@gmail.com
