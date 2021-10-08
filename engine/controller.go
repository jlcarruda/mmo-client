package engine

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/jlcarruda/mmo-client/engine/render"
)

var MOVE_SPEED float64 = 2.0

type Controller struct {
	obj *GameObject
	win *pixelgl.Window
	camera *render.Camera
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) SetObject(o *GameObject) {
	c.obj = o
	c.win = o.win
	c.camera = render.NewCamera(c.win, 0, 0)
}

func (c *Controller) MoveObject(vec Vector) {
	c.obj.Move(vec, MOVE_SPEED)
}

func (c *Controller) InputHandler() {
	
	scroll := c.win.MouseScroll()

	if scroll.Y != 0 {
		c.camera.Zoom(scroll.Y)
	}

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
	c.MoveCamera()
}

func (c *Controller) MoveCamera() {
	c.camera.Position = c.obj.Position()
	c.camera.Update()
}

func (c *Controller) Object() *GameObject {
	return c.obj
}

func (c *Controller) Camera() *render.Camera {
	return c.camera
}
