package OrdenAlfabetico

import (
	"Proyecto1/ListaDoble"
	"sort"
	"strings"
)

func Ordenar(informacion []ListaDoble.ListaDE){

	for a:=0;a<len(informacion);a++{
		if informacion[a].Tamanio>1{
			var orden []string
			aux:=informacion[a].Inicio
			for aux!=nil{
				orden=append(orden,aux.Datos.Nombre)
				aux=aux.Siguiente
			}
			sort.Strings(orden)
			var listaordenada ListaDoble.ListaDE

			for b:=0;b<len(orden);b++{
				aux2:=informacion[a].Inicio
			for aux2!=nil{
				if strings.EqualFold(orden[b],aux2.Datos.Nombre){

					ListaDoble.InsertarNuevaTienda(aux2.Datos,&listaordenada)
					aux2=nil
				}else{
				aux2=aux2.Siguiente}
			}

			}



			informacion[a].Inicio=listaordenada.Inicio
			informacion[a].Final=listaordenada.Final

		}

	}

}