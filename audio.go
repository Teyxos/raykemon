package main

import (
	"log"
	"os"
	"path/filepath"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var supportedAudioFormats = []string{".mp3", ".wav", ".ogg", ".flac"}
var directory = "/home/teyxos/Documents/raykemon/assets/audio/"

func getAudioFromFile(c chan string, files []os.DirEntry) {
	// Load audio files from the assets folder

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if slices.Contains(supportedAudioFormats, filepath.Ext(file.Name())) {
			filePath := filepath.Join(directory, file.Name())
			c <- filePath
		}
	}
}

var Audios = []rl.Music{}

func LoadAudio() {
	c := make(chan string)

	files, err := os.ReadDir(directory)
	log.Println(files)

	if err != nil {
		log.Fatal(err)
	}
	go getAudioFromFile(c, files)

	// Wait for the audio file to be sent through the channel

	for i := 0; i < len(files); i++ {
		filePath := <-c
		Audios = append(Audios, rl.LoadMusicStream(filePath))
	}
}

func UnloadAudio() {
	for _, audio := range Audios {
		rl.UnloadMusicStream(audio)
	}
}
