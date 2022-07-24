package main

import "fmt"

var (
	TiendaArg = "ARG"
	TiendaBrs = "BRZ"
)

type MiTiendaArg struct {
	Nombre    string
	Direccion string
	Moneda    string
}

func (t *MiTiendaArg) Presentacion() string {
	return fmt.Sprintf("\nMi tienda es: %s\nDirección: %s\nMoneda: %s\n", t.Nombre, t.Direccion, t.Moneda)
}

type MiTiendaBrazil struct {
	Nombre    string
	Direccion string
}

func (t *MiTiendaBrazil) Presentacion() string {
	return fmt.Sprintf("\nMi tienda es: %s\nDirección: %s\n", t.Nombre, t.Direccion)
}

// Interface Ecommerce
type Ecommerce interface {
	Presentacion() string
}

func New(tienda string) Ecommerce {
	switch tienda {
	case TiendaArg:
		return &MiTiendaArg{
			Nombre:    "Mi tienda ARG.",
			Direccion: "Buenos Aires",
			Moneda:    "Peso",
		}
	case TiendaBrs:
		return &MiTiendaBrazil{
			Nombre:    "Mi tienda Brazil",
			Direccion: "Rio de janeiro",
		}
	default:
		return nil
	}
}

func main() {
	t := New(TiendaArg)
	//t := New(TiendaBrs)

	DetalleTienda(t)
}

func DetalleTienda(e Ecommerce) {
	fmt.Println(e.Presentacion())
}
