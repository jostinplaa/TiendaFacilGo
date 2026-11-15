// Paquete carrito: módulo 3, arma la compra y calcula el total.
package carrito

import "tiendafacil/modelo"

// IVA de Ecuador: 15 %.
const IVA = 0.15

// Agregar devuelve un carrito NUEVO con el producto agregado.
// El carrito anterior no se modifica.
func Agregar(carrito []modelo.ItemCarrito, p modelo.Producto, cantidad int) []modelo.ItemCarrito {
	nuevo := []modelo.ItemCarrito{}
	for _, item := range carrito {
		nuevo = append(nuevo, item)
	}
	nuevo = append(nuevo, modelo.ItemCarrito{Producto: p, Cantidad: cantidad})
	return nuevo
}

// Quitar devuelve un carrito nuevo sin el producto de ese código.
func Quitar(carrito []modelo.ItemCarrito, codigo string) []modelo.ItemCarrito {
	nuevo := []modelo.ItemCarrito{}
	for _, item := range carrito {
		if item.Producto.Codigo != codigo {
			nuevo = append(nuevo, item)
		}
	}
	return nuevo
}

// Subtotal suma precio x cantidad de cada producto.
func Subtotal(carrito []modelo.ItemCarrito) float64 {
	suma := 0.0
	for _, item := range carrito {
		suma = suma + item.Producto.Precio*float64(item.Cantidad)
	}
	return suma
}

// Descuento crea y devuelve una función que quita un porcentaje.
// Ejemplo: Descuento(10) devuelve una función que quita el 10 %.
func Descuento(porcentaje float64) func(float64) float64 {
	return func(valor float64) float64 {
		return valor - valor*porcentaje/100
	}
}

// AplicarIVA suma el 15 % de IVA a un valor.
func AplicarIVA(valor float64) float64 {
	return valor + valor*IVA
}

// Total aplica primero el descuento y después el IVA.
// El descuento llega como función, así se puede usar cualquier descuento.
func Total(carrito []modelo.ItemCarrito, descuento func(float64) float64) float64 {
	subtotal := Subtotal(carrito)
	conDescuento := descuento(subtotal)
	return AplicarIVA(conDescuento)
}
