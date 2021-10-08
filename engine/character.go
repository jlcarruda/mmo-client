package engine

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Character struct {
	controller InputHandler
	obj *GameObject
}

func NewCharacter(sprite *pixel.Sprite, startPos pixel.Vec, scale float64, win *pixelgl.Window, controller InputHandler) *Character {
	obj := NewGameObject(sprite, startPos, scale, win, true)
	controller.SetObject(obj)
	return &Character{
		controller,
		obj,
	}
}

func (c *Character) Move(vec Vector){
	c.controller.MoveObject(vec)
}

func (c *Character) Render() {
	c.obj.DrawSelf()
}
