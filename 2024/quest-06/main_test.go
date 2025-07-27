package quest_06

import (
	"bufio"
	"iter"
	"strings"
	"sync"
	"testing"

	"github.com/mdw-go/must"
	"github.com/mdw-go/must/osmust"
)

func TestPart1(t *testing.T) {
	t.Log(FullPath(FindPathWithUniqueLength(YieldPaths(ParseTreeFile("part1-sample.txt")))))
	t.Log(FullPath(FindPathWithUniqueLength(YieldPaths(ParseTreeFile("part1-actual.txt")))))

	t.Log(FirstLettersOnly(FindPathWithUniqueLength(YieldPaths(ParseTreeFile("part2-actual.txt")))))
}

func FullPath(path string) string {
	return strings.ReplaceAll(path, "|", "")
}
func FirstLettersOnly(path string) (result string) {
	for element := range strings.SplitSeq(path, "|") {
		result += string(element[0])
	}
	return result
}

func FindPathWithUniqueLength(paths iter.Seq[string]) (result string) {
	lengths := make(map[int][]string)
	for path := range paths {
		lengths[len(path)] = append(lengths[len(path)], path)
	}
	for _, paths := range lengths {
		if len(paths) == 1 {
			return paths[0]
		}
	}
	panic("nope")
}
func YieldPaths(tree map[string][]string) iter.Seq[string] {
	return func(yield func(string) bool) {
		results := make(chan string)
		closer := new(sync.Once)
		go Traverse(closer, results, tree, "", "RR")
		for result := range results {
			if !yield(result) {
				return
			}
		}
	}
}
func Traverse(closer *sync.Once, yield chan string, tree map[string][]string, path, node string) {
	if len(tree) == 0 {
		closer.Do(func() { close(yield) })
		return
	}
	longerPath := strings.TrimPrefix(path+"|"+node, "|")
	if node == "@" {
		yield <- longerPath
		return
	}
	branches := tree[node]
	delete(tree, node)
	for _, branch := range branches {
		Traverse(closer, yield, tree, longerPath, branch)
	}
}
func ParseTreeFile(path string) map[string][]string {
	file := osmust.Open(path)
	defer must.Defer(file.Close)()
	return ParseTree(bufio.NewScanner(file))
}
func ParseTree(scanner *bufio.Scanner) (result map[string][]string) {
	result = make(map[string][]string)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		root, nodes, ok := strings.Cut(line, ":")
		if !ok {
			continue
		}
		result[root] = strings.Split(nodes, ",")
	}
	return result
}
