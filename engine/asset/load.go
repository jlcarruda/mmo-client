package asset

import (
	"encoding/json"
	"image"
	_ "image/png"
	"io/fs"
	"io/ioutil"
	"os"

	"github.com/faiface/pixel"
	"github.com/jstewart7/packer"
)

type Load struct {
	filesystem fs.FS
	imagesPath string
}

func NewLoad() *Load {
	filesystem := os.DirFS(".")
	return &Load{
		filesystem: filesystem,
		imagesPath: "/images",
	}
}

func (load *Load) Open(path string) (fs.File, error) {
	return load.filesystem.Open(path)
}

func (load *Load) Image(path string) (image.Image, error) {
	file, fileErr := load.filesystem.Open(path)
	if fileErr != nil {
		return nil, fileErr
	}
	
	img, _, decodeErr := image.Decode(file)
	if decodeErr != nil {
		return nil, decodeErr
	}

	defer file.Close()
	return img, nil
}

func (load *Load) Sprite(path string) (*pixel.Sprite, error) {
	img, err := load.Image(path)
	if err != nil {
		return nil, err
	}
	
	pic := pixel.PictureDataFromImage(img)
	
	return pixel.NewSprite(pic, pic.Bounds()), nil
}

// dat = aceita qualquer interface
func (load *Load) Json(path string, dat interface{}) error {
	file, err := load.filesystem.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	jsonData, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonData, &dat)
}

// Spritesheet Loader function
func (load *Load) Spritesheet(path string) (*SpriteSheet, error) {
	serializedSpritesheet := packer.SerializedSpritesheet{}
	
	err := load.Json(path, &serializedSpritesheet)
	if err != nil {
		return nil, err
	}

	img, err := load.Image(serializedSpritesheet.ImageName)
	if err != nil {
		return nil, err
	}

	pic := pixel.PictureDataFromImage(img)
	bounds := pic.Bounds()
	lookup := make(map[string]*pixel.Sprite)
	for k, v := range serializedSpritesheet.Frames {
		rect := pixel.R(
			v.Frame.X,
			bounds.H() - v.Frame.Y,
			v.Frame.X + v.Frame.W,
			bounds.H() - (v.Frame.Y + v.Frame.H),
		).Norm()

		lookup[k] = pixel.NewSprite(pic, rect)
	}

	return NewSpriteSheet(pic, lookup), nil
}

