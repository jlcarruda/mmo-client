package main

//go:generate packer --input images --stats

// Possible Game name? Spry

import (
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/jlcarruda/mmo/engine"
	"github.com/jlcarruda/mmo/engine/asset"
)

func main() {
	pixelgl.Run(runGame)
}

func runGame() {
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

	load := asset.NewLoad(os.DirFS("./images"))
	charSprite, err := load.Sprite("char.png")
	if err != nil {
		panic(err)
	}
	
	charPos := win.Bounds().Center()
	
	controller := engine.NewController()
	engine.NewCharacter(charSprite, charPos, 2.0, win, controller)

	for !win.JustPressed(pixelgl.KeyEscape) {
		win.Clear(pixel.RGB(0, 0, 0))

		controller.MovementListener()
		
		for _, obj := range engine.GameObjectMap {
			obj.DrawSelf()
		}

		win.Update()
	}
}
