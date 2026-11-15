// Package usuarios: módulo 2, registro y validación de clientes.
package usuarios

import (
	"errors"
	"regexp"
	"strings"
	"tiendafacil/fn"
	m "tiendafacil/modelo"
)

var (
	CedulaValida = func(c string) bool { return regexp.MustCompile(`^\d{10}$`).MatchString(c) }
	CorreoValido = func(c string) bool { return regexp.MustCompile(`^[\w.+-]+@[\w-]+\.[\w.]+$`).MatchString(c) }
)

func Registrar(cedula, nombre, correo, dir string) (m.Cliente, error) {
	if !CedulaValida(cedula) || !CorreoValido(correo) || strings.TrimSpace(nombre) == "" {
		return m.Cliente{}, errors.New("datos del cliente no válidos")
	}
	return m.Cliente{Cedula: cedula, Nombre: nombre, Correo: correo, Direccion: dir}, nil
}

func Buscar(cs []m.Cliente, cedula string) (m.Cliente, bool) {
	return fn.Find(cs, func(c m.Cliente) bool { return c.Cedula == cedula })
}
