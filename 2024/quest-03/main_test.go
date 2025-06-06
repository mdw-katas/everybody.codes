package quest_03

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func assertEqual(t *testing.T, a, b any) {
	if a != b {
		t.Fatalf("%v != %v", a, b)
	}
	t.Helper()
	t.Log(b)
}

func TestPart1(t *testing.T) {
	assertEqual(t, 35, Mine("part1-sample.txt", neighbors4))
	assertEqual(t, 117, Mine("part1-full.txt", neighbors4))
}
func TestPart2(t *testing.T) {
	assertEqual(t, 2701, Mine("part2-full.txt", neighbors4))
}
func TestPart3(t *testing.T) {
	assertEqual(t, 29, Mine("part1-sample.txt", neighbors8))
	assertEqual(t, 2005, Mine("part3-full.txt", neighbors8))
}

func Mine(filename string, neighbors Neighbors) int {
	input, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	total := 0
	points := Scan(string(input))
	for {
		total += len(points)
		points = Dig(points, neighbors)
		if len(points) == 0 {
			break
		}
	}
	return total
}

func Dig(points Field, neighbors Neighbors) (result Field) {
	result = make(Field)
	for point := range points {
		expected, actual := neighbors(point, points)
		if actual == expected {
			result[point] = struct{}{}
		}
	}
	return result
}

func neighbors8(point Point, field Field) (expected, actual int) {
	var (
		upperLeft  = contains(field, Point{Row: point.Row - 1, Col: point.Col - 1})
		upperRight = contains(field, Point{Row: point.Row - 1, Col: point.Col + 1})
		lowerLeft  = contains(field, Point{Row: point.Row + 1, Col: point.Col - 1})
		lowerRight = contains(field, Point{Row: point.Row + 1, Col: point.Col + 1})
	)
	expected, actual = neighbors4(point, field)
	return expected + 4, actual + bool2int(upperLeft, upperRight, lowerLeft, lowerRight)
}
func neighbors4(point Point, field Field) (expected, actual int) {
	_, left := field[Point{Row: point.Row, Col: point.Col - 1}]
	_, right := field[Point{Row: point.Row, Col: point.Col + 1}]
	_, above := field[Point{Row: point.Row - 1, Col: point.Col}]
	_, below := field[Point{Row: point.Row + 1, Col: point.Col}]
	return 4, bool2int(left, right, above, below)
}
func bool2int(bools ...bool) (result int) {
	for _, b := range bools {
		if b {
			result++
		}
	}
	return result
}
func contains(field Field, point Point) bool {
	_, ok := field[point]
	return ok
}

type Point struct{ Row, Col int }

func Scan(input string) (result Field) {
	result = make(Field)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()
		for col, char := range line {
			if char == '#' {
				result[Point{row, col}] = struct{}{}
			}
		}
	}
	return result
}

type Neighbors func(Point, Field) (expected, actual int)
type Field map[Point]struct{}
