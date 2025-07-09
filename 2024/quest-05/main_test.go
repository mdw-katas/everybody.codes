package quest_05

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/mdw-go/must"
	"github.com/mdw-go/printing"
	"github.com/mdw-go/testing/should"
)

func parseInt(s string) int {
	return must.Value(strconv.Atoi(s))
}

func TestPart1Sample(t *testing.T) {
	testPart1(t, "part1-sample.txt", 2323)
	testPart1(t, "part1-full.txt", 4422)
}
func testPart1(t *testing.T, inputFile string, expected int) {
	t.Run(inputFile, func(t *testing.T) {
		var field = parseField(inputFile)
		result := 0
		for round := range 10 {
			result = field.performRound(round)
			//t.Logf("Round: %d\nResult: %d\nField:\n%s",
			//	round+1,
			//	result,
			//	field.String(),
			//)
		}
		should.So(t, result, should.Equal, expected)
	})
}

type Field struct {
	columns []*List[int]
}

func parseField(filename string) (result *Field) {
	result = &Field{}
	file := must.Value(os.Open(filename))
	defer must.Defer(file.Close)()
	scanner := bufio.NewScanner(file)
	for scanned := 0; scanner.Scan(); scanned++ {
		line := scanner.Text()
		for f, field := range strings.Fields(line) {
			if scanned == 0 {
				column := NewList[int]()
				result.columns = append(result.columns, column)
			}
			result.columns[f].Append(parseInt(field))
		}
	}
	return result
}

func (this *Field) String() string {
	builder := printing.NewBuilder()
	for x := 0; ; x++ {
		finished := true
		for _, column := range this.columns {
			if x < len(column.values) {
				builder.Print(column.values[x])
				builder.Print(" ")
				finished = false
			} else {
				builder.Print("-")
				builder.Print(" ")
			}
		}
		builder.Println()
		if finished {
			return builder.Inner().String()
		}
	}
}
func (this *Field) performRound(roundCounter int) int {
	columns := this.columns
	from := roundCounter % len(columns)
	clapper := columns[from].Pop(0)
	to := (from + 1) % len(columns)
	Clap(clapper, 0, -1, 1, columns[to])
	builder := printing.NewBuilder()
	for _, column := range columns {
		builder.Printf("%d", column.Nth(0))
	}
	return parseInt(builder.Inner().String())
}
func Clap(clapper, claps, at, direction int, column *List[int]) {
	if clapper == claps {
		if direction == 1 {
			column.Insert(at, clapper)
		} else {
			column.Insert(at+1, clapper)
		}
		return
	}
	if claps < clapper {
		claps++
		at += direction
	}
	if at >= column.Len() {
		at--
		direction *= -1
	} else if at < 0 {
		at++
		direction *= -1
	}
	Clap(clapper, claps, at, direction, column)
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
