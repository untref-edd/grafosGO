package grafos

import (
	"fmt"
	"testing"
)

func TestDijkstra(t *testing.T){
	g := NuevoGrafo(true)
	g.AgregarArista("a", "b", 4)
	g.AgregarArista("a", "c", 1)
	g.AgregarArista("c", "b", 2)
	g.AgregarArista("c", "d", 2)
	g.AgregarArista("b", "e", 3)
	g.AgregarArista("d", "e", 3)
	distancia, padre, err := Dijkstra(g, "a")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Distancia")
	for k, v := range distancia {
		fmt.Printf("dist(%s) = %d\n",k, v)
	}
	fmt.Println("Padre")
	for k, v := range padre {
		fmt.Printf("padre(%s) = %s\n",k, v)
	}
}