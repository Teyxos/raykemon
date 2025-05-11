package screens

import (
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/teyxos/raykemon/lib"
)

var currentMenu = MainMenu
var txArray []string

func DrawMenuScreen(movs map[string]*lib.Moveable, tx map[string]rl.Texture2D) error {
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.RayWhite)

	currentSprite := movs["player"].SpriteSheet.Name
	currentIndex := 0

	for i := 0; i < len(txArray); i++ {
		if strings.Contains(txArray[i], currentSprite) {
			currentIndex = i
			break
		}
	}

	if len(txArray) == 0 {
		for k := range tx {
			txArray = append(txArray, k)
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
	}

	switch rl.GetKeyPressed() {
	case rl.KeyF1:
		currentMenu = SpriteSelector
	case rl.KeyM:
		currentMenu = MainMenu
	case rl.KeyEscape:
		lib.SetScreen(lib.WorldScreen)
	case rl.KeyW:
		if currentIndex > 0 {
			currentIndex--
		}
		movs["player"].SpriteSheet.Name = txArray[currentIndex]
	case rl.KeyS:
		if currentIndex < len(txArray)-1 {
			currentIndex++
		}
		movs["player"].SpriteSheet.Name = txArray[currentIndex]
	case rl.KeyEnter:
		if currentMenu == SpriteSelector {
			movs["player"].SpriteSheet = lib.NewSpriteSheet(tx[movs["player"].SpriteSheet.Name], lib.SpriteSheets[txArray[currentIndex]].Frames, movs["player"].SpriteSheet.Name)
			currentMenu = MainMenu
		}
	}

	return nil
}

type Menus int32

const (
	MainMenu Menus = iota
	SpriteSelector
	MusicSelector
)
