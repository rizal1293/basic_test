package util

import (
	"bytes"
	"strings"
)

type HeighestWord struct {
	Word  string
	Count int
}

type SmallestWord struct {
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

	chanCountOneWord := make(chan int, 0)
	chanHighestWords := make(chan HeighestWord)
	chanSmallestWords := make(chan SmallestWord)

	for _, word := range wordList {
		_, found := counts[word]
		if found {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}

	// Number of words that only show up once
	go func() {
		c := 0
		highest := HeighestWord{}

		for word, count := range counts {
			if count == 1 {
				c++
			}

			if highest == (HeighestWord{}) {
				highest = HeighestWord{
					Word:  word,
					Count: count,
				}
			} else {
				if count > highest.Count {
					highest = HeighestWord{
						Word:  word,
						Count: count,
					}
				}
			}
		}
		chanCountOneWord <- c
	}()

	// Word that has the highest count
	go func() {
		highest := HeighestWord{}
		for word, count := range counts {
			if highest == (HeighestWord{}) {
				highest = HeighestWord{
					Word:  word,
					Count: count,
				}
			} else {
				if count > highest.Count {
					highest = HeighestWord{
						Word:  word,
						Count: count,
					}
				}
			}
		}
		chanHighestWords <- highest
	}()

	// Word that has the smallest count
	go func() {
		smallest := SmallestWord{}
		for word, count := range counts {
			if smallest == (SmallestWord{}) {
				smallest = SmallestWord{
					Word:  word,
					Count: count,
				}
			} else {
				if count < smallest.Count {
					smallest = SmallestWord{
						Word:  word,
						Count: count,
					}
				}
			}
		}
		chanSmallestWords <- smallest
	}()

	result = map[string]interface{}{
		"TotalWordCounts":        len(wordList),
		"WordCountEveryWord":     counts,
		"NumberOfWordShowUpOnce": <-chanCountOneWord,
		"WordHighestCount":       <-chanHighestWords,
		"WordSmallestCount":      <-chanSmallestWords,
	}

	close(chanCountOneWord)
	close(chanHighestWords)
	close(chanSmallestWords)

	return result
}
