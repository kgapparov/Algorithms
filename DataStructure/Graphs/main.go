package main

import "fmt"

type Graph struct {
	adjacencyList map[interface{}][]interface{}
}

func (g *Graph) new() {
	g.adjacencyList = make(map[interface{}][]interface{})
}

func (g *Graph) addVertex(v interface{}) {
	if _, ok := g.adjacencyList[v]; !ok {
		g.adjacencyList[v] = []interface{}{}
	}
}
func (g *Graph) addEdge(v1, v2 interface{}) {
	if _, ok := g.adjacencyList[v1]; !ok {
		g.addVertex(v1)
		g.adjacencyList[v1] = append(g.adjacencyList[v1], v2)
	} else {
		g.adjacencyList[v1] = append(g.adjacencyList[v1], v2)
	}
	if _, ok := g.adjacencyList[v2]; !ok {
		g.addVertex(v2)
		g.adjacencyList[v2] = append(g.adjacencyList[v2], v1)
	} else {
		g.adjacencyList[v2] = append(g.adjacencyList[v2], v1)
	}
}

func (g *Graph) removeEdge(v1, v2 interface{}) {
	if _, ok := g.adjacencyList[v1]; ok {
		g.adjacencyList[v1] = filter(g.adjacencyList, v1, v2)
	}
	if _, ok := g.adjacencyList[v2]; ok {
		g.adjacencyList[v2] = filter(g.adjacencyList, v2, v1)
	}
}

func (g *Graph) removeVertex(v interface{}) {
	if _, ok := g.adjacencyList[v]; ok {
		for len(g.adjacencyList[v]) > 0 {
			v2 := g.adjacencyList[v][len(g.adjacencyList[v])-1]
			g.adjacencyList[v] = g.adjacencyList[v][:len(g.adjacencyList[v])-1]
			g.removeEdge(v, v2)
		}
	}
	delete(g.adjacencyList, v)
}
func filter(m map[interface{}][]interface{}, k1, k2 interface{}) []interface{} {
	for i := 0; i < len(m[k1]); i++ {
		if m[k1][i] == k2 {
			var tmp []interface{}
			tmp = append(tmp, m[k1][:i]...)
			tmp = append(tmp, m[k1][i+1:]...)
			m[k1] = tmp
		}
	}
	return m[k1]
}

func main() {
	var graph Graph
	graph.new()
	graph.addVertex("Almaty")
	graph.addVertex("Astana")
	graph.addVertex("Taraz")
	graph.addVertex("Kokshetau")
	graph.addEdge("Almaty", "Kokshetau")
	graph.addEdge("Taraz", "New York")
	graph.addEdge("Taraz", "Kostanai")
	fmt.Println(graph)
	graph.removeVertex("Taraz")
	fmt.Println(graph)
}
