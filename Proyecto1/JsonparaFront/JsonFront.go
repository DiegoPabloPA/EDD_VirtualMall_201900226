package JsonparaFront

import "Proyecto1/ListaDoble"

type EstructuraJsonFront struct {

	Fila,Columna,Nombre,Descripcion,Contacto,Logo string
	Calificacion int

}
func JsonFront(array []ListaDoble.ListaDE)(nuevo []EstructuraJsonFront){
	for a:=0;a<len(array);a++{
		if array[a].Tamanio>0{
			aux:=array[a].Inicio
			for aux!=nil{
				nuevo=append(nuevo,EstructuraJsonFront{
					Fila:         aux.Datos.Fila,
					Columna:      aux.Datos.Columna,
					Nombre:       aux.Datos.Nombre,
					Descripcion:  aux.Datos.Descripcion,
					Contacto:     aux.Datos.Contacto,
					Logo:         aux.Datos.Logo,
					Calificacion: aux.Datos.Calificacion,
				})
				aux=aux.Siguiente
			}
		}
	}
	return nuevo
}