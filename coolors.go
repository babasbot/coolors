package coolors

import "image/color"

func Red(r uint32) color.RGBA {
	return color.RGBA{uint8(r >> 8), 0, 0, 255}
}

func Green(g uint32) color.RGBA {
	return color.RGBA{0, uint8(g), 0, 255}
}

func Blue(b uint32) color.RGBA {
	return color.RGBA{0, 0, uint8(b >> 8), 255}
}

func Color(r, g, b uint32, col color.RGBA) color.RGBA {
	pixel := color.RGBA{}

	pixel.R = uint8(r>>8) & col.R
	pixel.G = uint8(g>>8) & col.G
	pixel.B = uint8(b>>8) & col.B
	pixel.A = 255

	return pixel
}
