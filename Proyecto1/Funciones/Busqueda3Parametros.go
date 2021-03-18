package Busqueda3Parametros

import (
	Inventarios2 "Proyecto1/Inventarios"
	"Proyecto1/ListaDoble"
	"encoding/json"
	"strings"
)

type EstructuraBusqueda struct {
	Departamento string `json:"Departamento"`
	Nombre string `json:"Nombre"`
	Calificacion int `json:"Calificacion"`

}




type EstructuraRespuesta struct {
	Nombre,Descripcion,Contacto string
	Calificacion int
	Logo string
}
type EstructuraFallido struct {
	Mensaje string
}


func Buscar3Param(parametro EstructuraBusqueda,informacion []ListaDoble.ListaDE)(resultado []byte){
	var AJson []byte
	for a:=0;a<len(informacion);a++{
		if informacion[a].Tamanio!=0{
			aux:=informacion[a].Inicio
			for aux!=nil{
				if strings.EqualFold(aux.Datos.Columna,parametro.Departamento){
					if parametro.Calificacion==aux.Datos.Calificacion{
						if strings.EqualFold(aux.Datos.Nombre,parametro.Nombre){
							a=len(informacion)
							envio:= EstructuraRespuesta{
								Nombre:       aux.Datos.Nombre,
								Descripcion:  aux.Datos.Descripcion,
								Contacto:     aux.Datos.Contacto,
								Calificacion: aux.Datos.Calificacion,
								Logo: aux.Datos.Logo,
							}
							AJson,_=json.Marshal(envio)
							aux=nil
						}else{
							aux=aux.Siguiente
						}
					}else{
						aux=aux.Siguiente
					}
				}else{
					aux=aux.Siguiente
				}
			}

		}
	}
	if len(AJson)==0{
		mensaje:=EstructuraFallido{Mensaje: "NO EXISTEN DATOS ACORDES A LOS PARAMETROS ENVIADOS"}
		AJson,_=json.Marshal(mensaje)
	}

	return AJson
}

func BuscarPedidos(parametro EstructuraBusqueda,informacion []ListaDoble.ListaDE)(resultado []byte){
	var AJson []byte
	for a:=0;a<len(informacion);a++{
		if informacion[a].Tamanio!=0{
			aux:=informacion[a].Inicio
			for aux!=nil{
				if strings.EqualFold(aux.Datos.Columna,parametro.Departamento){
					if parametro.Calificacion==aux.Datos.Calificacion{
						if strings.EqualFold(aux.Datos.Nombre,parametro.Nombre){
							a=len(informacion)
							var archie []Inventarios2.InventarioJson
							archie=Inventarios2.EnviarDatos(aux.Datos.Inventario.Raiz,archie)
							AJson,_=json.Marshal(archie)
							aux=nil
						}else{
							aux=aux.Siguiente
						}
					}else{
						aux=aux.Siguiente
					}
				}else{
					aux=aux.Siguiente
				}
			}

		}
	}
	if len(AJson)==0{
		mensaje:=EstructuraFallido{Mensaje: "NULO"}
		AJson,_=json.Marshal(mensaje)
	}

	return AJson
}
