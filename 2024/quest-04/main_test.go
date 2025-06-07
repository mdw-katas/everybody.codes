package quest_04

import (
	"bufio"
	"iter"
	"math"
	"os"
	"sort"
	"strconv"
	"testing"
)

func assertEqual(t *testing.T, a, b any) {
	if a != b {
		t.Helper()
		t.Errorf("%v != %v", a, b)
	}
}

func TestPart1(t *testing.T) {
	assertEqual(t, 10, levelToLowest("part1-sample.txt"))
	assertEqual(t, 76, levelToLowest("part1-full.txt"))
}
func TestPart2(t *testing.T) {
	assertEqual(t, 923412, levelToLowest("part2-full.txt"))
}
func TestPart3(t *testing.T) {
	assertEqual(t, 8, levelBruteForce("part3-sample.txt"))
	assertEqual(t, 126398193, levelBruteForce("part3-full.txt"))
}

func levelBruteForce(filename string) (result int) {
	result = 0xFFFFFFFF
	heights := parseHeights(filename)
	for _, height := range heights {
		candidate := countHitsToAlignHeights(heights, height)
		if candidate < result {
			result = candidate
		}
	}
	return result
}
func levelToLowest(filename string) int {
	heights := parseHeights(filename)
	sort.Ints(heights)
	return countHitsToAlignHeights(heights, heights[0])
}

func lines(path string) iter.Seq[string] {
	return func(yield func(string) bool) {
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer func() { _ = file.Close() }()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				break
			}
		}
	}
}
func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
func parseHeights(filename string) []int {
	var heights []int
	for line := range lines(filename) {
		heights = append(heights, parseInt(line))
	}
	return heights
}
func countHitsToAlignHeights(heights []int, target int) (result int) {
	for _, height := range heights {
		result += int(math.Abs(float64(height) - float64(target)))
	}
	return result
}
