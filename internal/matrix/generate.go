package matrix

import "fmt"

func (mat *Matrix) Fill(m, n uint) {
	var i uint
	mat.store = make([][]*Color, m)
	for i = 0; i < m; i++ {
		mat.store[i] = make([]*Color, n)
		var j uint
		for j = 0; j < n; j++ {
			mat.store[i][j] = NewRandomColor()
		}
		fmt.Println(mat.store[i])
	}
}
