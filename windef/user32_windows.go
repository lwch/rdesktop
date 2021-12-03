package windef

import "syscall"

var (
	libUser32, _ = syscall.LoadLibrary("user32.dll")
	// FuncGetDesktopWindow https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdesktopwindow
	FuncGetDesktopWindow, _ = syscall.GetProcAddress(syscall.Handle(libUser32), "GetDesktopWindow")
	// FuncGetDC https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdc
	FuncGetDC, _ = syscall.GetProcAddress(syscall.Handle(libUser32), "GetDC")
	// FuncReleaseDC https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-releasedc
	FuncReleaseDC, _ = syscall.GetProcAddress(syscall.Handle(libUser32), "ReleaseDC")
	// FuncGetCursorInfo https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcursorinfo
	FuncGetCursorInfo, _ = syscall.GetProcAddress(syscall.Handle(libUser32), "GetCursorInfo")
	// FuncGetIconInfo https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-geticoninfo
	FuncGetIconInfo, _ = syscall.GetProcAddress(syscall.Handle(libUser32), "GetIconInfo")
	// FuncDrawIcon https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawicon
	FuncDrawIcon, _ = syscall.GetProcAddress(syscall.Handle(libUser32), "DrawIcon")
)

type (
	// HANDLE handle object
	HANDLE uintptr
	// BOOL bool
	BOOL int32
	// DWORD double word
	DWORD uint32
	// LONG long
	LONG int32
	// HCURSOR cursor handle
	HCURSOR HANDLE
	// HBITMAP bitmap handle
	HBITMAP HANDLE
)

// POINT pointer
type POINT struct {
	X LONG
	Y LONG
}

// CURSORINFO cursor info
type CURSORINFO struct {
	CbSize      DWORD
	Flags       DWORD
	HCursor     HCURSOR
	PTScreenPos POINT
}

// ICONINFO icon info
type ICONINFO struct {
	FIcon    BOOL
	XHotspot DWORD
	YHotspot DWORD
	HbmMask  HBITMAP
	HbmColor HBITMAP
}

const (
	// CURSORSHOWING https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-cursorinfo
	CURSORSHOWING = 0x00000001
)
