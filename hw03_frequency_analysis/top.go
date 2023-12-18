package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(inputStr string) []string {
	if inputStr == "" {
		return []string{}
	}
	type words struct {
		word  string
		count int
	}
	strSlice := strings.Fields(inputStr)
	mapWords := make(map[string]int)

	for i, v := range strSlice {
		if _, ok := mapWords[strSlice[i]]; ok {
			mapWords[strSlice[i]]++
		} else {
			mapWords[v] = 1
		}
	}

	resSlice := make([]words, 0, len(mapWords))

	for key, v := range mapWords {
		word := words{word: key, count: v}
		resSlice = append(resSlice, word)
	}

	sort.Slice(resSlice, func(i, j int) bool {
		return resSlice[i].count > resSlice[j].count
	})

	s := make([][]string, 0, len(resSlice))
	subSlice := []string{}
	counter := 0

	for i := range resSlice {
		if i > 9 && resSlice[i].count != resSlice[i-1].count {
			break
		}
		if resSlice[i].count != resSlice[i+1].count {
			subSlice = append(subSlice, resSlice[i].word)
			s = append(s, subSlice)
			subSlice = []string{}
			continue
		} else {
			subSlice = append(subSlice, resSlice[i].word)
			counter++
		}
	}
	resSL := make([]string, 0)
	for _, v := range s {
		sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })
		resSL = append(resSL, v...)
	}

	return resSL[:10]
}
