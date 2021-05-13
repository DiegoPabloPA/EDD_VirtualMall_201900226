package ArregloEstatico

import (
	"Proyecto1/Comentarios"
	vector "Proyecto1/ListaDVec"
	"Proyecto1/ListaDoble"
	"strings"
)



func CrearArreglo(Datos vector.InfoVector)(a[]string,b[]string){

	var Filas []string
	var Columnas[]string

	for a:=0;a<len(Datos.Indice1);a++{
		if Filas==nil{
			Filas=append(Filas,Datos.Indice1[a].Fila)
		}else{
			bandera:=true
			for b:=0;b<len(Filas);b++{
				if strings.EqualFold(Filas[b],Datos.Indice1[a].Fila){
					bandera=false
				}
			}
			if bandera{
				Filas=append(Filas,Datos.Indice1[a].Fila)
			}
		}
	}

	for a:=0;a<len(Datos.Indice1);a++{
		for b:=0;b<len(Datos.Indice1[a].Indice2);b++{
			if Columnas==nil{
				Columnas=append(Columnas,Datos.Indice1[a].Indice2[b].Columna)
			}else{
				bandera:=true
				for c:=0;c<len(Columnas);c++{
					if strings.EqualFold(Columnas[c],Datos.Indice1[a].Indice2[b].Columna){
						bandera=false
					}
				}
				if bandera{
					Columnas=append(Columnas,Datos.Indice1[a].Indice2[b].Columna)
				}
			}
		}
	}

return Filas,Columnas
}
func ColumnMajor(array []ListaDoble.ListaDE,Filas []string,Columnas[]string,Datos vector.InfoVector){
	tamanioFila:=len(Filas)
	tamanioColumna:=len(Columnas)

	for a:=0;a<tamanioFila;a++{
		for b:=0;b<tamanioColumna;b++{
				for d:=0;d<len(Datos.Indice1[a].Indice2[b].Tienda);d++{
					Ingreso:= ListaDoble.Tienda{
						Fila:         Datos.Indice1[a].Fila,
						Columna:      Datos.Indice1[a].Indice2[b].Columna,
						Calificacion: Datos.Indice1[a].Indice2[b].Tienda[d].Calificacion,
						Nombre:       Datos.Indice1[a].Indice2[b].Tienda[d].Nombre,
						Descripcion:  Datos.Indice1[a].Indice2[b].Tienda[d].Descripcion,
						Contacto:     Datos.Indice1[a].Indice2[b].Tienda[d].Contacto,
						Logo: Datos.Indice1[a].Indice2[b].Tienda[d].Logo,
						Comentarios: make([]Comentarios.DatosComentarios,7),
					}

					tam:=Ingreso.Calificacion-1
					ListaDoble.InsertarNuevaTienda(&Ingreso,&array[a+tamanioFila*(b+tamanioColumna*tam)])

			}
		}
	}
}
