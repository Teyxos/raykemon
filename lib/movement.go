package lib

import (
	"reflect"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Moveable struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
	Speed  float32
	Sprite rl.Texture2D // Optional: Add a sprite for the moveable object
}

// MoveableFromTexture creates a new Moveable object from a texture
// It sets the width and height of the Moveable object to the width and height of the texture
func MoveableFromTexture(x float32, y float32, speed float32, texture rl.Texture2D) Moveable {
	return Moveable{
		X:      x,
		Y:      y,
		Width:  float32(texture.Width),
		Height: float32(texture.Height),
		Speed:  speed,
		Sprite: texture,
	}
}

// Moves the object up
func (m *Moveable) MoveUp() {
	// Return if the object is already at the top of the screen
	if m.Y-m.Speed < 0 {
		return
	}

	m.Y -= m.Speed
}

// Moves the object down
func (m *Moveable) MoveDown() {
	// Return if the object is already at the bottom of the screen
	if m.Y+m.Height+m.Speed > float32(rl.GetScreenHeight()) {
		return
	}

	m.Y += m.Speed
}

// Moves the object left
func (m *Moveable) MoveLeft() {
	// Return if the object is already at the left of the screen
	if m.X-m.Speed < 0 {
		return
	}
	m.X -= m.Speed
}

// Moves the object right
func (m *Moveable) MoveRight() {
	// Return if the object is already at the right of the screen
	if m.X+m.Width+m.Speed > float32(rl.GetScreenWidth()) {
		return
	}
	m.X += m.Speed
}

// DrawSelf draws the moveable object on the screen
// If the moveable has a sprite, it draws the sprite
// Otherwise, it draws a rectangle with the moveable's width and height
// The rectangle is filled with red color and has a black outline
// The X and Y coordinates are used to position the rectangle on the screen
func (m *Moveable) DrawSelf() {
	// Check if the moveable has a sprite
	if reflect.TypeOf(m.Sprite) == reflect.TypeOf(rl.Texture2D{}) {
		rl.DrawTexture(m.Sprite, int32(m.X), int32(m.Y), rl.White)
	} else {
		rl.DrawRectangle(int32(m.X), int32(m.Y), int32(m.Width), int32(m.Height), rl.Red)
		rl.DrawRectangleLines(int32(m.X), int32(m.Y), int32(m.Width), int32(m.Height), rl.Black)
	}
}
