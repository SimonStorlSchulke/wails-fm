package fileutil

import (
	"log"
	"runtime"
	"syscall"
)

func FileIsHidden(filename string) (bool, error) {

	if runtime.GOOS == "windows" {

		pointer, err := syscall.UTF16PtrFromString(filename)
		if err != nil {
			return false, err
		}
		attributes, err := syscall.GetFileAttributes(pointer)
		if err != nil {
			return false, err
		}
		return attributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0, nil
	} else {
		log.Fatal("Unable to invoke GetFileAttributes() function or FILE_ATTRIBUTE_HIDDEN property")
	}
	return false, nil
}
