package AVLMeses

import (
	"Proyecto1/MatrizDispersa"
	"fmt"
	"strings"
)

type NodoAVLMes struct {
	Izquierda, Derecha *NodoAVLMes
	Altura int
	Datos DatosMesAvl
}

type DatosMesAvl struct {
	NoMes int
	Mes string
	Matriz *MatrizDispersa.MatrizDispersa
}

type AVLMeses struct {
	Raiz* NodoAVLMes
}
func Mayor(numa int,numb int)int{
	if numa>numb{
		return numa
	}
	return numb

}
func Imprimir(nodo *NodoAVLMes){
	if nodo!=nil{
		nodo.Datos.Matriz.Imprimir()
		Imprimir(nodo.Izquierda)
		fmt.Println(nodo.Datos.Mes)
		nodo.Datos.Matriz.Imprimir()
		Imprimir(nodo.Derecha)
	}
}


func Altura(evaluar *NodoAVLMes)int{
	if evaluar==nil{
		return -1
	}
	return evaluar.Altura
}
func RotacionSimpleIzquierda(evaluar*NodoAVLMes)*NodoAVLMes{
	aux1:=evaluar.Izquierda
	evaluar.Izquierda=aux1.Derecha
	aux1.Derecha=evaluar
	evaluar.Altura=Mayor(Altura(evaluar.Izquierda),Altura(evaluar.Derecha))+1
	aux1.Altura=Mayor(Altura(evaluar.Izquierda),evaluar.Altura)+1
	evaluar=aux1
	return evaluar
}
func RotacionSimpleDerecha(evaluar *NodoAVLMes)*NodoAVLMes{
	aux1:=evaluar.Derecha
	evaluar.Derecha=aux1.Izquierda
	aux1.Izquierda=evaluar
	evaluar.Altura=Mayor(Altura(evaluar.Izquierda),Altura(evaluar.Derecha))+1
	aux1.Altura=Mayor(Altura(evaluar.Derecha),evaluar.Altura)+1
	evaluar=aux1
	return evaluar
}
func RotacionDobleIzquierda(evaluar *NodoAVLMes)*NodoAVLMes{
	evaluar.Izquierda=RotacionSimpleDerecha(evaluar.Izquierda)
	evaluar=RotacionSimpleIzquierda(evaluar)
	return evaluar
}
func RotacionDobleDerecha(evaluar *NodoAVLMes)*NodoAVLMes{
	evaluar.Derecha=RotacionSimpleIzquierda(evaluar.Derecha)
	evaluar=RotacionSimpleDerecha(evaluar)
	return evaluar
}
func SeleccionMes(mes string) int{
	if strings.EqualFold(mes,"Enero"){
		return 1
	}else if strings.EqualFold(mes,"Febrero"){
		return 2
	}else if strings.EqualFold(mes,"Marzo"){
		return 3
	}else if strings.EqualFold(mes,"Abril"){
		return 4
	}else if strings.EqualFold(mes,"Mayo"){
		return 5
	}else if strings.EqualFold(mes,"Junio"){
		return 6
	}else if strings.EqualFold(mes,"Julio"){
		return 7
	}else if strings.EqualFold(mes,"Agosto"){
		return 8
	}else if strings.EqualFold(mes,"Septiembre"){
		return 9
	}else if strings.EqualFold(mes,"Octubre"){
		return 10
	}else if strings.EqualFold(mes,"Noviembre"){
		return 11
	}else if strings.EqualFold(mes,"Diciembre"){
		return 12
	}else{
		return 0
	}

}
func ConvertirMes(mesIngles string)string{
	if strings.EqualFold(mesIngles,"January"){
		return "Enero"
	}else if strings.EqualFold(mesIngles,"February"){
		return "Febrero"
	}else if strings.EqualFold(mesIngles,"March"){
		return "Marzo"
	}else if strings.EqualFold(mesIngles,"April"){
		return "Abril"
	}else if strings.EqualFold(mesIngles,"May"){
		return "Mayo"
	}else if strings.EqualFold(mesIngles,"June"){
		return "Junio"
	}else if strings.EqualFold(mesIngles,"July"){
		return "Julio"
	}else if strings.EqualFold(mesIngles,"August"){
		return "Agosto"
	}else if strings.EqualFold(mesIngles,"September"){
		return "Septiembre"
	}else if strings.EqualFold(mesIngles,"October"){
		return "Octubre"
	}else if strings.EqualFold(mesIngles,"November"){
		return "Noviembre"
	}else if strings.EqualFold(mesIngles,"December"){
		return "Diciembre"
	}else{
		return ""
	}
}

func  Insertar(nodo *NodoAVLMes,datos MatrizDispersa.Pedido)*NodoAVLMes{
	nuevo:=&NodoAVLMes{
		Izquierda: nil,
		Derecha:   nil,
		Altura:    0,
		Datos:     DatosMesAvl{
			NoMes:  SeleccionMes(ConvertirMes(datos.Mes)),
			Mes:    ConvertirMes(datos.Mes),
			Matriz: &MatrizDispersa.MatrizDispersa{Nodoinicio: nil},
		},
	}
	if nodo==nil{
		nuevo.Datos.Matriz.Init()
		nuevo.Datos.Matriz.InsertarPedido(datos)
		return nuevo
	}

	if  nodo.Datos.NoMes>SeleccionMes(ConvertirMes(datos.Mes)){
		nodo.Izquierda=Insertar(nodo.Izquierda,datos)
		if Altura(nodo.Izquierda)-Altura(nodo.Derecha)==2{
			if SeleccionMes(ConvertirMes(datos.Mes))<nodo.Izquierda.Datos.NoMes{
				nodo=RotacionSimpleIzquierda(nodo)
			}else{
				nodo=RotacionDobleIzquierda(nodo)
			}
		}
	}else if nodo.Datos.NoMes<SeleccionMes(ConvertirMes(datos.Mes)){
		nodo.Derecha=Insertar(nodo.Derecha,datos)
		if Altura(nodo.Derecha)-Altura(nodo.Izquierda)==2{
			if SeleccionMes(ConvertirMes(datos.Mes))>nodo.Datos.NoMes{
				nodo=RotacionSimpleDerecha(nodo)

			}else{
				nodo=RotacionDobleDerecha(nodo)
			}
		}
	}else{
		nodo.Datos.Matriz.InsertarPedido(datos)
	}
	nodo.Altura=Mayor(Altura(nodo.Izquierda),Altura(nodo.Derecha))+1

	return nodo
}