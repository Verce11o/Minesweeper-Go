package game

import (
	"fmt"
)

const GridSize int = 8
const BombAmount = GridSize * GridSize / 6

func (p Player) NewGame() {
	board := BuildBoard()
	fmt.Println("Please consider using following example:\nx,y,flag(if necessary)")
	for i := 0; !board.isEnd(); i++ {
		fmt.Println(board)
		g := p.DoMove(&board)
		switch g {
		case InvalidCords:
			fmt.Println("Bad coords. Example: 5,9\nFlagging a cell: 5,9 Flag")
		case NotEnoughFlags:
			fmt.Println("You don't have enough flags. You can unflagg a cell by flagging it again.")
		case Over:
			fmt.Println("You lost! :(")
			board.revealBoard()
			fmt.Println(board)
			return
		}

	}
	fmt.Println("Congratulations, you won! :)")
	fmt.Println(board)
}
