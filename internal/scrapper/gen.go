package scrapper

import (
	"strconv"
	"strings"
)

func floatFromPtsHeader(s string) float32 {
	s = strings.TrimSpace(s)
	bef, _, f := strings.Cut(s, " ")
	if f {
		v, err := strconv.ParseFloat(bef, 32)
		if err == nil {
			return float32(v)
		}
	}
	return -1.0
}

func valueFromImages(imageVal string) int8 {
	switch imageVal {
	case "/gfx/picto_standard.svg":
		return 5
	case "/gfx/picto_safety.svg":
		return 4
	case "/gfx/picto_not-safety.svg":
		return 3
	case "/gfx/picto_not-available.svg":
		return 2
	case "/gfx/picto_not-applicable.svg":
		return 1
	default:
		return -1
	}
}
