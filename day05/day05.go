package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

func ReadInput(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic(err)
	}
	return strings.TrimSuffix(string(content), "\n")
}

func collides(a rune, b rune) bool {
	if unicode.IsLower(a) && unicode.IsUpper(b) {
		if unicode.ToUpper(a) == b {
			return true
		}
	} else if unicode.IsUpper(a) && unicode.IsLower(b) {
		if unicode.ToLower(a) == b {
			return true
		}
	}
	return false
}

func Part1(input string) int {
	charList := []rune(input)
	modified := true

	for modified {
		modified = false
		i := 0
		for i < len(charList) {
			if i+1 < len(charList) && collides(charList[i], charList[i+1]) {
				modified = true
				charList = append(charList[:i], charList[i+2:]...)
			}
			i++
		}
	}
	return len(charList)
}

func getInputChars(input string) map[rune]struct{} {
	m := map[rune]struct{}{}
	for _, char := range input {
		m[unicode.ToLower(char)] = struct{}{}
	}
	return m
}

func Part2(input string) int {
	min := Part1(input)
	charsLowercase := getInputChars(input)

	for charL := range charsLowercase {
		inputList := []rune(input)
		reducedInputList := []rune{}
		for _, c := range inputList {
			if c == charL || unicode.ToLower(c) == charL {
				continue
			}
			reducedInputList = append(reducedInputList, c)
		}
		reducedInput := string(reducedInputList)
		res := Part1(reducedInput)
		if res < min {
			min = res
		}
	}

	return min
}

func main() {
	input := ReadInput(os.Args[1])
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
