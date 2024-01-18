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
		if resSlice[i].count == resSlice[j].count {
			return resSlice[i].word < resSlice[j].word
		}
		return resSlice[i].count > resSlice[j].count
	})

	result := make([]string, 0, 10)
	for _, v := range resSlice[:10] {
		result = append(result, v.word)
	}
	return result
}
