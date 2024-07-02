package main

import "fmt"

func main() {
	graph := make(map[string][]string)

	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}
}

func bfs(graph map[string][]string, root string, target string) bool {
	visited := make(map[string]struct{})

	queue := []string{root}

	for len(queue) > 0 {
		node := queue[0]
		if node == target {
			return true
		}
		// dequeue
		queue = queue[1:]

		if _, ok := visited[node]; !ok {
			visited[node] = struct{}{}
			fmt.Printf("%s node has been visited", node)

			for _, v := range graph[node] {
				if _, ok := visited[v]; !ok {
					queue = append(queue, v)
				}
			}
		}
	}

	return false
}
