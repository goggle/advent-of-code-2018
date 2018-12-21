package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	children []*Node
	metadata []int
}

func ReadInput(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	scanner := bufio.NewScanner(file)

	parsedInput := []int{}
	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), " ")
		for _, entry := range splitted {
			x, err := strconv.Atoi(entry)
			if err != nil {
				log.Fatal(err)
			}
			parsedInput = append(parsedInput, x)
		}
	}
	return parsedInput
}

func BuildTree(input []int, index int) (*Node, int) {
	root := Node{}

	nChildren := input[index]
	root.children = make([]*Node, nChildren)
	index++

	nMetadata := input[index]
	root.metadata = make([]int, nMetadata)
	index++

	for i := 0; i < nChildren; i++ {
		root.children[i], index = BuildTree(input, index)
	}

	for i := 0; i < nMetadata; i++ {
		root.metadata[i] = input[index+i]
	}
	index += nMetadata

	return &root, index
}

func SumMetadata(tree *Node) int {
	sum := 0
	for _, meta := range tree.metadata {
		sum += meta
	}
	for _, child := range tree.children {
		sum += SumMetadata(child)
	}
	return sum
}

func SumMetadata2(tree *Node) int {
	sum := 0
	if len(tree.children) == 0 {
		for _, meta := range tree.metadata {
			sum += meta
		}
		return sum
	}
	nodesIndices := tree.metadata
	for _, ind := range nodesIndices {
		if ind >= 1 && ind <= len(tree.children) {
			sum += SumMetadata2(tree.children[ind-1])
		}
	}
	return sum
}

func Part1(input []int) int {
	rootPtr, _ := BuildTree(input, 0)
	return SumMetadata(rootPtr)
}

func Part2(input []int) int {
	rootPtr, _ := BuildTree(input, 0)
	return SumMetadata2(rootPtr)
}

func main() {
	input := ReadInput(os.Args[1])
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))

}
