package main

import (
	"fmt"

	"github.com/w1png/go-chan/types"
)

func main() {
	boards, err := types.GetBoards()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, board := range boards {
		fmt.Println(board)
	}
}
