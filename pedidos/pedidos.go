// Paquete pedidos: módulo 4, convierte el carrito en pedido y controla el stock.
//
// ESTADO: planificado. Se programará en la siguiente etapa.
//
// Funciones que se van a implementar:
//   - HayStockSuficiente(carrito) bool
//     Revisa que haya unidades para cada producto del carrito.
//   - Crear(numero, cliente, carrito, total) (Pedido, bool)
//     Arma el pedido en estado PENDIENTE solo si el carrito no está vacío y hay stock.
//   - DescontarStock(inventario, pedido) []Producto
//     Devuelve un inventario NUEVO con las unidades vendidas restadas
//     (sin modificar el inventario original).
package pedidos
