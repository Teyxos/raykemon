package lib

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Moveable struct {
	X           float32
	Y           float32
	Width       int32
	Height      int32
	Speed       float32
	Direction   Direction
	SpriteSheet SpriteSheet // Optional: Add a sprite for the moveable object
}

type Direction int32

const (
	DirectionUp Direction = iota
	DirectionDown
	DirectionLeft
	DirectionRight
)

// MoveableFromTexture creates a new Moveable object from a texture
// It sets the width and height of the Moveable object to the width and height of the texture
func MoveableFromTexture(x float32, y float32, speed float32, texture rl.Texture2D, frames int32) Moveable {
	return Moveable{
		X:           x,
		Y:           y,
		Width:       texture.Width,
		Height:      texture.Height,
		Speed:       speed,
		Direction:   DirectionUp,
		SpriteSheet: NewSpriteSheet(texture, frames),
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
	if int(m.Y)+int(m.Height)+int(m.Speed) > rl.GetScreenHeight() {
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
	if int(m.X)+int(m.Width)+int(m.Speed) > rl.GetScreenWidth() {
		return
	}
	m.X += m.Speed * dt
}

// DrawSelf draws the moveable object on the screen
// If the moveable has a sprite, it draws the sprite
// Otherwise, it draws a rectangle with the moveable's width and height
// The rectangle is filled with red color and has a black outline
// The X and Y coordinates are used to position the rectangle on the screen
func (m *Moveable) DrawSelf(firstFrame int32) {
	// Check if the moveable has a sprite
	if m.SpriteSheet.Texture.ID > 0 {
		// Draw the current frame of the sprite sheet

		// TODO: Make so the direction changes using m.Direction to match the spritesheet, depeding on the rows
		frame := m.SpriteSheet.GetFrame(int32(rl.GetTime()*10)%m.SpriteSheet.Frames + firstFrame - 1)
		rl.DrawTexturePro(
			m.SpriteSheet.Texture,
			frame,
			rl.Rectangle{X: m.X, Y: m.Y, Width: float32(m.Width), Height: float32(m.Height)},
			rl.Vector2{},
			0,
			rl.White,
		)

	} else {
		// Draw a red rectangle with a black outline
		rl.DrawRectangle(int32(m.X), int32(m.Y), int32(m.Width), int32(m.Height), rl.Red)
		rl.DrawRectangleLines(int32(m.X), int32(m.Y), int32(m.Width), int32(m.Height), rl.Black)
	}
}
