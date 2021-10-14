package render

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Camera struct {
	win *pixelgl.Window
	Position pixel.Vec
	ZoomValue float64
	ZoomSpeed float64
	mat pixel.Matrix
}

func NewCamera(win *pixelgl.Window, x, y float64) *Camera {
	return &Camera{
		win: win,
		Position: pixel.V(x, y),
		ZoomValue: 1.0,
		ZoomSpeed: 0.1,
		mat: pixel.IM,
	}
}

func (c *Camera) Update() {
	screenCenter := c.win.Bounds().Center()

	movePos := pixel.V(
		math.Floor(-c.Position.X),
		math.Floor(-c.Position.Y),
	).Add(screenCenter)

	c.mat = pixel.IM.Moved(movePos).Scaled(screenCenter, c.ZoomValue)
}

func (c *Camera) Mat() pixel.Matrix {
	return c.mat
}

func (c *Camera) Zoom(z float64) {
	c.ZoomValue += z * c.ZoomSpeed
}
