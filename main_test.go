package main

import "testing"

func TestInitSquare(t *testing.T) {
	var x byte
	var y byte
	x = 1
	y = 2
	square, err := InitSquare(x, y)
	if err != nil {
		t.Errorf("%v", err)
	} else if square.posX != x || square.posY != y || square.algebraicNotation != "A2" || square.path != "" {
		t.Error("Expected not returned")

	}
}
func TestInitSquareFromAlgebraicNotation(t *testing.T) {
	var algNotation = "A5"
	square, err := InitSquareFromAlgebraicNotation(algNotation)
	if err != nil {
		t.Errorf("%v", err)
	} else if square.posX != 1 || square.posY != 5 || square.algebraicNotation != "A5" || square.path != "" {
		t.Error("Expected not returned")
	}
}

func TestShortestPathToASquare(t *testing.T) {
	source, err := InitSquareFromAlgebraicNotation("A8")
	if err != nil {
		t.Errorf("%v", err)
	}
	target, err := InitSquareFromAlgebraicNotation("B7")
	if err != nil {
		t.Errorf("%v", err)
	}

	path, err := source.ShortestPathToASquare(target)
	if err != nil {
		t.Errorf("%v", err)
	}

	// Expected output will change based on the dx and dy arrays in main.go. Update accordingly.
	expected := "C7 B5 D6 B7"
	if path != expected {
		t.Errorf("Incorrect path. Expected %q got %q", expected, path)
		t.Error("If the solution is valid, change the dx and dy arrays in main.go or update test")
	}

	source2, err := InitSquareFromAlgebraicNotation("A1")
	if err != nil {
		t.Errorf("%v", err)
	}
	target2, err := InitSquareFromAlgebraicNotation("H8")
	if err != nil {
		t.Errorf("%v", err)
	}

	path2, err := source2.ShortestPathToASquare(target2)
	if err != nil {
		t.Errorf("%v", err)
	}
	// Expected output will change based on the dx and dy arrays in main.go. Update accordingly.
	// Too many possible paths, not practical to list them all.
	expected2 := "B3 D2 C4 D6 F7 H8"
	if path2 != expected2 {
		t.Errorf("Incorrect path. Expected %q got %q", expected2, path2)
		t.Error("If the solution is valid, change the direction matrix in main.go or update test")
	}
}
