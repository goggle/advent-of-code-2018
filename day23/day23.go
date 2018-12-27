package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Nanobot struct {
	x int
	y int
	z int
	r int
}

func (nb *Nanobot) Distance(other *Nanobot) int {
	return Abs(nb.x-other.x) + Abs(nb.y-other.y) + Abs(nb.z-other.z)
}

func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func ReadInput(path string) []Nanobot {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	scanner := bufio.NewScanner(file)
	regex := regexp.MustCompile(`pos\s*=\s*<\s*(-?\d+)\s*,\s*(-?\d+)\s*,\s*(-?\d+)\s*>\s*,\s*r\s*=\s*(\d+)`)

	out1, out2, out3, out4 := 0, 0, 0, 0
	nanobots := []Nanobot{}
	for scanner.Scan() {
		line := scanner.Text()
		match := regex.FindStringSubmatch(line)
		if len(match) > 4 {
			out1, _ = strconv.Atoi(match[1])
			out2, _ = strconv.Atoi(match[2])
			out3, _ = strconv.Atoi(match[3])
			out4, _ = strconv.Atoi(match[4])
		}
		nanobots = append(nanobots, Nanobot{out1, out2, out3, out4})
	}
	return nanobots
}

func getStrongest(nanobots []Nanobot) *Nanobot {
	maxRadius := 0
	var strongest *Nanobot
	for i, nanobot := range nanobots {
		if nanobot.r > maxRadius {
			maxRadius = nanobot.r
			strongest = &nanobots[i]
		}
	}
	return strongest
}

func Part1(nanobots []Nanobot) int {
	strongest := getStrongest(nanobots)
	count := 0
	for _, nb := range nanobots {
		if nb.Distance(strongest) <= strongest.r {
			count++
		}
	}
	return count
}

func Part2(nanobots []Nanobot) (int, int, int) {
	return 0, 0, 0
}

func main() {
	nanobots := ReadInput(os.Args[1])
	fmt.Println(Part1(nanobots))
	fmt.Println(Part2(nanobots))
}
