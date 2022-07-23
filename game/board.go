package game

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Board [GridSize][GridSize]Cell

func BuildBoard() Board {
	b := Board{}

	for x := 0; x < GridSize; x++ {
		for y := 0; y < GridSize; y++ {

			c := Cell{
				CellState: Closed,
				CellType:  Zero,
				X:         x,
				Y:         y,
			}
			b[x][y] = c

		}
	}
	b.placeBombs(BombAmount)
	b.placeCells()

	return b
}

func (b Board) String() string {
	var str string

	latest := strconv.Itoa(GridSize - 1)
	str += strings.Repeat(" ", len(latest)+1)

	for i := 0; i <= GridSize-1; i++ {
		str += fmt.Sprintf("%d%s", i, strings.Repeat(" ", 3-len(strconv.Itoa(i))))
	}
	str += "\n"

	for i, v := range b {
		str += fmt.Sprintf("%d%s|", i, strings.Repeat(" ", len(latest)-len(strconv.Itoa(i))))
		for _, j := range v {

			str += fmt.Sprintf("%v  ", j)

		}

		str += "\n"
	}
	return str
}

func (b *Board) placeBombs(amount int) {
	rand.Seed(time.Now().UnixNano())
	for bombs := 0; bombs < amount; bombs++ {
		max := int(math.Pow(float64(GridSize), 2)) - 1
		location := rand.Intn(max-1) + 1

		XCord := location / GridSize
		YCord := location % GridSize
		if b[XCord][YCord].CellType == Bomb {
			continue
		}
		c := Cell{
			CellState: Closed,
			CellType:  Bomb,
			X:         XCord,
			Y:         YCord,
		}

		b[XCord][YCord] = c

	}
}

func (b *Board) placeCells() {
	for x := 0; x < GridSize; x++ {
		for y := 0; y < GridSize; y++ {

			if b[x][y].CellType == Bomb {
				continue
			}

			b[x][y] = Cell{
				CellState: Closed,
				CellType:  b.getNeighborBombs(x, y),
				X:         x,
				Y:         y,
			}
		}
	}
}

func (b *Board) getNeighborBombs(x, y int) CellType {
	neighborBombs := 0
	r, c := float64(x), float64(y)

	for row := math.Max(0, r-1.0); row < math.Min(float64(GridSize)-1.0, r+1.0)+1.0; row++ {
		for col := math.Max(0, c-1.0); col < math.Min(float64(GridSize)-1.0, c+1.0)+1.0; col++ {

			if row == r && col == c {
				continue
			}

			if b[int(row)][int(col)].CellType == Bomb {

				neighborBombs += 1
			}
		}
	}

	return CellType(neighborBombs)
}

func (b *Board) isEnd() bool {
	for x := 0; x < GridSize; x++ {
		for y := 0; y < GridSize; y++ {

			if b[x][y].CellState == Closed {

				return false

			} else if b[x][y].CellState == Flagged && b[x][y].CellType != Bomb {

				return false
			}
		}
	}

	return true
}
func (b *Board) revealBoard() {
	for x := 0; x < GridSize; x++ {
		for y := 0; y < GridSize; y++ {
			b[x][y].CellState = Opened
		}
	}

}
