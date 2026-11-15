// Paquete reportes: módulo 6, información para que el dueño tome decisiones.
//
// ESTADO: planificado. Se programará en la siguiente etapa.
//
// Funciones que se van a implementar:
//   - TotalVendido(pedidos) float64
//     Suma el total de los pedidos que ya están PAGADOS.
//   - UnidadesPorProducto(pedidos) map[string]int
//     Cuenta cuántas unidades se vendieron de cada producto.
//   - StockBajo(inventario, minimo) []Producto
//     Reutilizará catalogo.Filtrar pasándole una función anónima
//     que se quede con los productos que tienen menos stock que el mínimo.
package reportes
