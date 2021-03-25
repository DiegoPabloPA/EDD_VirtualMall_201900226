package Inform

import (
	"Proyecto1/AVLAnio"
	"Proyecto1/AVLMeses"
	ArregloEstatico "Proyecto1/Arreglo"
	"Proyecto1/CambioInventario"
	"Proyecto1/CargaPedidos"
	"Proyecto1/Carrito"
	"Proyecto1/EnviarCarrito"
	"Proyecto1/EstructuraAVL"
	BusquedaNumero "Proyecto1/FuncionBusquedaNumero"
	Busqueda3Parametros "Proyecto1/Funciones"
	"Proyecto1/GenerarGraficoReportes"
	"Proyecto1/Graficar"
	"Proyecto1/GraficoAVL"
	"Proyecto1/GraficoMatrizDispersa"
	"Proyecto1/GuardarJson"
	Inventarios2 "Proyecto1/Inventarios"
	"Proyecto1/JsonparaFront"
	vector "Proyecto1/ListaDVec"
	"Proyecto1/ListaDoble"
	"Proyecto1/MatrizDispersa"
	"Proyecto1/MenuAnios"
	"Proyecto1/OrdenAlfabetico"
	"Proyecto1/ReportesDiaCategoria"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var Informacion vector.InfoVector
var Inventarios Inventarios2.CargadeInventario
var InventarioInd EstructuraAVL.InventarioIndividual
var CarritoConfirm Carrito.JsonPedidoIndv
var CarroActual []Carrito.JsonPedidoIndv
var Arreglo []ListaDoble.ListaDE
var fila[]string
var Pedidos AVLAnio.AVLAnio
var columna[]string
var ubicacionTienda EstructuraAVL.EstructuraBusqueda
func Path(){
	router:=mux.NewRouter()
	router.HandleFunc("/",PaginaInicio)
	router.HandleFunc("/cargartienda",AsignarInformacion).Methods("POST")
	router.HandleFunc("/getArreglo",GenerarGrafico).Methods("GET")
	router.HandleFunc("/TiendaEspecifica",Busqueda3).Methods("POST")
	router.HandleFunc("/AgregarCarrito",RecibirCarrito).Methods("POST")
	router.HandleFunc("/EnvioInventario",EnviarJsonInventario).Methods("POST")
	router.HandleFunc("/id/{ide}",BusquedaPosicion).Methods("GET")
	router.HandleFunc("/guardar",GuardarInformacion)
	router.HandleFunc("/Eliminar",Eliminar).Methods("DELETE")
	router.HandleFunc("/inventario",CargaMasivaInventario).Methods("POST")
	router.HandleFunc("/JsonFrontEnd",JsonFront).Methods("GET")
	router.HandleFunc("/ubicaTienda",UbicarTienda).Methods("POST")
	router.HandleFunc("/subirInventarioIndividual",CargaIndividualInventario).Methods("POST")
	router.HandleFunc("/InfoCompra",Compra).Methods("GET")
	router.HandleFunc("/EliminarCarrito",EliminarCarrito).Methods("POST")
	router.HandleFunc("/EjecutarCompra",RealizarPedido).Methods("POST")
	router.HandleFunc("/GrafAVLAnios",GraficoAVLAnios).Methods("GET")
	router.HandleFunc("/GrafAVLMeses",GraficoAVLMeses).Methods("GET")
	router.HandleFunc("/MenuAnios",JsonMenuReporte).Methods("GET")
	router.HandleFunc("/MenuDiasCategoria",JsonCategoriaDia).Methods("POST")
	router.HandleFunc("/InformeDia",JsonReporteDia).Methods("POST")
	router.HandleFunc("/PedidosMasivo",CargaMasivaPedidos).Methods("POST")
	router.HandleFunc("/DownAVLAnio",DescargaAVLAnio)
	router.HandleFunc("/DownAVLMes",DescargaAVLMeses)
	router.HandleFunc("/DownMatriz",DescargaMatriz)
	router.HandleFunc("/DownReporteDia",DescargaInformeDia)
	router.HandleFunc("/MostrarMatriz",MostrarMatriz)




	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))


}
func CargaMasivaPedidos(w http.ResponseWriter, req *http.Request){
	if len(Arreglo)>0{
		respuesta, _ := ioutil.ReadAll(req.Body)
		var busqueda CargaPedidos.JsonPedidosMasivo
		if err := json.Unmarshal(respuesta, &busqueda); err != nil {
			panic(err)

		}
		Pedidos.Raiz=CargaPedidos.CargaMasiva(busqueda,Arreglo,Pedidos.Raiz)

	}
}

