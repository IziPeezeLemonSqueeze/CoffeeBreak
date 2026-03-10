package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
	"unsafe"
)

var (
	user32              = syscall.NewLazyDLL("user32.dll")
	procGetLastInputInfo = user32.NewProc("GetLastInputInfo")
	procSendInput       = user32.NewProc("SendInput")
	kernel32            = syscall.NewLazyDLL("kernel32.dll")
	procGetTickCount64  = kernel32.NewProc("GetTickCount64")
)

type LASTINPUTINFO struct {
	CbSize uint32
	DwTime uint32
}

// Strutture per SendInput
const (
	INPUT_KEYBOARD    = 1
	KEYEVENTF_KEYUP   = 0x0002
	VK_SHIFT          = 0x10
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

func main() {
	debug := len(os.Args) > 1 && os.Args[1] == "--d"

	fmt.Println("CoffeeBreak avviato...")

	for {
		idleSeconds := getIdleSeconds()
		idleStatus := "active"
		if idleSeconds >= 300 {
			idleStatus = "idle"
		}

		if debug {
			fmt.Printf("System idle state: %s\n", idleStatus)
			fmt.Printf("  - Idle seconds: %d\n", idleSeconds)
		}

		if idleSeconds >= 290 {
			if debug {
				fmt.Println("  → Premo Shift...")
			}
			pressShift()
		}

		time.Sleep(5 * time.Second)
	}
}


