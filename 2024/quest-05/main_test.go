package quest_05

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func assertEqual(t *testing.T, a, b any) {
	if a != b {
		t.Helper()
		t.Errorf("%v != %v", a, b)
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
	var columns = parseRows("part1-sample.txt")
	for _, column := range columns {
		t.Log(column.values)
	}
	for round := range 10 {
		t.Log(performRound(round, columns))
	}
}
func parseRows(filename string) (result []*List[int]) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()
	scanner := bufio.NewScanner(file)
	for scanned := 0; scanner.Scan(); scanned++ {
		line := scanner.Text()
		for field := range strings.FieldsSeq(line) {
			if scanned == 0 {
				column := NewList[int]()
				result = append(result, column)
			}
			result[scanned].Append(parseInt(field))
		}
	}
	return result
}
func performRound(roundCounter int, columns []*List[int]) int {
	from := roundCounter % len(columns)
	clapper := columns[from].Pop(0)
	to := (from + 1) % len(columns)
	Clap(clapper, clapper, columns[to])
	var builder strings.Builder
	for _, column := range columns {
		builder.WriteString(strconv.Itoa(column.Nth(0)))
	}
	return parseInt(builder.String())
}
func Clap(clapper, claps int, column *List[int]) {
	cursor := 0
	for cursor < column.Len() && claps >= 0 {
		cursor, claps = cursor+1, claps-1
	}
	if claps == 0 { // on left side, insert 'in-front' or before
		column.Insert(cursor, clapper)
		return
	}
	for cursor >= 0 && claps >= 0 {
		cursor, claps = cursor-1, claps-1
	}
	if claps == 0 { // on right side, insert 'behind' or after
		column.Insert(cursor+1, clapper)
		return
	}
	if claps > 0 {
		Clap(clapper, claps, column)
	}
}

type List[T any] struct{ values []T }

func NewList[T any](values ...T) *List[T] {
	return &List[T]{values: values}
}
func (l *List[T]) Len() int {
	return len(l.values)
}
func (this *List[T]) Nth(n int) T {
	return this.values[n]
}
func (this *List[T]) Insert(at int, value T) {
	this.values = slices.Insert(this.values, at, value)
}
func (this *List[T]) Pop(at int) T {
	v := this.values[at]
	this.values = slices.Delete(this.values, at, at+1)
	return v
}
func (this *List[T]) Append(values ...T) {
	this.values = append(this.values, values...)
}
