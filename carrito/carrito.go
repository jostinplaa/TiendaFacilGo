// Package carrito: módulo 3. Cada operación devuelve un carrito nuevo.
package carrito

import (
	"tiendafacil/fn"
	m "tiendafacil/modelo"
)

func Agregar(c []m.ItemCarrito, p m.Producto, cant int) []m.ItemCarrito {
	mismo := func(i m.ItemCarrito) bool { return i.Producto.Codigo == p.Codigo }
	if fn.Any(c, mismo) {
		return fn.Map(c, func(i m.ItemCarrito) m.ItemCarrito {
			if mismo(i) {
				return m.ItemCarrito{Producto: p, Cantidad: i.Cantidad + cant}
			}
			return i
		})
	}
	return append(append([]m.ItemCarrito{}, c...), m.ItemCarrito{Producto: p, Cantidad: cant})
}

func Quitar(c []m.ItemCarrito, cod string) []m.ItemCarrito {
	return fn.Filter(c, func(i m.ItemCarrito) bool { return i.Producto.Codigo != cod })
}

func Subtotal(c []m.ItemCarrito) float64 {
	return fn.Reduce(c, 0.0, func(a float64, i m.ItemCarrito) float64 { return a + i.Subtotal() })
}

// ConIVA aplica el IVA de Ecuador (15 %).
var ConIVA = func(v float64) float64 { return v * 1.15 }

func Descuento(pct float64) func(float64) float64 {
	return func(v float64) float64 { return v * (1 - pct/100) }
}

func Total(c []m.ItemCarrito, desc func(float64) float64) float64 {
	return fn.Compose(desc, ConIVA)(Subtotal(c))
}
