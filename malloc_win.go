//go:build windows
// +build windows

package malloc

import (
	"syscall"
)

var (
	kernel32         = syscall.NewLazyDLL("kernel32.dll")
	procVirtualAlloc = kernel32.NewProc("VirtualAlloc")
	procVirtualFree  = kernel32.NewProc("VirtualFree")
)

const (
	MEM_COMMIT     = 0x1000
	MEM_RESERVE    = 0x2000
	PAGE_READWRITE = 0x04
)

func PlatformMalloc(size uintptr) (uintptr, error) {
	addr, _, err := procVirtualAlloc.Call(0, size, MEM_COMMIT|MEM_RESERVE, PAGE_READWRITE)
	if addr == 0 {
		return 0, err
	}
	return addr, nil
}

func PlatformFree(addr uintptr) error {
	_, _, err := procVirtualFree.Call(addr, 0, 0x8000) // MEM_RELEASE
	if err != syscall.Errno(0) {
		return err
	}
	return nil
}
