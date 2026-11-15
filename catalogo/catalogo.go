// Package catalogo: módulo 1, gestión de productos con funciones puras.
package catalogo

import (
	"sort"
	"strings"
	"tiendafacil/fn"
	m "tiendafacil/modelo"
)

type Criterio func(m.Producto) bool

func Agregar(ps []m.Producto, nuevo m.Producto) []m.Producto {
	return append(append([]m.Producto{}, ps...), nuevo)
}

func Filtrar(ps []m.Producto, c Criterio) []m.Producto { return fn.Filter(ps, c) }

func PorCategoria(cat string) Criterio {
	return func(p m.Producto) bool { return strings.EqualFold(p.Categoria, cat) }
}
func PrecioMaximo(max float64) Criterio { return func(p m.Producto) bool { return p.Precio <= max } }
func Disponible() Criterio              { return func(p m.Producto) bool { return p.Stock > 0 } }

func BuscarPorCodigo(ps []m.Producto, cod string) (m.Producto, bool) {
	return fn.Find(ps, func(p m.Producto) bool { return p.Codigo == cod })
}

func OrdenarPorPrecio(ps []m.Producto) []m.Producto {
	c := append([]m.Producto{}, ps...)
	sort.SliceStable(c, func(i, j int) bool { return c[i].Precio < c[j].Precio })
	return c
}

func AgruparPorCategoria(ps []m.Producto) map[string][]m.Producto {
	return fn.Reduce(ps, map[string][]m.Producto{}, func(acc map[string][]m.Producto, p m.Producto) map[string][]m.Producto {
		acc[p.Categoria] = append(acc[p.Categoria], p)
		return acc
	})
}
