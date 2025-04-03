package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/teyxos/raykemon/lib"
	"github.com/teyxos/raykemon/screens"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "Raykemon Tests") // Initialize window
	defer rl.CloseWindow()                                     // Close window

	rl.InitAudioDevice()        // Initialize audio device
	defer rl.CloseAudioDevice() // Close audio device

	LoadAudio()         // Load audio files
	defer UnloadAudio() // Unload audio files

	lib.SetBackgroundMusic(Audios[0]) // Set background music

	LoadTextures()         // Load textures
	defer UnloadTextures() // Unload textures

	rl.SetTargetFPS(60)

	lib.SetScreen(lib.WorldScreen)

	// Create a moveable object
	moveable := lib.Moveable{X: 100, Y: 100, Width: 50, Height: 50, Speed: 5}

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
		}

		rl.ClearBackground(rl.Black)

		rl.EndDrawing()
	}
}
