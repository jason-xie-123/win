// +build windows

package win

import "unsafe"

var (
	// Library
	wininet uintptr

	// Functions
	internetGetConnectedState uintptr
)

func init() {
	// Library
	wininet = doLoadLibrary("Wininet.dll")

	// Functions
	internetGetConnectedState = doGetProcAddress(wininet, "InternetGetConnectedState")
}

func InternetGetConnectedState(lpdwFlags LPDWORD, dwReserved DWORD) bool {
	ret1 := syscall3(internetGetConnectedState, 2,
		uintptr(unsafe.Pointer(lpdwFlags)),
		uintptr(dwReserved),
		0)
	return ret1 != 0
}
