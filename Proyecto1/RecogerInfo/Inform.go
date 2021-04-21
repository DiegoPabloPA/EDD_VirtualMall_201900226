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
	"Proyecto1/Tracking"
	Usuarios2 "Proyecto1/Usuarios"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/fernet/fernet-go"
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
var Usuarios Usuarios2.ArbolB
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
var NombreCliente string
var EstructuraFunciones Tracking.ListaVertices
var EstructuraGraphviz Tracking.ListaVertices
var GrafoAUtilizar string
var DatosGraph []Tracking.DatosGraphviz
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
	router.HandleFunc("/RegistrarUsuario",CrearUsuario).Methods("POST")
	router.HandleFunc("/NombreCliente",RecibirNombreCliente).Methods("POST")
	router.HandleFunc("/CargaNodosGrafo",CargadeNodos).Methods("POST")
	router.HandleFunc("/BuscarUsuario",BuscarUsuario).Methods("POST")
	router.HandleFunc("/UsuarioMasivo",InsertarUsuariosMasivos).Methods("POST")
	router.HandleFunc("/DescargaTrack/{ide}",DescargaTrack)
	router.HandleFunc("/BuscarTracks",ObtenerTrackings).Methods("GET")
	router.HandleFunc("/GenerarUsuariosNormal",GenerarArbolBNormal).Methods("GET")
	router.HandleFunc("/Encriptacion/{ide}",Encriptacion)
	router.HandleFunc("/ArbolNormalB",DescargaArbolBNormal)
	router.HandleFunc("/Cifrado2B",DescargaArbolBCifradoCompleto)
	router.HandleFunc("/Cifrado3B",DescargaArbolBCifradoSensible)




	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))


}
func GenerarArbolBNormal(w http.ResponseWriter, req *http.Request){
	if Usuarios.Raiz!=nil{
		info:=Usuarios2.GenerarGrafoArbolB(Usuarios.Raiz)
		conv:=[]byte(info)
		generacion:=ioutil.WriteFile("ReporteUsuarios.dot",conv,0644)
		if generacion!=nil{
			log.Fatal(generacion)
		}
		graph, _ := exec.LookPath("dot")
		direccion,_:=os.Getwd()
		consola, _ := exec.Command(graph, "-Tpng",direccion+"/ReporteUsuarios.dot").Output()
		ioutil.WriteFile("ReporteUsuarios.png", consola, 0777)

	}
}


func InsertarUsuariosMasivos(w http.ResponseWriter, req *http.Request){
	respuesta, _ := ioutil.ReadAll(req.Body)
	var nuevo Usuarios2.CargaMasivaUsuarios
	if err := json.Unmarshal(respuesta, &nuevo); err != nil {
		panic(err)

	}
	if Usuarios.Raiz==nil{
		Usuarios.Raiz=Usuarios.IniciarArbolB(Usuarios.Raiz)
		nuevo2:=Usuarios2.DatosUsuario{
			DPI:      1234567890101,
			Correo:   "auxiliar@edd.com",
			Password: "1234",
			Nombre:   "EDD2021",
			Cuenta: "Admin",
		}
		Usuarios.Raiz=Usuarios.Raiz.Insertar(nuevo2,Usuarios.Raiz)
		Usuarios.Raiz=Usuarios2.InsercionMasivaUsuarios(Usuarios.Raiz,nuevo)


	}else{
		Usuarios.Raiz=Usuarios2.InsercionMasivaUsuarios(Usuarios.Raiz,nuevo)

	}
}


func BuscarUsuario(w http.ResponseWriter, req *http.Request){
	respuesta, _ := ioutil.ReadAll(req.Body)
	var nuevo Usuarios2.SolicitudLogueo
	if err := json.Unmarshal(respuesta, &nuevo); err != nil {
		panic(err)

	}
	if Usuarios.Raiz==nil{
		Usuarios.Raiz=Usuarios.IniciarArbolB(Usuarios.Raiz)
		nuevo2:=Usuarios2.DatosUsuario{
			DPI:      1234567890101,
			Correo:   "auxiliar@edd.com",
			Password: "1234",
			Nombre:   "EDD2021",
			Cuenta: "Admin",
		}


		Usuarios.Raiz=Usuarios.Raiz.Insertar(nuevo2,Usuarios.Raiz)
		info:=Usuarios.Raiz.BuscarUsuario(nuevo)

		if info.DPI!=0{
			AJson,_:=json.Marshal(info)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(AJson)
		}else{
			info.DPI=-1
			AJson,_:=json.Marshal(info)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(AJson)
		}

	}else{
		info:=Usuarios.Raiz.BuscarUsuario(nuevo)
		if info.DPI!=0{
			AJson,_:=json.Marshal(info)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(AJson)
		}else{
			info.DPI=-1
			AJson,_:=json.Marshal(info)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(AJson)
		}

	}






}


