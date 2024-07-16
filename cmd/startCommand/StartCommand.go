package startCommand

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"

	Board "github.com/The-EpaG/GameOfLife-GO/internal/classes/board"
	"github.com/The-EpaG/GameOfLife-GO/internal/constants"
)

var step *int

func Flags() {
	step = flag.Int("generations", 1, "Number of generations to generate")
}

func rgbaToGray(img image.Image) *image.Gray {
	var (
		bounds = img.Bounds()
		gray   = image.NewGray(bounds)
	)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var rgba = img.At(x, y)
			gray.Set(x, y, rgba)
		}
	}
	return gray
}

func saveImage(image *image.Gray, index int) error {
	path := fmt.Sprintf(constants.OutputPath, index)
	photo, err := os.Create(path)
	if err != nil {
		return err
	}

	err = png.Encode(photo, image)

	return err
}

func getFirstImage() (*image.Gray, int, error) {

	if err := os.RemoveAll(constants.OutputFolder); err != nil {
		return nil, -1, err
	}

	if err := os.Mkdir(constants.OutputFolder, os.ModePerm); err != nil {
		return nil, -1, err
	}

	file, err := os.Open(constants.InputPath)
	if err != nil {
		return nil, -1, err
	}
	defer file.Close()

	img, imgType, err := image.Decode(file)
	if err != nil || imgType != "png" {
		return nil, -1, err
	}

	gr := rgbaToGray(img)

	return gr, 0, nil
}

func StartCommand() error {
	file, index, err := getFirstImage()
	if err != nil {
		return err
	}

	board, err := Board.FromImage(file)
	if err != nil {
		return err
	}

	for i := index; i < *step; i++ {
		newBoard, err := board.Next()
		if err != nil {
			return err
		}

		image, err := newBoard.ToImage()
		if err != nil {
			return err
		}

		err = saveImage(image, i)
		if err != nil {
			return err
		}

		if board.Equals(newBoard) {
			break
		}

		board = newBoard
	}

	return nil
}
