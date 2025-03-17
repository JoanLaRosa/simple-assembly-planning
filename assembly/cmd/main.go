package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/joan/assembly/pkg/assembly"
	pb "github.com/joan/assembly/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	// Read input from stdin or file
	var input []byte
	var err error
	if len(os.Args) > 1 {
		input, err = ioutil.ReadFile(os.Args[1])
	} else {
		input, err = ioutil.ReadAll(os.Stdin)
	}
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	// Parse JSON into protobuf message
	component := &pb.Component{}
	unmarshaler := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	if err := unmarshaler.Unmarshal(input, component); err != nil {
		log.Fatalf("Failed to parse input: %v", err)
	}

	// Compute assembly time
	totalTime := assembly.ComputeAssemblyTime(component)
	fmt.Printf("Minimum Assembly Time: %d minutes\n", totalTime)
}