func JsonReporteDia(w http.ResponseWriter, req *http.Request){
	if Pedidos.Raiz!=nil {
		respuesta, _ := ioutil.ReadAll(req.Body)
		var busqueda ReportesDiaCategoria.JsonReporteDiaSolicitado
		if err := json.Unmarshal(respuesta, &busqueda); err != nil {
			panic(err)

		}
		archivo:=ReportesDiaCategoria.GenerarReporte(Pedidos.Raiz,busqueda)
		info:=GenerarGraficoReportes.GenerarReporteDia(archivo)

		conv:=[]byte(info)
		generacion:=ioutil.WriteFile("ReporteDia.dot",conv,0644)
		if generacion!=nil{
			log.Fatal(generacion)
		}
		graph, _ := exec.LookPath("dot")
		direccion,_:=os.Getwd()
		consola, _ := exec.Command(graph, "-Tpdf",direccion+"/ReporteDia.dot").Output()
		ioutil.WriteFile("ReporteDia.pdf", consola, 0777)




		AJson,_:=json.Marshal(archivo)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(AJson)
	}
}


func JsonCategoriaDia(w http.ResponseWriter, req *http.Request){
	if Pedidos.Raiz!=nil{
	respuesta, _ := ioutil.ReadAll(req.Body)
	var busqueda MenuAnios.RespInfo
	var archivo []MenuAnios.DiasConCateogira
	if err := json.Unmarshal(respuesta, &busqueda); err != nil {
		panic(err)

	}
		aux:=Pedidos.Raiz
		for aux!=nil{
			if aux.Datos.Anio>busqueda.Anio{
				aux=aux.Izquierda
			}else if aux.Datos.Anio<busqueda.Anio{
				aux=aux.Derecha
			}else{
				aux2:=aux.Datos.AVLdeMes.Raiz
				for aux2!=nil{
					if aux2.Datos.NoMes>AVLMeses.SeleccionMes(busqueda.Mes){
						aux2=aux2.Izquierda
					}else if aux2.Datos.NoMes<AVLMeses.SeleccionMes(busqueda.Mes){
						aux2=aux2.Derecha
					}else{

						archivo=MenuAnios.DevolverDiasCategoria(aux2.Datos.Matriz.Nodoinicio)
						informe:=GraficoMatrizDispersa.GraficarMatriz(aux2.Datos.Matriz.Nodoinicio)
						conv:=[]byte(informe)
						generacion:=ioutil.WriteFile("diagramaMatriz.dot",conv,0644)
						if generacion!=nil{
							log.Fatal(generacion)
						}
						graph, _ := exec.LookPath("dot")
						direccion,_:=os.Getwd()
						consola, _ := exec.Command(graph, "-Tpng",direccion+"/diagramaMatriz.dot").Output()
						ioutil.WriteFile("Matriz.png", consola, 0777)

						aux2=aux2.Derecha
					}
				}


				aux=aux.Derecha
			}
		}

		AJson,_:=json.Marshal(archivo)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(AJson)
	}

}



func JsonMenuReporte(w http.ResponseWriter, req *http.Request){
	if Pedidos.Raiz!=nil{
	estructura:=MenuAnios.RecepAnios(Pedidos.Raiz)
	AJson,_:=json.Marshal(estructura)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(AJson)
	}
}



