package main

import (
	"os"
	"path/filepath"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var supportedAudioFormats = []string{".mp3", ".wav", ".ogg", ".flac"}
var sndDirectory = "./assets/audio/"

func getAudioFromFile(c chan string, files []os.DirEntry) {
	// Load audio files from the assets folder

	if len(files) == 0 {
		rl.TraceLog(rl.LogError, "No audio files found in the directory.")
		close(c)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if slices.Contains(supportedAudioFormats, filepath.Ext(file.Name())) {
			filePath := filepath.Join(sndDirectory, file.Name())
			c <- filePath
		}
	}

	close(c)
}

var Audios = map[string]rl.Music{}

func LoadAudio() {
	c := make(chan string)

	files, err := os.ReadDir(sndDirectory)
	rl.TraceLog(rl.LogInfo, "Reading directory: %s", sndDirectory)

	if err != nil {
		rl.TraceLog(rl.LogError, "Failed to read directory: %v", err)
	}
	go getAudioFromFile(c, files)

	// Wait for the audio file to be sent through the channel

	for filePath := range c {
		Audios[filePath] = rl.LoadMusicStream(filePath)
	}

}

func UnloadAudio() {
	for _, audio := range Audios {
		rl.UnloadMusicStream(audio)
	}
}
