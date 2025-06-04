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
	t.Log(runicWords)
	rows := lines[2:]
	t.Log(len(rows))
}

type StringLoop struct {
	haystack string
	reversed string
	marked   []int
}

func NewStringLoop(s string) *StringLoop {
	return &StringLoop{
		haystack: s + s,
		reversed: reverse(s + s),
		marked:   make([]int, len(s)),
	}
}

func (this *StringLoop) Find(needle string) {
	// TODO
}
