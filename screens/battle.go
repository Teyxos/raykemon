package screens

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawBattleScreen() error {
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.SkyBlue)

	return nil
}
