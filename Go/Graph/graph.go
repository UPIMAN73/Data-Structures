package Graph

import (
	"fmt"
	"reflect"
)

// Edges are a set of vertices that are connected
type Edge[T any] struct {
	V1 T // Vector 1
	V2 T // Vector 2
}

// Graph is a datstructure that contains vertices and edges
type Graph[T any] struct {
	Vertices []T       // Vertices are a collection of vertexes or objects that are associated by a identifier (typically a number)
	Edges    []Edge[T] // Edges are the connection between vertices
}

// Search for edge in graph
func GraphEdgeSearch[T any](graph *Graph[T], edge *Edge[T]) *Edge[T] {
	if graph == nil || edge == nil {
		return nil
	} else {
		for i := 0; i < len(graph.Edges); i++ {
			if reflect.DeepEqual(*edge, graph.Edges[i]) {
				return edge
			}
		}
		return nil
	}
}

// Checks to see if an edge exists in the graph
func GraphEdgeExist[T any](graph *Graph[T], edge *Edge[T]) bool {
	if graph == nil || edge == nil {
		return false
	} else {
		return (edge == GraphEdgeSearch[T](graph, edge))
	}
}

// Checks to see if a vertex exists in a graph
func GraphVertexExist[T any](graph *Graph[T], vertex T) bool {
	if graph == nil {
		return false
	} else {
		for i := 0; i < len(graph.Vertices); i++ {
			if reflect.DeepEqual(vertex, graph.Vertices[i]) {
				return true
			}
		}
		return false
	}
}

// Test function to demonstrate the usable components of this datastructure
func TestGraph() {
	// Description to user about the type of datatype being used for testing
	fmt.Println("\nGraph:")

	// Definitions
	var graph Graph[int]
	graph.Vertices = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	graph.Edges = []Edge[int]{{1, 2}, {2, 4}, {3, 4}, {4, 5}, {5, 8}, {5, 7}, {7, 9}, {9, 10}, {3, 6}, {6, 9}}

	// Printout
	fmt.Print("\tVertices:\n\t\t")
	fmt.Println(graph.Vertices)
	fmt.Print("\tEdges:\n\t\t")
	fmt.Println(graph.Edges)

	// Search for edge
	edgeA := graph.Edges[0]
	edgeB := Edge[int]{V1: 1, V2: 5}
	fmt.Println(fmt.Sprintf("\n\nDoes edge %v exist?: %t", edgeA, GraphEdgeExist[int](&graph, &edgeA)))
	fmt.Println(fmt.Sprintf("Does edge %v exist?: %t", edgeB, GraphEdgeExist[int](&graph, &edgeB)))
}
