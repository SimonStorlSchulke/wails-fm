package main

import (
	"bytes"
	"context"
	b64 "encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

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

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// Greet returns a greeting for the given name
func (a *App) GreetJohn(name string) string {
	return "Hello John It's show time!"
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

func (a *App) GetFolderFilepaths(path string) string {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return "<div class='filearea'></div>"
	}
	strDirs := "<div class='folderarea'>"
	strFiles := "<div class='filearea'>"
	for _, file := range files {
		filename := file.Name()
		filesize := file.Size()
		moddate := file.ModTime()
		extension := filepath.Ext(filename)

		if file.IsDir() {
			strDirs += fmt.Sprintf(`
			<div class='file-card dir onclick="enterFolder()"'>
				<p class="filename">%s</p>
				<p class="modified">%v</p>
			</div>`,
				filename, moddate.Format("2006-01-02"))
		} else {
			strFiles += fmt.Sprintf(`
					<div class='file-card ext-%s'>
						<p class="filename">%s</p>
						<p class="filesize">%v</p>
						<p class="modified">%v</p>
					</div>`,
				RemoveFirstChar(extension), filename, filesize, moddate.Format("2006-01-02"))
		}

	}
	strDirs += "</div>"
	strFiles += "</div>"
	return strDirs + strFiles
}

func (a *App) GetThumbnailAsBase64(requestedFilename string) string {
	return GetThumbnail(requestedFilename)
}
