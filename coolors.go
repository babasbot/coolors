package coolors

import "image/color"

func Red(r uint32) color.RGBA {
	return color.RGBA{uint8(r >> 8), 0, 0, 255}
}

func Green(g uint32) color.RGBA {
	return color.RGBA{0, uint8(g), 0, 255}
}
