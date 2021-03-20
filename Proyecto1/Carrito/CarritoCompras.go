package Carrito

import (
	"Proyecto1/ListaDoble"
	"strings"
)

type JsonPedidoIndv struct {
	Departamento string `json:"Departamento"`
	Nombre       string `json:"Nombre"`
	Calificacion int    `json:"Calificacion"`
	Codigo       int    `json:"Codigo"`
	Cantidad     int    `json:"Cantidad"`
}
type Respuesta struct {
	Res string
}


func ValidarPedido(datos JsonPedidoIndv,Arreglo []ListaDoble.ListaDE)bool{
	bandera:=false
	for a:=0;a<len(Arreglo);a++{
		if Arreglo[a].Tamanio>0{
			aux:=Arreglo[a].Inicio
			for aux!=nil{
				if strings.EqualFold(datos.Departamento,aux.Datos.Columna)&&strings.EqualFold(datos.Nombre,aux.Datos.Nombre)&&aux.Datos.Calificacion==datos.Calificacion{
					aux2:=aux.Datos.Inventario.Raiz
					for aux2!=nil{
						if datos.Codigo<aux2.Datos.Codigo{
							aux2=aux2.Izquierda
						}else if datos.Codigo>aux2.Datos.Codigo{
							aux2=aux2.Derecha
						}else{
							if aux2.Datos.Cantidad>=datos.Cantidad{
								bandera=true
								aux2=nil
							}else{
								bandera=false
								aux2=nil
							}
						}
					}

				}
				aux=aux.Siguiente
			}
		}
	}
	return bandera
}