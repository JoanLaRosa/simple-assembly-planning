package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestComputeAssemblyTime(t *testing.T) {
	tests := []struct {
		name     string
		jsonFile string
		want     int
	}{
		{
			name:     "Example Case",
			jsonFile: "./assembly/testdata/example.json",
			want:     11,
		},
		{
			name:     "Gadget X",
			jsonFile: "./assembly/testdata/gadgetX.json",
			want:     10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := os.ReadFile(tt.jsonFile)
			if err != nil {
				t.Fatalf("Failed to read test file: %v", err)
			}

			var root Component
			if err := json.Unmarshal(data, &root); err != nil {
				t.Fatalf("Failed to unmarshal JSON: %v", err)
			}

			got := computeAssemblyTime(root, true)
			if got != tt.want {
				t.Errorf("computeAssemblyTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchedule(t *testing.T) {
	tests := []struct {
		name   string
		times  []int
		maxPar int
		want   int
	}{
		{
			name:   "Single Task",
			times:  []int{5},
			maxPar: 1,
			want:   5,
		},
		{
			name:   "Two Tasks Parallel",
			times:  []int{3, 4},
			maxPar: 2,
			want:   4,
		},
		{
			name:   "Two Tasks Serial",
			times:  []int{3, 4},
			maxPar: 1,
			want:   7,
		},
		{
			name:   "Multiple Tasks Limited Parallelism",
			times:  []int{2, 4, 1, 3},
			maxPar: 2,
			want:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := schedule(tt.times, tt.maxPar)
			if got != tt.want {
				t.Errorf("schedule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "Single Element",
			arr:  []int{5},
			want: 5,
		},
		{
			name: "Multiple Elements",
			arr:  []int{3, 7, 2, 5},
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := max(tt.arr)
			if got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgMin(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "Single Element",
			arr:  []int{5},
			want: 0,
		},
		{
			name: "Multiple Elements",
			arr:  []int{3, 7, 2, 5},
			want: 2,
		},
		{
			name: "First Element Minimum",
			arr:  []int{1, 7, 2, 5},
			want: 0,
		},
		{
			name: "Last Element Minimum",
			arr:  []int{3, 7, 2, 1},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := argMin(tt.arr)
			if got != tt.want {
				t.Errorf("argMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
