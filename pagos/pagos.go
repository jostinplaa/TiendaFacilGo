// Paquete pagos: módulo 5, pagos simulados (no se cobra dinero real).
package pagos

import (
	"strings"
	"tiendafacil/modelo"
)

// Cada método de pago tiene su propia función que revisa el dato.
// Se guardan en un mapa: nombre del método -> función que valida.
var metodos = map[string]func(string) bool{
	"TARJETA": func(numero string) bool {
		numero = strings.ReplaceAll(numero, " ", "")
		return len(numero) == 16
	},
	"TRANSFERENCIA": func(comprobante string) bool {
		return len(comprobante) >= 6
	},
	"EFECTIVO": func(dato string) bool {
		return true
	},
}

// Pagar devuelve el pedido como PAGADO si el método existe y el dato es válido.
// Si no, devuelve el pedido igual que estaba.
func Pagar(pedido modelo.Pedido, metodo string, dato string) modelo.Pedido {
	validar, existe := metodos[metodo]
	if existe && validar(dato) {
		pedido.Estado = modelo.Pagado
	}
	return pedido
}
