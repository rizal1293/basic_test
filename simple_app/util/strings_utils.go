package util

import (
	"bytes"
	"strings"
)

type WordCounts struct {
	Word  string
	Count int
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

func WordStatistic(str string) map[string]interface{} {
	wordList := strings.Fields(str)
	counts := make(map[string]int)

	result := make(map[string]interface{}, 0)

	for _, word := range wordList {
		_, found := counts[word]
		if found {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}

	countOneWord := 0
	highest := WordCounts{}
	smallest := WordCounts{}

	for word, count := range counts {
		if count == 1 {
			countOneWord++
		}

		if highest == (WordCounts{}) {
			highest = WordCounts{
				Word:  word,
				Count: count,
			}
		} else {
			if count > highest.Count {
				highest = WordCounts{
					Word:  word,
					Count: count,
				}
			}
		}

		if smallest == (WordCounts{}) {
			smallest = WordCounts{
				Word:  word,
				Count: count,
			}
		} else {
			if count < smallest.Count {
				smallest = WordCounts{
					Word:  word,
					Count: count,
				}
			}
		}
	}

	result = map[string]interface{}{
		"TotalWordCounts":        len(wordList),
		"WordCountEveryWord":     counts,
		"NumberOfWordShowUpOnce": countOneWord,
		"WordHighestCount":       highest,
		"WordSmallestCount":      smallest,
	}

	return result
}
