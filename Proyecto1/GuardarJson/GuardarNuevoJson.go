package GuardarJson

import (
	"Proyecto1/ListaDoble"
	"strings"
)


type Tiendas2 struct {
	Nombre string
	Descripcion string
	Contacto string
	Calificacion int
}
type Depa struct {
	Nombre string
	Tiendas []Tiendas2
}
type Dat struct {
	Indice string
	Departamentos []Depa
}
type General struct {
	Datos []Dat
}



func Guardar(array []ListaDoble.ListaDE,fila []string, columna []string)(estructura General){
	var nuevo General
	var filas []Dat
	for a:=0; a<len(fila);a++{
		var columnas []Depa
		for b:=0;b<len(columna);b++{
			var tienda []Tiendas2
			for c:=0;c<len(array);c++{
				if array[c].Inicio!=nil{
					aux:=array[c].Inicio
					for aux!=nil{

						if strings.EqualFold(aux.Datos.Fila,fila[a]) &&strings.EqualFold(aux.Datos.Columna,columna[b]){
						tienda=append(tienda,Tiendas2{
							Nombre:       aux.Datos.Nombre,
							Descripcion:  aux.Datos.Descripcion,
							Contacto:     aux.Datos.Contacto,
							Calificacion: aux.Datos.Calificacion,
						})
						}
						aux=aux.Siguiente
					}

				}
			}
			if len(tienda)==0{
				tienda=[]Tiendas2{}
			}

		columnas=append(columnas,Depa{
			Nombre:  columna[b],
			Tiendas: tienda,
		})
		}
		filas=append(filas,Dat{
			Indice:        fila[a],
			Departamentos: columnas,
		})
	}
	nuevo=General{Datos: filas}



	return nuevo
}
