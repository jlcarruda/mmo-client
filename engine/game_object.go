package engine

import (
	"errors"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	uuid "github.com/nu7hatch/gouuid"
)

var GameObjectMap map[string]*GameObject = make(map[string]*GameObject)

type GameObject struct {
	id *uuid.UUID
	position pixel.Vec
	sprite *pixel.Sprite
	scale float64
	win *pixelgl.Window
	movable bool
}

func NewGameObject(sprite *pixel.Sprite, position pixel.Vec, scale float64, win *pixelgl.Window, movable bool) *GameObject {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	obj := &GameObject{
		id,
		position,
		sprite,
		scale,
		win,
		movable,
	}

	GameObjectMap[id.String()] = obj

	return obj
}

func (obj *GameObject) Remove() {
	_, ok := GameObjectMap[obj.id.String()]
	if ok {
		delete(GameObjectMap, obj.id.String())
	}
}

func (obj *GameObject) SetSprite(sprite *pixel.Sprite) {
	obj.sprite = sprite
}

func (obj *GameObject) DrawSelf() {
	obj.sprite.Draw(obj.win, pixel.IM.Scaled(pixel.ZV, obj.scale).Moved(obj.position))
}

func (obj *GameObject) Move(mVector Vector, speed float64) pixel.Vec {
	if !obj.movable {
		return obj.position
	}
	if mVector.X > 1 || mVector.X < -1 || mVector.Y > 1 || mVector.Y < -1 {
		panic(errors.New("movement vector is not a identity vector"))
	}

	obj.position.X += mVector.X * speed
	obj.position.Y += mVector.Y * speed
	return obj.position
}
