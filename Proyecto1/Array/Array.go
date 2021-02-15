package Array

import "Proyecto1/ListaDoble"

type Nodo struct{
	anterior,siguiente *Nodo
	ejex, ejey string
	ejez,posicion int
	tiendas ListaDoble.ListaDE
}
type Arreglo struct {
	inicio,final*Nodo
	tamanio int
}

func CrearArreglo(tamanio int,array *Arreglo){
	for a:=0;a<tamanio;a++{
		var Nuevo=&Nodo{nil,nil,nil,nil,nil,a,nil}
		if array.inicio==nil{
			array.inicio=Nuevo
			array.final=Nuevo
			array.tamanio++
		}else{
			array.final.siguiente=Nuevo
			array.final=Nuevo
			array.tamanio++
		}
	}
}