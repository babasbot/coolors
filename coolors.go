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

func Bright(r, g, b uint32, brightness int32) color.RGBA {
	normalize := func(n int32) uint8 {
		if n > 255 {
			return 255
		} else if n < 0 {
			return 0
		}

		return uint8(n)
	}

	pixel := color.RGBA{}

	pixel.R = normalize(int32(r>>8) + brightness)
	pixel.G = normalize(int32(g>>8) + brightness)
	pixel.B = normalize(int32(b>>8) + brightness)
	pixel.A = 255

	return pixel
}
