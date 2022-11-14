package matrix

type Matrix struct {
	store [][]*Color
}

func NewMatrix() *Matrix {
	var mat Matrix
	return &mat
}
