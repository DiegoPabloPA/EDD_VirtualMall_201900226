package ListaDoble

import "fmt"

type Nodo struct {
	Anterior, Siguiente *Nodo
	 Datos *Tienda

}

type Tienda struct{
	Fila string
	Columna string
	Calificacion int
	Nombre,Descripcion,Contacto string

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













