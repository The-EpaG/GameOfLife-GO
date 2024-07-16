package errors

import "fmt"

type PositionOutsideImageError struct {
	X int
	Y int
}

func (err *PositionOutsideImageError) Error() string {
	return fmt.Sprintln(err.X, "x", err.Y, " is outside the image")
}