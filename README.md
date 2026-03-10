# ☕ CoffeeBreak

> A smart app that keeps your computer awake while you relax with a coffee ☕

![Node.js](https://img.shields.io/badge/Node.js-v14+-339933?style=flat-square&logo=nodedotjs)
![License](https://img.shields.io/badge/License-See%20License%20File-blue?style=flat-square)
![Status](https://img.shields.io/badge/Status-Active-brightgreen?style=flat-square)

## 🎯 Description

**CoffeeBreak** is a lightweight Node.js application that monitors system inactivity and prevents your computer from going into sleep mode. When the system remains inactive for more than 5 minutes, the app automatically sends a keyboard command (Shift key press) to keep your computer active.

Perfect for those who work in environments where security policies force the system to sleep after a short period of inactivity, or for anyone who simply wants to keep their workstation active during breaks.

## ✨ Features

- 🔍 **Smart Monitoring**: Detects system inactivity state in real-time
- ⏱️ **Customizable Timer**: Configure inactivity time (default: 290 seconds)
- 🎮 **Discrete Automation**: Sends practically invisible keyboard commands
- 🔧 **Debug Mode**: Use `--d` flag to display inactivity status
- 🚀 **Lightweight & Fast**: Minimal resource consumption
- 🛡️ **Cross-Platform**: Compatible with Windows, macOS, and Linux

## 📋 Prerequisites

- Node.js v14 or higher
- npm or yarn

## 🚀 Installation

1. **Clone the repository**
   ```bash
   Download from Release
   ```

## 💻 Usage

### Standard Start
Launch the app in silent mode:
```bash
node index.js
```

### Debug Mode
Start with debug output to see inactivity status in real-time:
```bash
node index.js --d
```

**Debug output:**
```
System idle state: active
  - Idle seconds: 0
System idle state: idle
  - Idle seconds: 125
```

### Continuous Execution (Background)
On Windows, you can launch the app in background:
```bash
start "" "C:\Program Files\nodejs\node.exe" "path\to\index.js"
```

On Linux/macOS:
```bash
nohup node index.js &
```

## 📦 Dependencies

- **[@paymoapp/real-idle](https://www.npmjs.com/package/@paymoapp/real-idle)**: System inactivity state detection
- **[robotjs](https://github.com/octalmage/robotjs)**: Mouse and keyboard control

## 🔐 Security

The app does not log or send data to remote servers. It works completely locally on your computer.

## 📝 License

See the [License](./License) file for details.

## 👤 Author

**Stefano Pastore**
- Email: p.stefano92@gmail.com
- GitHub: [@IziPeezeLemonSqueeze](https://github.com/IziPeezeLemonSqueeze)

##  Contributing

No contribution is required. If you want to fork this project, send an email to p.stefano92@gmail.com

## 🐛 Bug Report

If you find a bug, please open an [Issue](https://github.com/IziPeezeLemonSqueeze/coffeebreak/issues) with:
- Clear description of the problem
- Steps to reproduce it
- Operating system and Node.js version

## 🎉 Enjoy Your Coffee Break! ☕

---

**Made with ❤️ by Stefano Pastore**
