package GraficoMatrizDispersa

import (
	"Proyecto1/MatrizDispersa"
	"strconv"
)
var contador1 int
var contador2 int
var grupo int
var memoria [] *MatrizDispersa.Nodo
var nombrenodo []string
var Columnas[]string
func GraficarMatriz(nodo *MatrizDispersa.Nodo)string{
	memoria=nil
	nombrenodo=nil
	Columnas=nil
	archie:="digraph g{\n"
	archie+="node [shape=box];\n"
	archie+="NodoRaiz[label=\"Raiz\",group=1];\n"
	contador1=1
	contador2=1
	grupo=2
	archie=RecorrerFilas("NodoRaiz",nodo,archie)
	archie=RecorrerColumnas("NodoRaiz",nodo,archie)
	archie+="{rank = same; NodoRaiz;"
	for b:=0;b<len(Columnas);b++{
		archie+=Columnas[b]+"; "
	}
	archie+="}\n"
	contador1=1
	grupo=2
	auxColumna:=nodo
	for auxColumna.Derecha!=nil{
		auxColumna=auxColumna.Derecha
		aux2:=auxColumna
		if aux2.Abajo!=nil{
			for aux2.Abajo!=nil{
				aux2=aux2.Abajo
				nodoActual:="Compra"+strconv.Itoa(contador1)
				archie+=nodoActual+"[group ="+strconv.Itoa(grupo)+"];\n"

				contador1+=1
				memoria=append(memoria,aux2)
				nombrenodo=append(nombrenodo,nodoActual)

				if aux2.Arriba!=nil{
					archie+=nombrenodo[ObtenerPosicion(aux2.Arriba,memoria)]+"->"+nodoActual+"[dir=both];\n"
				}
				if aux2.Izquierda!=nil{
					archie+=nombrenodo[ObtenerPosicion(aux2.Izquierda,memoria)]+"->"+nodoActual+"[dir=both];\n"
					archie+="{rank = same; "+nombrenodo[ObtenerPosicion(aux2.Izquierda,memoria)]+"; "+nodoActual+"; }\n"
				}
			}
		}
	grupo+=1
	}


	archie+="}\n"
	return archie
}
func ObtenerPosicion(direccion *MatrizDispersa.Nodo,arreglo []*MatrizDispersa.Nodo)int{
	numero:=-1
	for a:=0;a<len(arreglo);a++{
		if direccion==arreglo[a]{
			numero=a
		}
	}
	return numero
}


func RecorrerColumnas(NodoPrecedente string, nodo*MatrizDispersa.Nodo,informacion string)string{
	if nodo.Derecha!=nil{
	NodoActual:="C"+strconv.Itoa(contador1)
	contador1+=1
	informacion+=NodoActual+"[label=\""+strconv.Itoa(nodo.Derecha.Informacion.Dia)+"\",group ="+strconv.Itoa(grupo)+"];\n"
	grupo+=1
	memoria=append(memoria, nodo.Derecha)
	nombrenodo=append(nombrenodo,NodoActual)
	Columnas=append(Columnas,NodoActual)

	informacion+=NodoPrecedente+"->"+NodoActual+"[dir=both];\n"
	informacion=RecorrerColumnas(NodoActual,nodo.Derecha,informacion)
	}
	return informacion
}
func RecorrerFilas(NodoPrecedente string, nodo*MatrizDispersa.Nodo,informacion string)string{
	if nodo.Abajo!=nil{
		NodoActual:="F"+strconv.Itoa(contador2)
		contador2+=1
		informacion+=NodoActual+"[label=\""+nodo.Abajo.Informacion.Departamento+"\", group = 1];\n"
		memoria=append(memoria, nodo.Abajo)
		nombrenodo=append(nombrenodo,NodoActual)
		informacion+=NodoPrecedente+"->"+NodoActual+"[dir=both];\n"
		informacion=RecorrerFilas(NodoActual,nodo.Abajo,informacion)
	}
	return informacion
}
