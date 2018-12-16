package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

func ReadInput(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func countTwoThree(hash string) (bool, bool) {
	m := map[rune]int{}
	for _, char := range hash {
		m[char]++
	}

	two, three := false, false
	for _, v := range m {
		if v == 2 {
			two = true
		} else if v == 3 {
			three = true
		}
	}
	return two, three
}

func Part1(hashes []string) int {
	twoS, threeS := 0, 0
	for _, hash := range hashes {
		two, three := countTwoThree(hash)
		if two {
			twoS++
		}
		if three {
			threeS++
		}
	}
	return twoS * threeS
}

func difference(hash1 string, hash2 string) int {
	diff := 0
	hash2Runes := []rune(hash2)
	for i, char1 := range hash1 {
		if i >= utf8.RuneCountInString(hash2) {
			diff += utf8.RuneCountInString(hash1) - utf8.RuneCountInString(hash2)
			break
		}
		if char1 != hash2Runes[i] {
			diff++
		}
	}
	return diff
}

func Part2(hashes []string) string {
	if len(hashes) < 2 {
		return ""
	}
	diff := difference(hashes[0], hashes[1])
	i1, i2 := 0, 1
	for i, hash1 := range hashes {
		for j, hash2 := range hashes[i+1:] {
			d := difference(hash1, hash2)
			if d < diff {
				diff = d
				i1 = i
				i2 = j + i + 1
			}
		}
	}

	hash1 := hashes[i1]
	hash2 := hashes[i2]
	hash2Runes := []rune(hash2)
	output := ""
	for i, r1 := range hash1 {
		if i >= utf8.RuneCountInString(hash2) {
			break
		}
		if r1 == hash2Runes[i] {
			output += string(r1)
		}
	}
	return output
}

func main() {
	hashes := ReadInput(os.Args[1])

	fmt.Println(Part1(hashes))
	fmt.Println(Part2(hashes))

}
