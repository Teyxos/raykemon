package screens

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawWorldScreen() error {
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.Green)
	rl.DrawText("World Screen", 10, 10, 20, rl.Black)

	return nil
}
