package ListaDoble

import "fmt"

type Nodo struct {
	anterior, siguiente *Nodo
	 datos *Tienda

}

type Tienda struct{
	Nombre,Descripcion,Contacto string
	Calificacion int
}

type ListaDE struct{
	inicio *Nodo
	final *Nodo
	tamanio int
}
func NuevaLDE()*ListaDE{
	return &ListaDE{nil,nil,0}
}


func InsertarNuevaTienda(infor *Tienda, almacen *ListaDE){
	var nuevo = &Nodo{nil,nil,infor}
	if almacen.inicio==nil{
		almacen.inicio=nuevo
		almacen.final=nuevo
		almacen.tamanio++
		}else{
			almacen.final.siguiente=nuevo
			almacen.final=nuevo
			almacen.tamanio++
	}
}
func Imprimir(lista*ListaDE){
	fmt.Println(lista.inicio.datos.Nombre)
}













