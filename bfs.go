package grafos

import (
	"fmt"

	Cola "github.com/untref-ed/colaGO"
)

// BFS realiza un recorrido BFS sobre el grafo g comenzando en el nodo s.
// Procesa todos los vértices con la función Visitar.
// Complejidad en tiempo: O(|V| + |A|)


func BFS(g *Grafo, s string, Visitar func(string)) (error) {
	if !g.ExisteVertice(s) {
		return fmt.Errorf("El nodo %s no existe", s)
	}
	visitados := make(map[string]bool)
	cola := Cola.NuevaCola[string]()
	cola.Encolar(s)
	visitados[s] = true
	for !cola.EstaVacia() {
		v, _ := cola.Desencolar()
		Visitar(v)
		for _, w := range g.Adyacentes(v) {
			if !visitados[w] {
				cola.Encolar(w)
				visitados[w] = true
			}
		}
	}
	return nil
}
