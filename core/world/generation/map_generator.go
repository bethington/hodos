package generation

import (
	d2asset "nostos/core/asset"

	"nostos/common/enum"
	d2geom "nostos/common/geometry"
	d2util "nostos/common/util"
	d2mapengine "nostos/core/world/engine"
	d2mapstamp "nostos/core/world/stamp"
)

const (
	logPrefix = "Map Generator"
)

// NewMapGenerator creates a map generator instance
func NewMapGenerator(a *d2asset.AssetManager, l d2util.LogLevel, e *d2mapengine.MapEngine) (*MapGenerator, error) {
	generator := &MapGenerator{
		asset:  a,
		engine: e,
	}

	generator.Logger = d2util.NewLogger()
	generator.Logger.SetLevel(l)
	generator.Logger.SetPrefix(logPrefix)

	return generator, nil
}

// MapGenerator generates maps for the map engine
type MapGenerator struct {
	asset  *d2asset.AssetManager
	engine *d2mapengine.MapEngine

	*d2util.Logger
}

func (g *MapGenerator) loadPreset(id, index int) *d2mapstamp.Stamp {
	for _, file := range g.asset.Records.LevelPreset(id).Files {
		g.engine.AddDS1(file)
	}

	return g.engine.LoadStamp(enum.RegionAct1Wilderness, id, index)
}

func areaEmpty(mapEngine *d2mapengine.MapEngine, rect d2geom.Rectangle) bool {
	mapHeight := mapEngine.Size().Height
	mapWidth := mapEngine.Size().Width

	if rect.Top < 0 || rect.Left < 0 || rect.Bottom() >= mapHeight || rect.Right() >= mapWidth {
		return false
	}

	for y := rect.Top; y <= rect.Bottom(); y++ {
		for x := rect.Left; x <= rect.Right(); x++ {
			if len(mapEngine.Tile(x, y).Components.Floors) == 0 {
				continue
			}

			floor := mapEngine.Tile(x, y).Components.Floors[0]

			if floor.Style != 0 || floor.Sequence != 0 || floor.Prop1 != 1 {
				return false
			}
		}
	}

	return true
}
