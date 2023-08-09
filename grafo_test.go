package grafos

import (
	"fmt"
	"testing"
)

func TestNuevoGrafo(t *testing.T) {
	g := NuevoGrafo(false)
	if g == nil {
		t.Error("Se esperaba un grafo")
	}
}

func TestAgregarAristaNoDirigido(t *testing.T) {
	// Grafo No Dirigido
	g := NuevoGrafo(false)
	err := g.AgregarArista("a", "b", 0)
	if err != nil {
		t.Error(err)
	}
	err = g.AgregarArista("b", "c", 0)
	if err != nil {
		t.Error(err)
	}
	err = g.AgregarArista("c", "d", 0)
	if err != nil {
		t.Error(err)
	}
	err = g.AgregarArista("d", "e", 0)
	if err != nil {
		t.Error(err)
	}
	err = g.AgregarArista("e", "f", 0)
	if err != nil {
		t.Error(err)
	}
	err = g.AgregarArista("f", "g", 0)
	if err != nil {
		t.Error(err)
	}

	if g.CantidadVertices() != 7 {
		t.Errorf("Se esperaban 7 vértices y se obtuvieron %d", g.CantidadVertices())
	}
	a := g.Adyacentes("a")
	if len(a) != 1 && a[0] != "b" {
		t.Errorf("Se esperaba que b sea adyacente a a")
	}
	b := g.Adyacentes("b")
	if len(b) != 2 && (b[0] != "c" || b[1] != "a") {
		t.Errorf("Se esperaba que a y c sean adyacente a b")
	}
	c := g.Adyacentes("c")
	if len(c) != 2 && (c[0] != "d" || c[1] != "b") {
		t.Errorf("Se esperaba que b y d sean adyacentes a c")
	}
	d := g.Adyacentes("d")
	if len(d) != 2 && (d[0] != "e" || d[1] != "c") {
		t.Errorf("Se esperaba que c y e sean adyacentes a d")
	}
	e := g.Adyacentes("e")
	if len(e) != 2 && (e[0] != "f" || e[1] != "d") {
		t.Errorf("Se esperaba que d y f sean adyacentes a e")
	}
	f := g.Adyacentes("f")
	if len(f) != 2 && (f[0] != "g" || f[1] != "e") {
		t.Errorf("Se esperaba que e y g sean adyacentes a f")
	}

	fmt.Println(g)
}

func TestAgregarAristaDirigido(t *testing.T) {
	// Grafo Dirigido
	g := NuevoGrafo(true)
	err := g.AgregarArista("a", "b", 0)
	if err != nil {
		t.Error(err)
	}
	err = g.AgregarArista("b", "c", 0)
	if err != nil {
		t.Error(err)
	}
	err = g.AgregarArista("c", "d", 0)
	if err != nil {
		t.Error(err)
	}
	err = g.AgregarArista("d", "e", 0)
	if err != nil {
		t.Error(err)
	}
	err = g.AgregarArista("e", "f", 0)
	if err != nil {
		t.Error(err)
	}
	err = g.AgregarArista("f", "g", 0)
	if err != nil {
		t.Error(err)
	}

	if g.CantidadVertices() != 7 {
		t.Errorf("Se esperaban 7 vértices y se obtuvieron %d", g.CantidadVertices())
	}
	a := g.Adyacentes("a")
	if len(a) != 1 && a[0] != "b" {
		t.Errorf("Se esperaba que b sea adyacente a a")
	}
	b := g.Adyacentes("b")
	if len(b) != 1 && b[0] != "a" {
		t.Errorf("Se esperaba que a sea adyacente a b")
	}
	c := g.Adyacentes("c")
	if len(c) != 1 && c[0] != "b" {
		t.Errorf("Se esperaba que b sea adyacente a c")
	}
	d := g.Adyacentes("d")
	if len(d) != 1 && d[0] != "c" {
		t.Errorf("Se esperaba que c sea adyacente a d")
	}
	e := g.Adyacentes("e")
	if len(e) != 1 && e[0] != "d" {
		t.Errorf("Se esperaba que d sea adyacente a e")
	}
	f := g.Adyacentes("f")
	if len(f) != 1 && f[0] != "e" {
		t.Errorf("Se esperaba que e sea adyacente a f")
	}
	fmt.Println(g)
}
