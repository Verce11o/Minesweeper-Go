package game

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Game int

const (
	Over Game = iota
	InvalidCords
	NotEnoughFlags
	Success
)

type Player struct {
	Flags int
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if strings.ToLower(v) == str {
			return true
		}
	}
	return false
}
func (p *Player) DoMove(b *Board) (g Game) {
	var move string
	for {
		fmt.Print("Your move: ")
		fmt.Scanln(&move)

		move := strings.Split(move, ",")
		if len(move) >= 2 {
			row, _ := strconv.Atoi(move[0])
			col, _ := strconv.Atoi(move[1])

			if contains(move, "flag") {
				flag := p.Flag(b, row, col)
				return flag
			}
			dig := p.Dig(b, row, col)
			return dig

		} else {
			return InvalidCords
		}
	}
}
func (p *Player) Dig(b *Board, x, y int) Game {
	if x < 0 || x >= GridSize || y < 0 || y >= GridSize {
		return InvalidCords

	} else if b[x][y].CellType == Bomb {
		b[x][y].CellState = Opened
		return Over
	}

	b[x][y].CellState = Opened
	if b[x][y].CellType == Zero {
		r, c := float64(x), float64(y)
		for row := math.Max(0, r-1.0); row < math.Min(float64(GridSize)-1.0, r+1.0)+1.0; row++ {
			for col := math.Max(0, c-1.0); col < math.Min(float64(GridSize)-1.0, c+1.0)+1.0; col++ {
				if row == r && col == c {
					continue
				}
				intr, intc := int(row), int(col)
				currentCell := b[intr][intc]
				b[intr][intc].CellState = Opened
				if currentCell.CellType == Zero && currentCell.CellState == Closed {
					p.Dig(b, intr, intc)
				}

			}
		}
	}
	return Success

}

func (p *Player) Flag(b *Board, x, y int) Game {
	if b[x][y].CellState == Flagged {
		b[x][y].CellState = Closed
		p.Flags += 1
		return Success
	}

	if p.Flags == 0 {
		return NotEnoughFlags
	} else if x < 0 || x > GridSize || y < 0 || y > GridSize || b[x][y].CellState == Opened {
		return InvalidCords
	}

	b[x][y].CellState = Flagged
	p.Flags -= 1
	fmt.Println(fmt.Sprintf("Remaining flags: %d", p.Flags))
	return Success
}
