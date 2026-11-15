# TiendaFácil (Go)

Sistema de Gestión de E-commerce hecho en Go para la materia de Programación Orientada a Objetos (UIDE, AA1, Etapa 1: Planeación del software).

Permite a una tienda pequeña manejar su catálogo, registrar clientes, armar el carrito con descuento e IVA del 15 %, crear pedidos, registrar pagos simulados y ver reportes de ventas.

## Cómo ejecutarlo

Necesitas Go 1.21 o superior.

```bash
go run .
```

## Módulos (paquetes)

| Paquete    | Qué hace | Estado |
|------------|----------|--------|
| `modelo`   | Los datos: Producto, Cliente, ItemCarrito y Pedido | ✅ Hecho |
| `catalogo` | Agregar, buscar, filtrar, ordenar y agrupar productos | ✅ Hecho |
| `usuarios` | Registrar clientes validando cédula y correo | ✅ Hecho |
| `carrito`  | Agregar y quitar productos, descuento e IVA 15 % | ✅ Hecho |
| `pedidos`  | Crear el pedido si hay stock y descontar el inventario | 📝 Planificado |
| `pagos`    | Pago simulado con tarjeta, transferencia o efectivo | 📝 Planificado |
| `reportes` | Total vendido, unidades por producto y stock bajo | 📝 Planificado |
| `main.go`  | Muestra lo que ya funciona: catálogo, cliente y carrito | ✅ Hecho |

## Avance

Esta es la **Etapa 1 (planeación)**. Ya está programada la base del sistema
(modelo, catálogo, usuarios y carrito). Los paquetes `pedidos`, `pagos` y
`reportes` ya existen, pero por ahora solo tienen, en comentarios, las funciones
que se van a programar en las siguientes etapas. El software completo se
presenta al final del curso.

## Programación funcional

- **Funciones que reciben funciones:** `catalogo.Filtrar` recibe la condición como función.
- **Funciones que devuelven funciones:** `carrito.Descuento(10)` devuelve una función que quita el 10 %.
- **Funciones anónimas:** en `catalogo.OrdenarPorPrecio` (con `sort.Slice`); después también en `reportes.StockBajo` y en los métodos de `pagos`.
- **No se modifican los datos:** `catalogo.Agregar`, `carrito.Agregar` y `carrito.Quitar` devuelven listas nuevas.

Solo usa la librería estándar de Go (`fmt`, `strings`, `sort`).
