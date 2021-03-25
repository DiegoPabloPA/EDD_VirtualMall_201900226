package CambioInventario

import (
	"Proyecto1/Carrito"
	"Proyecto1/EstructuraAVL"
	"Proyecto1/ListaDoble"
	"strings"
)

func EfectuarCompra(arreglo []ListaDoble.ListaDE,carrito []Carrito.JsonPedidoIndv)[]ListaDoble.ListaDE{
aux:=arreglo
	for a:=0;a<len(carrito);a++{
		for b:=0;b<len(aux);b++{
			if aux[b].Tamanio>0{
					aux2:=aux[b].Inicio

					for aux2!=nil{
						if strings.EqualFold(carrito[a].Nombre,aux2.Datos.Nombre)&&strings.EqualFold(carrito[a].Departamento,aux2.Datos.Columna)&&aux2.Datos.Calificacion==carrito[a].Calificacion{
							RestarInventario(aux2.Datos.Inventario.Raiz,carrito[a].Codigo,carrito[a].Cantidad)
							for aux2.Anterior!=nil{
								aux2=aux2.Anterior
							}
							aux[b].Inicio=aux2
							for aux2.Siguiente!=nil{
								aux2=aux2.Siguiente
							}
							aux[b].Final=aux2

						}
						aux2=aux2.Siguiente
					}



			}
		}
	}


return aux
}

func RestarInventario(Raiz *EstructuraAVL.NodoArbol,codigo int, cantidad int){
	for Raiz!=nil{
		if Raiz.Datos.Codigo>codigo{
			Raiz=Raiz.Izquierda
		}else if Raiz.Datos.Codigo<codigo{
			Raiz=Raiz.Derecha
		}else{
			Raiz.Datos.Cantidad=Raiz.Datos.Cantidad-cantidad
			Raiz=Raiz.Derecha
		}
	}

}
