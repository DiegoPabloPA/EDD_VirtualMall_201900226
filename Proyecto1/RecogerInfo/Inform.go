package Inform

import (
	ArregloEstatico "Proyecto1/Arreglo"
	"Proyecto1/EstructuraAVL"
	BusquedaNumero "Proyecto1/FuncionBusquedaNumero"
	Busqueda3Parametros "Proyecto1/Funciones"
	"Proyecto1/Graficar"
	"Proyecto1/GuardarJson"
	Inventarios2 "Proyecto1/Inventarios"
	"Proyecto1/JsonparaFront"
	"github.com/gorilla/handlers"
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
var Inventarios Inventarios2.CargadeInventario
var InventarioInd EstructuraAVL.InventarioIndividual
var Arreglo []ListaDoble.ListaDE
var fila[]string
var columna[]string
var ubicacionTienda EstructuraAVL.EstructuraBusqueda
func Path(){
	router:=mux.NewRouter()
	router.HandleFunc("/",PaginaInicio)
	router.HandleFunc("/cargartienda",AsignarInformacion).Methods("POST")
	router.HandleFunc("/getArreglo",GenerarGrafico).Methods("GET")
	router.HandleFunc("/TiendaEspecifica",Busqueda3).Methods("POST")
	router.HandleFunc("/EnvioInventario",EnviarJsonInventario).Methods("POST")
	router.HandleFunc("/id/{ide}",BusquedaPosicion).Methods("GET")
	router.HandleFunc("/guardar",GuardarInformacion)
	router.HandleFunc("/Eliminar",Eliminar).Methods("DELETE")
	router.HandleFunc("/inventario",CargaMasivaInventario).Methods("POST")
	router.HandleFunc("/JsonFrontEnd",JsonFront).Methods("GET")
	router.HandleFunc("/ubicaTienda",UbicarTienda).Methods("POST")
	router.HandleFunc("/subirInventarioIndividual",CargaIndividualInventario).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))


}
func UbicarTienda(w http.ResponseWriter, req *http.Request){
	respuesta, _ := ioutil.ReadAll(req.Body)

	if err := json.Unmarshal(respuesta, &ubicacionTienda); err != nil {
		panic(err)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println(ubicacionTienda)
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
func EnviarJsonInventario(w http.ResponseWriter, req *http.Request){
	if len(Arreglo)==0{
		fmt.Fprintln(w,"Error aun no se han cargado datos")
	}else {
		respuesta, _ := ioutil.ReadAll(req.Body)
		var busqueda Busqueda3Parametros.EstructuraBusqueda
		if err := json.Unmarshal(respuesta, &busqueda); err != nil {
			panic(err)

		}
		Respuesta := Busqueda3Parametros.BuscarPedidos(busqueda, Arreglo)

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
func JsonFront(w http.ResponseWriter, req *http.Request){
	if len(Arreglo)==0{
		fmt.Fprintln(w,"ERROR, Aun no se han cargado datos")
	}else{
		estructura:=JsonparaFront.JsonFront(Arreglo)
		AJson,_:=json.Marshal(estructura)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(AJson)
	}
}

func CargaMasivaInventario(w http.ResponseWriter, req *http.Request){
	if len(Arreglo)==0{
		fmt.Fprintln(w,"ERROR, Aun no se han cargado datos")
	}else{
		respuesta,_:=ioutil.ReadAll(req.Body)

		if err := json.Unmarshal(respuesta, &Inventarios); err != nil {
			panic(err)

		}

		Arreglo=Inventarios2.CargaInventariosMasivo(Arreglo,Inventarios)

		for a:=0;a<len(Arreglo);a++{
			if Arreglo[a].Tamanio!=0{
				aux4:=Arreglo[a].Inicio
				for aux4!=nil{
					Inventarios2.Imprimir(aux4.Datos.Inventario.Raiz)
					aux4=aux4.Siguiente
				}
			}
		}
	}
}
func CargaIndividualInventario(w http.ResponseWriter, req *http.Request){
	if len(Arreglo)==0{
		fmt.Fprintln(w,"ERROR, Aun no se han cargado datos")
	}else{
		respuesta,_:=ioutil.ReadAll(req.Body)

		if err := json.Unmarshal(respuesta, &InventarioInd); err != nil {
			panic(err)

		}

		Arreglo=Inventarios2.CargaInventariosIndividual(Arreglo,InventarioInd,ubicacionTienda)

		for a:=0;a<len(Arreglo);a++{
			if Arreglo[a].Tamanio!=0{
				aux4:=Arreglo[a].Inicio
				for aux4!=nil{
					Inventarios2.Imprimir(aux4.Datos.Inventario.Raiz)
					aux4=aux4.Siguiente
				}
			}
		}
	}
}
