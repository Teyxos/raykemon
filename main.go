package main

import (
	"strconv"
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

	rl.SetTargetFPS(0) // This will be a settings option later
	rl.SetExitKey(0)   // So user doesnt accidently close the window

	lib.SetScreen(lib.WorldScreen)

	moveables := make(map[string]*lib.Moveable)

	parser, err := lib.NewLineParser("assets/resources/raykemon.txt")
	if err != nil {
		rl.TraceLog(rl.LogError, "Failed to create line parser: %v", err)
		return
	}
	defer parser.Close()
	// Read lines from the file
	for line, ok := parser.Next(); ok; line, ok = parser.Next() {
		words := strings.Fields(line)
		// var action lib.ParserAction

		switch words[0] {
		case "Player":
			frames, err := strconv.Atoi(words[2])
			if err != nil {
				rl.TraceLog(rl.LogError, "Failed to convert string to int: %v", err)
				return
			}

			width, _ := strconv.Atoi(words[3])
			height, _ := strconv.Atoi(words[4])

			moveables["player"] = lib.MoveableFromTexture(100, 100, int32(width), int32(height), 2.5, Textures[strings.Join([]string{"assets/textures/", words[1]}, "")], int32(frames), words[1])
			lib.SpriteSheets[strings.Join([]string{"assets/textures/", words[1]}, "")] = lib.NewSpriteSheet(Textures[strings.Join([]string{"assets/textures/", words[1]}, "")], int32(frames), words[1])
		case "BGMusic":
			lib.SetBackgroundMusic(Audios[strings.Join([]string{"assets/audio/", words[1]}, "")])
		case "SpriteSheet":
			frames, err := strconv.Atoi(words[2])
			if err != nil {
				rl.TraceLog(rl.LogError, "Failed to convert string to int: %v", err)
				return
			}
			lib.SpriteSheets[strings.Join([]string{"assets/textures/", words[1]}, "")] = lib.NewSpriteSheet(Textures[strings.Join([]string{"assets/textures/", words[1]}, "")], int32(frames), words[1])
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
			screens.DrawWorldScreen(moveables, dt)

			moveables["player"].DrawSelf()

		} else if currentScreen == lib.BattleScreen {
			screens.DrawBattleScreen()
		} else if currentScreen == lib.MenuScreen {
			screens.DrawMenuScreen(moveables, Textures, Audios)
		}

		switch rl.GetKeyPressed() {
		case rl.KeyF1:
			lib.SetScreen(lib.WorldScreen)
		case rl.KeyF2:
			lib.SetScreen(lib.BattleScreen)
		case rl.KeyM:
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

		UI()

		rl.ClearBackground(rl.Black)
		rl.EndDrawing()
	}
}

func UI() {
	rl.DrawText("FPS: "+strconv.Itoa(int(rl.GetFPS())), 10, 10, 20, rl.Black)
}
