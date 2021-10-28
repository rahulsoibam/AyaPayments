package main

import (
	"container/list"
	"errors"
	"fmt"
	"log"
)

type Square struct {
	posX              int
	posY              int
	algebraicNotation string
	path              string
}

func InitSquare(x int, y int) (*Square, error) {
	// Check if Square is inside board
	if (x >= 1 && x <= 8) && (y >= 1 && y <= 8) {
		return &Square{
			posX:              x,
			posY:              y,
			algebraicNotation: string('A'+x-1) + string(y+'0'),
			path:              "",
		}, nil
	}
	return nil, errors.New("Square is not inside board")
}

func InitSquareFromAlgebraicNotation(algNotation string) (*Square, error) {
	if len(algNotation) != 2 {
		return nil, errors.New("Invalid algebraic notation")
	} else {
		x, y := int(algNotation[0]-'A'+1), int(algNotation[1]-'0')
		// fmt.Println("x: ", x, " y: ", y)
		square, err := InitSquare(x, y)
		return square, err
	}
}

func (source *Square) ShortestPathToaSquare(target *Square) string {
	// Possible knight moves in X and Y direction
	dx := []int{-2, -1, 1, 2, -2, -1, 1, 2}
	dy := []int{-1, -2, -2, -1, 1, 2, 2, 1}

	// Store visited squares
	visited := [][]bool{
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
	}

	// Queue for Breadth First Search
	// a slice can be used but will be prone to memory leaks
	queue := list.New()

	// mark the source square as visited
	visited[source.posX-1][source.posY-1] = true

	// push source square into the queue
	queue.PushBack(source)

	// until the queue is empty, do breadth first search until
	// the target square is reached
	for queue.Len() > 0 {
		// pop the fist element of queue
		e := queue.Front()
		current := queue.Remove(e).(*Square)

		// Check if the current square is the target
		if current.posX == target.posX && current.posY == target.posY {
			return current.path
		}

		for i := 0; i < 8; i++ {
			x := int(current.posX) + dx[i]
			y := int(current.posY) + dy[i]

			if (x >= 1 && x <= 8) && (y >= 1 && y <= 8) && !visited[x-1][y-1] {
				visited[x-1][y-1] = true
				visiting, err := InitSquare(x, y)
				if err != nil {
					log.Println(err)
					continue
				}
				if current.path == "" {
					visiting.path = visiting.algebraicNotation
				} else {
					visiting.path = current.path + " " + visiting.algebraicNotation
				}
				queue.PushBack(visiting)
			}
		}
	}
	return ""

}

func main() {
	// square, err := InitSquareFromAlgebraicNotation("A8")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(square)

	source, err := InitSquareFromAlgebraicNotation("A8")
	if err != nil {
		log.Println(err)
		return
	}
	target, err := InitSquareFromAlgebraicNotation("B7")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(source.ShortestPathToaSquare(target))

}
