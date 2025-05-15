package screens

import (
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/teyxos/raykemon/lib"
)

var currentMenu = MainMenu
var txArray []string
var aArray []string
var runes = make([]rune, 0)

func DrawMenuScreen(movs map[string]*lib.Moveable, tx map[string]rl.Texture2D, a map[string]rl.Music) error {
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.RayWhite)

	currentSprite := movs["player"].SpriteSheet.Name
	// var currentMusic string
	currentSIndex := 0
	// currentAIndex := 0

	for i := 0; i < len(txArray); i++ {
		if strings.Contains(txArray[i], currentSprite) {
			currentSIndex = i
			break
		}
	}

	if len(txArray) == 0 {
		for k := range tx {
			txArray = append(txArray, k)
		}
	}

	if len(aArray) == 0 {
		for k := range a {
			aArray = append(aArray, k)
		}
	}

	switch currentMenu {
	case MainMenu:
		rl.DrawText("Current sprite: "+movs["player"].SpriteSheet.Name, 10, 30, 20, rl.Black)
		rl.DrawText("Press 'F1' to change to sprite", 10, 60, 20, rl.Black)
	case SpriteSelector:
		rl.DrawText("Press 'Enter' to save the new sprite", 10, 30, 20, rl.Black)
		for i := 0; i < len(txArray); i++ {
			if strings.Contains(txArray[i], currentSprite) {
				rl.DrawText("> "+txArray[i], 10, 60+30*int32(i), 20, rl.Black)
			} else {
				rl.DrawText(txArray[i], 10, 60+30*int32(i), 20, rl.Black)
			}
		}

		switch rl.GetKeyPressed() {
		case rl.KeyEnter:
			currentMenu = SpriteSaving
		case rl.KeyW:
			if currentSIndex > 0 {
				currentSIndex--
			}
			movs["player"].SpriteSheet.Name = txArray[currentSIndex]
		case rl.KeyS:
			if currentSIndex < len(txArray)-1 {
				currentSIndex++
			}
			movs["player"].SpriteSheet.Name = txArray[currentSIndex]
		}

		// case MusicSelector:
		// 	rl.DrawText("Press 'Enter' to save the new music", 10, 30, 20, rl.Black)
		// 	rl.TraceLog(rl.LogInfo, "current index: %v", currentAIndex)
		// 	for i := 0; i < len(aArray); i++ {
		// 		if strings.Contains(aArray[i], currentMusic) {
		// 			rl.DrawText("> "+aArray[i], 10, 60+30*int32(i), 20, rl.Black)
		// 		} else {
		// 			rl.DrawText(aArray[i], 10, 60+30*int32(i), 20, rl.Black)
		// 		}
		// 	}

	case SpriteSaving:
		input := string(runes)
		rl.DrawText("Please state the number of frames the sprite animetion have: ", 10, 30, 20, rl.Black)
		rl.DrawText(string(input), 10, 60, 20, rl.Black)

		key := rl.GetCharPressed()
		for key > 0 {
			if key >= 32 && key <= 125 {
				runes = append(runes, rune(key))
				rl.TraceLog(rl.LogInfo, "Runes: %v", runes)
			}
			key = rl.GetCharPressed()
		}
		if rl.IsKeyPressed(rl.KeyBackspace) && len(runes) > 0 {
			runes = runes[:len(runes)-1]
		}
		if rl.IsKeyPressed(rl.KeyEnter) {
			if string(runes) != "" {
				frames, err := strconv.Atoi(string(input))
				if err != nil {
					rl.TraceLog(rl.LogError, "Error converting string to int: %v", err)
					currentMenu = SpriteSelector
				}
				movs["player"].SpriteSheet = lib.NewSpriteSheet(tx[movs["player"].SpriteSheet.Name], int32(frames), movs["player"].SpriteSheet.Name)
				currentMenu = MainMenu
				runes = []rune{}
			}
		}
	}

	switch rl.GetKeyPressed() {
	case rl.KeyF1:
		currentMenu = SpriteSelector
	case rl.KeyF2:
		currentMenu = MusicSelector
	case rl.KeyM:
		currentMenu = MainMenu
	case rl.KeyEscape:
		lib.SetScreen(lib.WorldScreen)
	}

	return nil
}

type Menus int32

const (
	MainMenu Menus = iota
	SpriteSelector
	MusicSelector
	SpriteSaving
)
