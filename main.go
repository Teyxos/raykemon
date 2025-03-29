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
	defer rl.CloseAudioDevice()                                // Close audio device

	rl.InitAudioDevice()   // Initialize audio device
	defer rl.CloseWindow() // Close window

	// logic to load tiled maps (not working yet)
	// gameMap, err := tiled.LoadFile("map/map.tmx")
	// if err != nil {
	// 	fmt.Printf("error parsing map: %s", err.Error())
	// 	os.Exit(2)
	// }

	// err = lib.ParseTiledMap(*gameMap)

	// if err != nil {
	// 	fmt.Printf("error parsing map: %s", err.Error())
	// 	os.Exit(2)
	// }

	LoadAudio()
	defer UnloadAudio()
	LoadTextures()

	rl.SetTargetFPS(60)

	lib.SetScreen(lib.WorldScreen)

	for !rl.WindowShouldClose() {
		currentScreen := lib.GetScreen()
		currentMusic := lib.GetMusic()

		rl.UpdateMusicStream(currentMusic)

		rl.BeginDrawing()

		// Logic to draw the current screen
		if currentScreen == lib.WorldScreen {
			screens.DrawWorldScreen()
		} else if currentScreen == lib.BattleScreen {
			screens.DrawBattleScreen()
		} else if currentScreen == lib.MenuScreen {
			screens.DrawMenuScreen()
		}

		if rl.IsKeyPressed(rl.KeyF1) {
			lib.SetScreen(lib.WorldScreen)
		}

		if rl.IsKeyPressed(rl.KeyF2) {
			lib.SetScreen(lib.BattleScreen)
		}

		if rl.IsKeyPressed(rl.KeyF3) {
			lib.SetScreen(lib.MenuScreen)
		}

		rl.ClearBackground(rl.Black)

		rl.EndDrawing()
	}
}
