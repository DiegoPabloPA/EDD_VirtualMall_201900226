package Inform

import (
	ArregloEstatico "Proyecto1/Arreglo"
	BusquedaNumero "Proyecto1/FuncionBusquedaNumero"
	Busqueda3Parametros "Proyecto1/Funciones"
	"Proyecto1/Graficar"
	"Proyecto1/GuardarJson"
	vector "Proyecto1/ListaDVec"
	"Proyecto1/ListaDoble"
	"Proyecto1/OrdenAlfabetico"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var Informacion vector.InfoVector
var Arreglo []ListaDoble.ListaDE
var fila[]string
var columna[]string
func Path(){
	router:=mux.NewRouter()
	router.HandleFunc("/",PaginaInicio)
	router.HandleFunc("/cargartienda",AsignarInformacion).Methods("POST")
	router.HandleFunc("/getArreglo",GenerarGrafico).Methods("GET")
	router.HandleFunc("/TiendaEspecifica",Busqueda3).Methods("POST")
	router.HandleFunc("/id/{ide}",BusquedaPosicion).Methods("GET")
	router.HandleFunc("/guardar",GuardarInformacion)
	router.HandleFunc("/Eliminar",Eliminar).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", router))
}
func Eliminar(w http.ResponseWriter, req *http.Request){
	if len(Arreglo)==0{
		fmt.Fprintln(w,"Error aun no se han cargado datos")
	}else{
		respuesta, _ := ioutil.ReadAll(req.Body)
		var busqueda ListaDoble.EstructuraEliminacion
		if err := json.Unmarshal(respuesta, &busqueda); err != nil {
			panic(err)

		}
		condicion:=ListaDoble.EliminarTienda(busqueda,Arreglo)
		if condicion{
			fmt.Fprintln(w,"Eliminacion exitosa!")
		}else{
			fmt.Fprintln(w,"No se encontraron tiendas con esas caracteristicas")
		}
	}
}


func Busqueda3(w http.ResponseWriter, req *http.Request){
	if len(Arreglo)==0{
		fmt.Fprintln(w,"Error aun no se han cargado datos")
	}else {
		respuesta, _ := ioutil.ReadAll(req.Body)
		var busqueda Busqueda3Parametros.EstructuraBusqueda
		if err := json.Unmarshal(respuesta, &busqueda); err != nil {
			panic(err)

		}
		Respuesta := Busqueda3Parametros.Buscar3Param(busqueda, Arreglo)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(Respuesta)
	}
}
func BusquedaPosicion(w http.ResponseWriter, req *http.Request){
	if len(Arreglo)!=0{
Dato:=mux.Vars(req)
No:=Dato["ide"]

a,_:=strconv.Atoi(No)
a=a-1
if a<len(Arreglo){
Encontradas:=BusquedaNumero.Buscar(a,Arreglo)
if len(Encontradas)==0{
	fmt.Fprintln(w,"No se encontraron tiendas en dicho indice")
}else {
	AJson,_:=json.Marshal(Encontradas)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(AJson)

}}else{
	fmt.Fprintln(w,"No se encontraron tiendas en dicho indice")
}
}else{
		fmt.Fprintln(w,"Error aun no se han cargado datos")
	}
}



func PaginaInicio(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w,"TODOS LOS DERECHOS RESERVADOS")
}





func AsignarInformacion(w http.ResponseWriter, req *http.Request){
	respuesta,_:=ioutil.ReadAll(req.Body)

	if err := json.Unmarshal(respuesta, &Informacion); err != nil {
		panic(err)

	}


	fila,columna=ArregloEstatico.CrearArreglo(Informacion)
	tamanio:=len(fila)*len(columna)*5
	Arreglo=make([]ListaDoble.ListaDE,tamanio)
	ArregloEstatico.ColumnMajor(Arreglo,fila,columna,Informacion)
	OrdenAlfabetico.Ordenar(Arreglo)
	if len(Arreglo)==0{
		fmt.Fprintln(w,"ERROR CON EL ARCHIVO DE ENTRADA")
	}else{
		fmt.Fprintln(w,"Archivo linealizado")
	}

}
func GenerarGrafico(w http.ResponseWriter, req *http.Request){
if len(Arreglo)==0{
	fmt.Fprintln(w,"ERROR, DEBE INSERTAR UN ARCHIVO ANTES DE GENERAR EL GRÁFICO")
}else{
	Graficar.GenerarArregloGrafico(Arreglo)
	fmt.Fprintln(w,"Gráfico generado exitosamente!")
}
}

func GuardarInformacion(w http.ResponseWriter, req *http.Request){
	if len(Arreglo)==0{
		fmt.Fprintln(w,"ERROR, Aun no se han cargado datos")
	}else{
	estructura:=GuardarJson.Guardar(Arreglo,fila,columna)
	AJson,_:=json.Marshal(estructura)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(AJson)
	}

}
