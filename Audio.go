package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/teyxos/raykemon/lib"
)

var (
	bgMusic rl.Music
)

func LoadAudio() {
	bgMusic = rl.LoadMusicStream("assets/country.mp3") // Load music file               // Unload music stream

	lib.SetBackgroundMusic(bgMusic)
}

func UnloadAudio() {
	rl.UnloadMusicStream(bgMusic)
}
