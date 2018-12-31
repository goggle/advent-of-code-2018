package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type State struct {
	state []bool
	zero  int
}

func (s *State) Init(state []bool, zero int) {
	s.state = []bool{}
	for _, v := range state {
		s.state = append(s.state, v)
	}
	s.zero = zero
}

func (s *State) GetSlices() []string {
	slices := []string{}
	for i := -4; i < len(s.state)+4; i++ {
		slice := []rune{}
		for j := 0; j < 5; j++ {
			if i+j < 0 {
				slice = append(slice, '.')
			} else if i+j >= len(s.state) {
				slice = append(slice, '.')
			} else {
				r := '.'
				if s.state[i+j] {
					r = '#'
				}
				slice = append(slice, r)
			}
		}
		slices = append(slices, string(slice))
	}
	return slices
}

func (s *State) Trim() {
	nLeftEmpty := 0
	for i := 0; i < len(s.state); i++ {
		if !s.state[i] {
			nLeftEmpty++
		} else {
			break
		}
	}
	s.state = s.state[nLeftEmpty:]
	s.zero -= nLeftEmpty

	nRightEmpty := 0
	for i := len(s.state) - 1; i >= 0; i-- {
		if !s.state[i] {
			nRightEmpty++
		} else {
			break
		}
	}
	s.state = s.state[:len(s.state)-nRightEmpty]
}

func (s *State) Transform(previous *State, transformationMap map[string]rune) {
	s.state = []bool{}
	s.zero = previous.zero

	slices := previous.GetSlices()
	for _, slice := range slices {
		r := transformationMap[slice]
		put := false
		if r == '#' {
			put = true
		}
		s.state = append(s.state, put)
	}
	s.zero += 2
}

func (s *State) GetCorrectedPlantIndices() []int {
	indices := []int{}
	for i, v := range s.state {
		if v {
			indices = append(indices, i-s.zero)
		}
	}
	return indices
}

func (s State) String() string {
	out := ""
	for _, v := range s.state {
		if v {
			out += "#"
		} else {
			out += "."
		}
	}
	return out
}

func ReadInput(path string) ([]bool, []string, []bool) {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	scanner := bufio.NewScanner(file)
	regex1 := regexp.MustCompile(`initial\s+state:\s*([#\.]+)`)
	regex2 := regexp.MustCompile(`([#\.]+)\s*=>\s*([#\.])`)

	initialState := []bool{}
	patterns := []string{}
	results := []bool{}
	for scanner.Scan() {
		line := scanner.Text()
		match1 := regex1.FindStringSubmatch(line)
		match2 := regex2.FindStringSubmatch(line)
		if len(match1) > 1 {
			for _, char := range match1[1] {
				state := false
				if char == '#' {
					state = true
				}
				initialState = append(initialState, state)
			}
		} else if len(match2) > 2 {
			pattern := ""
			for _, char := range match2[1] {
				pattern += string(char)
			}
			patterns = append(patterns, pattern)
			for _, char := range match2[2] {
				state := false
				if char == '#' {
					state = true
				}
				results = append(results, state)
			}
		}
	}
	return initialState, patterns, results
}

func Part1(initialState []bool, patternMap map[string]rune) int {
	state := State{}
	state.Init(initialState, 0)

	newState := State{}
	for i := 0; i < 20; i++ {
		newState.Transform(&state, patternMap)
		newState.Trim()

		state = newState
	}

	indices := newState.GetCorrectedPlantIndices()
	sum := 0
	for _, ind := range indices {
		sum += ind
	}

	return sum
}

func Part2(initialState []bool, patternMap map[string]rune) int {
	state := State{}
	state.Init(initialState, 0)

	newState := State{}
	sumDiff := 0
	count := 0
	for i := 0; i < 50000000000; i++ {
		newState.Transform(&state, patternMap)
		newState.Trim()

		if state.String() == newState.String() {
			sumOld := 0
			sumNew := 0
			indicesOld := state.GetCorrectedPlantIndices()
			for _, ind := range indicesOld {
				sumOld += ind
			}
			indicesnew := newState.GetCorrectedPlantIndices()
			for _, ind := range indicesnew {
				sumNew += ind
			}
			sumDiff = sumNew - sumOld
			count = i + 1
			break
		}

		state = newState
	}
	indices := newState.GetCorrectedPlantIndices()
	sum := 0
	for _, ind := range indices {
		sum += ind
	}
	sum += (50000000000 - count) * sumDiff
	return sum
}

func main() {
	initialState, patterns, results := ReadInput(os.Args[1])

	patternMap := map[string]rune{}
	for i, pattern := range patterns {
		result := '.'
		if results[i] {
			result = '#'
		}
		patternMap[pattern] = result
	}

	pMap := map[uint8]bool{}
	for i, pattern := range patterns {
		var code uint8
		var pot uint8 = 1
		for _, char := range pattern {
			if char == '#' {
				code += pot
			}
			pot *= 2
		}
		pMap[code] = results[i]
	}

	fmt.Println(Part1(initialState, patternMap))
	fmt.Println(Part2(initialState, patternMap))

}
