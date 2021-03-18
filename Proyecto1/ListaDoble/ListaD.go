package ListaDoble

import (
	 "Proyecto1/EstructuraAVL"
	"fmt"
	"strings"
)

type Nodo struct {
	Anterior, Siguiente *Nodo
	 Datos *Tienda

}

type Tienda struct{
	Fila string
	Columna string
	Calificacion int
	Nombre,Descripcion,Contacto,Logo string
	Inventario EstructuraAVL.AVL

}

type ListaDE struct{
	Inicio *Nodo
	Final *Nodo
	Tamanio int
}
func NuevaLDE()*ListaDE{
	return &ListaDE{nil,nil,0}
}


func InsertarNuevaTienda(infor *Tienda, almacen *ListaDE){
	var nuevo = &Nodo{nil,nil,infor}
	if almacen.Inicio==nil{
		almacen.Inicio=nuevo
		almacen.Final=nuevo
		almacen.Tamanio++
		}else{
			nuevo.Anterior=almacen.Final
			almacen.Final.Siguiente=nuevo

			almacen.Final=nuevo

			almacen.Tamanio++
	}
}
func Imprimir(lista*ListaDE){
	fmt.Println(lista.Inicio.Datos.Nombre)
}

func EliminarTienda(parametro EstructuraEliminacion,informacion []ListaDE)(resp bool){
	resp=false
	for a:=0;a<len(informacion);a++{
		if informacion[a].Tamanio!=0{
			aux:=informacion[a].Inicio
			for aux!=nil{
				if strings.EqualFold(aux.Datos.Columna,parametro.Categoria){
					if parametro.Calificacion==aux.Datos.Calificacion{
						if strings.EqualFold(aux.Datos.Nombre,parametro.Nombre){
							if aux==informacion[a].Inicio&&informacion[a].Tamanio==1{
								informacion[a].Inicio=nil
								informacion[a].Final=nil
								informacion[a].Tamanio=0
								resp=true
							}else if aux==informacion[a].Inicio&&informacion[a].Tamanio>1{
								aux2:=informacion[a].Inicio.Siguiente
								aux2.Anterior=nil
								informacion[a].Inicio=aux2
								aux3:=informacion[a].Inicio
								for aux3!=informacion[a].Final{
									aux3=aux3.Siguiente
								}
								informacion[a].Final=aux3

								informacion[a].Tamanio-=1
								resp=true
							}else if aux==informacion[a].Final&&informacion[a].Tamanio>1{
								aux2:=informacion[a].Final.Anterior
								aux2.Siguiente=nil

								informacion[a].Final=aux2
								informacion[a].Tamanio-=1
								aux3:=informacion[a].Final
								for aux3!=informacion[a].Inicio{
									aux3=aux3.Anterior
								}
								informacion[a].Inicio=aux3
								resp=true
							}else{
								aux2:=aux.Siguiente
								aux3:=aux.Anterior
								aux.Anterior.Siguiente=aux2
								aux.Siguiente.Anterior=aux3
								for aux!=informacion[a].Inicio{
									aux=aux.Anterior
								}
								informacion[a].Inicio=aux
								for aux!=informacion[a].Final{
									aux=aux.Siguiente
								}
								informacion[a].Final=aux
								informacion[a].Tamanio-=1
								resp=true
							}

							a=len(informacion)
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
	return resp

}


type EstructuraEliminacion struct {
	Nombre string `json:"Nombre"`
	Categoria string `json:"Categoria"`
	Calificacion int `json:"Calificacion"`

}










