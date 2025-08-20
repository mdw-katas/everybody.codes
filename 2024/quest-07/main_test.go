package quest_07

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/mdw-go/generic-list/list"
	"github.com/mdw-go/testing/v2/assert"
	"github.com/mdw-go/testing/v2/should"
)

func readFile(filename string) *bytes.Buffer {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(content)
}

func TestPart1(t *testing.T) {
	assert.So(t, Part1(readFile("part1-sample.txt")), should.Equal, "BDCA")
	assert.So(t, Part1(readFile("part1-full.txt")), should.Equal, "HKFJGDBAI")
}

func Part1(reader *bytes.Buffer) string {
	var tracks []*Track
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		track := ParseTrack(scanner.Text())
		track.Run(10, 10)
		tracks = append(tracks, track)
	}
	sort.Slice(tracks, func(i, j int) bool {
		return tracks[i].Result > tracks[j].Result
	})
	var rank strings.Builder
	for _, track := range tracks {
		rank.WriteString(track.Name)
	}
	return rank.String()
}

type Track struct {
	Name   string
	Steps  []int
	Result int
}

func ParseTrack(line string) *Track {
	name, rawSteps, ok := strings.Cut(line, ":")
	if !ok {
		panic("nope")
	}
	result := &Track{Name: name}
	for _, rawStep := range strings.Split(rawSteps, ",") {
		switch rawStep {
		case "+":
			result.Steps = append(result.Steps, 1)
		case "-":
			result.Steps = append(result.Steps, -1)
		default:
			result.Steps = append(result.Steps, 0)
		}
	}
	return result
}

func (this *Track) Run(stepCount, power int) {
	for x := range stepCount {
		power += this.Steps[x%len(this.Steps)]
		this.Result += power
	}
}

var sampleTrack = strings.NewReader(strings.TrimSpace(`
S+===
-   +
=+=-+
`))

func TestPart2(t *testing.T) {
	log.SetOutput(t.Output())
	assert.So(t, ParseLoopTrack(sampleTrack), should.Equal, "S+===++-=+=-")
}

func ParseLoopTrack(track io.Reader) string {
	segments := list.New[rune]()
	scanner := bufio.NewScanner(track)
	for scanner.Scan() {
		line := scanner.Text()
		if segments.Len() == 0 {
			for _, c := range line {
				segments.PushBack(c)
			}
		} else if strings.Count(line, " ") > 0 {
			segments.PushFront(rune(line[0]))
			segments.PushBack(rune(line[len(line)-1]))
			log.Println(render(segments))
		} else {
			for _, c := range line {
				segments.PushFront(c)
			}
		}
	}
	result := render(segments)
	return result
}

func render(segments *list.List[rune]) string {
	before := new(strings.Builder)
	after := new(strings.Builder)
	target := after
	for e := segments.Front(); e != nil; e = e.Next() {
		if e.Value == 'S' {
			target = before
		}
		target.WriteRune(e.Value)
	}
	before.WriteString(after.String())
	return before.String()
}
