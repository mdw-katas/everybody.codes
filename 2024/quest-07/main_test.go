package quest_07

import (
	"bufio"
	"bytes"
	"os"
	"sort"
	"strings"
	"testing"

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
