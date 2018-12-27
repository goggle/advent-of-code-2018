package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Recipies struct {
	scores          []int
	currentIndexOne int
	currentIndexTwo int
}

func (rec *Recipies) Init(cap int) {
	rec.scores = make([]int, 0, cap)
	rec.scores = append(rec.scores, []int{3, 7}...)
	rec.currentIndexOne = 0
	rec.currentIndexTwo = 1
}

func (rec *Recipies) GenerateNewRecipies() {
	score := rec.scores[rec.currentIndexOne] + rec.scores[rec.currentIndexTwo]
	newScores := make([]int, 0, 2)
	if score >= 10 {
		newScores = append(newScores, 1)
		newScores = append(newScores, score-10)
	} else {
		newScores = append(newScores, score)
	}
	rec.scores = append(rec.scores, newScores...)
	rec.currentIndexOne = (rec.currentIndexOne + 1 + rec.scores[rec.currentIndexOne]) % len(rec.scores)
	rec.currentIndexTwo = (rec.currentIndexTwo + 1 + rec.scores[rec.currentIndexTwo]) % len(rec.scores)
}

func (rec *Recipies) Count() int {
	return len(rec.scores)
}

func (rec *Recipies) Match(sequence []int, index int) bool {
	for i, v := range sequence {
		if v != rec.scores[index+i] {
			return false
		}
	}
	return true
}

func (rec *Recipies) Print() {
	s := ""
	for i, v := range rec.scores {
		if i != 0 {
			s += " "
		}
		if i == rec.currentIndexOne {
			s += "("
		} else if i == rec.currentIndexTwo {
			s += "["
		}
		s += fmt.Sprintf("%d", v)
		if i == rec.currentIndexOne {
			s += ")"
		} else if i == rec.currentIndexTwo {
			s += "]"
		}
	}
	fmt.Println(s)
}

func Part1(serial int) string {
	recipies := Recipies{}
	recipies.Init(serial + 20)
	for recipies.Count() < serial+10 {
		recipies.GenerateNewRecipies()
	}
	output := ""
	for i := 0; i < 10; i++ {
		output += fmt.Sprintf("%d", recipies.scores[serial+i])
	}
	return output
}

func Part2(serial []int) int {
	recipies := Recipies{}
	recipies.Init(1000)
	index := 0
	for {
		for recipies.Count() < index+len(serial) {
			recipies.GenerateNewRecipies()
		}
		if recipies.Match(serial, index) {
			break
		}
		index++
	}
	return index
}

func main() {
	serial, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(Part1(serial))

	serialSlice := []int{}
	serialString := strconv.Itoa(serial)
	for _, v := range serialString {
		serialSlice = append(serialSlice, int(v-'0'))
	}

	fmt.Println(serialSlice)

	fmt.Println(Part2(serialSlice))
}
