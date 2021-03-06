package engine

import "github.com/faiface/pixel/pixelgl"

type DrawableObject interface {
	DrawSelf(win *pixelgl.Window)
}

type InputHandler interface {
	MoveObject(vec Vector)
	SetObject(obj *GameObject)
	MovementListener()
}

type Vector struct {
	X, Y float64
}
