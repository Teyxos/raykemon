package screens

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/teyxos/raykemon/lib"
)

func DrawWorldScreen(movs map[string]*lib.Moveable, dt float32) error {
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.Green)

	if rl.IsKeyDown(rl.KeyUp) {
		movs["player"].MoveUp(dt)
		movs["player"].Direction = lib.DirectionUp
	}
	if rl.IsKeyDown(rl.KeyDown) {
		movs["player"].MoveDown(dt)
		movs["player"].Direction = lib.DirectionDown
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		movs["player"].MoveLeft(dt)
		movs["player"].Direction = lib.DirectionLeft
	}
	if rl.IsKeyDown(rl.KeyRight) {
		movs["player"].MoveRight(dt)
		movs["player"].Direction = lib.DirectionRight
	}

	return nil
}
