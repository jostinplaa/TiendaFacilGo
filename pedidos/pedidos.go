// Paquete pedidos: módulo 4, convierte el carrito en pedido y controla el stock.
package pedidos

import "tiendafacil/modelo"

// HayStockSuficiente revisa que haya unidades para cada producto del carrito.
func HayStockSuficiente(carrito []modelo.ItemCarrito) bool {
	for _, item := range carrito {
		if item.Producto.Stock < item.Cantidad {
			return false
		}
	}
	return true
}

// Crear arma el pedido solo si el carrito no está vacío y hay stock.
func Crear(numero int, cliente modelo.Cliente, carrito []modelo.ItemCarrito, total float64) (modelo.Pedido, bool) {
	if len(carrito) == 0 || !HayStockSuficiente(carrito) {
		return modelo.Pedido{}, false
	}
	pedido := modelo.Pedido{
		Numero:  numero,
		Cliente: cliente,
		Items:   carrito,
		Total:   total,
		Estado:  modelo.Pendiente,
	}
	return pedido, true
}

// DescontarStock devuelve un inventario NUEVO con las unidades vendidas restadas.
// "p" es una copia de cada producto, por eso el inventario original no cambia.
func DescontarStock(inventario []modelo.Producto, pedido modelo.Pedido) []modelo.Producto {
	nuevo := []modelo.Producto{}
	for _, p := range inventario {
		for _, item := range pedido.Items {
			if item.Producto.Codigo == p.Codigo {
				p.Stock = p.Stock - item.Cantidad
			}
		}
		nuevo = append(nuevo, p)
	}
	return nuevo
}
