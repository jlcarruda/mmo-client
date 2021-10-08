package engine

import (
	"github.com/faiface/pixel/pixelgl"
)

var MOVE_SPEED float64 = 2.0

type Controller struct {
	obj *GameObject
	win *pixelgl.Window
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) SetObject(o *GameObject) {
	c.obj = o
	c.win = o.win
}

func (c *Controller) MoveObject(vec Vector) {
	c.obj.Move(vec, MOVE_SPEED)
}

func (c *Controller) InputHandler() {
	vec := Vector{0,0}

	if c.win.Pressed(pixelgl.KeyLeft) {
		vec.X = -1
	}
	if c.win.Pressed(pixelgl.KeyRight) {
		vec.X = 1
	}
	if c.win.Pressed(pixelgl.KeyDown) {
		vec.Y = -1
	}
	if c.win.Pressed(pixelgl.KeyUp) {
		vec.Y = 1
	}

	c.MoveObject(vec)
}
