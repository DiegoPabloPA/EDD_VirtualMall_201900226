package Tracking

import (
	"Proyecto1/Carrito"
	"Proyecto1/EstructuraAVL"
	"Proyecto1/ListaDoble"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type ArchivoGrafoJson struct {
	Nodos []struct {
		Nombre  string `json:"Nombre"`
		Enlaces []struct {
			Nombre    string `json:"Nombre"`
			Distancia int    `json:"Distancia"`
		} `json:"Enlaces"`
	} `json:"Nodos"`
	Posicioninicialrobot string `json:"PosicionInicialRobot"`
	Entrega              string `json:"Entrega"`
}
type Direccion struct {
	NodoOrigen string
	NodoDestino string
	Tamanio int
}
type NombreImagen struct {
	Imagen string
}

type Nodo struct {
	Siguiente *Nodo
	NombreVertice string
	VerticesAnexos ListaAdyacencia
}
type Nodo2 struct {
	Siguiente *Nodo2
	Nombre string
	Tamanio int
}

type ListaAdyacencia struct {
	Inicio *Nodo2
}


type ListaVertices struct {
	Inicio *Nodo
	PuntoInicio string
	PuntoEntrega string
}

func IniciarListaVertices()*Nodo{
	Inicio:=&Nodo{
		Siguiente:      nil,
		NombreVertice:  "",
		VerticesAnexos: ListaAdyacencia{},
	}
	return Inicio
}

func InciarListaAnexos()*Nodo2{
	Nuevo:=&Nodo2{
		Siguiente: nil,
		Nombre:    "",
		Tamanio:   0,
	}
	return Nuevo
}

func(lista *ListaVertices) ListaGrafo(datos ArchivoGrafoJson){
	var apuntadores []Direccion
	lista.PuntoInicio=datos.Posicioninicialrobot
	lista.PuntoEntrega=datos.Entrega
	aux:=lista.Inicio
	for a:=0;a<len(datos.Nodos);a++{
		if a<len(datos.Nodos)-1{
			aux.Siguiente=IniciarListaVertices()
		}
		aux.NombreVertice=datos.Nodos[a].Nombre
		if len(datos.Nodos[a].Enlaces)>0{
			aux.VerticesAnexos.Inicio=InciarListaAnexos()
		}
		aux2:=aux.VerticesAnexos.Inicio
		for b:=0;b<len(datos.Nodos[a].Enlaces);b++{
			apuntadores=append(apuntadores,Direccion{
				NodoOrigen:  datos.Nodos[a].Enlaces[b].Nombre,
				NodoDestino: datos.Nodos[a].Nombre,
				Tamanio: datos.Nodos[a].Enlaces[b].Distancia,
			})
			if b<len(datos.Nodos[a].Enlaces)-1 {
				aux2.Siguiente = InciarListaAnexos()
			}
			aux2.Nombre=datos.Nodos[a].Enlaces[b].Nombre
			aux2.Tamanio=datos.Nodos[a].Enlaces[b].Distancia
			aux2=aux2.Siguiente
		}
		aux=aux.Siguiente
	}}

func(lista *ListaVertices) AnadirInformacionTrack(datos ArchivoGrafoJson){
	var apuntadores []Direccion
	lista.PuntoInicio=datos.Posicioninicialrobot
	lista.PuntoEntrega=datos.Entrega
	aux:=lista.Inicio
	for a:=0;a<len(datos.Nodos);a++{
		if a<len(datos.Nodos)-1{
			aux.Siguiente=IniciarListaVertices()
		}
		aux.NombreVertice=datos.Nodos[a].Nombre
		if len(datos.Nodos[a].Enlaces)>0{
			aux.VerticesAnexos.Inicio=InciarListaAnexos()
		}
		aux2:=aux.VerticesAnexos.Inicio
		for b:=0;b<len(datos.Nodos[a].Enlaces);b++{
			apuntadores=append(apuntadores,Direccion{
				NodoOrigen:  datos.Nodos[a].Enlaces[b].Nombre,
				NodoDestino: datos.Nodos[a].Nombre,
				Tamanio: datos.Nodos[a].Enlaces[b].Distancia,
			})
			if b<len(datos.Nodos[a].Enlaces)-1 {
				aux2.Siguiente = InciarListaAnexos()
			}
			aux2.Nombre=datos.Nodos[a].Enlaces[b].Nombre
			aux2.Tamanio=datos.Nodos[a].Enlaces[b].Distancia
			aux2=aux2.Siguiente
		}
		aux=aux.Siguiente
	}


	for a:=0;a<len(apuntadores);a++{
		aux=lista.Inicio
		for aux!=nil{
			if aux.NombreVertice==apuntadores[a].NodoOrigen{
				aux2:=aux.VerticesAnexos.Inicio
				bandera:=true
				for aux2!=nil{
					if strings.EqualFold(aux2.Nombre,apuntadores[a].NodoDestino){
						bandera=false
					}
					aux2=aux2.Siguiente
				}
				if bandera{
					aux2=aux.VerticesAnexos.Inicio
					if aux2==nil{
						aux.VerticesAnexos.Inicio=InciarListaAnexos()
						aux2=aux.VerticesAnexos.Inicio
						aux2.Nombre=apuntadores[a].NodoDestino
						aux2.Tamanio=apuntadores[a].Tamanio

					}else{



						for aux2!=nil{
							if aux2.Siguiente==nil{

								aux2.Siguiente=InciarListaAnexos()
								aux2=aux2.Siguiente
								aux2.Nombre=apuntadores[a].NodoDestino
								aux2.Tamanio=apuntadores[a].Tamanio
							}

							aux2=aux2.Siguiente
						}}


				}

			}
			aux=aux.Siguiente
		}

	}


}
func(lista ListaVertices)  Imprimir(){
	aux:=lista.Inicio
	for aux!=nil{

		fmt.Print(aux,"-->")
		aux2:=aux.VerticesAnexos.Inicio
		for aux2!=nil{
			fmt.Print(aux2,"--")
			aux2=aux2.Siguiente
		}
		aux=aux.Siguiente
		fmt.Println()

	}

}
var cont int
type AlmacenGrafo struct {
	Nombre string
	ValorGrafo string
}


func(lista ListaVertices)Grafo(info string)string{
	info+="digraph G{\n"
	info+="node[shape=circle];\n"
	info+="edge[dir=\"both\"];\nrankdir=LR\n"
	aux:=lista.Inicio
	var almacen []AlmacenGrafo
	cont=0
	for aux!=nil{
		info+="Nodo"+strconv.Itoa(cont)+"[label=\""+aux.NombreVertice+"\"];\n"
		almacen=append(almacen,AlmacenGrafo{
			Nombre:     aux.NombreVertice,
			ValorGrafo: "Nodo"+strconv.Itoa(cont),
		})
		aux=aux.Siguiente
		cont+=1
	}
	aux=lista.Inicio
	for aux!=nil{
		aux2:=aux.VerticesAnexos.Inicio
		var a int
		for a=0;a<len(almacen);a++{
			if strings.EqualFold(aux.NombreVertice,almacen[a].Nombre){
				break
			}
		}
		for aux2!=nil{
			var b int
			for b=0;b<len(almacen);b++{
				if strings.EqualFold(aux2.Nombre,almacen[b].Nombre){
					break
				}
			}
			info+=almacen[a].ValorGrafo+"->"+almacen[b].ValorGrafo+"[label=\""+strconv.Itoa(aux2.Tamanio)+"\"];\n"
			aux2=aux2.Siguiente
		}


		aux=aux.Siguiente
	}




	return info
}
type EsqueletoCaminoCorto struct{
	Vertice string
	PesoFinal int
	PesoTemporal int
}
type TemporalIn struct{
	Vertice string
	Valor int
}
type RespuestaAlgoritmo struct {
	ValorRecorrido int
	Recorrido []string
}


func(lista ListaVertices) CaminoCorto(nodoInicio string,nodoFin string)RespuestaAlgoritmo{
	aux:=lista.Inicio
	var VerticesCerrados []string
	var CuadroInformacion []EsqueletoCaminoCorto
	for aux!=nil{
		if strings.EqualFold(nodoInicio,aux.NombreVertice){
			CuadroInformacion=append(CuadroInformacion,EsqueletoCaminoCorto{
				Vertice:      nodoInicio,
				PesoFinal:    0,
				PesoTemporal: 0,
			})
			VerticesCerrados=append(VerticesCerrados,nodoInicio)
		}else{
			CuadroInformacion=append(CuadroInformacion,EsqueletoCaminoCorto{
				Vertice: aux.NombreVertice,
				PesoFinal:    -1,
				PesoTemporal: -1,
			})
		}
		aux=aux.Siguiente
	}
	nodoActual:=nodoInicio

	for !strings.EqualFold(nodoActual,nodoFin){
		aux2:=lista.Inicio
		var aux3 *Nodo2
		for aux2!=nil{
			if strings.EqualFold(aux2.NombreVertice,nodoActual){
				aux3=aux2.VerticesAnexos.Inicio
			}
			aux2=aux2.Siguiente
		}

		var c int
		for c=0;c<len(CuadroInformacion);c++{
			if strings.EqualFold(CuadroInformacion[c].Vertice,nodoActual){
				break
			}
		}
		var temporal []TemporalIn

		for aux3!=nil{
			for b:=0;b<len(CuadroInformacion);b++{
				if strings.EqualFold(CuadroInformacion[b].Vertice,aux3.Nombre)&&CuadroInformacion[b].PesoTemporal==-1&&!SeEncuentraContenido(VerticesCerrados,aux3.Nombre){
					CuadroInformacion[b].PesoTemporal=aux3.Tamanio+CuadroInformacion[c].PesoFinal

				}else if strings.EqualFold(CuadroInformacion[b].Vertice,aux3.Nombre)&&!SeEncuentraContenido(VerticesCerrados,aux3.Nombre)&&CuadroInformacion[b].PesoTemporal>aux3.Tamanio+CuadroInformacion[c].PesoFinal{
					CuadroInformacion[b].PesoTemporal=aux3.Tamanio+CuadroInformacion[c].PesoFinal
				}



			}

			aux3=aux3.Siguiente
		}
		for b:=0;b<len(CuadroInformacion);b++{
			if !SeEncuentraContenido(VerticesCerrados,CuadroInformacion[b].Vertice)&&CuadroInformacion[b].PesoTemporal!=-1{
				temporal=append(temporal,TemporalIn{
					Vertice: CuadroInformacion[b].Vertice,
					Valor:  CuadroInformacion[b].PesoTemporal,
				})
			}
		}


		temporal=Ordenar(temporal)

		for m:=0;m<len(CuadroInformacion);m++{
			if strings.EqualFold(CuadroInformacion[m].Vertice,temporal[0].Vertice){
				CuadroInformacion[m].PesoFinal=CuadroInformacion[m].PesoTemporal
				nodoActual=CuadroInformacion[m].Vertice
				VerticesCerrados=append(VerticesCerrados,nodoActual)
			}
		}




	}
	arreglo:=lista.RecorridoMasCorto(CuadroInformacion,nodoInicio,nodoFin)
	var z int
	for z=0;z<len(CuadroInformacion);z++{
		if strings.EqualFold(CuadroInformacion[z].Vertice,nodoFin){
			break
		}
	}


	regreso:=RespuestaAlgoritmo{
		ValorRecorrido: CuadroInformacion[z].PesoFinal,
		Recorrido:      arreglo,
	}
	return regreso
}
func Ordenar(informacion []TemporalIn)[]TemporalIn{
	var aux TemporalIn
	for a:=0;a<len(informacion);a++{
		for b:=0;b<len(informacion);b++{
			if informacion[a].Valor<informacion[b].Valor{
				aux=informacion[a]
				informacion[a]=informacion[b]
				informacion[b]=aux
			}
		}
	}
	return informacion
}


func SeEncuentraContenido( historial []string ,busqueda string)bool{
	condicion:=false
	for a:=0;a<len(historial);a++{
		if strings.EqualFold(busqueda,historial[a]){
			condicion=true
			break
		}
	}
	return condicion
}
func(lista ListaVertices) RecorridoMasCorto(Datos []EsqueletoCaminoCorto,nodoInicio string,nodoFinal string)(info []string){
	nodoActual:=nodoFinal

	for!strings.EqualFold(nodoInicio,nodoActual){
		var a int
		for a=0;a<len(Datos);a++{
			if strings.EqualFold(nodoActual,Datos[a].Vertice){
				break
			}
		}
		aux:=lista.Inicio
		var aux2 *Nodo2
		for aux!=nil{
			if strings.EqualFold(aux.NombreVertice,nodoActual){
				aux2=aux.VerticesAnexos.Inicio
			}
			aux=aux.Siguiente
		}
		bandera:=true
		for aux2!=nil{
			if bandera{
				var b int
				for b=0;b<len(Datos);b++{
					if strings.EqualFold(Datos[b].Vertice,aux2.Nombre){
						break
					}
				}
				if Datos[a].PesoFinal-aux2.Tamanio==Datos[b].PesoFinal&&Datos[b].PesoFinal!=-1{
					bandera=false
					info=append(info,nodoActual)
					nodoActual=Datos[b].Vertice

				}
			}

			aux2=aux2.Siguiente
		}


	}
	info=append(info,nodoInicio)
	var aux3 []string
	for c:=len(info)-1;c>-1;c--{
		aux3=append(aux3,info[c])
	}
	info=aux3

	return info
}

func GenerarGrafoSegunPedido(nombreCliente string, carrodeCompra[]Carrito.JsonPedidoIndv,Arreglo[]ListaDoble.ListaDE)(regreso PedidoConNombreClienteGrafo){
	aux:=Arreglo
	var datos []ProductosConNodoInicio
	for a:=0;a<len(carrodeCompra);a++{

		for b:=0;b<len(aux);b++{
			if aux[b].Tamanio>0{
				aux2:=aux[b].Inicio
				for aux2!=nil{
					if strings.EqualFold(carrodeCompra[a].Nombre,aux2.Datos.Nombre)&&strings.EqualFold(carrodeCompra[a].Departamento,aux2.Datos.Columna)&&aux2.Datos.Calificacion==carrodeCompra[a].Calificacion{
						temporal:=EncontrarNodoporProducto(aux2.Datos.Inventario.Raiz,carrodeCompra[a].Codigo)
						datos=append(datos,ProductosConNodoInicio{
							NodoInicio: temporal,
							Codigo:     carrodeCompra[a].Codigo,
						})

					}


					aux2=aux2.Siguiente
				}
			}
		}
	}
	regreso.NombreCliente=nombreCliente
	regreso.Descrip=datos
	return regreso
}
type PedidoConNombreClienteGrafo struct {
	NombreCliente string
	Descrip []ProductosConNodoInicio
}


type ProductosConNodoInicio struct {
	NodoInicio string
	Codigo int
}
type AlmacendeGrafos struct{
	Cliente string
	Grafo string
	NombreGrafo string
}
type EstructuraListaGrafo struct{
	codigo int
	valor RespuestaAlgoritmo
}
func ListadeGrafos (pedido PedidoConNombreClienteGrafo,Estructura ListaVertices)(retorno AlmacenGrafo){
	var grafo string

	retorno.Nombre=pedido.NombreCliente
	nodoInicio:=Estructura.PuntoInicio
	bandera:=true
	var registro []int
	for bandera{

		var memoria []EstructuraListaGrafo
		for a:=0;a<len(pedido.Descrip);a++{
			if !BusquedaRegistro(registro,pedido.Descrip[a].Codigo){

			m:=Estructura.CaminoCorto(nodoInicio,pedido.Descrip[a].NodoInicio)

			memoria=append(memoria,EstructuraListaGrafo{
				codigo: pedido.Descrip[a].Codigo,
				valor:  m,
			})
			}
		}
		memoria=OrdenarMemoria(memoria)

		registro=append(registro,memoria[0].codigo)
		nodoInicio=memoria[0].valor.Recorrido[len(memoria[0].valor.Recorrido)-1]
		for t:=0;t<len(memoria[0].valor.Recorrido);t++{
			grafo+=memoria[0].valor.Recorrido[t]+" "
		}
		grafo+="|||"

		if len(registro)==len(pedido.Descrip){
			bandera=false
		}

	}
	n:=Estructura.CaminoCorto(nodoInicio,Estructura.PuntoEntrega)
	for t:=0;t<len(n.Recorrido);t++{
		grafo+=n.Recorrido[t]+" "
	}
	grafo+=" "

	retorno.ValorGrafo=grafo



	return retorno
}
func BusquedaRegistro(registro []int,codigo int)bool{
condicion:=false
for a:=0;a<len(registro);a++{
		if registro[a]==codigo{
			condicion=true

	}
}


return condicion
}


func OrdenarMemoria(datos []EstructuraListaGrafo)[]EstructuraListaGrafo{
	var aux EstructuraListaGrafo
	for a:=0;a<len(datos);a++{
		for b:=0;b<len(datos);b++{
			if datos[a].valor.ValorRecorrido<datos[b].valor.ValorRecorrido{
				aux=datos[a]
				datos[a]=datos[b]
				datos[b]=aux
			}
		}
	}
	return datos
}


func EncontrarNodoporProducto(Raiz *EstructuraAVL.NodoArbol,codigo int)string{
	var informacion string
	for Raiz!=nil{
		if Raiz.Datos.Codigo>codigo{
			Raiz=Raiz.Izquierda
		}else if Raiz.Datos.Codigo<codigo{
			Raiz=Raiz.Derecha
		}else{
			informacion=Raiz.Datos.Almacenamiento
			Raiz=Raiz.Derecha
		}
	}
	return informacion
}

type DatosGraphviz struct{
	NombreCliente string
	NombreArchivo string
}

func GeneracionGraphviz(informacion AlmacenGrafo,grafoincompleto string,tamanio string)DatosGraphviz{
	var nuevo DatosGraphviz
	nombre:="Tracking"+tamanio+".png"
	nuevo.NombreCliente=informacion.Nombre

	grafoincompleto+="n[shape=none label=<<table><tr>\n<td>Recorrido</td>\n</tr>\n"
	recoleccion:=strings.Split(informacion.ValorGrafo,"|||")
	for a:=0;a<len(recoleccion);a++{
		grafoincompleto+="<tr>\n<td>\n"+recoleccion[a]+"\n</td>\n</tr>\n"
	}

	grafoincompleto+="</table>>];\n}"

	conv:=[]byte(grafoincompleto)
	generacion:=ioutil.WriteFile("Tracking"+tamanio+".dot",conv,0644)
	if generacion!=nil{
		log.Fatal(generacion)
	}
	graph, _ := exec.LookPath("dot")
	direccion,_:=os.Getwd()
	consola, _ := exec.Command(graph, "-Tpng",direccion+"/"+"Tracking"+tamanio+".dot").Output()
	ioutil.WriteFile(nombre, consola, 0777)
	nuevo.NombreArchivo=nombre
return nuevo
}
