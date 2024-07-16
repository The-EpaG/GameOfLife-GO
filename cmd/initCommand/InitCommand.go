package initCommand

import (
	"image"
	"image/png"
	"os"

	"flag"

	"github.com/The-EpaG/GameOfLife-GO/internal/constants"
	"github.com/The-EpaG/GameOfLife-GO/internal/errors"
)

var width *int
var height *int

func Flags() {
	width = flag.Int("width", constants.Width, "width")
	height = flag.Int("height", constants.Height, "height")
}

func colorImage(image *image.Gray) *image.Gray {
	for i := 0; i < *width**height; i += 1 {
		image.SetGray(
			i%*width,
			i / *width,
			constants.Dead,
		)
	}

	return image
}

func createImage() error {
	var upLeft image.Point = image.Point{0, 0}
	var lowRight image.Point = image.Point{*width, *height}

	image := image.NewGray(image.Rectangle{upLeft, lowRight})
	image = colorImage(image)

	if err := os.RemoveAll(constants.InputFolder); err != nil {
		return err
	}

	if err := os.Mkdir(constants.InputFolder, os.ModePerm); err != nil {
		return err
	}

	photo, err := os.Create(constants.InputPath)
	if err != nil {
		return err
	}

	png.Encode(photo, image)

	return nil
}

func InitCommand() error {
	if *width <= 0 || *height <= 0 {
		return &errors.ParamError{}
	}
	err := createImage()
	return err
}
