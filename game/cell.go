package game

import "fmt"

type (
	CellType  int
	CellState int
)

const (
	Closed CellState = iota
	Opened
	Flagged
)

const (
	Zero CellType = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Bomb
)

type Cell struct {
	CellState CellState
	CellType  CellType
	X         int
	Y         int
}

func (c Cell) String() string {
	colors := map[CellType]string{
		Zero:  "\033[1;37m%d\u001B[0m",
		One:   "\033[1;36m%d\u001B[0m",
		Two:   "\033[1;32m%d\u001B[0m",
		Three: "\033[1;35m%d\u001B[0m",
		Four:  "\033[1;34m%d\u001B[0m",
		Five:  "\033[1;95m%d\u001B[0m",
		Six:   "\033[1;96m%d\u001B[0m",
		Seven: "\033[1;30m%d\u001B[0m",
		Eight: "\033[1;33m%d\u001B[0m",
	}
	if c.CellState == Closed {
		return string(rune(0x25A1))
	} else if c.CellState == Flagged {
		return string(rune(0x2691))
	} else if c.CellType == Bomb {
		return "\033[1;31mB\u001B[0m"
	}

	return fmt.Sprintf(colors[c.CellType], int(c.CellType))
}
