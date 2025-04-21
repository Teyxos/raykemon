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
func (m *Moveable) MoveUp(dt float32) {
	// Return if the object is already at the top of the screen
	if m.Y-m.Speed < 0 {
		return
	}

	m.Y -= m.Speed * dt
}

// Moves the object down
func (m *Moveable) MoveDown(dt float32) {
	// Return if the object is already at the bottom of the screen
	if m.Y+m.Height+m.Speed > float32(rl.GetScreenHeight()) {
		return
	}

	m.Y += m.Speed * dt
}

// Moves the object left
func (m *Moveable) MoveLeft(dt float32) {
	// Return if the object is already at the left of the screen
	if m.X-m.Speed < 0 {
		return
	}
	m.X -= m.Speed * dt
}

// Moves the object right
func (m *Moveable) MoveRight(dt float32) {
	// Return if the object is already at the right of the screen
	if m.X+m.Width+m.Speed > float32(rl.GetScreenWidth()) {
		return
	}
	m.X += m.Speed * dt
}

// DrawSelf draws the moveable object on the screen
// If the moveable has a sprite, it draws the sprite
// Otherwise, it draws a rectangle with the moveable's width and height
// The rectangle is filled with red color and has a black outline
// The X and Y coordinates are used to position the rectangle on the screen
func (m *Moveable) DrawSelf() {
	// Check if the moveable has a sprite
	if reflect.TypeOf(m.Sprite) == reflect.TypeOf(rl.Texture2D{}) {
		// TODO: Pass to the right each frame to make animations
		rl.DrawTexturePro(m.Sprite, rl.Rectangle{X: 0 + (24 * 0) /* per frame */, Y: 24, Width: 24, Height: 24}, rl.Rectangle{X: m.X, Y: m.Y, Width: m.Width, Height: m.Height}, rl.Vector2{}, 0, rl.White)
	} else {
		rl.DrawRectangle(int32(m.X), int32(m.Y), int32(m.Width), int32(m.Height), rl.Red)
		rl.DrawRectangleLines(int32(m.X), int32(m.Y), int32(m.Width), int32(m.Height), rl.Black)
	}
}
