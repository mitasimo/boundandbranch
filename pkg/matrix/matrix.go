package matrix

import "errors"

var (
	errorColumnIndexOutOfBound = errors.New("column index out of bound")
	errorRowIndexOutOfBound    = errors.New("row index out of bound")
)

func New(xMax, yMax int) Matrix {
	return Matrix{x: xMax, y: yMax, data: make([]int, xMax*yMax)}
}

type Matrix struct {
	x, y int
	data []int
}

func (m Matrix) At(x, y int) (int, error) {
	if err := m.checkIndex(x, y); err != nil {
		return 0, err
	}

	return m.data[x+m.x*y], nil
}

func (m Matrix) Set(x, y, val int) error {
	if err := m.checkIndex(x, y); err != nil {
		return err
	}

	m.data[x+m.x*y] = val
	return nil
}

func (m Matrix) checkIndex(x, y int) error {
	if x >= m.x {
		return errorColumnIndexOutOfBound
	}
	if y >= m.y {
		return errorRowIndexOutOfBound
	}
	return nil
}
