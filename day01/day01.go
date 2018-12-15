package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ParseInput(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	scanner := bufio.NewScanner(file)
	numbers := []int{}
	for scanner.Scan() {
		number, errN := strconv.Atoi(scanner.Text())
		if errN != nil {
			log.Panic(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func Part1(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func Part2(numbers []int) int {
	currentFrequency := 0
	freqCount := map[int]int{}
	freqCount[0] = 1
	for {
		for _, v := range numbers {
			currentFrequency += v
			freqCount[currentFrequency]++
			if freqCount[currentFrequency] > 1 {
				return currentFrequency
			}
		}

	}
}

func main() {
	numbers := ParseInput(os.Args[1])
	fmt.Println(Part1(numbers))
	fmt.Println(Part2(numbers))
}
