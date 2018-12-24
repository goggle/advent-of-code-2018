package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Circle struct {
	circle  map[int]*Node
	current *Node
	start   *Node
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (c Circle) Print() {
	if len(c.circle) == 0 {
		return
	}
	l := []int{c.start.value}
	next := c.start.right
	for next != c.start {
		l = append(l, next.value)
		next = next.right
	}
	s := ""
	for i, val := range l {
		if i != 0 {
			s += " "
		}
		if val == c.current.value {
			s += fmt.Sprintf("(%d)", val)
		} else {
			s += fmt.Sprintf("%d", val)
		}
	}
	fmt.Println(s)
}

func (c *Circle) Init() {
	c.circle = map[int]*Node{}
	c.current = nil
	c.start = nil
}

func (c *Circle) Insert(key int) {
	if len(c.circle) == 0 {
		node := Node{key, nil, nil}
		node.left = &node
		node.right = &node
		c.circle[key] = &node
		c.current = &node
		c.start = &node
		return
	}
	onePtr := c.current.right
	twoPtr := onePtr.right

	newNode := Node{key, onePtr, twoPtr}
	onePtr.right = &newNode
	twoPtr.left = &newNode

	c.current = &newNode
}

func (c *Circle) Remove() int {
	remNodePtr := c.current
	// fmt.Println(remNodePtr)
	for i := 0; i < 7; i++ {
		// fmt.Println(remNodePtr.value)
		remNodePtr = remNodePtr.left
	}
	lPtr := remNodePtr.left
	rPtr := remNodePtr.right
	lPtr.right = rPtr
	rPtr.left = lPtr

	val := remNodePtr.value
	// delete(c.circle, val)
	c.current = rPtr

	return val
}

func ReadInput(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	scanner := bufio.NewScanner(file)
	regex := regexp.MustCompile(`(\d+).+\s+(\d+)`)

	out1, out2 := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		match := regex.FindStringSubmatch(line)
		if len(match) > 2 {
			out1, _ = strconv.Atoi(match[1])
			out2, _ = strconv.Atoi(match[2])
		}
		return out1, out2
	}
	return out1, out2
}

// func Part1(players int, worth int) int {
// 	playerScores := make([]int, players)
// 	// circle := []int{0}
// 	circle := make([]int, 0, worth)
// 	circle = append(circle, 0)
// 	playerIndex := 0
// 	currentIndex := 0

// 	nextMarble := 1
// 	for nextMarble <= worth {
// 		insertPosition := (currentIndex + 2) % (len(circle))
// 		if nextMarble%23 == 0 {
// 			playerScores[playerIndex] += nextMarble
// 			remIndex := (currentIndex - 7)
// 			for remIndex < 0 {
// 				remIndex += len(circle)
// 			}
// 			playerScores[playerIndex] += circle[remIndex]
// 			circle = append(circle[:remIndex], circle[remIndex+1:]...)
// 			currentIndex = remIndex
// 			playerIndex = (playerIndex + 1) % players
// 			nextMarble++
// 			continue
// 		}
// 		if insertPosition == 0 {
// 			insertPosition = len(circle)
// 			circle = append(circle, nextMarble)
// 		} else {
// 			circle = append(circle[:insertPosition], append([]int{nextMarble}, circle[insertPosition:]...)...)
// 		}
// 		currentIndex = insertPosition
// 		nextMarble++
// 		playerIndex = (playerIndex + 1) % players
// 	}

// 	highScore := 0
// 	for _, score := range playerScores {
// 		if score > highScore {
// 			highScore = score
// 		}
// 	}

// 	return highScore
// }

func Part1(players int, worth int) int {
	circle := Circle{}
	circle.Init()

	playerScores := make([]int, players)
	playerIndex := 0

	circle.Insert(0)
	nextMarble := 1
	for nextMarble <= worth {
		if nextMarble%23 != 0 {
			circle.Insert(nextMarble)
		} else {
			playerScores[playerIndex] += nextMarble
			playerScores[playerIndex] += circle.Remove()
		}
		nextMarble++
		playerIndex = (playerIndex + 1) % players
	}

	highestScore := 0
	for _, score := range playerScores {
		if score > highestScore {
			highestScore = score
		}
	}

	return highestScore
}

func Part2(players int, worth int) int {
	worth *= 100
	return Part1(players, worth)
}

func main() {
	players, worth := ReadInput(os.Args[1])
	fmt.Println(Part1(players, worth))
	fmt.Println(Part2(players, worth))

}
