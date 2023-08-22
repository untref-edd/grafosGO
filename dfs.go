package grafos


// Recorrido en profundidad (DFS)
// Complejidad en tiempo: O(|V| + |A|)
func DFS(g *Grafo, Visitar func(string)){
	visitados := make(map[string]bool)
	for _, v := range g.Vertices() {
		if !visitados[v] {
			dfs(g, v, visitados, Visitar)
		}
	}
}

func dfs(g *Grafo, v string, visitados map[string]bool, Visitar func(string)) {
	visitados[v] = true
	Visitar(v)
	for _, w := range g.Adyacentes(v) {
		if !visitados[w] {
			dfs(g, w, visitados, Visitar)
		}
	}
}
