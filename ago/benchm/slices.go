package caching

const (
	rows = 2 * 1024
	cols = 2 * 1024
)

var matrix [rows][cols]byte

type data struct {
	v byte
	p *data
}

var list *data

func init() {
	var last *data
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			var d data
			if list == nil {
				list = &d
			}
			if last != nil {
				last.p = &d
			}
			last = &d
			if row%2 == 0 {
				matrix[row][col] = 0xFF
				d.v = 0xFF
			}
		}
	}
}

func LinkedListTraverse() int {
	var ctr int
	d := list
	for d != nil {
		if d.v == 0xFF {
			ctr++
		}
		d = d.p
	}
	return ctr
}

func ColumnTraverse() int {
	ctr := 0
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			if matrix[row][col] == 0xFF {
				ctr++
			}
		}
	}
	return ctr
}

func RowTraverse() int {
	ctr := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == 0xFF {
				ctr++
			}
		}
	}
	return ctr
}
