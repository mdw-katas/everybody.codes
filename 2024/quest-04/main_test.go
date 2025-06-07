package quest_04

import (
	"bufio"
	"iter"
	"os"
	"sort"
	"strconv"
	"testing"
)

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

func TestPart1(t *testing.T) {
	t.Log(level("part1-sample.txt"))
	t.Log(level("part1-full.txt"))
}
func TestPart2(t *testing.T) {
	t.Log(level("part2-full.txt"))
}

func level(filename string) int {
	var heights []int
	for line := range lines(filename) {
		heights = append(heights, parseInt(line))
	}
	sort.Ints(heights)
	target := heights[0]
	heights = heights[1:]
	hits := 0
	for _, height := range heights {
		hits += height - target
	}
	return hits
}
