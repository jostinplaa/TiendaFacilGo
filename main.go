// TiendaFácil (versión Go): demostración del flujo completo.
package main

import (
	"fmt"
	"tiendafacil/carrito"
	"tiendafacil/catalogo"
	"tiendafacil/fn"
	m "tiendafacil/modelo"
	"tiendafacil/pagos"
	"tiendafacil/pedidos"
	"tiendafacil/reportes"
	"tiendafacil/usuarios"
)

func main() {
	inv := []m.Producto{
		{Codigo: "P01", Nombre: "Audífonos Bluetooth", Categoria: "Electrónica", Precio: 25.00, Stock: 10},
		{Codigo: "P02", Nombre: "Mouse inalámbrico", Categoria: "Electrónica", Precio: 12.50, Stock: 3},
		{Codigo: "P03", Nombre: "Camiseta UIDE", Categoria: "Ropa", Precio: 15.00, Stock: 20},
		{Codigo: "P04", Nombre: "Termo de acero", Categoria: "Hogar", Precio: 9.90, Stock: 0},
	}

	fmt.Println("=== Catálogo disponible (ordenado por precio) ===")
	for _, p := range catalogo.OrdenarPorPrecio(catalogo.Filtrar(inv, catalogo.Disponible())) {
		fmt.Printf("%s - %s $%.2f (stock %d)\n", p.Codigo, p.Nombre, p.Precio, p.Stock)
	}

	cli, err := usuarios.Registrar("2200000001", "Jostin Herrera", "jostin@ejemplo.com", "El Coca")
	if err != nil {
		panic(err)
	}
	fmt.Println("\nCliente registrado:", cli.Nombre)

	audif, _ := catalogo.BuscarPorCodigo(inv, "P01")
	cami, _ := catalogo.BuscarPorCodigo(inv, "P03")
	car := carrito.Agregar(carrito.Agregar(nil, audif, 1), cami, 2)
	total := carrito.Total(car, carrito.Descuento(10))
	fmt.Printf("Subtotal: $%.2f | Total con 10%% desc. + IVA 15%%: $%.2f\n", carrito.Subtotal(car), total)

	ped, err := pedidos.Crear(1, cli, car, total)
	if err != nil {
		panic(err)
	}
	ped = pagos.Pagar(ped, "TARJETA", "4111 1111 1111 1111")
	fmt.Printf("Pedido #%d -> %s\n", ped.Numero, ped.Estado)

	inv = pedidos.DescontarStock(inv, ped)
	fmt.Printf("\n=== Reportes ===\nTotal vendido: $%.2f\n", reportes.TotalVendido([]m.Pedido{ped}))
	fmt.Println("Unidades por producto:", reportes.UnidadesPorProducto([]m.Pedido{ped}))
	fmt.Println("Stock bajo (<5):", fn.Map(reportes.StockBajo(inv, 5), func(p m.Producto) string { return p.Nombre }))
}
