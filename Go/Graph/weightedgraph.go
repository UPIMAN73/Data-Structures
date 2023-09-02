package Graph

import "fmt"

// Weighted edges have a specific use case in certain problems
type WeightedEdge[T any] struct {
	V1     T // Vector 1
	V2     T // Vector 2
	Weight T // Weight of the edge based on set parameter
}

// Weighted graph uses the weighted edge to describe the environment
type WeightedGraph[T any] struct {
	Vertices []T               // Vertices are a collection of vertexes or objects that are associated by a identifier (typically a number)
	Edges    []WeightedEdge[T] // Edges are the connection between vertices (with weights)
}

// Test function to demonstrate the usable components of this datastructure
func TestWeightedGraph() {
	// Description to user about the type of datatype being used for testing
	fmt.Println("\nWeighted Graph:")

	// Definitions
	var graph WeightedGraph[int]
	graph.Vertices = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	graph.Edges = []WeightedEdge[int]{{1, 3, 1}}

	// Printout
	fmt.Print("\tVertices:\n\t\t")
	fmt.Println(graph.Vertices)
	fmt.Print("\tEdges:\n\t\t")
	fmt.Println(graph.Edges)
}
