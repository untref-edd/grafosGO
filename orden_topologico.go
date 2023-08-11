package grafos

import (
	"fmt"
	Cola "github.com/untref-ed/colaGO"
)

// Calcula el orden topológico de un grafo dirigido.
func OrdenTopologico(g *Grafo) ([]string, error){
	orden := make([]string, 0)
	if !g.dirigido {
		return nil, fmt.Errorf("OrdenTopologico: el grafo debe ser dirigido")
	}
	cola := Cola.NuevaCola[string]()
	// Usamos una copia de los grados de entrada para no modificar el grafo.
	// Si el grado de entrada es cero lo agregamos a la cola.
	gradoEntrada := make(map[string]int)
	for _, v := range g.vertices {
		gradoEntrada[v.nombre] = v.gradoEntrada
		fmt.Println(v.nombre, v.gradoEntrada)
		if gradoEntrada[v.nombre] == 0 {
			cola.Encolar(v.nombre)
		}
	}
	
	// Recorremos la cola.
	for !cola.EstaVacia() {
		nombre,_ := cola.Desencolar()
		orden = append(orden, nombre)
		for _, adyacente := range g.Adyacentes(nombre) {
			gradoEntrada[adyacente]--
			if gradoEntrada[adyacente] == 0 {
				cola.Encolar(adyacente)
				fmt.Println(adyacente)
			}
		}
	}
	return orden, nil
}