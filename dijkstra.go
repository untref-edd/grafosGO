package grafos

import (
	"fmt"
	"math"
	Heap "github.com/untref-ed/heapGO"
)

// Algoritmo de Dijkstra
// Camino mínimo en grafos dirigidos con pesos no negativos
func Dijkstra(g *Grafo, origen string) (map[string]int, map[string]string, error) {
	if !g.dirigido {
		return nil, nil, fmt.Errorf("Dijkstra: el grafo debe ser no dirigido")
	}
	// Inicializamos las estructuras auxiliares
	distancia := make(map[string]int)
	padre := make(map[string]string)
	visitado := make(map[string]bool)
	heap := Heap.NuevoHeapMin(func(i, j string) int {
		return distancia[i] - distancia[j]
	})
	for _, v := range g.vertices {
		distancia[v.nombre] = math.MaxInt32
		padre[v.nombre] = ""
		visitado[v.nombre] = false
	}
	distancia[origen] = 0
	heap.Insertar(origen)
	// Recorremos el heap
	for !heap.EstaVacio() {
		v, _ := heap.Extraer()
		visitado[v] = true
		// Recorremos los adyacentes
		for _, arista := range g.vertices[v].aristas {
			if distancia[arista.Destino()] > distancia[v]+arista.Peso() {
				distancia[arista.Destino()] = distancia[v] + arista.Peso()
				padre[arista.Destino()] = v
				heap.Insertar(arista.Destino())
			}	
		}
	}
	return distancia, padre, nil
}
