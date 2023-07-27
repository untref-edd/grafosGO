package grafos

import (
	"fmt"
	"math"
)

// Algoritmo de Bellman-Ford
// Camino mínimo en grafos dirigidos con pesos negativos, pero sin ciclos negativos
func BellmanFord(g *Grafo, origen string) (map[string]int, map[string]string, error) {
	if !g.dirigido {
		return nil, nil, fmt.Errorf("BellmanFord: el grafo debe ser dirigido")
	}
	// Inicializamos las estructuras auxiliares
	distancia := make(map[string]int)
	padre := make(map[string]string)
	for _, v := range g.vertices {
		distancia[v.nombre] = math.MaxInt32
		padre[v.nombre] = ""
	}
	distancia[origen] = 0
	// Recorremos todas las aristas del grafo
	for i := 0; i < len(g.vertices)-1; i++ {
		for _, v := range g.vertices {
			for _, arista := range v.aristas {
				if distancia[arista.Destino()] > distancia[v.nombre]+arista.Peso() {
					distancia[arista.Destino()] = distancia[v.nombre] + arista.Peso()
					padre[arista.Destino()] = v.nombre
				}
			}
		}
	}
	// Chequeamos si hay ciclos negativos
	for _, v := range g.vertices {
		for _, arista := range v.aristas {
			if distancia[arista.Destino()] > distancia[v.nombre]+arista.Peso() {
				return nil, nil, fmt.Errorf("BellmanFord: el grafo tiene ciclos negativos")
			}
		}
	}
	return distancia, padre, nil
}
