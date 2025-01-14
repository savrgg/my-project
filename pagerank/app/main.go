package main

func main() {
	graph := NewGraph()
	graph.AddNode("A", []string{"B", "C"})
	graph.AddNode("B", []string{"C"})
	graph.AddNode("C", []string{"A"})

	dampingFactor := 0.85
	epsilon := 0.0001
	maxIterations := 100

	graph.PageRank(dampingFactor, epsilon, maxIterations)

	// Print the ranks
	graph.PrintRanks()
}