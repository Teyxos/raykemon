package lib

import rl "github.com/gen2brain/raylib-go/raylib"

type SpriteSheet struct {
	Width   int32
	Height  int32
	Frames  int32
	Columns int32
	Rows    int32
	Texture rl.Texture2D
}

func NewSpriteSheet(texture rl.Texture2D, frames int32) SpriteSheet {
	columns := frames
	rows := int32(texture.Height / (texture.Width / frames))

	return SpriteSheet{
		Width:   texture.Width,
		Height:  texture.Height,
		Frames:  frames,
		Columns: columns,
		Rows:    rows,
		Texture: texture,
	}
}
func (s *SpriteSheet) GetFrame(frame int32) rl.Rectangle {
	// Calculate the width and height of each frame
	frameWidth := s.Width / s.Columns
	frameHeight := s.Height / s.Rows

	// Calculate the x and y position of the frame in the sprite sheet
	x := (frame % s.Columns) * frameWidth
	y := (frame / s.Columns) * frameHeight

	return rl.Rectangle{
		X:      float32(x),
		Y:      float32(y),
		Width:  float32(frameWidth),
		Height: float32(frameHeight),
	}
}
