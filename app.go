package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	b64 "encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/djherbis/atime"
	"github.com/skratchdot/open-golang/open"
	"github.com/wailsapp/mimetype"

	fileutil "wails-fm/core"

	"github.com/nfnt/resize"
	"github.com/shirou/gopsutil/disk"
)

// App struct
type App struct {
	ctx context.Context
}

type FileLoader struct {
	http.Handler
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) GetHomeDir() string {
	path, _ := os.UserHomeDir()
	return path
}

var thumbCacheFolder string

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	thumbCacheFolder, _ = os.UserHomeDir()
	thumbCacheFolder += "/AppData/Roaming/wails-fm/thumbnail_cache/"

	thbErr := os.MkdirAll(thumbCacheFolder, os.ModePerm)
	if thbErr != nil {
		println("Could not create thumbnailfolder at " + thumbCacheFolder + " " + thbErr.Error())
	}

}

func (a *App) GetLocalFile(requestedFilename string) string {
	var err error
	println("Requesting file:", requestedFilename)
	fileData, err := os.ReadFile(requestedFilename)
	if err != nil {
		println("Error: ", err)
	}
	sEnc := b64.StdEncoding.EncodeToString(fileData)
	return sEnc
}

func RemoveFirstChar(input string) string {
	if len(input) <= 1 {
		return ""
	}
	return input[1:]
}

func GenerateThumbnailFilename(path string) string {
	path = strings.ReplaceAll(path, ":", "$D%")
	path = strings.ReplaceAll(path, "\\", "$S%")
	path = strings.ReplaceAll(path, "/", "$S%")

	thumbCacheFolder, _ := os.UserHomeDir()
	thumbCacheFolder += "/AppData/Roaming/wails-fm/thumbnail_cache/"
	return thumbCacheFolder + path + ".png"
}

func ThumbnailInCache(path string) (bool, string) {
	path = GenerateThumbnailFilename(path)
	if _, err := os.Stat(path); err == nil {
		return true, path
	} else {
		return false, ""
	}
}

func GetMountPoints() []string {

	partitions, _ := disk.Partitions(false)
	partitionMountpoints := make([]string, len(partitions))
	for i, partition := range partitions {
		partitionMountpoints[i] = partition.Mountpoint
	}
	return partitionMountpoints
}

func (a *App) GetSubDirPaths(path string) []string {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return []string{}
	}

	dirs := []string{}

	for i := range files {
		if !files[i].IsDir() {
			continue
		}

		dirPath := filepath.Join(path, files[i].Name())

		if files[i].IsDir() {
			dirs = append(dirs, dirPath)
		}
	}

	return dirs
}

func GetThumbnail(path string) string {

	isPng := strings.HasSuffix(strings.ToLower(path), ".png")
	//isWebp := strings.HasSuffix(strings.ToLower(path), ".webp")
	isJpg := strings.HasSuffix(strings.ToLower(path), ".jpg") || strings.HasSuffix(strings.ToLower(path), ".jpeg")

	if !(isPng || isJpg) {
		return "invalid filetype, only jpg and png are supported for now"
	}

	thumbExists, thumbPath := ThumbnailInCache(path)

	if thumbExists {

		// Open file on disk.
		f, _ := os.Open(thumbPath)
		defer f.Close()

		// Read entire JPG into byte slice.
		reader := bufio.NewReader(f)
		content, _ := ioutil.ReadAll(reader)

		// Encode as base64.
		return base64.StdEncoding.EncodeToString(content)
	}

	imagePath, openErr := os.Open(path)
	if openErr != nil {
		println(openErr.Error())
		return ""
	}
	defer imagePath.Close()
	srcImage, _, decodeErr := image.Decode(imagePath)
	if decodeErr != nil {
		println(decodeErr.Error())
		return ""
	}

	thumbnail := resize.Thumbnail(128, 128, srcImage, resize.NearestNeighbor)

	buf := new(bytes.Buffer)

	var err error = nil

	filePath := GenerateThumbnailFilename(path)

	if isPng {
		err = png.Encode(buf, thumbnail)

		out, _ := os.Create(filePath)
		defer out.Close()
		err = png.Encode(out, thumbnail)
	} else if isJpg {
		err = jpeg.Encode(buf, thumbnail, nil)

		out, _ := os.Create(filePath)
		defer out.Close()
		err = jpeg.Encode(out, thumbnail, nil)
	}

	if err != nil {
		return ""
	}
	send_s3 := buf.Bytes()

	return b64.StdEncoding.EncodeToString(send_s3)
}

func imgIsSupported(extension string) bool {
	switch extension {
	case ".png":
		return true

	case ".jpg":
		return true
	default:
		return false
	}
}

func (a *App) OpenWithDefaultApp(path string) {
	open.Run(path)
}

func (a *App) FormatDate(date time.Time, format string) string {
	return date.Format(format)
}

type Filestruct struct {
	FullPath  string
	Name      string
	Size      int64
	ModDate   time.Time
	Extension string
	MimeType  string
	Hidden    bool
}

type Folderstruct struct {
	FullPath string
	Name     string
	Size     int64
	ModDate  time.Time
	Hidden   bool
	IsLink   bool
}

