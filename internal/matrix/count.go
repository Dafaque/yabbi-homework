package matrix

import "fmt"

// Start Subsequence Ident
type SSID struct {
	color *Color
	atX   int
	atY   int
}

func (mat *Matrix) PrintMaxEqualsInRow() {
	var hash map[SSID]uint = make(map[SSID]uint)
	for i := range mat.store {
		row := mat.store[i]
		var ssid *SSID
		for j := range row {
			val := row[j]
			if ssid == nil || !ssid.color.IsEqual(val) {
				ssid = &SSID{color: val, atX: i, atY: j}
			}
			if _, exist := hash[*ssid]; !exist {
				hash[*ssid] = 1
			}
			if j > 0 && val.IsEqual(row[j-1]) {
				hash[*ssid]++
			}
		}
	}
	var mostInRowID SSID
	var mostInRow uint

	for k, v := range hash {
		if v > mostInRow {
			mostInRowID = k
			mostInRow = v
		}
	}

	fmt.Printf(
		"The color that repeats the most in a row is %v at %d:%d %d times\n",
		mostInRowID.color,
		mostInRowID.atX,
		mostInRowID.atY,
		mostInRow,
	)
}
