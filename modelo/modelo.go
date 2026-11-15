// Paquete modelo: aquí están los datos que usa toda la tienda.
package modelo

// Producto es algo que la tienda vende.
type Producto struct {
	Codigo    string
	Nombre    string
	Categoria string
	Precio    float64
	Stock     int
}

// Cliente es la persona que compra.
type Cliente struct {
	Cedula string
	Nombre string
	Correo string
}

// ItemCarrito es un producto dentro del carrito, con la cantidad que se quiere comprar.
type ItemCarrito struct {
	Producto Producto
	Cantidad int
}

// EstadoPedido indica en qué paso está un pedido.
type EstadoPedido string

const (
	Pendiente EstadoPedido = "PENDIENTE"
	Pagado    EstadoPedido = "PAGADO"
)

// Pedido es la compra ya confirmada.
type Pedido struct {
	Numero  int
	Cliente Cliente
	Items   []ItemCarrito
	Total   float64
	Estado  EstadoPedido
}
