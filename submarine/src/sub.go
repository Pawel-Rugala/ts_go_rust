package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func getInput() string {
	return `forward 5
down 5
forward 8
up 3
down 8
forward 2`
}

type Point struct {
	x int
	y int
}

func parseLIne(line string) Point {
	parts := strings.Split(line, " ")
	amount, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal("this should never ever happen")
	}

	if parts[0] == "forward" {
		return Point{amount, 0}
	} else if parts[0] == "up" {
		return Point{0, -amount}
	}
	return Point{0, amount}
}

func main() {
	lines := strings.Split(getInput(), "\n")

	pos := Point{0, 0}
	for _, line := range lines {
		amount := parseLIne(line)
		pos.x += amount.x
		pos.y += amount.y
	}

	fmt.Printf("point: %+v", pos)
}
