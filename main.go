package main

import (
	"container/list"
	"errors"
	"fmt"
	"log"
	"strings"
)

type Square struct {
	posX              byte
	posY              byte
	algebraicNotation string
	path              string
}

// Factory pattern used because in Go, struct does not have constructors
// Create a Square using the numerical representation of the knight's position
func InitSquare(x byte, y byte) (*Square, error) {
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

// Create a Square using the algebraic notation of the knight's position
func InitSquareFromAlgebraicNotation(algNotation string) (*Square, error) {
	// Convert to uppercase
	algNotation = strings.ToUpper(algNotation)

	// if the algebraic notation contains more than 2 components, return error
	if len(algNotation) != 2 {
		return nil, errors.New("Algebraic expression too long")
	} else {
		// Simple and readable way to convert the byte into a number.
		x, y := algNotation[0]-'A'+1, algNotation[1]-'0'
		// fmt.Println("x: ", x, " y: ", y)
		square, err := InitSquare(x, y)
		return square, err
	}
}

func (source *Square) ShortestPathToASquare(target *Square) (string, error) {
	// Possible knight moves in X and Y direction
	// DO NOT CHANGE THIS AS THERE ARE MULTIPLE SOLUTIONS
	// AND CHANGING THIS WILL GIVE DIFFERENT SOLUTIONS.
	// IF CHANGE IS NECESSARY, UPDATE THE TESTS ACCORDINGLY
	dx := []int{-1, 2, 1, -2, -2, -1, 1, 2}
	dy := []int{-2, -1, -2, -1, 1, 2, 2, 1}

	// Store visited squares. Default values should be false as none of the squares are visited.
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
	// a slice can be used like so
	// queue := []Square{}
	// to dequeue: current, queue := queue[0], queue[1:]
	// could be used. But this slice approach will cause
	// MEMORY LEAKS
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
			return current.path, nil
		}

		for i := 0; i < 8; i++ {
			x := int(current.posX) + dx[i]
			y := int(current.posY) + dy[i]

			if (x >= 1 && x <= 8) && (y >= 1 && y <= 8) && !visited[x-1][y-1] {
				visited[x-1][y-1] = true
				// byte(x) and byte(y) is the simplest here it is safe
				// because both x and y cannot be less than 1 or more than 8
				visiting, err := InitSquare(byte(x), byte(y))
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
	return "", errors.New("Given the right inputs, a solution is always possible. But no solution is found.")

}

func main() {
	// square, err := InitSquareFromAlgebraicNotation("A8")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(square)

	var sourceInputString string
	var targetInputString string
	fmt.Println("Enter the source and target squares in algebraic notation")
	_, err := fmt.Scan(&sourceInputString)
	if err != nil {
		log.Println(err)
	}
	_, err = fmt.Scan(&targetInputString)
	if err != nil {
		log.Println(err)
	}

	source, err := InitSquareFromAlgebraicNotation(sourceInputString)
	if err != nil {
		log.Println(err)
		return
	}
	target, err := InitSquareFromAlgebraicNotation(targetInputString)
	if err != nil {
		log.Println(err)
		return
	}

	path, err := source.ShortestPathToASquare(target)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(path)

}
