// Paquete pagos: módulo 5, pagos simulados (no se cobra dinero real).
//
// ESTADO: planificado. Se programará en la siguiente etapa.
//
// Idea: guardar cada método de pago en un mapa de funciones, donde cada
// función anónima revisa el dato de ese método:
//   - TARJETA: el número debe tener 16 dígitos (sin contar espacios).
//   - TRANSFERENCIA: el comprobante debe tener al menos 6 caracteres.
//   - EFECTIVO: siempre es válido.
//
// Funciones que se van a implementar:
//   - Pagar(pedido, metodo, dato) Pedido
//     Si el dato es válido, devuelve una copia del pedido en estado PAGADO.
package pagos
