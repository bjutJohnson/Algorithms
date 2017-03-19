package graph

import (
	"fmt"
)

func Test_BFS() {
	pGm := NewGraphManager(false)
	for i := 0; i < 8; i++ {
		pGm.AddNode()
	}

	pGm.AddEdge(0, 1)
	pGm.AddEdge(1, 2)
	pGm.AddEdge(2, 3)
	pGm.AddEdge(3, 4)
	pGm.AddEdge(3, 5)
	pGm.AddEdge(4, 5)
	pGm.AddEdge(4, 6)
	pGm.AddEdge(4, 7)
	pGm.AddEdge(5, 6)
	pGm.AddEdge(6, 7)

	//pGm.PrintEdges()

	result, err := pGm.BFS(2)
	fmt.Println("<==========>")
	if err == nil {
		for _, v := range result {
			fmt.Println(v.GetId())
		}
	}
	fmt.Println("<===========>")

	fmt.Println("stop")
}

func Test_DFS() {
	pGm := NewGraphManager(false)
	for i := 0; i < 8; i++ {
		pGm.AddNode()
	}

	pGm.AddEdge(0, 1)
	pGm.AddEdge(1, 2)
	pGm.AddEdge(2, 3)
	pGm.AddEdge(3, 4)
	pGm.AddEdge(3, 5)
	pGm.AddEdge(4, 5)
	pGm.AddEdge(4, 6)
	pGm.AddEdge(4, 7)
	pGm.AddEdge(5, 6)
	pGm.AddEdge(6, 7)

	pGm.DFS()
	pGm.PrintNode()
}
