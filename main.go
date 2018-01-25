package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var table = map[string]int{}
var words = []string{}
var (
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

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
		table[scanner.Text()] = len(scanner.Text())
	}

	for i := 0; i < len(words)-1; i++ {
		if len(words[i+1]) < len(words[i]) {
			continue
		}

		if strings.HasPrefix(words[i+1], words[i]) {
			if scanSuffix(strings.TrimPrefix(words[i+1], words[i])) == "" {
				totalWords++
				if len(words[i+1]) >= lenLong {
					lenLong = len(words[i+1])
					longestWord = words[i+1]
					continue
				}
				if len(words[i+1]) >= lenNotSoLong {
					lenNotSoLong = len(words[i+1])
					preLongestWord = words[i+1]
					continue
				}
			}
		}
	}

	fmt.Printf("the longest concatenated word: %s\n", longestWord)
	fmt.Printf("2nd longest concatenated word: %s\n", preLongestWord)
	fmt.Printf("total count of concatenated words - %d\n", totalWords)
}

func scanSuffix(s string) string {
	if len(s) == 0 {
		return ""
	}
	l, ok := table[s]
	if ok {
		return scanSuffix(string(s[l:]))
	}
	if len(s) == 1 {
		return " "
	}
	return scanSuffix(string(s[:len(s)-1]))
}
