package bulksearch

import (
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

var readFile = func(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	return string(bytes), err
}

var getAbs = filepath.Abs
type Result struct {
	Term string
	FileCount int
}

func Search(searchTerms []string, dir string) []Result {
	results := []Result{}
	for _, term := range searchTerms {
		results = append(results, Result{Term: term, FileCount: SearchFor(term, dir)})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].FileCount > results[j].FileCount
	})
	return results
}

func SearchFor(term string, dir string) int {
	modTerm := strings.Replace(term, "-", "\\-", -1)

	out, _ := exec.Command("rg", "-Ic", modTerm, dir).Output()
	return len(strings.Split(string(out), "\n")) - 1
}
