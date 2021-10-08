package asset

import (
	"errors"

	"github.com/faiface/pixel"
)

type SpriteSheet struct {
	picture pixel.Picture
	lookup map[string]*pixel.Sprite
}

func NewSpriteSheet(picture pixel.Picture, lookup map[string]*pixel.Sprite) *SpriteSheet {
	return &SpriteSheet{
		picture,
		lookup,
	}
}

func (s *SpriteSheet) Get(name string) (*pixel.Sprite , error){
	sprite, ok := s.lookup[name]
	if !ok {
		return nil, errors.New("invalid sprite name")
	}

	return sprite, nil
}

func (s *SpriteSheet) Picture() pixel.Picture {
	return s.picture
}
