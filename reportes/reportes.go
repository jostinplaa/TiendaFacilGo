// Paquete reportes: módulo 6, información para que el dueño tome decisiones.
package reportes

import (
	"tiendafacil/catalogo"
	"tiendafacil/modelo"
)

// TotalVendido suma el total de los pedidos que ya están pagados.
func TotalVendido(pedidos []modelo.Pedido) float64 {
	suma := 0.0
	for _, p := range pedidos {
		if p.Estado == modelo.Pagado {
			suma = suma + p.Total
		}
	}
	return suma
}

// UnidadesPorProducto cuenta cuántas unidades se vendieron de cada producto.
func UnidadesPorProducto(pedidos []modelo.Pedido) map[string]int {
	unidades := map[string]int{}
	for _, p := range pedidos {
		for _, item := range p.Items {
			unidades[item.Producto.Nombre] = unidades[item.Producto.Nombre] + item.Cantidad
		}
	}
	return unidades
}

// StockBajo devuelve los productos con menos unidades que el mínimo.
// Reutiliza catalogo.Filtrar pasándole una función anónima.
func StockBajo(inventario []modelo.Producto, minimo int) []modelo.Producto {
	return catalogo.Filtrar(inventario, func(p modelo.Producto) bool {
		return p.Stock < minimo
	})
}
