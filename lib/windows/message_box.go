package windows

import (
	"errors"
	"io/fs"
	"log"
	"syscall"
	"unsafe"
)

// MessageBoxYesNo YesNo 类型弹窗
func MessageBoxYesNo(text, caption string) bool {
	ret := MessageBox(text, caption, uintptr(MB_YESNO))
	if ret == IDYES {
		return true
	}
	return false
}

// MessageBox windows 弹窗
func MessageBox(text, caption string, mbType uintptr) uintptr {
	user32, _ := syscall.LoadLibrary("user32.dll")
	messageBox, _ := syscall.GetProcAddress(user32, "MessageBoxW")
	defer syscall.FreeLibrary(user32)

	t, _ := syscall.UTF16FromString(text)
	c, _ := syscall.UTF16FromString(caption)

	ret, _, errno := syscall.SyscallN(messageBox, 0,
		uintptr(unsafe.Pointer(&t[0])),
		uintptr(unsafe.Pointer(&c[0])),
		mbType)
	if errors.Is(errno, fs.ErrNotExist) {
		log.Println(errno)
	}

	return ret
}

// message box return value
const (
	IDABORT    = 3  // The Abort button was selected.
	IDCANCEL   = 2  // The Cancel button was selected.
	IDCONTINUE = 11 // The Continue button was selected.
	IDIGNORE   = 5  // The Ignore button was selected.
	IDNO       = 7  // The No button was selected.
	IDOK       = 1  // The OK button was selected.
	IDRETRY    = 4  // The Retry button was selected.
	IDTRYAGAIN = 10 // The Try Again button was selected.
	IDYES      = 6  // The Yes button was selected.
)

// message box type
const (
	MB_OK = 0x00000000

	MB_OKCANCEL = 0x00000001

	MB_ABORTRETRYIGNORE = 0x00000002

	MB_YESNOCANCEL = 0x00000003

	MB_YESNO = 0x00000004

	MB_RETRYCANCEL = 0x00000005

	MB_CANCELTRYCONTINUE = 0x00000006

	MB_ICONHAND = 0x00000010

	MB_ICONQUESTION = 0x00000020

	MB_ICONEXCLAMATION = 0x00000030

	MB_ICONASTERISK = 0x00000040

	MB_USERICON = 0x00000080

	MB_ICONWARNING = MB_ICONEXCLAMATION

	MB_ICONERROR = MB_ICONHAND

	MB_ICONINFORMATION = MB_ICONASTERISK

	MB_ICONSTOP = MB_ICONHAND

	MB_DEFBUTTON1 = 0x00000000

	MB_DEFBUTTON2 = 0x00000100

	MB_DEFBUTTON3 = 0x00000200

	MB_DEFBUTTON4 = 0x00000300
)