func RealizarPedido(w http.ResponseWriter, req *http.Request){
	if len(CarroActual)>0{
		var reserva []string
		var identificador string
		t:=time.Now()
		for a:=0;a<len(CarroActual);a++{
			nuevo:=MatrizDispersa.Pedido{
				Siguiente:    nil,
				Dia:          t.Day(),
				Mes:          t.Month().String(),
				Anio:         t.Year(),
				Cliente:      "",
				Direccion:    "Guatemala",
				NombreTienda: "",
				Departamento: CarroActual[a].Departamento,
				Calificacion: CarroActual[a].Calificacion,
				Descripcion:  nil,
			}

			if len(reserva)==0{
			identificador=CarroActual[a].Departamento
			reserva=append(reserva,CarroActual[a].Departamento)
			}else{
				identificador=CarroActual[a].Departamento
				for c:=0;c<len(reserva);c++{
					if strings.EqualFold(identificador,reserva[c]){
						identificador="rep123"
					}
				}
				if strings.EqualFold(identificador,CarroActual[a].Departamento){
					reserva=append(reserva,identificador)
				}
			}
			for b:=0;b<len(CarroActual);b++{
				if strings.EqualFold(identificador,CarroActual[b].Departamento){
					nuevo.Descripcion=append(nuevo.Descripcion,MatrizDispersa.DescripcionPedido{
						Codigo:   CarroActual[b].Codigo,
						Cantidad: CarroActual[b].Cantidad,
					})
				}
			}
			Pedidos.Raiz=AVLAnio.Insertar(Pedidos.Raiz,nuevo)



		}
	}
	Arreglo=CambioInventario.EfectuarCompra(Arreglo,CarroActual)
	CarroActual=nil
}

