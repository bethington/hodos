package ebiten

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"nostos/common/enum"
)

type surfaceState struct {
	x              int
	y              int
	filter         ebiten.Filter
	color          color.Color
	brightness     float64
	saturation     float64
	effect         enum.DrawEffect
	skewX, skewY   float64
	scaleX, scaleY float64
}
