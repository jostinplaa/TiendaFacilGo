// Package pedidos: módulo 4, creación de pedidos y control de stock.
package pedidos

import (
	"errors"
	"tiendafacil/fn"
	m "tiendafacil/modelo"
	"time"
)

func HayStock(c []m.ItemCarrito) bool {
	return fn.All(c, func(i m.ItemCarrito) bool { return i.Producto.Stock >= i.Cantidad })
}

func Crear(num int, cli m.Cliente, c []m.ItemCarrito, total float64) (m.Pedido, error) {
	if len(c) == 0 || !HayStock(c) {
		return m.Pedido{}, errors.New("carrito vacío o sin stock")
	}
	return m.Pedido{Numero: num, Cliente: cli, Items: c, Total: total, Estado: m.Pendiente, Fecha: time.Now()}, nil
}

func DescontarStock(inv []m.Producto, p m.Pedido) []m.Producto {
	return fn.Map(inv, func(prod m.Producto) m.Producto {
		if it, ok := fn.Find(p.Items, func(i m.ItemCarrito) bool { return i.Producto.Codigo == prod.Codigo }); ok {
			return prod.ConStock(prod.Stock - it.Cantidad)
		}
		return prod
	})
}
