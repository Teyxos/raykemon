package main

import (
	"log"
	"os"
	"path/filepath"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var supportedTextureFormats = []string{".png", ".jpg", ".bmp", ".tga"}
var txtDirectory = "./assets/textures/"

// ! There is an error when there is no textures file in the directory. It only happpens when either there are no files of textures or audio (all threads are offline)
func getTexturefromFile(c chan string, files []os.DirEntry) {
	// Load textures files from the assets folder

	if len(files) == 0 {
		log.Println("No audio files found in the directory.")
		close(c)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if slices.Contains(supportedTextureFormats, filepath.Ext(file.Name())) {
			filePath := filepath.Join(txtDirectory, file.Name())
			c <- filePath
		}
	}

	close(c)
}

var Textures = map[string]rl.Texture2D{}

func LoadTextures() {
	c := make(chan string)

	files, err := os.ReadDir(txtDirectory)
	log.Println(files)

	if err != nil {
		log.Fatal(err)
	}
	go getTexturefromFile(c, files)

	// Wait for the audio file to be sent through the channel

	for filePath := range c {
		Textures[filePath] = rl.LoadTexture(filePath)
	}
}

func UnloadTextures() {
	for _, texture := range Textures {
		rl.UnloadTexture(texture)
	}
}
