# TiendaFácil (Go)

Sistema de Gestión de E-commerce hecho en Go para la materia de Programación Orientada a Objetos (UIDE, AA1, Etapa 1: Planeación del software).

Permite a una tienda pequeña manejar su catálogo, registrar clientes, armar el carrito con descuento e IVA del 15 %, crear pedidos, registrar pagos simulados y ver reportes de ventas.

## Cómo ejecutarlo

Necesitas Go 1.21 o superior.

```bash
go run .
```

## Módulos (paquetes)

| Paquete    | Qué hace |
|------------|----------|
| `modelo`   | Los datos: Producto, Cliente, ItemCarrito y Pedido |
| `catalogo` | Agregar, buscar, filtrar, ordenar y agrupar productos |
| `usuarios` | Registrar clientes validando cédula y correo |
| `carrito`  | Agregar y quitar productos, descuento e IVA 15 % |
| `pedidos`  | Crear el pedido si hay stock y descontar el inventario |
| `pagos`    | Pago simulado con tarjeta, transferencia o efectivo |
| `reportes` | Total vendido, unidades por producto y stock bajo |
| `main.go`  | Muestra el flujo completo de una venta |

## Programación funcional

- **Funciones que reciben funciones:** `catalogo.Filtrar` recibe la condición como función.
- **Funciones que devuelven funciones:** `carrito.Descuento(10)` devuelve una función que quita el 10 %.
- **Funciones anónimas:** en `reportes.StockBajo` y en los métodos de `pagos`.
- **No se modifican los datos:** `Agregar`, `Quitar` y `DescontarStock` devuelven listas nuevas.

Solo usa la librería estándar de Go (`fmt`, `strings`, `sort`).
