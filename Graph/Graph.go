package main

import "fmt"

type Graph map[string]map[string]int

func main() {
	graph := Graph{
		"a": {"b": 1, "c": 1},
		"b": {"d": 1},
		"c": {"e": 1},
		"d": {"f": 1},
		"e": {},
		"f": {},
	}

	// graph["a"] = make(map[string]int)
	// graph["b"] = make(map[string]int)
	// graph["c"] = make(map[string]int)
	// graph["d"] = make(map[string]int)
	// graph["e"] = make(map[string]int)
	// graph["f"] = make(map[string]int)

	// graph["a"]["c"] = 1
	// graph["a"]["b"] = 1
	// graph["b"]["d"] = 1
	// graph["c"]["e"] = 1
	// graph["d"]["f"] = 1
	// graph["f"][""] = 0
	// fmt.Printf("%v", graph)

	Dfs(graph, "a")
	DfsWithStack(graph, "a")
	Bfs(graph, "a")
}

func Bfs(graph Graph, node string) {
	queue := SimpleQueue{}
	queue.enqueue(node)

	for !queue.isEmpty() {
		newNode, _ := queue.dequeue()
		fmt.Printf(" %s -->", newNode)
		neighbors := graph[newNode]
		for neighbor := range neighbors {
			queue.enqueue(neighbor)
		}
	}
}

func Dfs(graph Graph, node string) {
	fmt.Printf(" %s -->", node)

	neighbors := graph[node]
	for neighbor := range neighbors {
		Dfs(graph, neighbor)
	}
}

func DfsWithStack(graph Graph, node string) {
	stack := SimpleStack{}
	stack.push(node)

	for !stack.isEmpty() {
		newNode, _ := stack.pop()
		fmt.Printf(" %s -->", newNode)
		neighbors := graph[newNode]
		for neighbor := range neighbors {
			stack.push(neighbor)
		}
	}
}
