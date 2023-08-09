package grafos

import (
	"fmt"
	"testing"
)

func TestBFS(t *testing.T) {
	g := NuevoGrafo(false)
	g.AgregarArista("A", "B", 1)
	g.AgregarArista("A", "G", 1)
	g.AgregarArista("A", "F", 1)
	g.AgregarArista("B", "C", 1)
	g.AgregarArista("B", "D", 1)
	g.AgregarArista("B", "E", 1)
	g.AgregarArista("D", "H", 1)
	g.AgregarArista("D", "I", 1)
	g.AgregarArista("D", "J", 1)

	err := BFS(g, "A", func(v string) {
		fmt.Println(v)
	})

	if err != nil {
		t.Error(err)
	}
}
