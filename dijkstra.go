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
	heap := Heap.NuevoHeapMin(func(i, j string) int {
		return distancia[i] - distancia[j]
	})
	for _, v := range g.vertices {
		distancia[v.nombre] = math.MaxInt32
		padre[v.nombre] = ""
	}
	distancia[origen] = 0
	heap.Insertar(origen)
	// Recorremos el heap
	for !heap.EstaVacio() {
		nombre, _ := heap.Extraer()
		for _, arista := range g.vertices[nombre].aristas {
			if distancia[arista.Destino()] > distancia[nombre]+arista.Peso() {
				distancia[arista.Destino()] = distancia[nombre] + arista.Peso()
				padre[arista.Destino()] = nombre
				heap.Insertar(arista.Destino())
			}
		}
	}
	return distancia, padre, nil
}
