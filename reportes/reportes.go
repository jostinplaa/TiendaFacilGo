// Package reportes: módulo 6, reportes con map, filter y reduce.
package reportes

import (
	"tiendafacil/fn"
	m "tiendafacil/modelo"
)

func TotalVendido(ps []m.Pedido) float64 {
	pagados := fn.Filter(ps, func(p m.Pedido) bool { return p.Estado == m.Pagado })
	return fn.Reduce(pagados, 0.0, func(a float64, p m.Pedido) float64 { return a + p.Total })
}

func UnidadesPorProducto(ps []m.Pedido) map[string]int {
	items := fn.Reduce(ps, []m.ItemCarrito{}, func(a []m.ItemCarrito, p m.Pedido) []m.ItemCarrito { return append(a, p.Items...) })
	return fn.Reduce(items, map[string]int{}, func(acc map[string]int, i m.ItemCarrito) map[string]int {
		acc[i.Producto.Nombre] += i.Cantidad
		return acc
	})
}

func StockBajo(inv []m.Producto, min int) []m.Producto {
	return fn.Filter(inv, func(p m.Producto) bool { return p.Stock < min })
}
