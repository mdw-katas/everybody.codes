package quest_03

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func assertEqual(t *testing.T, a, b any) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func TestPart1(t *testing.T) {
	assertEqual(t, 35, Mine(t, "part1-sample.txt"))
	assertEqual(t, 117, Mine(t, "part1-full.txt"))
}

func Mine(t *testing.T, filename string) int {
	input, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	total := 0
	points := Scan(string(input))
	for {
		total += len(points)
		points = Dig(points)
		if len(points) == 0 {
			break
		}
	}
	return total
}

func Dig(points map[Point]struct{}) (result map[Point]struct{}) {
	result = make(map[Point]struct{})
	for point := range points {
		if neighbors(point, points) == 4 {
			result[point] = struct{}{}
		}
	}
	return result
}

func neighbors(point Point, points map[Point]struct{}) (count int) {
	_, left := points[Point{Row: point.Row, Col: point.Col - 1}]
	_, right := points[Point{Row: point.Row, Col: point.Col + 1}]
	_, above := points[Point{Row: point.Row - 1, Col: point.Col}]
	_, below := points[Point{Row: point.Row + 1, Col: point.Col}]
	if left {
		count++
	}
	if right {
		count++
	}
	if above {
		count++
	}
	if below {
		count++
	}
	return count
}

type Point struct{ Row, Col int }

func Scan(input string) (result map[Point]struct{}) {
	result = make(map[Point]struct{})
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
