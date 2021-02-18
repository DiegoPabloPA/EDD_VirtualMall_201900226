package Graficar

import (
	"Proyecto1/ListaDoble"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GenerarArregloGrafico(Arr []ListaDoble.ListaDE){
archie:="digraph G{\nrankdir=TB;\n"

archie+="node[shape=record,width=2.5,height=.1];"

for a:=0;a<len(Arr);a++{
	archie+="\nf"+strconv.Itoa(a)+"[ranksep=0.3 label=\""+strconv.Itoa(a)+"\"];"
}

	for a:=0;a<len(Arr);a++{
		aux:=Arr[a].Inicio
		if aux!=nil{
			for aux!=nil{
			if aux==Arr[a].Inicio{
				archie+=strings.Replace(aux.Datos.Nombre, " ", "", -1)+"[label=\""+aux.Datos.Nombre+"\"];\n"
				archie+="f"+strconv.Itoa(a)+"->"+strings.Replace(aux.Datos.Nombre, " ", "", -1)+" [dir=\"both\"]"+";\n"
				aux=aux.Siguiente
			}else{
				archie+=strings.Replace(aux.Datos.Nombre, " ", "", -1)+"[label=\""+aux.Datos.Nombre+"\"];\n"
				archie+=strings.Replace(aux.Anterior.Datos.Nombre, " ", "", -1)+"->"+strings.Replace(aux.Datos.Nombre, " ", "", -1)+";\n"
				archie+=strings.Replace(aux.Datos.Nombre, " ", "", -1)+"->"+strings.Replace(aux.Anterior.Datos.Nombre, " ", "", -1)+";\n"
				aux=aux.Siguiente
			}
		}}
	}


archie+="}"


conv:=[]byte(archie)
generacion:=ioutil.WriteFile("diagrama.dot",conv,0644)
if generacion!=nil{
	log.Fatal(generacion)
}
	graph, _ := exec.LookPath("dot")
	direccion,_:=os.Getwd()
	consola, _ := exec.Command(graph, "-Tpng",direccion+"/diagrama.dot").Output()
	ioutil.WriteFile("Grafico_Arreglo.png", consola, 0777)


}