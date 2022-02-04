package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Word struct {
	Val       string
	Frequency uint
}

type Words []Word

func (w Words) Len() int      { return len(w) }
func (w Words) Swap(i, j int) { w[i], w[j] = w[j], w[i] }
func (w Words) Less(i, j int) bool {
	return w[i].Frequency > w[j].Frequency || w[i].Frequency == w[j].Frequency && w[i].Val < w[j].Val
}

func (w Words) getValues(n int) []string {
	if n > w.Len() {
		n = w.Len()
	}
	values := make([]string, n)
	for i := 0; i < n; i++ {
		values[i] = w[i].Val
	}
	return values
}

func Top10(str string) []string {
	data := make(map[string]uint)
	for _, val := range strings.Fields(str) {
		data[val]++
	}
	words := make(Words, 0, len(data))
	for k, v := range data {
		words = append(words, Word{k, v})
	}
	sort.Sort(words)
	return words.getValues(10)
}
