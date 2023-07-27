package grafos

import (
	"fmt"
	"testing"
)

func TestOrdenTopologico(t *testing.T) {
	g := NuevoGrafo(true)
	g.AgregarArista("a", "b", 1)
	g.AgregarArista("a", "c", 1)
	g.AgregarArista("b", "d", 1)
	g.AgregarArista("c", "d", 1)
	g.AgregarArista("d", "e", 1)
	g.AgregarArista("d", "f", 1)
	g.AgregarArista("e", "g", 1)
	g.AgregarArista("f", "g", 1)
	orden, err := OrdenTopologico(g)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(orden)
}