func CrearUsuario(w http.ResponseWriter, req *http.Request){
	respuesta, _ := ioutil.ReadAll(req.Body)
	var nuevo Usuarios2.UsuarioNormal
	if err := json.Unmarshal(respuesta, &nuevo); err != nil {
		panic(err)

	}
	nuevo3:=Usuarios2.DatosUsuario{
		DPI:      nuevo.Dpi,
		Correo:   nuevo.Correo,
		Password: nuevo.Password,
		Nombre:   nuevo.Nombre,
		Cuenta: nuevo.Cuenta,
	}


	if Usuarios.Raiz==nil{
	Usuarios.Raiz=Usuarios.IniciarArbolB(Usuarios.Raiz)
	nuevo2:=Usuarios2.DatosUsuario{
		DPI:      1234567890101,
		Correo:   "auxiliar@edd.com",
		Password: "1234",
		Nombre:   "EDD2021",
		Cuenta: "Admin",
	}


	Usuarios.Raiz=Usuarios.Raiz.Insertar(nuevo2,Usuarios.Raiz)
	Usuarios.Raiz=Usuarios.Raiz.Insertar(nuevo3,Usuarios.Raiz)
}else{
		Usuarios.Raiz=Usuarios.Raiz.Insertar(nuevo3,Usuarios.Raiz)
	}


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
func Encriptacion(w http.ResponseWriter, req *http.Request){
	if Usuarios.Raiz!=nil{
		Dato:=mux.Vars(req)
		No:=Dato["ide"]
		generador:=sha256.New()
		generador.Write([]byte(No))
		preliminar:=generador.Sum(nil)
		preliminar2:=hex.EncodeToString(preliminar)
		llave:=fernet.MustDecodeKeys(preliminar2)
		grafo1:=Usuarios2.GenerarGrafoArbolBCifradoCompleto(Usuarios.Raiz,llave)
		grafo2:=Usuarios2.GenerarGrafoArbolBCifradoSensible(Usuarios.Raiz,llave)
		conv:=[]byte(grafo1)
		generacion:=ioutil.WriteFile("CifradoCompleto.dot",conv,0644)
		if generacion!=nil{
			log.Fatal(generacion)
		}
		graph, _ := exec.LookPath("dot")
		direccion,_:=os.Getwd()
		consola, _ := exec.Command(graph, "-Tpdf",direccion+"/CifradoCompleto.dot").Output()
		ioutil.WriteFile("CifradoCompleto.pdf", consola, 0777)

		conv2:=[]byte(grafo2)
		generacion2:=ioutil.WriteFile("CifradoSensible.dot",conv2,0644)
		if generacion2!=nil{
			log.Fatal(generacion2)
		}
		graph2, _ := exec.LookPath("dot")
		direccion2,_:=os.Getwd()
		consola2, _ := exec.Command(graph2, "-Tpdf",direccion2+"/CifradoSensible.dot").Output()
		ioutil.WriteFile("CifradoSensible.pdf", consola2, 0777)


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
				Cliente:      NombreCliente,
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
	Aux:=CarroActual
	CarroActual=nil
	Arreglo=CambioInventario.EfectuarCompra(Arreglo,CarroActual)
	Nuevo:=Tracking.GenerarGrafoSegunPedido(NombreCliente,Aux,Arreglo)
	d:=Tracking.ListadeGrafos(Nuevo,EstructuraFunciones)
	DatosGraph=append(DatosGraph,Tracking.GeneracionGraphviz(d,GrafoAUtilizar,strconv.Itoa(len(DatosGraph))))

}
func CargadeNodos(w http.ResponseWriter, req *http.Request){
	respuesta, _ := ioutil.ReadAll(req.Body)
	var datos Tracking.ArchivoGrafoJson
	if err := json.Unmarshal(respuesta, &datos); err != nil {
		panic(err)

	}
	EstructuraFunciones.Inicio=Tracking.IniciarListaVertices()
	EstructuraGraphviz.Inicio=Tracking.IniciarListaVertices()
	EstructuraFunciones.AnadirInformacionTrack(datos)
	EstructuraGraphviz.ListaGrafo(datos)
	GrafoAUtilizar=EstructuraGraphviz.Grafo(GrafoAUtilizar)
}
func ObtenerTrackings(w http.ResponseWriter, req *http.Request){
	if len(DatosGraph)>0{
		AJson,_:=json.Marshal(DatosGraph)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(AJson)
	}
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
func RecibirNombreCliente(w http.ResponseWriter, req *http.Request){
	var NomCliente Carrito.NombreCliente
	respuesta,_:=ioutil.ReadAll(req.Body)
	if err := json.Unmarshal(respuesta, &NomCliente); err != nil {
		panic(err)

	}
	NombreCliente=NomCliente.Nombre
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
func DescargaTrack(w http.ResponseWriter, req *http.Request){
	Dato:=mux.Vars(req)
	No:=Dato["ide"]
	w.Header().Set("Content-Disposition", "attachment; filename="+No)
	w.Header().Set("Content-Type", "application/image")
	http.ServeFile(w, req, No)

}
func DescargaArbolBNormal(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Disposition", "attachment; filename="+"ReporteUsuarios.png")
	w.Header().Set("Content-Type", "application/image")
	http.ServeFile(w, req, "ReporteUsuarios.png")

}
func DescargaArbolBCifradoCompleto(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Disposition", "attachment; filename="+"CifradoCompleto.pdf")
	w.Header().Set("Content-Type", "application/pdf")
	http.ServeFile(w, req, "CifradoCompleto.pdf")
}
func DescargaArbolBCifradoSensible(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Disposition", "attachment; filename="+"CifradoSensible.pdf")
	w.Header().Set("Content-Type", "application/pdf")
	http.ServeFile(w, req, "CifradoSensible.pdf")
}
