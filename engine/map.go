package engine

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/jlcarruda/mmo-client/engine/tilemaps"
)

type Map struct {
	tmap tilemaps.Tilemap
}

func New(tmap tilemaps.Tilemap) *Map {
	return &Map{
		tmap: tmap,
	}
}

func (m *Map) DrawSelf(win *pixelgl.Window) {
	m.tmap.DrawSelf(win)
}

