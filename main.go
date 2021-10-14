package main

//go:generate packer --input images --stats

// Possible Game name? Spry

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/jlcarruda/mmo-client/engine"
	"github.com/jlcarruda/mmo-client/engine/asset"
	tilemap "github.com/jlcarruda/mmo-client/engine/tilemaps"
)

func main() {
	pixelgl.Run(runGame)
}

func runGame() {
	mapSize := 100
	tileSize := 16

	cfg := pixelgl.WindowConfig{
		Title: "MMO",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync: true,
		Resizable: true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(false)

	load := asset.NewLoad()
	spritesheet, err := load.Spritesheet("packed.json")
	if err != nil {
		panic(err)
	}
	charSprite, err := spritesheet.Get("char.png")
	if err != nil {
		panic(err)
	}
	
	charPos := win.Bounds().Center()
	
	controller := engine.NewController()
	engine.NewCharacter(charSprite, charPos, 1.0, win, controller)

	tiles := make(tilemap.TileMatrix, mapSize)
	grassSprite, err := spritesheet.Get("grass.png")
	if err != nil {
		panic(err)
	}
	
	for x := range tiles {
		tiles[x] = make([]tilemap.Tile, mapSize)
		for y := range tiles[x] {
			
			tiles[x][y] = tilemap.Tile{
				Type: 0,
				Sprite: grassSprite,
			}
		}
	}
	batch := pixel.NewBatch(&pixel.TrianglesData{}, spritesheet.Picture())
	tmap := tilemap.New(tiles, batch, tileSize)
	tmap.Rebatch()
	
	for !win.JustPressed(pixelgl.KeyEscape) {
		win.Clear(pixel.RGB(0, 0, 0))

		controller.InputHandler()
		
		
		win.SetMatrix(controller.Camera().Mat())
		tmap.DrawSelf(win)
		for _, obj := range engine.GameObjectMap {
			obj.DrawSelf()
		}
		win.SetMatrix(pixel.IM)

		win.Update()
	}
}
