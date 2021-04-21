package GraficoAVL

import (
	"Proyecto1/AVLAnio"
	"Proyecto1/AVLMeses"
	"strconv"
)
var cont int
var contp int
func GrafoAVLAnios(nodo*AVLAnio.NodoAVLAnio)string{
	archie:="digraph G{\n"
	archie+="node[shape=circle];\n"
	archie+="NodoRaiz[label=\""+strconv.Itoa(nodo.Datos.Anio)+"\"];\n"
	cont=1
	archie=RecorrerAVLAnios("NodoRaiz",nodo,archie)
	archie+="\n}"
	return archie
}
func RecorrerAVLAnios(nodoProcedente string,nodo*AVLAnio.NodoAVLAnio,info string)string{
	if nodo.Izquierda!=nil{
		nodoActual:="Nodo"+strconv.Itoa(cont)
		cont+=1
		info+=nodoActual+"[label=\""+strconv.Itoa(nodo.Izquierda.Datos.Anio)+"\"];\n"
		info+=nodoProcedente+"->"+nodoActual+";\n"
		info=RecorrerAVLAnios(nodoActual,nodo.Izquierda,info)
	}
	if nodo.Derecha!=nil{
		nodoActual:="Nodo"+strconv.Itoa(cont)
		cont+=1
		info+=nodoActual+"[label=\""+strconv.Itoa(nodo.Derecha.Datos.Anio)+"\"];\n"
		info+=nodoProcedente+"->"+nodoActual+";\n"
		info=RecorrerAVLAnios(nodoActual,nodo.Derecha,info)
	}
	return info

}
func GrafoAVLMeses(nodo*AVLAnio.NodoAVLAnio)string{
	contp=1
	archie:="digraph G{\n"
	archie+="node[shape=circle];\n"
	archie+="NodoRaiz[label=\""+strconv.Itoa(nodo.Datos.Anio)+"\"];\n"
	archie+="subgraph meses{\n    style = rounded;\n"
	archie+="NodRaiz[label=\""+nodo.Datos.AVLdeMes.Raiz.Datos.Mes+"\"];\n"
	cont=1
	archie=RecorrerAVLMeses("NodRaiz",nodo.Datos.AVLdeMes.Raiz,archie)
	archie+="\n}"
	archie+="NodoRaiz-> NodRaiz [lhead = meses];\n"
	archie=RecorrerAVLAniosMeses(nodo,archie)
	archie+="}\n"
	return archie
}
func RecorrerAVLAniosMeses(nodo *AVLAnio.NodoAVLAnio,info string)string{
 if nodo.Izquierda!=nil{
 	info+="NodoRaiz"+strconv.Itoa(contp)+"[label=\""+strconv.Itoa(nodo.Izquierda.Datos.Anio)+"\"];\n"
	 info+="subgraph meses"+strconv.Itoa(contp)+"{\n    style = rounded;\n"
	 info+="NodRaiz"+strconv.Itoa(contp)+"[label=\""+nodo.Izquierda.Datos.AVLdeMes.Raiz.Datos.Mes+"\"];\n"
	 cont+=1
	 info=RecorrerAVLMeses("NodRaiz"+strconv.Itoa(contp),nodo.Izquierda.Datos.AVLdeMes.Raiz,info)
	 info+="\n}"
	 info+="NodoRaiz"+strconv.Itoa(contp)+"->"+"NodRaiz"+strconv.Itoa(contp)+"[lhead = meses"+strconv.Itoa(contp)+"];\n"
	 contp+=1
	 info=RecorrerAVLAniosMeses(nodo.Izquierda,info)
 }
 if nodo.Derecha!=nil{
	 info+="NodoRaiz"+strconv.Itoa(contp)+"[label=\""+strconv.Itoa(nodo.Derecha.Datos.Anio)+"\"];\n"
	 info+="subgraph meses"+strconv.Itoa(contp)+"{\n    style = rounded;\n"
	 info+="NodRaiz"+strconv.Itoa(contp)+"[label=\""+nodo.Derecha.Datos.AVLdeMes.Raiz.Datos.Mes+"\"];\n"
	 cont+=1
	 info=RecorrerAVLMeses("NodRaiz"+strconv.Itoa(contp),nodo.Derecha.Datos.AVLdeMes.Raiz,info)
	 info+="\n}"
	 info+="NodoRaiz"+strconv.Itoa(contp)+"->"+"NodRaiz"+strconv.Itoa(contp)+"[lhead = meses"+strconv.Itoa(contp)+"];\n"
	 contp+=1
	 info=RecorrerAVLAniosMeses(nodo.Derecha,info)

 }
 return info
}


func RecorrerAVLMeses(nodoProcedente string,nodo*AVLMeses.NodoAVLMes,info string)string{
	if nodo.Izquierda!=nil{
		nodoActual:="Nodo"+strconv.Itoa(cont)
		cont+=1
		info+=nodoActual+"[label=\""+nodo.Izquierda.Datos.Mes+"\"];\n"
		info+=nodoProcedente+"->"+nodoActual+";\n"
		info=RecorrerAVLMeses(nodoActual,nodo.Izquierda,info)
	}
	if nodo.Derecha!=nil{
		nodoActual:="Nodo"+strconv.Itoa(cont)
		cont+=1
		info+=nodoActual+"[label=\""+nodo.Derecha.Datos.Mes+"\"];\n"
		info+=nodoProcedente+"->"+nodoActual+";\n"
		info=RecorrerAVLMeses(nodoActual,nodo.Derecha,info)
	}
	return info

}