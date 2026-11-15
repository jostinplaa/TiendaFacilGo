// TiendaFácil: sistema de gestión de e-commerce.
// main.go muestra el avance de la Etapa 1: catálogo, clientes y carrito.
// Los pedidos, pagos y reportes están planificados para la siguiente etapa.
package main

import (
	"fmt"
	"tiendafacil/carrito"
	"tiendafacil/catalogo"
	"tiendafacil/modelo"
	"tiendafacil/usuarios"
)

func main() {
	// 1. Inventario inicial de la tienda
	inventario := []modelo.Producto{
		{Codigo: "P01", Nombre: "Audífonos Bluetooth", Categoria: "Electrónica", Precio: 25.00, Stock: 10},
		{Codigo: "P02", Nombre: "Mouse inalámbrico", Categoria: "Electrónica", Precio: 12.50, Stock: 3},
		{Codigo: "P03", Nombre: "Camiseta UIDE", Categoria: "Ropa", Precio: 15.00, Stock: 20},
		{Codigo: "P04", Nombre: "Termo de acero", Categoria: "Hogar", Precio: 9.90, Stock: 0},
	}

	// 2. Catálogo: solo productos con stock, del más barato al más caro
	fmt.Println("=== 1. Catálogo disponible ===")
	disponibles := catalogo.Filtrar(inventario, catalogo.HayStock)
	for _, p := range catalogo.OrdenarPorPrecio(disponibles) {
		fmt.Printf("%s - %s: $%.2f (quedan %d)\n", p.Codigo, p.Nombre, p.Precio, p.Stock)
	}

	fmt.Println("\nSolo Electrónica:")
	for _, p := range catalogo.Filtrar(inventario, catalogo.DeCategoria("Electrónica")) {
		fmt.Println("-", p.Nombre)
	}

	// 3. Registrar al cliente
	fmt.Println("\n=== 2. Registro del cliente ===")
	cliente, ok := usuarios.Registrar("2200000001", "Jostin Herrera", "jostin@ejemplo.com")
	if !ok {
		fmt.Println("Los datos del cliente no son válidos")
		return
	}
	fmt.Println("Cliente registrado:", cliente.Nombre)

	// 4. Armar el carrito: 1 audífono y 2 camisetas
	fmt.Println("\n=== 3. Carrito de compras ===")
	audifonos, _ := catalogo.Buscar(inventario, "P01")
	camiseta, _ := catalogo.Buscar(inventario, "P03")

	miCarrito := []modelo.ItemCarrito{}
	miCarrito = carrito.Agregar(miCarrito, audifonos, 1)
	miCarrito = carrito.Agregar(miCarrito, camiseta, 2)

	total := carrito.Total(miCarrito, carrito.Descuento(10))
	fmt.Printf("Subtotal: $%.2f\n", carrito.Subtotal(miCarrito))
	fmt.Printf("Total con 10%% de descuento + IVA 15%%: $%.2f\n", total)

	// Lo que falta (siguiente etapa)
	fmt.Println("\n=== Próximo avance ===")
	fmt.Println("- pedidos: crear el pedido y descontar el stock")
	fmt.Println("- pagos: pago simulado con tarjeta, transferencia o efectivo")
	fmt.Println("- reportes: total vendido, unidades por producto y stock bajo")
}
