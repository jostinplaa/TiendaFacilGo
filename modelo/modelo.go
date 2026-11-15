// Package modelo define los datos del negocio (se tratan como inmutables).
package modelo

import "time"

type Producto struct {
	Codigo, Nombre, Categoria string
	Precio                    float64
	Stock                     int
}

// ConStock devuelve una copia del producto con otro stock.
func (p Producto) ConStock(n int) Producto { p.Stock = n; return p }

type Cliente struct{ Cedula, Nombre, Correo, Direccion string }

type ItemCarrito struct {
	Producto Producto
	Cantidad int
}

func (i ItemCarrito) Subtotal() float64 { return i.Producto.Precio * float64(i.Cantidad) }

type EstadoPedido string

const (
	Pendiente EstadoPedido = "PENDIENTE"
	Pagado    EstadoPedido = "PAGADO"
	Enviado   EstadoPedido = "ENVIADO"
	Entregado EstadoPedido = "ENTREGADO"
	Cancelado EstadoPedido = "CANCELADO"
)

type Pedido struct {
	Numero  int
	Cliente Cliente
	Items   []ItemCarrito
	Total   float64
	Estado  EstadoPedido
	Fecha   time.Time
}

// ConEstado devuelve una copia del pedido con otro estado.
func (p Pedido) ConEstado(e EstadoPedido) Pedido { p.Estado = e; return p }
