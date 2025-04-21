package lib

import rl "github.com/gen2brain/raylib-go/raylib"

var currentAudio rl.Sound
var backgroundMusic rl.Music

func SetSound(s rl.Sound) {
	currentAudio = s
}

func PlaySound() {
	rl.PlaySound(currentAudio)
}

func StopSound() {
	rl.StopSound(currentAudio)
}

func SetVolume(volume float32) {
	rl.SetSoundVolume(currentAudio, volume)
}

func GetMusic() rl.Music {
	return backgroundMusic
}

func SetBackgroundMusic(s rl.Music) {
	backgroundMusic = s
	rl.SetMusicVolume(backgroundMusic, 0.5)
	rl.PlayMusicStream(backgroundMusic)
}
