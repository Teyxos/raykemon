package main

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/teyxos/raykemon/lib"
	"github.com/teyxos/raykemon/screens"
)

var ScreenWidth = int32(800)
var ScreenHeight = int32(450)

func main() {

	rl.InitWindow(ScreenWidth, ScreenHeight, "Raykemon Tests") // Initialize window
	defer rl.CloseWindow()                                     // Close window

	rl.InitAudioDevice()        // Initialize audio device
	defer rl.CloseAudioDevice() // Close audio device

	LoadAudio()         // Load audio files
	defer UnloadAudio() // Unload audio files

	if len(Audios) == 0 {
		rl.TraceLog(rl.LogWarning, "No audio files found in the directory.")
	} else {
		lib.SetBackgroundMusic(Audios["assets/audio/country.mp3"]) // Set background music
	}

	if len(Textures) == 0 {
		rl.TraceLog(rl.LogWarning, "No texture files found in the directory.")
	} else {

	}

	LoadTextures()         // Load textures
	defer UnloadTextures() // Unload textures

	rl.SetTargetFPS(60)

	lib.SetScreen(lib.WorldScreen)

	// Create a moveable object
	moveable := lib.MoveableFromTexture(100, 100, 5, Textures["assets/textures/character.png"])
	log.Println(moveable)

	for !rl.WindowShouldClose() {
		currentScreen := lib.GetScreen()
		currentMusic := lib.GetMusic()

		rl.UpdateMusicStream(currentMusic)

		rl.BeginDrawing()

		// Logic to draw the current screen
		if currentScreen == lib.WorldScreen {

			if rl.IsKeyDown(rl.KeyUp) {
				moveable.MoveUp()
			}
			if rl.IsKeyDown(rl.KeyDown) {
				moveable.MoveDown()
			}
			if rl.IsKeyDown(rl.KeyLeft) {
				moveable.MoveLeft()
			}
			if rl.IsKeyDown(rl.KeyRight) {
				moveable.MoveRight()
			}

			screens.DrawWorldScreen()

			moveable.DrawSelf()

		} else if currentScreen == lib.BattleScreen {
			screens.DrawBattleScreen()
		} else if currentScreen == lib.MenuScreen {
			screens.DrawMenuScreen()
		}

		switch rl.GetKeyPressed() {
		case rl.KeyF1:
			lib.SetScreen(lib.WorldScreen)
		case rl.KeyF2:
			lib.SetScreen(lib.BattleScreen)
		case rl.KeyF3:
			lib.SetScreen(lib.MenuScreen)
		case rl.KeyF11:
			rl.ToggleFullscreen()
		}

		rl.ClearBackground(rl.Black)

		rl.EndDrawing()
	}
}
