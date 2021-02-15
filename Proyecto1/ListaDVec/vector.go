package vector

import "Proyecto1/ListaDoble"


type InfoVector struct {
	Indice1 []struct{
		Fila string `json:"Indice"`
			Indice2 []struct{
				Columna string `json:"Nombre"`
				Tienda []struct{
					Nombre string `json:"Nombre"`
					Descripcion string `json:"Descripcion"`
					Contacto string `json:"Contacto"`
					Calificacion int  `json:"Calificacion"`
				}`json:"Tiendas"`

	}`json:"Departamentos"`
	}`json:"Datos"`
}
var ListaTienda ListaDoble.ListaDE


func CrearArreglo(vector InfoVector){
	a:=len(vector.Indice1)
	b:=0
	for t:=0; t<a;t++{
		if b<len(vector.Indice1[t].Indice2){
			b=len(vector.Indice1[t].Indice2)
		}
	}





}

