package Inform

import (
	vector "Proyecto1/ListaDVec"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

var Informacion vector.InfoVector

func Path(){
	router:=mux.NewRouter()
	router.HandleFunc("/",PaginaInicio)
	router.HandleFunc("/cargartienda",AsignarInformacion).Methods("POST")
	router.HandleFunc("/mostrar",Mostrar)
	log.Fatal(http.ListenAndServe(":3000", router))
}

func PaginaInicio(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w,"TODOS LOS DERECHOS RESERVADOS")
}

func Mostrar(w http.ResponseWriter, req *http.Request){

	if len(Informacion.Indice1)==0{
		fmt.Fprintf(w,"Error aun no se han cargado datos")
	}
}



func AsignarInformacion(w http.ResponseWriter, req *http.Request){
	respuesta,_:=ioutil.ReadAll(req.Body)

	if err := json.Unmarshal(respuesta, &Informacion); err != nil {
		panic(err)

	}}
