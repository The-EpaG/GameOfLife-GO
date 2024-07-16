package Board

import (
	"image"

	"github.com/The-EpaG/GameOfLife-GO/internal/classes/cell"
	"github.com/The-EpaG/GameOfLife-GO/internal/constants"
	"github.com/The-EpaG/GameOfLife-GO/internal/errors"
)

type Board struct {
	Width  int
	Height int
	cells  [][]cell.Cell
}

func New(width int, height int) Board {
	cells := make([][]cell.Cell, width)

	for i := range cells {
		cells[i] = make([]cell.Cell, 0, height)

		for j := 0; j < height; j++ {
			cells[i] = append(cells[i], cell.Cell{
				State: cell.Dead,
				X:     i,
				Y:     j,
			})
		}
	}

	return Board{
		Width:  width,
		Height: height,
		cells:  cells,
	}
}

func (board *Board) GetCell(x int, y int) (*cell.Cell, error) {
	if x < 0 || x >= board.Width || y < 0 || y >= board.Height {
		return nil, &errors.PositionOutsideImageError{
			X: x,
			Y: y,
		}
	}
	return &board.cells[x][y], nil
}

func (board *Board) SetCell(cell *cell.Cell) {
	board.cells[cell.X][cell.Y] = *cell
}

func FromImage(image *image.Gray) (*Board, error) {
	board := New(image.Bounds().Max.X, image.Bounds().Max.Y)
	for y := 0; y < image.Bounds().Max.Y; y++ {
		for x := 0; x < image.Bounds().Max.X; x++ {
			pixel, err := board.GetCell(x, y)
			if err != nil {
				return &Board{}, err
			}

			if image.GrayAt(x, y).Y == 0 {
				pixel.State = cell.Dead
			} else {
				pixel.State = cell.Alive
			}

			board.SetCell(pixel)
		}
	}

	return &board, nil
}

func (board *Board) ToImage() (*image.Gray, error) {
	image := image.NewGray(image.Rect(0, 0, board.Width, board.Height))

	for y := 0; y < board.Height; y++ {
		for x := 0; x < board.Width; x++ {
			pixel, err := board.GetCell(x, y)
			if err != nil {
				return image, err
			}

			if pixel.State == cell.Alive {
				image.SetGray(x, y, constants.Alive)
			} else {
				image.SetGray(x, y, constants.Dead)
			}
		}
	}
	return image, nil
}

func (board *Board) IsAlive(x int, y int) (bool, error) {
	pixel, err := board.GetCell(x, y)
	if err != nil {
		return false, err
	}
	return pixel.IsAlive(), nil
}

func (board *Board) getNeighbors(x int, y int) ([]cell.Cell, error) {
	var neighbors []cell.Cell
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if x+i < 0 || x+i >= board.Width {
				break
			}

			if y+j < 0 || y+j >= board.Height {
				continue
			}

			if i == 0 && j == 0 {
				continue
			}

			pixel, err := board.GetCell(x+i, y+j)
			if err != nil {
				return []cell.Cell{}, err
			}
			neighbors = append(neighbors, *pixel)
		}
	}
	return neighbors, nil
}

func (board *Board) WillSurvive(x int, y int) (bool, error) {
	pixel, err := board.GetCell(x, y)
	if err != nil {
		return false, err
	}

	neighbors, err := board.getNeighbors(x, y)

	if err != nil {
		return false, err
	}

	return pixel.WillSurvive(neighbors), nil
}

func (board *Board) Next() (*Board, error) {
	newBoard := New(board.Width, board.Height)
	for y := 0; y < board.Height; y++ {
		for x := 0; x < board.Width; x++ {
			pixel, err := board.GetCell(x, y)
			if err != nil {
				return &Board{}, err
			}

			neighbors, err := board.getNeighbors(x, y)
			if err != nil {
				return &Board{}, err
			}

			if pixel.IsAlive() && !pixel.WillSurvive(neighbors) {
				newPixel, err := newBoard.GetCell(x, y)
				if err != nil {
					return &Board{}, err
				}

				newPixel.State = cell.Dead
				newBoard.SetCell(newPixel)
			} else if !pixel.IsAlive() && pixel.WillRevive(neighbors) {
				newPixel, err := newBoard.GetCell(x, y)
				if err != nil {
					return &Board{}, err
				}

				newPixel.State = cell.Alive
				newBoard.SetCell(newPixel)
			} else {
				newBoard.SetCell(pixel)
			}
		}
	}
	return &newBoard, nil
}

func (a *Board) Equals(b *Board) bool {
	if a == b {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Width != b.Width || a.Height != b.Height {
		return false
	}

	for y := 0; y < a.Height; y++ {
		for x := 0; x < a.Width; x++ {
			if a.cells[x][y].State != b.cells[x][y].State {
				return false
			}
		}
	}
	return true
}
