package tilemap

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type TyleType uint8

type Tile struct {
	Type TyleType
	Sprite *pixel.Sprite
}

type TileMatrix [][]Tile

type Tilemap struct {
	TileSize int
	Tiles TileMatrix
	Batch *pixel.Batch
}

func New(tiles TileMatrix, Batch *pixel.Batch, tileSize int) *Tilemap {
	return &Tilemap{
		Batch: Batch,
		TileSize: tileSize,
		Tiles: tiles,
	}
}

func (t *Tilemap) Rebatch() {
	for x := range t.Tiles {
		for y, tile := range t.Tiles[x] {
			pos := pixel.V(
				float64(x * t.TileSize),
				float64(y * t.TileSize),
			)

			mat := pixel.IM.Moved(pos)
			tile.Sprite.Draw(t.Batch, mat)
		}
	}
}

func (t *Tilemap) DrawSelf(win *pixelgl.Window) {
	t.Batch.Draw(win)
}
