// package lib

// import (
// 	"errors"

// 	tl "github.com/lafriks/go-tiled"
// )

// func ParseTiledMap(m tl.Map) error {
// 	var xs, xe, xi, ys, ye, yi int
// 	if m.RenderOrder == "" || m.RenderOrder == "right-down" {
// 		xs = 0
// 		xe = m.Width
// 		xi = 1
// 		ys = 0
// 		ye = m.Height
// 		yi = 1
// 	} else {
// 		return errors.New("unsupported render order")
// 	}

// 	i := 0
// 	for y := ys; y*yi < ye; y = y + yi {
// 		for x := xs; x*xi < xe; x = x + xi {
// 			println("layer", i, "x", x, "y", y)

// 			i++
// 		}
// 	}

//		return nil
//	}
package lib

// File under maintenance
