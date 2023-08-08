package grafos

// Vértices es una estructura que representa a los vértices de un grafo.
type Vertice struct {
	// Nombre del vértice.
	nombre string
	// Lista de aristas que salen del vértice.
	aristas []Arista
	// Grado de entrada del vértice. Para Orden Topológico.
	gradoEntrada int
}

func nuevoVertice(nombre string) *Vertice {
	return &Vertice{
		nombre:       nombre,
		aristas:      make([]Arista, 0),
		gradoEntrada: 0,
	}
}

func (v *Vertice) Nombre() string {
	return v.nombre
}

func (v *Vertice) GradoDeEntrada() int {
	return v.gradoEntrada
}

func (v *Vertice) String() string {
	return v.nombre
}
