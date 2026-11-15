// Paquete catalogo: módulo 1, maneja los productos de la tienda.
package catalogo

import (
	"sort"
	"tiendafacil/modelo"
)

// Agregar devuelve una lista NUEVA con el producto agregado.
// La lista original no se modifica.
func Agregar(productos []modelo.Producto, nuevo modelo.Producto) []modelo.Producto {
	resultado := []modelo.Producto{}
	for _, p := range productos {
		resultado = append(resultado, p)
	}
	resultado = append(resultado, nuevo)
	return resultado
}

// Buscar devuelve el producto con ese código y true si lo encontró.
func Buscar(productos []modelo.Producto, codigo string) (modelo.Producto, bool) {
	for _, p := range productos {
		if p.Codigo == codigo {
			return p, true
		}
	}
	return modelo.Producto{}, false
}

// Filtrar recibe la lista y una FUNCIÓN que decide qué productos se quedan.
// Como recibe una función, sirve para cualquier filtro (función de orden superior).
func Filtrar(productos []modelo.Producto, condicion func(modelo.Producto) bool) []modelo.Producto {
	resultado := []modelo.Producto{}
	for _, p := range productos {
		if condicion(p) {
			resultado = append(resultado, p)
		}
	}
	return resultado
}

// HayStock es una condición para Filtrar: true si quedan unidades.
func HayStock(p modelo.Producto) bool {
	return p.Stock > 0
}

// DeCategoria crea y devuelve una condición para una categoría.
func DeCategoria(categoria string) func(modelo.Producto) bool {
	return func(p modelo.Producto) bool {
		return p.Categoria == categoria
	}
}

// OrdenarPorPrecio devuelve una copia ordenada del más barato al más caro.
func OrdenarPorPrecio(productos []modelo.Producto) []modelo.Producto {
	copia := []modelo.Producto{}
	for _, p := range productos {
		copia = append(copia, p)
	}
	// sort.Slice recibe una función anónima que dice cuál va primero
	sort.Slice(copia, func(i, j int) bool {
		return copia[i].Precio < copia[j].Precio
	})
	return copia
}

// AgruparPorCategoria arma un mapa: categoría -> productos de esa categoría.
func AgruparPorCategoria(productos []modelo.Producto) map[string][]modelo.Producto {
	grupos := map[string][]modelo.Producto{}
	for _, p := range productos {
		grupos[p.Categoria] = append(grupos[p.Categoria], p)
	}
	return grupos
}
