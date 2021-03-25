package EnviarCarrito

import (
	"Proyecto1/Carrito"
	"Proyecto1/ListaDoble"
	"strings"
)

type DescripcionCompra struct {
	Codigo int
	Nombre string
	Tienda string
	Cantidad int
	Total float64
}

type EliminacionProd struct {
	Producto string `json:"Producto"`
	Nombre   string `json:"Nombre"`
	Codigo   int    `json:"Codigo"`
}
func EliminardelCarrito(datos EliminacionProd,carroActual[]Carrito.JsonPedidoIndv)(carroNuevo[]Carrito.JsonPedidoIndv){
	for a:=0;a<len(carroActual);a++{
		if strings.EqualFold(datos.Nombre,carroActual[a].Nombre)&&datos.Codigo==carroActual[a].Codigo{
			carroNuevo=append(carroActual[:a],carroActual[a+1:]...)
		}
	}
	return carroNuevo
}




func ConvertirCarritoDesc(carroActual []Carrito.JsonPedidoIndv,arreglo []ListaDoble.ListaDE)(descJsonCompra []DescripcionCompra){
	for a:=0; a<len(carroActual);a++{
		for b:=0;b<len(arreglo);b++{
			if arreglo[b].Tamanio>0{
				aux:=arreglo[b].Inicio
				for aux!=nil{
					if strings.EqualFold(carroActual[a].Departamento,aux.Datos.Columna)&&strings.EqualFold(carroActual[a].Nombre,aux.Datos.Nombre)&&carroActual[a].Calificacion==aux.Datos.Calificacion{
						aux2:=aux.Datos.Inventario.Raiz
						for aux2!=nil{
							if carroActual[a].Codigo<aux2.Datos.Codigo{
								aux2=aux2.Izquierda
							}else if carroActual[a].Codigo>aux2.Datos.Codigo{
								aux2=aux2.Derecha
							}else if carroActual[a].Codigo==aux2.Datos.Codigo{
								nuevo:=DescripcionCompra{
									Codigo: carroActual[a].Codigo,
									Nombre:   aux2.Datos.Nombre,
									Tienda: carroActual[a].Nombre,
									Cantidad: carroActual[a].Cantidad,
									Total:    float64(carroActual[a].Cantidad)*aux2.Datos.Precio,
								}
								descJsonCompra=append(descJsonCompra,nuevo)
								aux2=aux2.Derecha
							}
						}
					}
					aux=aux.Siguiente
				}
			}
		}
	}
	return descJsonCompra
}