func Compra(w http.ResponseWriter, req *http.Request){
	if len(Arreglo)==0{
		fmt.Fprintln(w,"ERROR, Aun no se han cargado datos")
	}else{
		estructura:=EnviarCarrito.ConvertirCarritoDesc(CarroActual,Arreglo)
		AJson,_:=json.Marshal(estructura)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(AJson)

	}
}
func EliminarCarrito(w http.ResponseWriter, req *http.Request){
	if len(Arreglo)==0{
		fmt.Fprintln(w,"Error aun no se han cargado datos")
	}else{
		respuesta, _ := ioutil.ReadAll(req.Body)
		var busqueda EnviarCarrito.EliminacionProd
		if err := json.Unmarshal(respuesta, &busqueda); err != nil {
			panic(err)

		}
		CarroActual=EnviarCarrito.EliminardelCarrito(busqueda,CarroActual)
		fmt.Println(CarroActual)
	}
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

func RecibirCarrito(w http.ResponseWriter, req *http.Request){
	if len(Arreglo)==0{
		fmt.Fprintln(w,"ERROR, Aun no se han cargado datos")
	}else{
		respuesta,_:=ioutil.ReadAll(req.Body)

		if err := json.Unmarshal(respuesta, &CarritoConfirm); err != nil {
			panic(err)

		}


		if len(CarroActual)==0{
		condicion:=Carrito.ValidarPedido(CarritoConfirm,Arreglo)
			if condicion{
				mensaje:=Carrito.Respuesta{Res: "Si"}
				AJson,_:=json.Marshal(mensaje)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(AJson)
				CarroActual=append(CarroActual,CarritoConfirm)

			}else{
				mensaje:=Carrito.Respuesta{Res: "No"}
				AJson,_:=json.Marshal(mensaje)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(AJson)
			}

		}else{
			condicion2:=true
			for a:=0;a<len(CarroActual);a++{
				if strings.EqualFold(CarritoConfirm.Nombre,CarroActual[a].Nombre)&&strings.EqualFold(CarritoConfirm.Departamento,CarroActual[a].Departamento)&&CarritoConfirm.Calificacion==CarroActual[a].Calificacion&&CarritoConfirm.Codigo==CarroActual[a].Codigo{
					CarritoConfirm.Cantidad+=CarroActual[a].Cantidad
					condicion2=false
				}
			}
			condicion:=Carrito.ValidarPedido(CarritoConfirm,Arreglo)
			for a:=0;a<len(CarroActual);a++{
				if strings.EqualFold(CarritoConfirm.Nombre,CarroActual[a].Nombre)&&strings.EqualFold(CarritoConfirm.Departamento,CarroActual[a].Departamento)&&CarritoConfirm.Calificacion==CarroActual[a].Calificacion&&CarritoConfirm.Codigo==CarroActual[a].Codigo{
					CarritoConfirm.Cantidad-=CarroActual[a].Cantidad

				}
			}


			if condicion{
				mensaje:=Carrito.Respuesta{Res: "Si"}
				AJson,_:=json.Marshal(mensaje)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(AJson)
				if condicion2{
				CarroActual=append(CarroActual,CarritoConfirm)
				}else{
					for a:=0;a<len(CarroActual);a++{
						if strings.EqualFold(CarritoConfirm.Nombre,CarroActual[a].Nombre)&&strings.EqualFold(CarritoConfirm.Departamento,CarroActual[a].Departamento)&&CarritoConfirm.Calificacion==CarroActual[a].Calificacion&&CarritoConfirm.Codigo==CarroActual[a].Codigo{
							CarroActual[a].Cantidad+=CarritoConfirm.Cantidad

						}
					}
				}

			}else{
				mensaje:=Carrito.Respuesta{Res: "No"}
				AJson,_:=json.Marshal(mensaje)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(AJson)
			}



		}



	}
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
func GraficoAVLAnios(w http.ResponseWriter, req *http.Request){
	if Pedidos.Raiz!=nil{
	archivo:=GraficoAVL.GrafoAVLAnios(Pedidos.Raiz)
	conv:=[]byte(archivo)
	generacion:=ioutil.WriteFile("diagramaAVLAnios.dot",conv,0644)
	if generacion!=nil{
		log.Fatal(generacion)
	}
	graph, _ := exec.LookPath("dot")
	direccion,_:=os.Getwd()
	consola, _ := exec.Command(graph, "-Tpng",direccion+"/diagramaAVLAnios.dot").Output()
	ioutil.WriteFile("AVLAnios.png", consola, 0777)

	}

}

func GraficoAVLMeses(w http.ResponseWriter, req *http.Request){
	if Pedidos.Raiz!=nil{
		archivo:=GraficoAVL.GrafoAVLMeses(Pedidos.Raiz)
		conv:=[]byte(archivo)
		generacion:=ioutil.WriteFile("diagramaAVLMeses.dot",conv,0644)
		if generacion!=nil{
			log.Fatal(generacion)
		}
		graph, _ := exec.LookPath("dot")
		direccion,_:=os.Getwd()
		consola, _ := exec.Command(graph, "-Tpng",direccion+"/diagramaAVLMeses.dot").Output()
		ioutil.WriteFile("AVLMeses.png", consola, 0777)

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
func MostrarMatriz(w http.ResponseWriter, req *http.Request){
	http.ServeFile(w, req, "Matriz.png")
}

func DescargaAVLAnio(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Disposition", "attachment; filename="+"AVLAnios.png")
	w.Header().Set("Content-Type", "application/image")
	http.ServeFile(w, req, "AVLAnios.png")
}
func DescargaAVLMeses(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Disposition", "attachment; filename="+"AVLMeses.png")
	w.Header().Set("Content-Type", "application/image")
	http.ServeFile(w, req, "AVLMeses.png")
}
func DescargaMatriz(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Disposition", "attachment; filename="+"Matriz.png")
	w.Header().Set("Content-Type", "application/image")
	http.ServeFile(w, req, "Matriz.png")
}
func DescargaInformeDia(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Disposition", "attachment; filename="+"Informe.pdf")
	w.Header().Set("Content-Type", "application/pdf")
	http.ServeFile(w, req, "ReporteDia.pdf")
}