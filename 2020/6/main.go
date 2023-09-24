package main

import (
	"bufio"
	"fmt"
	"os"
)

func countUniqueQuestions(questions [26]int) int {
	var count int
	for _, q := range questions {
		if q > 0 {
			count++
		}
	}
	return count
}

func countCommonQuestions(questions [26]int, people int) int {
	var count int
	for _, q := range questions {
		if q == people {
			count++
		}
	}
	return count
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var countAny, countAll, currPeople int
	var questions [26]int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			countAny += countUniqueQuestions(questions)
			countAll += countCommonQuestions(questions, currPeople)
			clear(questions[:])
			currPeople = 0
		} else {
			currPeople++
			for i := 0; i < len(line); i++ {
				questions[line[i]-'a']++
			}
		}
	}
	countAny += countUniqueQuestions(questions)
	countAll += countCommonQuestions(questions, currPeople)
	fmt.Println("[1] count unique:", countAny)
	fmt.Println("[2] count common:", countAll)
}
