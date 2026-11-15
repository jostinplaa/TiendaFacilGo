// Paquete usuarios: módulo 2, registra a los clientes.
package usuarios

import (
	"strings"
	"tiendafacil/modelo"
)

// CedulaValida revisa que la cédula tenga 10 dígitos y que todos sean números.
func CedulaValida(cedula string) bool {
	if len(cedula) != 10 {
		return false
	}
	for _, letra := range cedula {
		if letra < '0' || letra > '9' {
			return false
		}
	}
	return true
}

// CorreoValido revisa, de forma sencilla, que el correo tenga @ y un punto.
func CorreoValido(correo string) bool {
	return strings.Contains(correo, "@") && strings.Contains(correo, ".")
}

// Registrar crea el cliente solo si los datos son válidos.
// Devuelve el cliente y true, o un cliente vacío y false.
func Registrar(cedula, nombre, correo string) (modelo.Cliente, bool) {
	if !CedulaValida(cedula) || !CorreoValido(correo) || nombre == "" {
		return modelo.Cliente{}, false
	}
	return modelo.Cliente{Cedula: cedula, Nombre: nombre, Correo: correo}, true
}
