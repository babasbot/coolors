package coolors

import "image/color"

func Red(r uint32) color.RGBA {
	return color.RGBA{uint8(r >> 8), 0, 0, 255}
}
