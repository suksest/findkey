package main

import (
	"fmt"
)

type rule struct {
	north uint64
	east  uint64
	south uint64
}

type cursor struct {
	curRow       int
	curCol       int
	ellapsedDots uint64
}

type result struct {
	aRule rule
	dots  uint64
}

const row, col = 6, 8

type constraint struct {
	maxNorth uint64
	maxEast  uint64
	maxSouth uint64
}

type board struct {
	elements [row][col]string
}

var possibleRules []result

func (b *board) initBoard() {
	b.elements = [row][col]string{
		{"#", "#", "#", "#", "#", "#", "#", "#"},
		{"#", ".", ".", ".", ".", ".", ".", "#"},
		{"#", ".", "#", "#", "#", ".", ".", "#"},
		{"#", ".", ".", ".", "#", ".", "#", "#"},
		{"#", "X", "#", ".", ".", ".", ".", "#"},
		{"#", "#", "#", "#", "#", "#", "#", "#"},
	}
}

func (b *board) display() {
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Print(b.elements[i][j])
			if j == col-1 {
				fmt.Println("")
			}
		}
	}
}

func (c *cursor) isPossibleStep(b board, direction string) bool {
	switch direction {
	case "north":
		c.curRow--
	case "east":
		c.curCol++
	case "south":
		c.curRow++
	}

	if c.curRow < row && c.curCol < col {
		if b.elements[c.curRow][c.curCol] != "#" {
			return true
		}
	}
	return false
}

func (c *cursor) step(b board, r rule) {
	for i := uint64(0); i < r.north; i++ {
		if c.isPossibleStep(b, "north") {
			c.ellapsedDots += r.north
		}
	}
	for i := uint64(0); i < r.east; i++ {
		if c.isPossibleStep(b, "east") {
			c.ellapsedDots += r.east
		}
	}
	for i := uint64(0); i < r.south; i++ {
		if c.isPossibleStep(b, "south") {
			c.ellapsedDots += r.south
		}
	}
}

func (c *cursor) isPossiblePosition(b board) bool {
	if b.elements[c.curRow][c.curCol] != "#" {
		return true
	}
	return false
}

func (c *cursor) walk(b board, r rule) {
	c.step(b, r)
}

func (co *constraint) initConstraint(c cursor) {
	co.maxNorth = uint64(c.curRow - 1)
	co.maxEast = uint64(col - 2 - c.curCol)
	co.maxSouth = uint64(row - 2 - c.curRow)
}

func genPossibleRules(co constraint, c cursor) (rules []rule) {
	for i := uint64(1); i < co.maxNorth; i++ {
		for j := uint64(1); j < co.maxEast; j++ {
			for k := uint64(1); k < co.maxSouth; k++ {
				aRule := rule{i, j, k}
				rules = append(rules, aRule)
			}
		}
	}
	return rules
}

func main() {
	var aBoard board
	aBoard.initBoard()
	var player cursor
	player.curRow, player.curCol = 4, 1
	var aConstraint constraint
	aConstraint.initConstraint(player)
	rules := genPossibleRules(aConstraint, player)
	fmt.Println(rules)
	for _, aRule := range rules {
		player.walk(aBoard, aRule)
		if player.isPossiblePosition(aBoard) {
			var posResult result
			posResult.aRule = aRule
			posResult.dots = player.ellapsedDots
			possibleRules = append(possibleRules, posResult)
		}
	}
	aBoard.display()
	fmt.Println(possibleRules)
}
