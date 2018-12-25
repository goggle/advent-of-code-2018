package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func ReadInput(path string) [][4]int {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	scanner := bufio.NewScanner(file)
	regex := regexp.MustCompile(`position\s*=\s*<\s*(-?\d+)\s*,\s*(-?\d+)\s*>\s*velocity\s*=\s*<\s*(-?\d+)\s*,\s*(-?\d+)\s*>`)

	out1, out2, out3, out4 := 0, 0, 0, 0
	output := [][4]int{}
	for scanner.Scan() {
		line := scanner.Text()
		match := regex.FindStringSubmatch(line)
		if len(match) > 2 {
			out1, _ = strconv.Atoi(match[1])
			out2, _ = strconv.Atoi(match[2])
			out3, _ = strconv.Atoi(match[3])
			out4, _ = strconv.Atoi(match[4])
			out := [4]int{out1, out2, out3, out4}
			output = append(output, out)
		}
	}
	return output
}

// func Part1([][4]input]) int {
// 	return 0

// }

// func Part2(p[][4]input]) int {
// 	return 0

// }

func main() {
	input := ReadInput(os.Args[1])
	fmt.Println(input)
	fmt.Println(len(input))

}
