package lib

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Moveable struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
	Speed  float32
}

func (m *Moveable) MoveUp() {
	m.Y -= m.Speed
}

func (m *Moveable) MoveDown() {
	m.Y += m.Speed
}

func (m *Moveable) MoveLeft() {
	m.X -= m.Speed
}

func (m *Moveable) MoveRight() {
	m.X += m.Speed
}

func (m *Moveable) DrawSelf() {
	rl.DrawRectangle(int32(m.X), int32(m.Y), int32(m.Width), int32(m.Height), rl.Red)
	rl.DrawRectangleLines(int32(m.X), int32(m.Y), int32(m.Width), int32(m.Height), rl.Black)
}
