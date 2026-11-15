// Package pagos: módulo 5, pagos simulados. Cada método es una función validadora.
package pagos

import (
	"regexp"
	"strings"
	m "tiendafacil/modelo"
)

var Metodos = map[string]func(string) bool{
	"TARJETA":       func(d string) bool { return regexp.MustCompile(`^\d{16}$`).MatchString(strings.ReplaceAll(d, " ", "")) },
	"TRANSFERENCIA": func(d string) bool { return regexp.MustCompile(`^\d{6,}$`).MatchString(d) },
	"EFECTIVO":      func(string) bool { return true },
}

func Pagar(p m.Pedido, metodo, dato string) m.Pedido {
	if v, ok := Metodos[metodo]; ok && v(dato) {
		return p.ConEstado(m.Pagado)
	}
	return p
}
