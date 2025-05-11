package lib

import rl "github.com/gen2brain/raylib-go/raylib"

var SpriteSheets = make(map[string]SpriteSheet)

type SpriteSheet struct {
	Width   int32
	Height  int32
	Frames  int32
	Columns int32
	Rows    int32
	Name    string
	Texture rl.Texture2D
}

func NewSpriteSheet(t rl.Texture2D, f int32, n string) SpriteSheet {
	columns := f
	rows := int32(t.Height / (t.Width / f))

	return SpriteSheet{
		Width:   t.Width,
		Height:  t.Height,
		Frames:  f,
		Columns: columns,
		Rows:    rows,
		Name:    n,
		Texture: t,
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
