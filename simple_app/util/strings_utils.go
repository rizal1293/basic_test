package util

import (
	"bytes"
	"strings"
)

type WorldStatisticResult struct {
	TotalWordCounts        int
	WordCountEveryWord     map[string]int
	NumberOfWordShowUpOnce int
	WordHighestCount       string
	WordSmallestCount      string
}

func TitleLower(str string) string {
	if str != "" {
		fields := strings.Fields(str)

		var buffer bytes.Buffer
		for i, c := range fields {
			lower := strings.ReplaceAll(c, c[:1], strings.ToLower(c[:1]))
			upper := strings.ReplaceAll(lower, lower[1:], strings.ToUpper(lower[1:]))

			if i == len(str)-1 {
				buffer.WriteString(upper)
			} else {
				buffer.WriteString(upper + " ")
			}
		}
		return buffer.String()
	}
	return ""
}

func WordStatistic(str string) WorldStatisticResult {
	wordList := strings.Fields(str)
	counts := make(map[string]int)

	result := WorldStatisticResult{}

	countOneWord := 0

	for _, word := range wordList {
		_, ok := counts[word]
		if ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}

	result.TotalWordCounts = len(wordList)
	result.WordCountEveryWord = counts

	highestCount := 0
	smallestCount := 0

	wordHighest := ""
	wordSmallest := ""

	index := 0

	for word, count := range counts {

		// words that only show up once
		if count == 1 {
			countOneWord++
		}

		if index == 0 {
			highestCount = count
			smallestCount = count

			wordHighest = word
			wordSmallest = word
		}

		// find highest count of word
		if count > highestCount {
			highestCount = count
			wordHighest = word
		}

		// find smmallest count of word
		if count < smallestCount {
			smallestCount = count
			wordSmallest = word
		}

		index++
	}

	result.NumberOfWordShowUpOnce = countOneWord
	result.WordHighestCount = wordHighest
	result.WordSmallestCount = wordSmallest

	return result
}
