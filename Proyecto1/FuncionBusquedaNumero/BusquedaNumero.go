package BusquedaNumero

import (
	Busqueda3Parametros "Proyecto1/Funciones"
	"Proyecto1/ListaDoble"
)

func Buscar(a int,lista []ListaDoble.ListaDE)(retorno []Busqueda3Parametros.EstructuraRespuesta){
	var TiendasEncontradas []Busqueda3Parametros.EstructuraRespuesta
	aux:=lista[a].Inicio

	if aux!=nil{
	for aux!=nil{
		TiendasEncontradas=append(TiendasEncontradas,Busqueda3Parametros.EstructuraRespuesta{
			Nombre:       aux.Datos.Nombre,
			Descripcion:  aux.Datos.Descripcion,
			Contacto:     aux.Datos.Contacto,
			Calificacion: aux.Datos.Calificacion,
		})
		aux=aux.Siguiente
		}
	}

return TiendasEncontradas
}