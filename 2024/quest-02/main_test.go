package quest_02

import (
	"os"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	raw, err := os.ReadFile("notes-1.txt")
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(string(raw), "\n")
	_, joinedRunicWords, _ := strings.Cut(lines[0], ":")
	runicWords := strings.Split(joinedRunicWords, ",")
	t.Log(runicWords)
	count := 0
	corpus := lines[2]
	for _, word := range runicWords {
		count += strings.Count(corpus, word)
	}
	if count != 32 {
		t.Error("expected 32 runic words, got ", count)
	}
}

func TestPart2(t *testing.T) {
	raw, err := os.ReadFile("notes-2.txt")
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(string(raw), "\n")
	_, joinedRunicWords, _ := strings.Cut(lines[0], ":")
	runicWords := strings.Split(joinedRunicWords, ",")
	t.Log(runicWords)
	count := 0
	for _, line := range lines[2:] {
		runic := make(map[int]struct{})
		for _, word := range runicWords {
			for n := range len(line) {
				if n+len(word) > len(line) {
					continue
				}
				if line[n:n+len(word)] == word {
					for offset := range len(word) {
						runic[n+offset] = struct{}{}
					}
				}
				if line[n:n+len(word)] == reverse(word) {
					for offset := range len(word) {
						runic[n+offset] = struct{}{}
					}
				}
			}
		}
		count += len(runic)
	}
	if count != 5237 {
		t.Error("expected 5237 runic symbols, got ", count)
	}
}

func reverse(s string) string {
	var result []rune
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, rune(s[i]))
	}
	return string(result)
}

func TestPart3(t *testing.T) {
	raw, err := os.ReadFile("notes-3.txt")
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(string(raw), "\n")
	_, joinedRunicWords, _ := strings.Cut(lines[0], ":")
	runicWords := strings.Split(joinedRunicWords, ",")
	rows := lines[2:]
	left2right := make([]*Loop, len(rows))
	right2left := make([]*Loop, len(rows))
	ups := make([]*Vertical, len(rows[0]))
	downs := make([]*Vertical, len(rows[0]))
	var all []*Char
	for r, row := range rows {
		left2right[r] = &Loop{}
		right2left[r] = &Loop{}
		for c, char := range row {
			v := &Char{Row: r, Col: c, Rune: char}
			all = append(all, v)
			if r == 0 {
				ups[c] = &Vertical{}
				downs[c] = &Vertical{}
			}
			left2right[r].chars = append(left2right[r].chars, v)
			right2left[r].chars = append([]*Char{v}, right2left[r].chars...) // reverse order
			ups[c].chars = append(ups[c].chars, v)
			downs[c].chars = append([]*Char{v}, downs[c].chars...) // reverse order
		}
	}
	for _, word := range runicWords {
		for _, loop := range left2right {
			loop.Find(word)
		}
		for _, loop := range right2left {
			loop.Find(word)
		}
		for _, column := range ups {
			column.Find(word)
		}
		for _, column := range downs {
			column.Find(word)
		}
	}
	count := 0
	for _, c := range all {
		if c.Runic {
			count++
		}
	}
	if count != 11967 {
		t.Error("expected 11967 runic symbols, got ", count)
	}
}

type Char struct {
	Row   int
	Col   int
	Rune  rune
	Runic bool
}

type Loop struct {
	chars []*Char
}

func (this *Loop) Find(word string) {
	s := this.String()
	s = s + s // double the string to account for wrap-around
	for c := range len(s)/2 + 1 {
		if s[c:c+len(word)] == word {
			for cc := c; cc < c+len(word); cc++ {
				this.chars[cc%len(this.chars)].Runic = true
			}
		}
	}
}
func (this *Loop) String() string {
	return String(this.chars)
}

type Vertical struct {
	chars []*Char
}

func (this *Vertical) Find(word string) {
	s := this.String()
	for c := range len(s) - len(word) + 1 {
		if s[c:c+len(word)] == word {
			for cc := c; cc < c+len(word); cc++ {
				this.chars[cc].Runic = true
			}
		}
	}
}
func (this *Vertical) String() string {
	return String(this.chars)
}

func String(chars []*Char) string {
	result := ""
	for _, char := range chars {
		result += string(char.Rune)
	}
	return result
}
