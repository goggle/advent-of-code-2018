package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Rectangle struct {
	id     int
	x      int
	y      int
	width  int
	height int
}

func ReadInput(path string) []Rectangle {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	scanner := bufio.NewScanner(file)
	rectangles := []Rectangle{}
	regex := regexp.MustCompile(`#(\d+)\s*@\s*(\d+),(\d+)\s*:\s*(\d+)x(\d+)`)
	for scanner.Scan() {
		line := scanner.Text()
		match := regex.FindStringSubmatch(line)
		if len(match) > 5 {
			vals := [5]int{}
			for i := 0; i < 5; i++ {
				vals[i], _ = strconv.Atoi(match[i+1])
			}
			rectangle := Rectangle{}
			rectangle.id = vals[0]
			rectangle.x = vals[1]
			rectangle.y = vals[2]
			rectangle.width = vals[3]
			rectangle.height = vals[4]
			rectangles = append(rectangles, rectangle)
		}
	}
	return rectangles
}

func createOrganization(rectangles []Rectangle) map[[2]int][]int {
	org := map[[2]int][]int{}
	for _, rect := range rectangles {
		for x := rect.x; x < rect.x+rect.width; x++ {
			for y := rect.y; y < rect.y+rect.height; y++ {
				org[[2]int{x, y}] = append(org[[2]int{x, y}], rect.id)
			}
		}
	}

	return org
}

func Part1(rectangles []Rectangle) int {
	org := createOrganization(rectangles)
	sum := 0
	for _, v := range org {
		if len(v) > 1 {
			sum++
		}
	}
	return sum
}

func Part2(rectangles []Rectangle) int {
	org := createOrganization(rectangles)
	single_ids := map[int]bool{}
	for _, rect := range rectangles {
		single_ids[rect.id] = true
	}

	for _, v := range org {
		if len(v) > 1 {
			for _, id := range v {
				single_ids[id] = false
			}
		}
	}

	for id, single := range single_ids {
		if single {
			return id
		}
	}

	return 0
}

func main() {
	rectangles := ReadInput(os.Args[1])

	fmt.Println(Part1(rectangles))
	fmt.Println(Part2(rectangles))

}
