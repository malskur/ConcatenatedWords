package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var table = map[string]int{}
var words = []string{}

var (
	wordLen        = 0
	lenLong        = 0
	lenNotSoLong   = 0
	totalWords     = 0
	longestWord    string
	preLongestWord string
)

func main() {
	f, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer f.Close()

	//read file, create map (dictionary) table and array words
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
		table[scanner.Text()] = len(scanner.Text())
	}

	//measure execution time (just for fun)
	start := time.Now().Nanosecond()

	//search every word
	for i := 0; i < len(words); i++ {
		wordLen = len(words[i])
		if wordLen == 0 {
			continue
		}

		//scan subwords in word
		if scanPrefix(words[i]) {
			totalWords++

			//compare words lengths
			if wordLen >= lenLong {
				lenNotSoLong = lenLong
				preLongestWord = longestWord
				lenLong = wordLen
				longestWord = words[i]
				continue
			}
			if len(words[i]) >= lenNotSoLong {
				lenNotSoLong = wordLen
				preLongestWord = words[i]
				continue
			}
		}
	}

	finish := time.Now().Nanosecond()
	fmt.Printf("executed in: %d\n", finish-start)

	fmt.Printf("the longest concatenated word: %s\n", longestWord)
	fmt.Printf("2nd longest concatenated word: %s\n", preLongestWord)
	fmt.Printf("total count of concatenated words - %d\n", totalWords)
}

//recursively compare every prefix to words in dictionary
func scanPrefix(s string) bool {
	if len(s) == 0 {
		return true
	}
	for i := 1; i <= len(s); i++ {
		_, ok := table[s[:i]]
		if ok {
			if i == wordLen {
				return false
			}
			if !scanPrefix(string(s[i:])) {
				continue
			} else {
				return scanPrefix(s[i:])
			}
		}
	}
	return false
}
