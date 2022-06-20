package main

import "fmt"

type Producto interface {
	Print()
	UpdatePrecio(float64)
}

type producto struct {
	Nombre string
	Stock  int
	Precio float64
}

/*
// Usar este ejemplo para mostrar que no se actualiza

func (p producto) Print() {
	fmt.Println(fmt.Sprintf("Producto: %s - Stock: %d - Precio: %.2f", p.Nombre, p.Stock, p.Precio))
}

func (p producto) UpdatePrecio(precio float64) {
	p.Precio = precio
}

func NewProducto(nombre string, stock int, precio float64) Producto {
	return producto{
		Nombre: nombre,
		Stock: stock,
		Precio: precio,
	}
}
*/

func NewProducto(nombre string, stock int, precio float64) Producto {
	return &producto{
		Nombre: nombre,
		Stock:  stock,
		Precio: precio,
	}
}

func (p *producto) Print() {
	fmt.Println(fmt.Sprintf("Producto: %s - Stock: %d - Precio: %.2f", p.Nombre, p.Stock, p.Precio))
}

func (p *producto) UpdatePrecio(precio float64) {
	p.Precio = precio
}

func main() {

	producto := NewProducto("Coca Cola", 100, 500)

	producto.Print()
	producto.UpdatePrecio(200) // Si los metodos no están especificados como puntero a la estructura. La actualización del precio no ocurre.
	producto.Print()

}
