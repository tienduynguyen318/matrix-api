package constant

import "fmt"

var (
	// ErrEmptyMatrix --
	ErrEmptyMatrix = fmt.Errorf("Matrix is empty")
	// ErrDifferentDimension --
	ErrDifferentDimension = fmt.Errorf("Matrix has different dimension")
	// ErrNotInteger --
	ErrNotInteger = fmt.Errorf("A cell in matrix is not an integer")
	// ErrOverFlow --
	ErrOverFlow = fmt.Errorf("Number is overflow")
)
