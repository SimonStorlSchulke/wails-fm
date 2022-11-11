package fileutil

import (
	"log"
	"os"
	"path"
	"runtime"
	"syscall"
	"time"
	"unicode/utf8"

	"golang.org/x/tools/godoc/vfs"
)

func CreationTme(fi os.FileInfo) time.Time {
	return time.Unix(0, fi.Sys().(*syscall.Win32FileAttributeData).CreationTime.Nanoseconds())
}

func GenerateFileMetaPath(qid string) string {
	thumbCacheFolder, _ := os.UserHomeDir()
	thumbCacheFolder += "/AppData/Roaming/wails-fm/meta/"
	return thumbCacheFolder + qid + ".wmeta"
}

func CreatrionTimeFromPath(path string) (time.Time, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return time.Time{}, err
	}
	fi.Sys()
	return CreationTme(fi), nil
}

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

// FileQUI generates a more-or-less unique ID for each file based on its creation time.
// This is orders of magnitude faster than checking the actual File ID, but may result in bugs when files have the exact same creation time
func FileQUI(path string) string {
	/*t, err := CreatrionTimeFromPath(path)
	if err != nil {
		return "invalid QID - cannot get creation Time"
	}
	return t.Format(time.RFC3339Nano)*/
	return "a"
}

// FileQUI generates a more-or-less unique ID for each file based on its creation time.
// This is orders of magnitude faster than checking the actual File ID, but may result in bugs when files have the exact same creation time
func QIDFromTime(t time.Time, name string, isDir bool) string {
	if !isDir {
		return t.Format("2006-01-02T15_04_05.999999999Z07_00")
	} else {
		return t.Format("2006-01-02T15_04_05.999999999Z07_00") + name
	}
}

// IsText reports whether a significant prefix of s looks like correct UTF-8;
// that is, if it is likely that s is human-readable text.
func IsText(s []byte) bool {
	const max = 1024 // at least utf8.UTFMax
	if len(s) > max {
		s = s[0:max]
	}
	for i, c := range string(s) {
		if i+utf8.UTFMax > len(s) {
			// last char may be incomplete - ignore
			break
		}
		if c == 0xFFFD || c < ' ' && c != '\n' && c != '\t' && c != '\f' {
			// decoding error or control character - not a text file
			return false
		}
	}
	return true
}

var textExt = map[string]bool{
	".css":    false,
	".js":     false,
	".svg":    false,
	".md":     false,
	".txt":    false,
	".html":   false,
	".svelte": false,
	".go":     false,
}

// IsTextFile reports whether the file has a known extension indicating
// a text file, or if a significant chunk of the specified file looks like
// correct UTF-8; that is, if it is likely that the file contains human-
// readable text.
func IsTextFile(fs vfs.Opener, filename string) bool {
	// if the extension is known, use it for decision making
	if isText, found := textExt[path.Ext(filename)]; found {
		return isText
	}

	// the extension is not known; read an initial chunk
	// of the file and check if it looks like text
	f, err := fs.Open(filename)
	if err != nil {
		return false
	}
	defer f.Close()

	var buf [1024]byte
	n, err := f.Read(buf[0:])
	if err != nil {
		return false
	}
	return IsText(buf[0:n])
}
