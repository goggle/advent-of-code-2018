package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
)

type WorkerQueue struct {
	numberOfWorkers       int
	additionalTimePerTask int
	currentState          []rune
	remainingTimes        []int
	queue                 []rune
	pop                   []rune
	currentTime           int
}

func (wq *WorkerQueue) Init(numberOfWorkers int, additionalTimePerTask int) {
	wq.numberOfWorkers = numberOfWorkers
	wq.additionalTimePerTask = additionalTimePerTask
	wq.currentState = []rune{}
	wq.remainingTimes = []int{}
	wq.queue = []rune{}
	wq.pop = []rune{}
	wq.currentTime = 0
}

func (wq WorkerQueue) Print() {
	state_string := "Current state:"
	for i := range wq.currentState {
		state_string += fmt.Sprintf(" (%c, %d)", wq.currentState[i], wq.remainingTimes[i])
	}
	queue_string := "Queue:"
	for i := range wq.queue {
		queue_string += fmt.Sprintf(" %c", wq.queue[i])
	}
	time_string := fmt.Sprintf("Current time: %d", wq.currentTime)
	fmt.Println(state_string)
	fmt.Println(queue_string)
	fmt.Println(time_string)
	fmt.Println()
}

func (wq *WorkerQueue) NotEmpty() bool {
	if len(wq.currentState) > 0 || len(wq.queue) > 0 {
		return true
	}
	return false
}

func (wq *WorkerQueue) Add(r rune) {
	if len(wq.currentState) < wq.numberOfWorkers {
		wq.currentState = append(wq.currentState, r)
		wq.remainingTimes = append(wq.remainingTimes, additionalTime(r)+wq.additionalTimePerTask)
	} else {
		wq.queue = append(wq.queue, r)
	}
}

func (wq *WorkerQueue) Pop() []rune {
	// Transfer from queue to currentState
	i := 0
	for len(wq.currentState) < wq.numberOfWorkers && len(wq.queue) > i {
		wq.currentState = append(wq.currentState, wq.queue[i])
		wq.remainingTimes = append(wq.remainingTimes, additionalTime(wq.queue[i])+wq.additionalTimePerTask)
		i++
	}
	wq.queue = wq.queue[i:]

	output := []rune{}
	if len(wq.currentState) == 0 {
		return output
	}

	// Find the shortest remaining time:
	minTime := wq.remainingTimes[0]
	for _, t := range wq.remainingTimes {
		if t < minTime {
			minTime = t
		}
	}

	// Add this time to the elapsed time:
	wq.currentTime += minTime

	// Find the indices to remove from wq.remainingTimes and wq.currentState
	remIndices := []int{}
	for i := range wq.remainingTimes {
		wq.remainingTimes[i] -= minTime
		if wq.remainingTimes[i] == 0 {
			remIndices = append(remIndices, i)
		}
	}

	// Reduce the state and generate the output
	for i := len(remIndices) - 1; i >= 0; i-- {
		index := remIndices[i]
		output = append(output, wq.currentState[index])
		wq.currentState = append(wq.currentState[:index], wq.currentState[index+1:]...)
		wq.remainingTimes = append(wq.remainingTimes[:index], wq.remainingTimes[index+1:]...)

	}
	return output
}

func (wq *WorkerQueue) Time() int {
	return wq.currentTime
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

type DependencyMap map[rune][]rune

func (dm DependencyMap) Print() {
	for k, v := range dm {
		s := fmt.Sprintf("%c: [", k)
		for _, c := range v {
			s += fmt.Sprintf("%c ", c)
		}
		s += "]"
		fmt.Println(s)
	}
	fmt.Println()

}

func (dm DependencyMap) Init(input [][2]rune) {
	for _, inp := range input {
		for i := 0; i < 2; i++ {
			_, ok := dm[inp[i]]
			if !ok {
				dm[inp[i]] = []rune{}
			}
		}
		dm[inp[1]] = append(dm[inp[1]], inp[0])
	}
}

func (dm DependencyMap) Pop() []rune {
	availableStates := sortRunes{}

	for k, v := range dm {
		if len(v) == 0 {
			availableStates = append(availableStates, k)
			delete(dm, k)
		}
	}
	sort.Sort(availableStates)
	return availableStates
}

func (dm DependencyMap) RemoveDependency(r rune) {
	for k, v := range dm {
		for i, dep := range v {
			if dep == r {
				dm[k] = append(v[:i], v[i+1:]...)
			}
		}
	}
}

func ReadInput(path string) [][2]rune {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	scanner := bufio.NewScanner(file)
	regex := regexp.MustCompile(`Step\s+([[:alpha:]]).+step\s+([[:alpha:]])`)

	parsedInput := [][2]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		match := regex.FindStringSubmatch(line)
		if len(match) > 2 {
			r1 := []rune(match[1])[0]
			r2 := []rune(match[2])[0]
			parsedInput = append(parsedInput, [2]rune{r1, r2})
		}
	}
	return parsedInput
}

func createDependencyMap(input [][2]rune) map[rune][]rune {
	dependencyMap := map[rune][]rune{}
	for _, inp := range input {
		for i := 0; i < 2; i++ {
			_, ok := dependencyMap[inp[i]]
			if !ok {
				dependencyMap[inp[i]] = []rune{}
			}
		}
		dependencyMap[inp[1]] = append(dependencyMap[inp[1]], inp[0])
	}

	return dependencyMap
}

func additionalTime(r rune) int {
	return int(r-65) + 1
}

func retrieveAvailableStates(dMap map[rune][]rune) []rune {
	availableStates := sortRunes{}
	for k, v := range dMap {
		if len(v) == 0 {
			availableStates = append(availableStates, k)
			delete(dMap, k)
		}
	}
	sort.Sort(availableStates)
	return availableStates
}

func Part1(input [][2]rune) string {
	dependencyMap := DependencyMap{}
	dependencyMap.Init(input)
	order := []rune{}

	available := dependencyMap.Pop()
	for len(available) > 0 {
		order = append(order, available...)
		for _, k := range available {
			dependencyMap.RemoveDependency(k)
		}
		available = dependencyMap.Pop()
	}

	return string(order)
}

func Part2(input [][2]rune) int {
	dependencyMap := DependencyMap{}
	dependencyMap.Init(input)

	workerQueue := WorkerQueue{}
	workerQueue.Init(5, 60)

	availableStates := dependencyMap.Pop()
	for _, state := range availableStates {
		workerQueue.Add(state)
	}

	for workerQueue.NotEmpty() {
		finished := workerQueue.Pop()
		for _, state := range finished {
			dependencyMap.RemoveDependency(state)
		}
		availableStates = dependencyMap.Pop()
		for _, state := range availableStates {
			workerQueue.Add(state)
		}
	}

	return workerQueue.Time()
}

func main() {
	parsedInput := ReadInput(os.Args[1])
	fmt.Println(Part1(parsedInput))
	fmt.Println(Part2(parsedInput))
}
