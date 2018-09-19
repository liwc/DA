package main

import (
	"DA/7_graph/graph"
	"DA/7_graph/usage"
)

func main() {
	graph.AdjMatrixTest()
	graph.AdjListTest()
	graph.FsTest()
	usage.TopoTest()
	usage.DijkstraTest()
	usage.FloydTest()
}