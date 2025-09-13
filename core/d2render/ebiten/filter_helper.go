package ebiten

import (
	"github.com/hajimehoshi/ebiten/v2"

	"nostos/common/enum"
)

func d2ToEbitenFilter(filter enum.Filter) ebiten.Filter {
	switch filter {
	case enum.FilterNearest:
		return ebiten.FilterNearest
	default:
		return ebiten.FilterLinear
	}
}

// func ebitenToD2Filter(filter ebiten.Filter) enum.Filter {
// 	switch filter {
// 	case ebiten.FilterDefault:
// 		return enum.FilterDefault
// 	case ebiten.FilterLinear:
// 		return enum.FilterLinear
// 	case ebiten.FilterNearest:
// 		return enum.FilterNearest
// 	}
//
// 	return enum.FilterDefault
// }
