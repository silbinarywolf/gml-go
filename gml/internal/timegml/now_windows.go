// +build windows

package timegml

import (
	"syscall"
	"unsafe"
)

var (
	qpc  *syscall.Proc
	freq uint64
)

func init() {
	var err error
	dll, err := syscall.LoadDLL("kernel32.dll")
	if err != nil {
		panic(err)
	}
	qpc, err = dll.FindProc("QueryPerformanceCounter")
	if err != nil {
		panic(err)
	}

	// Get frequency once at initialization
	// docs: https://docs.microsoft.com/en-us/windows/desktop/SysInfo/acquiring-high-resolution-time-stamps
	qpf, err := dll.FindProc("QueryPerformanceFrequency")
	if err != nil {
		panic(err)
	}
	if ret, _, err := qpf.Call(uintptr(unsafe.Pointer(&freq))); ret == 0 {
		panic(err)
	}
}

func Now() int64 {
	var ctr uint64
	if ret, _, err := qpc.Call(uintptr(unsafe.Pointer(&ctr))); ret == 0 {
		panic(err)
	}
	// 1.0e9 to convert seconds to nanoseconds
	res := int64((1.0e9 * (ctr)) / freq)
	return res
}
