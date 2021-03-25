package AVLAnio

import (
	"Proyecto1/AVLMeses"
	"Proyecto1/MatrizDispersa"
	"fmt"
)

type NodoAVLAnio struct {
	Izquierda, Derecha *NodoAVLAnio
	Altura int
	Datos DatosMesAvl
}

type DatosMesAvl struct {
	Anio int
	AVLdeMes *AVLMeses.AVLMeses
}

type AVLAnio struct {
	Raiz* NodoAVLAnio
}
func Mayor(numa int,numb int)int{
	if numa>numb{
		return numa
	}
	return numb

}
func Altura(evaluar *NodoAVLAnio)int{
	if evaluar==nil{
		return -1
	}
	return evaluar.Altura
}
func RotacionSimpleIzquierda(evaluar*NodoAVLAnio)*NodoAVLAnio{
	aux1:=evaluar.Izquierda
	evaluar.Izquierda=aux1.Derecha
	aux1.Derecha=evaluar
	evaluar.Altura=Mayor(Altura(evaluar.Izquierda),Altura(evaluar.Derecha))+1
	aux1.Altura=Mayor(Altura(evaluar.Izquierda),evaluar.Altura)+1
	evaluar=aux1
	return evaluar
}
func RotacionSimpleDerecha(evaluar *NodoAVLAnio)*NodoAVLAnio{
	aux1:=evaluar.Derecha
	evaluar.Derecha=aux1.Izquierda
	aux1.Izquierda=evaluar
	evaluar.Altura=Mayor(Altura(evaluar.Izquierda),Altura(evaluar.Derecha))+1
	aux1.Altura=Mayor(Altura(evaluar.Derecha),evaluar.Altura)+1
	evaluar=aux1
	return evaluar
}
func RotacionDobleIzquierda(evaluar *NodoAVLAnio)*NodoAVLAnio{
	evaluar.Izquierda=RotacionSimpleDerecha(evaluar.Izquierda)
	evaluar=RotacionSimpleIzquierda(evaluar)
	return evaluar
}
func RotacionDobleDerecha(evaluar *NodoAVLAnio)*NodoAVLAnio{
	evaluar.Derecha=RotacionSimpleIzquierda(evaluar.Derecha)
	evaluar=RotacionSimpleDerecha(evaluar)
	return evaluar
}
func Imprimir(nodo *NodoAVLAnio){

	if nodo!=nil{
		Imprimir(nodo.Izquierda)
		fmt.Print(nodo.Datos.Anio," ")
		Imprimir(nodo.Derecha)
	}
}

func  Insertar(nodo *NodoAVLAnio,datos MatrizDispersa.Pedido)*NodoAVLAnio{
	nuevo:=&NodoAVLAnio{
		Izquierda: nil,
		Derecha:   nil,
		Altura:    0,
		Datos: DatosMesAvl{
			Anio:     datos.Anio,
			AVLdeMes: &AVLMeses.AVLMeses{Raiz: nil},
		},

	}
	if nodo==nil{
		nuevo.Datos.AVLdeMes.Raiz=AVLMeses.Insertar(nuevo.Datos.AVLdeMes.Raiz,datos)
		return nuevo
	}
	if  nodo.Datos.Anio>datos.Anio{
		nodo.Izquierda=Insertar(nodo.Izquierda,datos)
		if Altura(nodo.Izquierda)-Altura(nodo.Derecha)==2{
			if datos.Anio<nodo.Izquierda.Datos.Anio{
				nodo=RotacionSimpleIzquierda(nodo)
			}else{
				nodo=RotacionDobleIzquierda(nodo)
			}
		}
	}else if nodo.Datos.Anio<datos.Anio{
		nodo.Derecha=Insertar(nodo.Derecha,datos)
		if Altura(nodo.Derecha)-Altura(nodo.Izquierda)==2{
			if datos.Anio>nodo.Derecha.Datos.Anio{
				nodo=RotacionSimpleDerecha(nodo)

			}else{
				nodo=RotacionDobleDerecha(nodo)
			}
		}
	}else{
		nodo.Datos.AVLdeMes.Raiz=AVLMeses.Insertar(nodo.Datos.AVLdeMes.Raiz,datos)

	}
	nodo.Altura=Mayor(Altura(nodo.Izquierda),Altura(nodo.Derecha))+1

	return nodo
}