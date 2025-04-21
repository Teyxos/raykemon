package main

import (
	"os"
	"path/filepath"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var supportedTextureFormats = []string{".png", ".jpg", ".bmp", ".tga"}
var txtDirectory = "./assets/textures/"

func getTexturefromFile(c chan string, files []os.DirEntry) {
	// Load textures files from the assets folder

	if len(files) == 0 {
		rl.TraceLog(rl.LogError, "No texture files found in the directory.")
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
	rl.TraceLog(rl.LogInfo, "Reading directory: %s", txtDirectory)

	if err != nil {
		rl.TraceLog(rl.LogError, "Failed to read directory: %v", err)
		return
	}
	go getTexturefromFile(c, files)

	// Wait for the audio file to be sent through the channel

	for filePath := range c {
		// While the key right now is the file name, its gonna be user defined
		Textures[filePath] = rl.LoadTexture(filePath)
	}
}

func UnloadTextures() {
	for _, texture := range Textures {
		rl.UnloadTexture(texture)
	}
}
