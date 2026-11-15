# TiendaFácil (versión Go)

Versión en Go del Sistema de Gestión de E-commerce (UIDE, POO, AA1). Tiene los mismos 6 módulos que la versión en Java, escritos con estilo funcional: genéricos `Filter`, `Map`, `Reduce` y `Compose`, funciones como valores y copias en lugar de modificar los datos.

## Ejecutar (Go 1.21+)
```bash
go run .
```

Estructura: `modelo`, `fn` (utilidades funcionales), `catalogo`, `usuarios`, `carrito`, `pedidos`, `pagos`, `reportes`, `main.go`.
Solo se usa la librería estándar (sin paquetes de terceros).
