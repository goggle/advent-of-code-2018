package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Map [][]Cell

type Cell struct {
	id        int
	claimedBy map[int]struct{}
}

func ReadInput(path string) [][2]int {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	scanner := bufio.NewScanner(file)
	regex := regexp.MustCompile(`(\d+)\s*,\s*(\d+)`)

	parsedInput := [][2]int{}
	for scanner.Scan() {
		line := scanner.Text()
		match := regex.FindStringSubmatch(line)
		coordinates := [2]int{}
		if len(match) > 2 {
			for i := 0; i < 2; i++ {
				coordinates[i], _ = strconv.Atoi(match[i+1])
			}
			parsedInput = append(parsedInput, coordinates)
		}
	}
	return parsedInput
}

func createMap(coordinates [][2]int) Map {
	maxX, maxY := 0, 0
	for _, c := range coordinates {
		if c[0] > maxX {
			maxX = c[0]
		}
		if c[1] > maxY {
			maxY = c[1]
		}
	}

	// fmt.Println(maxX, maxY)

	m := make([][]Cell, maxX+1)
	for i := range m {
		m[i] = make([]Cell, maxY+1)
		for j := range m[i] {
			m[i][j].claimedBy = map[int]struct{}{}
		}
	}
	return m
}

func initMap(coordinates [][2]int) Map {
	m := createMap(coordinates)
	id := 1
	for _, cord := range coordinates {
		m[cord[0]][cord[1]].id = id
		id++
	}
	return m
}

func (m *Map) getNeighbours(x int, y int) [][2]int {
	neighbours := [][2]int{}
	if x >= 0 && x < len(*m) {
		if y >= 0 && y < len((*m)[x]) {
			possibleNeighbours := [][2]int{
				[2]int{x - 1, y},
				[2]int{x + 1, y},
				[2]int{x, y - 1},
				[2]int{x, y + 1},
			}
			for _, c := range possibleNeighbours {
				if c[0] >= 0 && c[0] < len(*m) && c[1] >= 0 && c[1] < len((*m)[c[0]]) {
					neighbours = append(neighbours, [2]int{c[0], c[1]})
				}
			}
		}
	}
	return neighbours
}

func (m *Map) claim() {
	for i := range *m {
		for j, cell := range (*m)[i] {
			if cell.id > 0 {
				neighbours := m.getNeighbours(i, j)
				for _, coord := range neighbours {
					// fmt.Println(coord[0], coord[1])
					if (*m)[coord[0]][coord[1]].id == 0 {
						(*m)[coord[0]][coord[1]].claimedBy[cell.id] = struct{}{}
					}
				}
			}
		}
	}
}

func (m *Map) resolve() bool {
	resolved := false
	for i := range *m {
		for j, cell := range (*m)[i] {
			if cell.id == 0 {
				if len(cell.claimedBy) == 1 {
					for k := range cell.claimedBy {
						// fmt.Println(cell.id)
						(*m)[i][j].id = k
						resolved = true
					}
				}
			}
		}
	}
	return resolved
}

func countFiniteAreas(m Map) map[int]int {
	countMap := map[int]int{}
	boundaryIds := map[int]struct{}{}
	for i := range m {
		for j, cell := range m[i] {
			if cell.id != 0 {
				countMap[cell.id]++
			}
			if i == 0 || j == 0 || i == len(m) || j == len(m[i]) {
				boundaryIds[cell.id] = struct{}{}
			}
		}
	}
	for k := range boundaryIds {
		countMap[k] = 0
	}
	return countMap
}

func Part1(coordinates [][2]int) int {
	m := initMap(coordinates)
	for {
		m.claim()
		resolved := m.resolve()
		if !resolved {
			break
		}
	}
	countMap := countFiniteAreas(m)

	max := 0
	for _, v := range countMap {
		if v > max {
			max = v
		}
	}
	return max
}

func ManhattenDistance(p1 [2]int, p2 [2]int) int {
	v1 := p1[0] - p2[0]
	v2 := p1[1] - p2[1]
	if v1 < 0 {
		v1 = -v1
	}
	if v2 < 0 {
		v2 = -v2
	}
	return v1 + v2

}

func Part2(coordinates [][2]int, limit int) int {
	m := initMap(coordinates)
	sum := 0
	for i := range m {
		for j := range m[i] {
			dist := 0
			for _, c := range coordinates {
				dist += ManhattenDistance(c, [2]int{i, j})
			}
			if dist < limit {
				sum++
			}
		}
	}

	return sum
}

func main() {
	parsedInput := ReadInput(os.Args[1])
	fmt.Println(Part1(parsedInput))
	fmt.Println(Part2(parsedInput, 10000))
}