type FolderData struct {
	Files   []Filestruct
	Folders []Folderstruct
}

type Tab struct {
	Address        string
	AddressHistory []string
	SelectedFiles  []string
}

type AppState struct {
	Tabs       []Tab
	currentTab Tab
}

type FileDetailsSingle struct {
	Name         string
	Data         string
	Size         int64
	MimeType     string
	CreationTime time.Time
	ModifiedTime time.Time
	AccessTime   time.Time
	Owner        string
}

func (a *App) GetFileDetailsSingle(path string) FileDetailsSingle {
	mType := "unknown"
	mimetype, err := mimetype.DetectFile(path)
	if err == nil {
		mType = mimetype.String()
	}

	fi, err := os.Lstat(path)
	at, err := atime.Stat(path)
	crTime, err := fileutil.CreatrionTimeFromPath(path)
	if err != nil {

	}

	return FileDetailsSingle{
		Name:         fi.Name(),
		Size:         fi.Size(),
		MimeType:     mType,
		CreationTime: crTime,
		ModifiedTime: fi.ModTime(),
		AccessTime:   at,
	}
}

func (a *App) GetFolderAPI(path string) FolderData {
	files, _ := ioutil.ReadDir(path)

	var FileList []Filestruct
	var FolderList []Folderstruct
	for _, file := range files {
		filename := file.Name()
		hidden, _ := fileutil.FileIsHidden(path + "/" + filename)
		fullPath := path + "/" + filename

		//Check if file is Symlink
		link, errIsLink := os.Readlink(path + "/" + filename)
		isLink := errIsLink == nil
		if isLink {
			fullPath = link
		}

		if filepath.Ext(filename) == ".lnk" {

		}

		if file.IsDir() || isLink {
			FolderList = append(FolderList, Folderstruct{
				FullPath: fullPath,
				Name:     filename,
				Size:     file.Size(),
				ModDate:  file.ModTime(),
				Hidden:   hidden,
				IsLink:   isLink,
			})
		} else {
			FileList = append(FileList, Filestruct{
				FullPath:  fullPath,
				Name:      filename,
				Size:      file.Size(),
				ModDate:   file.ModTime(),
				Hidden:    hidden,
				Extension: filepath.Ext(filename),
			})
		}
	}
	return FolderData{FileList, FolderList}
}

func (a *App) GetThumbnailAsBase64(requestedFilename string) string {
	return GetThumbnail(requestedFilename)
}

type Fs map[string]Fs

type fol struct {
	Name     string
	FullPath string
	childs   []fol
}

func RequestTreeExpand() {}

func GetSubDirPaths(path string) []string {
	f, err := os.ReadDir(path)

	if err != nil {
		return []string{}
	}

	dirs := []string{}

	for _, d := range f {
		if d.IsDir() {
			dirs = append(dirs, d.Name())
		}
	}
	return dirs
}

func (a *App) GetListView(path string) string {
	subDirs := []string{}
	if path == "" {
		subDirs = GetMountPoints()
	} else {
		subDirs = GetSubDirPaths(path)
	}

	html := ""

	for _, dir := range subDirs {
		html += fmt.Sprintf("<li>%s</li>", dir)
	}
	return html
}

func (a *App) GetTree(drive string) map[string]Fs {

	var tree map[string]Fs = map[string]Fs{}

	mounts := GetMountPoints()
	mounts = []string{"A:"}
	fullPath := ""
	for _, mount := range mounts {
		tree[mount] = map[string]Fs{}
		fullPath += mount + "/"

		for _, path := range GetSubDirPaths(mount + "/") {
			fullPath += path + "/"
			tree[mount][fullPath] = map[string]Fs{}

			for _, subPath := range GetSubDirPaths(mount + "/" + path) {
				nextPath := fullPath + "/" + subPath
				tree[mount][fullPath][nextPath] = map[string]Fs{}
			}
		}
	}

	return tree
}

func (a *App) GetTreeHTML(path string) []string {

	subDirs := []string{}
	if path == "" {
		subDirs = GetMountPoints()
	} else {
		subDirs = GetSubDirPaths(path)
	}

	htmlStr := []string{}

	for _, dir := range subDirs {
		//Read Subdirs to check wether they contain subfolders and can be expanded... kinda uselessly expensive? TODO.
		subSubDirs := GetSubDirPaths(path + dir + "/")
		subdirLen := len(subSubDirs)

		if subdirLen > 0 {
			htmlStr = append(htmlStr, fmt.Sprintf("<ul  data-path='%s'><li data-path='%s' data-empty='false' ><span class='tree-expander' data-path='%s'>%v</span><span data-path='%s' class='tree-foldertext'>%s</span> </li></ul>", path+dir+"/", path+dir+"/", path+dir+"/", subdirLen, path+dir+"/", dir))
		} else {
			htmlStr = append(htmlStr, fmt.Sprintf("<ul  data-path='%s'><li data-path='%s' data-empty='true' ></span><span data-path='%s' class='tree-foldertext'>%s</span></li></ul>", path+dir+"/", path+dir+"/", path+dir+"/", dir))
		}
	}
	return htmlStr
}
