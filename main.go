package main

import (
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/teyxos/raykemon/lib"
	"github.com/teyxos/raykemon/screens"
)

var ScreenWidth = int32(800)
var ScreenHeight = int32(450)

// TODO: Add functionallity to change the animation depeding on the direction without the need for user input just go to the row the specify when creating a texture

func main() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "Raykemon Tests") // Initialize window
	defer rl.CloseWindow()                                     // Close window

	rl.InitAudioDevice()        // Initialize audio device
	defer rl.CloseAudioDevice() // Close audio device

	LoadAudio()         // Load audio files
	defer UnloadAudio() // Unload audio files

	if len(Audios) == 0 {
		rl.TraceLog(rl.LogWarning, "No audio files found in the directory.")
	}

	if len(Textures) == 0 {
		rl.TraceLog(rl.LogWarning, "No texture files found in the directory.")
	}

	LoadTextures()         // Load textures
	defer UnloadTextures() // Unload textures

	var monitor = rl.GetCurrentMonitor()

	rl.SetTargetFPS(60) // This will be a settings option later
	rl.SetExitKey(0)    // So user doesnt accidently close the window

	lib.SetScreen(lib.WorldScreen)

	moveables := make(map[string]*lib.Moveable)

	parser, err := lib.NewLineParser("assets/resources/example.txt")
	if err != nil {
		rl.TraceLog(rl.LogError, "Failed to create line parser: %v", err)
		return
	}
	defer parser.Close()
	// Read lines from the file
	for line, ok := parser.Next(); ok; line, ok = parser.Next() {
		words := parser.SplitLine(line)
		var action lib.ParserAction

		for _, word := range words {
			if word == "Player" {
				action = lib.PlayerAction
				continue
			}

			if word == "BGMusic" {
				action = lib.BGMusicAction
				continue
			}

			if action == lib.PlayerAction {
				moveables["player"] = lib.MoveableFromTexture(100, 100, 2.5, Textures[strings.Join([]string{"assets/textures/", word}, "")], 3)
			}

			if action == lib.BGMusicAction {
				lib.SetBackgroundMusic(Audios[strings.Join([]string{"assets/audio/", word}, "")])
			}
		}
	}

	// Create a moveable object

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime() * 100 // Get delta time to make it be 1 speed = 100px per second
		currentScreen := lib.GetScreen()
		currentMusic := lib.GetMusic()

		rl.UpdateMusicStream(currentMusic)

		rl.BeginDrawing()

		// Logic to draw the current screen
		if currentScreen == lib.WorldScreen {

			if rl.IsKeyDown(rl.KeyUp) {
				moveables["player"].MoveUp(dt)
				moveables["player"].Direction = lib.DirectionUp
			}
			if rl.IsKeyDown(rl.KeyDown) {
				moveables["player"].MoveDown(dt)
				moveables["player"].Direction = lib.DirectionDown
			}
			if rl.IsKeyDown(rl.KeyLeft) {
				moveables["player"].MoveLeft(dt)
				moveables["player"].Direction = lib.DirectionLeft
			}
			if rl.IsKeyDown(rl.KeyRight) {
				moveables["player"].MoveRight(dt)
				moveables["player"].Direction = lib.DirectionRight
			}

			screens.DrawWorldScreen()

			moveables["player"].DrawSelf(7)

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
			// Test if this works on not wayland compositors
			if rl.IsWindowFullscreen() {
				rl.SetWindowSize(int(ScreenWidth), int(ScreenHeight))
			} else {
				width := rl.GetMonitorPhysicalWidth(monitor)
				height := rl.GetMonitorPhysicalHeight(monitor)
				rl.SetWindowSize(width, height)
			}
			rl.ToggleFullscreen()
		}

		rl.ClearBackground(rl.Black)

		rl.EndDrawing()
	}
}
