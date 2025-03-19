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
