package grafos

import (
	"fmt"
)

// Aristas es una estructura que representa a las aristas de un grafo.
type Arista struct {
	// Vértice al que apunta la arista.
	destino string
	// Peso de la arista.
	peso int
}
func nuevaArista(destino string, peso int) Arista {
	return Arista{
		destino: destino,
		peso: peso,
	}
}

func (a *Arista) Destino() string {
	return a.destino
}

func (a *Arista) Peso() int {
	return a.peso
}

func (a *Arista) String() string {
	return fmt.Sprintf("(%s, %d)", a.destino, a.peso)
}
