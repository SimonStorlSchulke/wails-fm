package main

import (
	"bytes"
	"context"
	b64 "encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/skratchdot/open-golang/open"

	fileutil "wails-fm/core"

	"github.com/nfnt/resize"
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

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
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

func GetThumbnail(path string) string {

	imagePath, _ := os.Open(path)
	defer imagePath.Close()
	srcImage, _, _ := image.Decode(imagePath)

	thumbnail := resize.Thumbnail(64, 64, srcImage, resize.NearestNeighbor)

	buf := new(bytes.Buffer)

	var err error = nil

	if filepath.Ext(path) == ".png" {
		err = png.Encode(buf, thumbnail)
	} else {
		err = jpeg.Encode(buf, thumbnail, nil)
	}

	if err != nil {
		return "ERROR: " + err.Error()
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
	Hidden    bool
}

type Folderstruct struct {
	FullPath string
	Name     string
	Size     int64
	ModDate  time.Time
	Hidden   bool
}

type FolderData struct {
	Files   []Filestruct
	Folders []Folderstruct
}

func (a *App) GetFolderAPI(path string) FolderData {
	files, _ := ioutil.ReadDir(path)

	var FileList []Filestruct
	var FolderList []Folderstruct
	for _, file := range files {
		filename := file.Name()
		hidden, _ := fileutil.FileIsHidden(path + "/" + filename)

		if file.IsDir() {
			FileList = append(FileList, Filestruct{
				FullPath:  path + "/" + filename,
				Name:      filename,
				Size:      file.Size(),
				ModDate:   file.ModTime(),
				Extension: filepath.Ext(filename),
				Hidden:    hidden,
			})
		} else {
			FolderList = append(FolderList, Folderstruct{
				FullPath: path + "/" + filename,
				Name:     filename,
				Size:     file.Size(),
				ModDate:  file.ModTime(),
				Hidden:   hidden,
			})
		}
	}
	return FolderData{FileList, FolderList}
}

func (a *App) GetThumbnailAsBase64(requestedFilename string) string {
	return GetThumbnail(requestedFilename)
}
