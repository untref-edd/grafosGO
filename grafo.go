package grafos

import (
	"fmt"
)


// Grafo es una estructura que representa a un grafo.
type Grafo struct {
	// Lista de vértices del grafo.
	vertices map[string]*Vertice
	// Indica si el grafo es dirigido o no.
	dirigido bool
}

// NuevoGrafo crea un nuevo grafo.
func NuevoGrafo(dirigido bool) *Grafo {
	return &Grafo{
		vertices: make(map[string]*Vertice),
		dirigido: dirigido,
	}
}

// Vertices devuelve la lista de vértices del grafo.
func (g *Grafo) Vertices() []*Vertice {
	vertices := make([]*Vertice, 0)
	for _, v := range g.vertices {
		vertices = append(vertices, v)
	}
	return vertices
}

// AgregarArista agrega una arista al grafo.
func (g *Grafo) AgregarArista(origen, destino string, peso int) error {
	// Si el vértice origen no existe, lo creamos.
	if _, ok := g.vertices[origen]; !ok {
		g.vertices[origen] = nuevoVertice(origen)
	}
	// Si el vértice destino no existe, lo creamos.
	if _, ok := g.vertices[destino]; !ok {
		g.vertices[destino] = nuevoVertice(destino)
	}
	// Agregamos la arista al vértice origen.
	err := g.agregarArista(origen, destino, peso)
	
	// Si el grafo no es dirigido, agregamos la arista inversa
	if err == nil && !g.dirigido {
		err = g.agregarArista(destino, origen, peso)
	}
	return err
}

// agregarArista agrega una arista al grafo.
func (g *Grafo) agregarArista(origen, destino string, peso int) error {
	// chequeamos si la arista ya existe
	for _, arista := range g.vertices[origen].aristas {
		if arista.destino == destino {
			return fmt.Errorf("La arista %s-%s ya existe", origen, destino)
		}
	}
	// agregamos la arista
	a := nuevaArista(destino, peso)
	g.vertices[origen].aristas = append(g.vertices[origen].aristas, a)
	g.vertices[destino].gradoEntrada++
	return nil
}


// Cantidad de Vértices
func (g *Grafo) CantidadVertices() int {
	return len(g.vertices)
}

// Adyacentes de un vértice
func (g *Grafo) Adyacentes(nombre string) []string {
	var adyacentes []string
	for _, arista := range g.vertices[nombre].aristas {
		adyacentes = append(adyacentes, arista.destino)
	}
	return adyacentes
}

// ExisteNodo devuelve true si el nodo existe en el grafo.
func (g *Grafo) ExisteVertice(nombre string) bool {
	_, ok := g.vertices[nombre]
	return ok
}

// String devuelve una representación del grafo en formato string.
func (g *Grafo) String() string {
	s := ""
	s = s + "Cantidad de vértices: " + fmt.Sprint(g.CantidadVertices()) + "\n"
	s = s +"----------------------------------\n"


	for nombre := range g.vertices {
		s += nombre + " -> "
		for _, arista := range g.vertices[nombre].aristas {
			s += fmt.Sprintf("%v, ", arista)
		}
		s += "\n"
	}
	return s
}
