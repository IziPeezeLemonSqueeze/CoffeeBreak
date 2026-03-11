package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
	"unicode/utf16"
	"unsafe"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procGetLastInputInfo = user32.NewProc("GetLastInputInfo")
	procSendInput        = user32.NewProc("SendInput")
	kernel32             = syscall.NewLazyDLL("kernel32.dll")
	procGetTickCount64   = kernel32.NewProc("GetTickCount64")
)

type LASTINPUTINFO struct {
	CbSize uint32
	DwTime uint32
}

// Strutture per SendInput
const (
	INPUT_KEYBOARD  = 1
	KEYEVENTF_KEYUP = 0x0002
	VK_SHIFT        = 0x10
)

type KEYBDINPUT struct {
	WVk         uint16
	WScan       uint16
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
}

type INPUT struct {
	Type uint32
	Ki   KEYBDINPUT
	_    [8]byte // padding
}

func hasArg(name string) bool {
	for _, arg := range os.Args[1:] {
		if arg == name {
			return true
		}
	}
	return false
}

func getIdleSeconds() int {
	var lii LASTINPUTINFO
	lii.CbSize = uint32(unsafe.Sizeof(lii))
	procGetLastInputInfo.Call(uintptr(unsafe.Pointer(&lii)))

	tickCount, _, _ := procGetTickCount64.Call()
	idleMs := uint64(tickCount) - uint64(lii.DwTime)
	return int(idleMs / 1000)
}

func pressShift() {
	inputs := []INPUT{
		// Key down
		{Type: INPUT_KEYBOARD, Ki: KEYBDINPUT{WVk: VK_SHIFT}},
		// Key up
		{Type: INPUT_KEYBOARD, Ki: KEYBDINPUT{WVk: VK_SHIFT, DwFlags: KEYEVENTF_KEYUP}},
	}
	procSendInput.Call(
		uintptr(len(inputs)),
		uintptr(unsafe.Pointer(&inputs[0])),
		uintptr(unsafe.Sizeof(inputs[0])),
	)
}

func psSingleQuoted(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "''") + "'"
}

func psEncodedCommand(script string) string {
	u16 := utf16.Encode([]rune(script))
	buf := make([]byte, len(u16)*2)
	for i, v := range u16 {
		buf[2*i] = byte(v)
		buf[2*i+1] = byte(v >> 8)
	}
	return base64.StdEncoding.EncodeToString(buf)
}

func runToastScript(title, message, appID string) error {
	script := fmt.Sprintf(`
$ErrorActionPreference = 'Stop'
Set-StrictMode -Version Latest
$title = %s
$message = %s
$appId = %s

[Windows.UI.Notifications.ToastNotificationManager, Windows.UI.Notifications, ContentType = WindowsRuntime] > $null
[Windows.Data.Xml.Dom.XmlDocument, Windows.Data.Xml.Dom.XmlDocument, ContentType = WindowsRuntime] > $null

$xml = @"
<toast>
  <visual>
    <binding template="ToastGeneric">
      <text>$([System.Security.SecurityElement]::Escape($title))</text>
      <text>$([System.Security.SecurityElement]::Escape($message))</text>
    </binding>
  </visual>
</toast>
"@

$doc = New-Object Windows.Data.Xml.Dom.XmlDocument
$doc.LoadXml($xml)

$toast = [Windows.UI.Notifications.ToastNotification]::new($doc)
[Windows.UI.Notifications.ToastNotificationManager]::CreateToastNotifier($appId).Show($toast)
# Some systems drop the toast if the host exits immediately.
Start-Sleep -Milliseconds 250
Write-Output 'TOAST_SENT'
`, psSingleQuoted(title), psSingleQuoted(message), psSingleQuoted(appID))

	encoded := psEncodedCommand(script)
	cmd := exec.Command(
		"powershell",
		"-NoProfile",
		"-NonInteractive",
		"-ExecutionPolicy", "Bypass",
		"-EncodedCommand", encoded,
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		msg := strings.TrimSpace(string(out))
		if msg == "" {
			return fmt.Errorf("toast error: %w", err)
		}
		return fmt.Errorf("toast error: %w: %s", err, msg)
	}
	if !strings.Contains(string(out), "TOAST_SENT") {
		msg := strings.TrimSpace(string(out))
		if msg == "" {
			msg = "missing TOAST_SENT marker"
		}
		return fmt.Errorf("toast error: %s", msg)
	}
	return nil
}

func showModernToast(title, message string) error {
	// Try app IDs that are accepted on non-packaged Win32 contexts.
	appIDs := []string{
		"Windows.SystemToast.SecurityAndMaintenance",
		"Windows PowerShell",
	}

	var lastErr error
	for _, appID := range appIDs {
		if err := runToastScript(title, message, appID); err == nil {
			return nil
		} else {
			lastErr = err
		}
	}
	return lastErr
}

var totalCount = 0.0

func main() {
	debug := hasArg("--d")

	if hasArg("--toast") {
		if err := showModernToast("Coffee Break!", "You need to take a coffee break!"); err != nil {
			fmt.Fprintf(os.Stderr, "Errore toast: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Toast inviato.")
		return
	}

	fmt.Println("CoffeeBreak started...")

	for {
		idleSeconds := getIdleSeconds()

		if debug {
			fmt.Printf("System idle state:\n  - Idle seconds: %d\n", idleSeconds)
		}

		if idleSeconds >= 290 {
			if debug {
				fmt.Println("  → preventing idle...")
			}
			pressShift()
		}
		totalCount += float64(5 * time.Second)

		if totalCount >= float64(2*time.Hour) {
			totalCount = 0.0
			if debug {
				fmt.Println("You need to take a coffee break!")
			}
			if err := showModernToast("Coffee Break!", "You need to take a coffee break!"); err != nil {
				fmt.Fprintf(os.Stderr, "Errore toast: %v\n", err)
				os.Exit(1)
			}
		}

		time.Sleep(5 * time.Second)
	}
}
