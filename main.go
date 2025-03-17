package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
)

// Component mirrors the protobuf message.
type Component struct {
	Name         string      `json:"name"`
	AssemblyTime int         `json:"assembly_time"`
	MaxParallel  int         `json:"max_parallel"`
	Children     []Component `json:"children"`
}

// computeAssemblyTime recursively computes the time required to assemble a component.
// For leaf nodes, it returns the component's AssemblyTime.
// For non-root (intermediate) nodes, we assume that the time is the total time to assemble its children.
// At the root, we add the component's own AssemblyTime.
func computeAssemblyTime(c Component, isRoot bool) int {
	// Base case: leaf node.
	if len(c.Children) == 0 {
		return c.AssemblyTime
	}

	// Recursively compute times for all children.
	childTimes := make([]int, 0, len(c.Children))
	for _, child := range c.Children {
		childTimes = append(childTimes, computeAssemblyTime(child, false))
	}

	// Schedule the children tasks using the parent's max_parallel constraint.
	scheduled := schedule(childTimes, c.MaxParallel)

	if isRoot {
		return scheduled + c.AssemblyTime
	}
	return scheduled
}

// schedule computes the total time required to finish tasks with given durations
// subject to a maximum parallelism of maxPar. Here we use a simple greedy algorithm:
// if the number of tasks is less than or equal to maxPar, assume they run concurrently (time = max),
// otherwise, assign tasks to "machines" to simulate a makespan schedule.
func schedule(times []int, maxPar int) int {
	if len(times) == 0 {
		return 0
	}

	// If we have fewer tasks than parallel capacity, run them concurrently.
	if len(times) <= maxPar {
		return max(times)
	}

	// Otherwise, use a greedy approach (Longest Processing Time first)
	// to schedule the tasks on maxPar "machines".
	machines := make([]int, maxPar)
	// Sort durations in descending order.
	sort.Sort(sort.Reverse(sort.IntSlice(times)))
	for _, t := range times {
		// assign t to the machine with the smallest load.
		i := argMin(machines)
		machines[i] += t
	}
	return max(machines)
}

// argMin returns the index of the minimum element in slice.
func argMin(arr []int) int {
	minIdx := 0
	for i, v := range arr {
		if v < arr[minIdx] {
			minIdx = i
		}
	}
	return minIdx
}

// max returns the maximum value in a slice.
func max(arr []int) int {
	m := arr[0]
	for _, v := range arr {
		if v > m {
			m = v
		}
	}
	return m
}

func main() {
	// Parse the input filename from command-line.
	filename := flag.String("input", "", "Path to the input JSON file")
	flag.Parse()

	if *filename == "" {
		fmt.Println("Please provide an input file using -input")
		os.Exit(1)
	}

	// Read JSON file.
	data, err := os.ReadFile(*filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Unmarshal into our Component struct.
	var root Component
	if err := json.Unmarshal(data, &root); err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		os.Exit(1)
	}

	// Compute the total assembly time.
	totalTime := computeAssemblyTime(root, true)
	fmt.Printf("Minimum Assembly Time: %d minutes\n", totalTime)
}
