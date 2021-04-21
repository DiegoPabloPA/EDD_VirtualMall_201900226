package Usuarios

import (
	"github.com/fernet/fernet-go"
	"strconv"
	"strings"
)

type UsuarioNormal struct {
	Dpi      int    `json:"Dpi"`
	Nombre   string `json:"Nombre"`
	Correo   string `json:"Correo"`
	Password string `json:"Password"`
	Cuenta   string `json:"Cuenta"`
}



type DatosUsuario struct {
	DPI int
	Correo string
	Password string
	Nombre string
	Cuenta string
}
type CargaMasivaUsuarios struct {
	Usuarios []struct {
		Dpi      int  `json:"Dpi"`
		Nombre   string `json:"Nombre"`
		Correo   string `json:"Correo"`
		Password string `json:"Password"`
		Cuenta   string `json:"Cuenta"`
	} `json:"Usuarios"`
}
func InsercionMasivaUsuarios(pag *Pagina,datos CargaMasivaUsuarios)*Pagina{
	for a:=0;a<len(datos.Usuarios);a++{
		temporal:=DatosUsuario{
			DPI:      datos.Usuarios[a].Dpi,
			Correo:   datos.Usuarios[a].Correo,
			Password: datos.Usuarios[a].Password,
			Nombre:   datos.Usuarios[a].Nombre,
			Cuenta:   datos.Usuarios[a].Cuenta,
		}
		pag=pag.Insertar(temporal,pag)

	}

	return pag
}


type Pagina struct{
	CondicionHoja bool
	CantidadElementos int
	Datos [5]DatosUsuario
	PaginasHijas [5]*Pagina
	PaginaPadre *Pagina
}
type ArbolB struct {
	Raiz *Pagina
}

type SolicitudLogueo struct {
	Dpi      int  `json:"Dpi"`
	Password string `json:"Password"`
}



func (Raiz Pagina)BuscarUsuario(dato SolicitudLogueo)(info DatosUsuario){
	bandera:=true
	for a:=0;a<Raiz.CantidadElementos;a++{
		if Raiz.Datos[a].DPI==dato.Dpi&&strings.EqualFold(Raiz.Datos[a].Password,dato.Password){
			bandera=false
			info=DatosUsuario{
				DPI:      Raiz.Datos[a].DPI,
				Correo:   Raiz.Datos[a].Correo,
				Password: Raiz.Datos[a].Password,
				Nombre:   Raiz.Datos[a].Nombre,
				Cuenta:   Raiz.Datos[a].Cuenta,
			}
			break
		}

	}
	if bandera&&!Raiz.CondicionHoja{
		bandera2:=true
		for b:=0;b<Raiz.CantidadElementos;b++{
			if Raiz.Datos[b].DPI>dato.Dpi{
				info=Raiz.PaginasHijas[b].BuscarUsuario(dato)
				bandera2=false
				break
			}
		}
		if bandera2{
			info=Raiz.PaginasHijas[Raiz.CantidadElementos].BuscarUsuario(dato)
		}

	}
	return info
}


func (ArbolB) IniciarArbolB(Arbol *Pagina)(*Pagina){
	nuevo:=&Pagina{
		CondicionHoja:     true,
		CantidadElementos: 0,
		Datos:             [5]DatosUsuario{},
		PaginasHijas:      [5]*Pagina{},
		PaginaPadre:       nil,
	}
	Arbol=nuevo
	return Arbol
}
func IniciarPag(Pag *Pagina)(*Pagina){
	nuevo:=&Pagina{
		CondicionHoja:     true,
		CantidadElementos: 0,
		Datos:             [5]DatosUsuario{},
		PaginasHijas:      [5]*Pagina{},
		PaginaPadre:       nil,
	}
	Pag=nuevo
	return Pag
}
func IniciarPagConPadre(Pag *Pagina,padre *Pagina)*Pagina{
	nuevo:=&Pagina{
		CondicionHoja:     true,
		CantidadElementos: 0,
		Datos:             [5]DatosUsuario{},
		PaginasHijas:      [5]*Pagina{},
		PaginaPadre:       padre,
	}
	Pag=nuevo
	return Pag
}

func Ordenar(pag *Pagina)(*Pagina){
	var a int
	for a=pag.CantidadElementos/2;a>0;a=a/2{
		for b:=a;b<pag.CantidadElementos;b++{
			almacen:=pag.Datos[b]
			c:=b
			for c=b;c>=a&&pag.Datos[c-a].DPI>almacen.DPI;c=c-a{
				pag.Datos[c]=pag.Datos[c-a]
			}
			pag.Datos[c]=almacen
		}
	}
	return pag
}





func (*Pagina)InsertarUsuario(Datos DatosUsuario,pag *Pagina)(*Pagina){



	if pag.CondicionHoja {

		pag.Datos[pag.CantidadElementos] = Datos
		pag.CantidadElementos += 1

		if pag.CantidadElementos > 1 {
			pag = Ordenar(pag)
		}

	} else {

		bandera := false
		for a := 0; a < pag.CantidadElementos-1; a++ {

			if Datos.DPI < pag.Datos[a].DPI {
				pag.PaginasHijas[a].InsertarUsuario(Datos, pag.PaginasHijas[a])
				bandera = true
				break
			}

		}

		if !bandera {

			pag.PaginasHijas[pag.CantidadElementos].InsertarUsuario(Datos, pag.PaginasHijas[pag.CantidadElementos])
		}

	}

	if pag.CantidadElementos == 5 {

		if pag.PaginaPadre == nil {

			nuevo := pag

			pag = IniciarPag(pag)

			pag.InsertarUsuario(nuevo.Datos[2], pag)

			pag.PaginasHijas[0] = IniciarPagConPadre(pag.PaginasHijas[0], pag)
			pag.PaginasHijas[1] = IniciarPagConPadre(pag.PaginasHijas[1], pag)

			for a := 0; a < 2; a++ {

				pag.PaginasHijas[0] = pag.PaginasHijas[0].InsertarUsuario(nuevo.Datos[a], pag.PaginasHijas[0])
			}
			for a := 3; a < 5; a++ {
				pag.PaginasHijas[1] = pag.PaginasHijas[1].InsertarUsuario(nuevo.Datos[a], pag.PaginasHijas[1])
			}
			pag.CondicionHoja = false

		} else {
			dato2 := pag.Datos[2]

			pag.PaginaPadre.Datos[pag.PaginaPadre.CantidadElementos] = dato2
			pag.PaginaPadre.CantidadElementos++

			if pag.PaginaPadre.CantidadElementos > 1 {
				pag.PaginaPadre = Ordenar(pag.PaginaPadre)
			}

			var a int
			for a = 0; a < pag.PaginaPadre.CantidadElementos; a++ {

				if pag.PaginaPadre.Datos[a].DPI == dato2.DPI {

					break
				}
			}
			for c := pag.PaginaPadre.CantidadElementos; c > a+1; c-- {

				pag.PaginaPadre.PaginasHijas[c] = pag.PaginaPadre.PaginasHijas[c-1]
			}

			pag.PaginaPadre.PaginasHijas[a+1] = IniciarPagConPadre(pag.PaginaPadre.PaginasHijas[a+1], pag.PaginaPadre)

			for m := 5/2 + 1; m < 5; m++ {
				pag.PaginaPadre.PaginasHijas[a+1] = pag.PaginaPadre.PaginasHijas[a+1].InsertarUsuario(pag.Datos[m], pag.PaginaPadre.PaginasHijas[a+1])

			}

			aux := pag
			pag.PaginaPadre.PaginasHijas[a] = IniciarPagConPadre(pag.PaginaPadre.PaginasHijas[a], pag.PaginaPadre)
			for n := 0; n < 2; n++ {
				pag.PaginaPadre.PaginasHijas[a] = pag.PaginaPadre.PaginasHijas[a].InsertarUsuario(aux.Datos[n], pag.PaginaPadre.PaginasHijas[a])

			}

		}

	}





	return pag
}
func (*Pagina)Insertar(Datos DatosUsuario,pag *Pagina)(*Pagina){
	if pag.ComprobarRamas(Datos.DPI)==5&&ComprobarCantElementos(pag,Datos)+1==5{
		pag=pag.InsertarConRompimientoIntermedio(Datos,pag)
		pag=pag.RompimientoIntermedio(pag)
	}else{

		pag=pag.InsertarUsuario(Datos,pag)

	}
	return pag
}
func (*Pagina)InsertarConRompimientoIntermedio(Dato DatosUsuario,pag *Pagina)*Pagina{
	pag=pag.InsertarNodoTemporal(Dato,pag)


	return pag
}
func (*Pagina)RompimientoIntermedio(pag *Pagina)*Pagina{
	if pag.CantidadElementos == 5 {

		if pag.PaginaPadre == nil {

			nuevo := pag

			pag = IniciarPag(pag)

			pag.InsertarUsuario(nuevo.Datos[2], pag)

			pag.PaginasHijas[0] = IniciarPagConPadre(pag.PaginasHijas[0], pag)
			pag.PaginasHijas[1] = IniciarPagConPadre(pag.PaginasHijas[1], pag)

			for a := 0; a < 2; a++ {

				pag.PaginasHijas[0] = pag.PaginasHijas[0].InsertarUsuario(nuevo.Datos[a], pag.PaginasHijas[0])

			}
			for a := 3; a < 5; a++ {
				pag.PaginasHijas[1] = pag.PaginasHijas[1].InsertarUsuario(nuevo.Datos[a], pag.PaginasHijas[1])


			}


			pag.PaginasHijas[0].CondicionHoja=false
			pag.PaginasHijas[1].CondicionHoja=false
			pag.CondicionHoja = false
			for a := 0; a < 3; a++ {

				pag.PaginasHijas[0].PaginasHijas[a]=IniciarPagConPadre(pag.PaginasHijas[0].PaginasHijas[a],pag.PaginasHijas[0])

			}
			for a := 0; a < 3; a++ {

				pag.PaginasHijas[1].PaginasHijas[a]=IniciarPagConPadre(pag.PaginasHijas[0].PaginasHijas[a],pag.PaginasHijas[1])

			}

			for a := 0; a < 3; a++ {
				for b:=0;b<nuevo.PaginasHijas[a].CantidadElementos;b++ {

					pag.PaginasHijas[0].PaginasHijas[a] = pag.PaginasHijas[0].PaginasHijas[a].InsertarUsuario(nuevo.PaginasHijas[a].Datos[b], pag.PaginasHijas[0].PaginasHijas[a])
				}
			}

			for a := 3; a < 5; a++ {
				if nuevo.PaginasHijas[a].CantidadElementos!=5 {
					for b:=0;b<nuevo.PaginasHijas[a].CantidadElementos;b++ {
						pag.PaginasHijas[1].PaginasHijas[0] = pag.PaginasHijas[1].PaginasHijas[0].InsertarUsuario(nuevo.PaginasHijas[a].Datos[b], pag.PaginasHijas[1].PaginasHijas[0])
					}
				}else{

					for c:=0;c<2;c++{
						pag.PaginasHijas[1].PaginasHijas[1]=pag.PaginasHijas[1].PaginasHijas[1].InsertarUsuario(nuevo.PaginasHijas[a].Datos[c],pag.PaginasHijas[1].PaginasHijas[1])
					}
					for c:=3;c<5;c++{
						pag.PaginasHijas[1].PaginasHijas[2]=pag.PaginasHijas[1].PaginasHijas[2].InsertarUsuario(nuevo.PaginasHijas[a].Datos[c],pag.PaginasHijas[1].PaginasHijas[2])
					}

				}

			}



		}
	}

	return pag
}


func (*Pagina)InsertarNodoTemporal(Dato DatosUsuario,pag*Pagina)*Pagina{
	if !pag.CondicionHoja{
		var a int
		bandera:=true
		for a=0;a<5;a++{
			if Dato.DPI<pag.Datos[a].DPI{
				bandera=false
				break

			}
		}
		if bandera{
			pag.PaginasHijas[a-1]=pag.PaginasHijas[a-1].InsertarNodoTemporal(Dato,pag.PaginasHijas[a-1])
			pag.PaginasHijas[a-1]=Ordenar(pag.PaginasHijas[a-1])
			pag.PaginasHijas[a-1].PaginaPadre.Datos[pag.PaginasHijas[a-1].PaginaPadre.CantidadElementos]=pag.PaginasHijas[a-1].Datos[2]
			pag.PaginasHijas[a-1].PaginaPadre.CantidadElementos++
			pag.PaginasHijas[a-1].CantidadElementos++
		}else{
			pag.PaginasHijas[a]=pag.PaginasHijas[a].InsertarNodoTemporal(Dato,pag.PaginasHijas[a])
			pag.PaginasHijas[a]=Ordenar(pag.PaginasHijas[a])
			pag.PaginasHijas[a].PaginaPadre.Datos[pag.PaginasHijas[a].PaginaPadre.CantidadElementos]=pag.PaginasHijas[a].Datos[2]
			pag.PaginasHijas[a].PaginaPadre.CantidadElementos++
			pag.PaginasHijas[a].CantidadElementos++
		}

	}else{
		for b:=0;b<5;b++{
			if pag.Datos[b].DPI==0{
				pag.Datos[b]=Dato
				break
			}
		}
	}


	return pag
}


func ComprobarCantElementos(pag *Pagina,dato DatosUsuario)(tamanio int){
	if !pag.CondicionHoja{

		var a int
		bandera:=false
		for a=0;a<5;a++{
			if dato.DPI<pag.Datos[a].DPI{
				bandera=true
				break
			}
		}

		if bandera{
			tamanio=ComprobarCantElementos(pag.PaginasHijas[a],dato)
		}else{

			tamanio=ComprobarCantElementos(pag.PaginasHijas[4],dato)
		}


	}else{
		tamanio=pag.CantidadElementos
	}



	return tamanio
}




func(pag Pagina)ComprobarRamas(Dato int ) (cantidad int){
	if !pag.CondicionHoja{
		if !pag.PaginasHijas[0].CondicionHoja{

			var b int
			var cont int
			for b=0;b<5;b++{
				if pag.PaginasHijas[b]!=nil{
					cont++
				}
			}
			var m int
			for c:=0;c<cont;c++{
				for v:=0;v<pag.PaginasHijas[c].CantidadElementos;v++{
					if pag.PaginasHijas[c].ComprobarRamas(Dato)==5&&Dato<pag.PaginasHijas[c].Datos[v].DPI{
						m=c
					}


				}
			}

			cantidad=pag.PaginasHijas[m].ComprobarRamas(Dato)
		}else{
			for a:=0;a<5;a++{
				if pag.PaginasHijas[a]!=nil{
					cantidad++
				}
			}
		}
	}

	return cantidad
}
var contGrafo int

func GenerarGrafoArbolB(pag *Pagina)(info string){
	contGrafo=0
	nombre:="Nodo"+strconv.Itoa(contGrafo)
	contGrafo++
	info+="digraph G{\n"
	info+="node[shape=none];\n"
	info+=nombre+"[label=<<table><tr><td>Nombre</td>\n"

	for a:=0;a<pag.CantidadElementos;a++{
		info+="<td>"+pag.Datos[a].Nombre+"</td>\n"
	}
	info+="</tr><tr><td>DPI</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{
		info+="<td>"+strconv.Itoa(pag.Datos[a].DPI)+"</td>\n"
	}
	info+="</tr><tr><td>Correo</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{
		info+="<td>"+pag.Datos[a].Correo+"</td>\n"
	}
	info+="</tr><tr><td>Password</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{
		info+="<td>"+pag.Datos[a].Password+"</td>\n"
	}
	info+="</tr><tr><td>Cuenta</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{
		info+="<td>"+pag.Datos[a].Cuenta+"</td>\n"
	}
	info+="</tr></table>>];\n"

	if !pag.CondicionHoja{
		var contarHijos int
		for a:=0;a<5;a++{
			if pag.PaginasHijas[a]!=nil{
				contarHijos++
			}
		}
		for a:=0;a<contarHijos;a++{
			info=RecorrerHijos(nombre,pag.PaginasHijas[a],info)
		}
	}


	info+="\n}"

	return info
}
func RecorrerHijos(NodoPadre string,pag *Pagina,info string)string{
	nombre:="Nodo"+strconv.Itoa(contGrafo)
	contGrafo++
	info+=nombre+"[label=<<table><tr><td>Nombre</td>\n"

	for a:=0;a<pag.CantidadElementos;a++{
		info+="<td>"+pag.Datos[a].Nombre+"</td>\n"
	}
	info+="</tr><tr><td>DPI</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{
		info+="<td>"+strconv.Itoa(pag.Datos[a].DPI)+"</td>\n"
	}
	info+="</tr><tr><td>Correo</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{
		info+="<td>"+pag.Datos[a].Correo+"</td>\n"
	}
	info+="</tr><tr><td>Password</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{
		info+="<td>"+pag.Datos[a].Password+"</td>\n"
	}
	info+="</tr><tr><td>Cuenta</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{
		info+="<td>"+pag.Datos[a].Cuenta+"</td>\n"
	}
	info+="</tr></table>>];\n"
	info+=NodoPadre+"->"+nombre+"\n"

	if !pag.CondicionHoja{
		var contarHijos int
		for a:=0;a<5;a++{
			if pag.PaginasHijas[a]!=nil{
				contarHijos++
			}
		}
		for a:=0;a<contarHijos;a++{
			info=RecorrerHijos(nombre,pag.PaginasHijas[a],info)
		}
	}


	return info
}

func GenerarGrafoArbolBCifradoCompleto(pag *Pagina,llave []*fernet.Key)(info string){
	contGrafo=0
	nombre:="Nodo"+strconv.Itoa(contGrafo)
	contGrafo++
	info+="digraph G{\n"
	info+="node[shape=none];\n"
	info+=nombre+"[label=<<table><tr><td>Nombre</td>\n"

	for a:=0;a<pag.CantidadElementos;a++{
		tok,_:= fernet.EncryptAndSign([]byte(pag.Datos[a].Nombre), llave[0])
		info+="<td>"+string(tok)+"</td>\n"

	}
	info+="</tr><tr><td>DPI</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{
		tok,_:= fernet.EncryptAndSign([]byte(strconv.Itoa(pag.Datos[a].DPI)), llave[0])
		info+="<td>"+string(tok)+"</td>\n"
	}
	info+="</tr><tr><td>Correo</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{
		tok,_:= fernet.EncryptAndSign([]byte(pag.Datos[a].Correo), llave[0])
		info+="<td>"+string(tok)+"</td>\n"
	}
	info+="</tr><tr><td>Password</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{
		tok,_:= fernet.EncryptAndSign([]byte(pag.Datos[a].Password), llave[0])
		info+="<td>"+string(tok)+"</td>\n"
	}
	info+="</tr><tr><td>Cuenta</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{
		tok,_:= fernet.EncryptAndSign([]byte(pag.Datos[a].Cuenta), llave[0])
		info+="<td>"+string(tok)+"</td>\n"
	}
	info+="</tr></table>>];\n"

	if !pag.CondicionHoja{
		var contarHijos int
		for a:=0;a<5;a++{
			if pag.PaginasHijas[a]!=nil{
				contarHijos++
			}
		}
		for a:=0;a<contarHijos;a++{
			info=RecorrerHijosCifradoCompleto(nombre,pag.PaginasHijas[a],info,llave)
		}
	}


	info+="\n}"

	return info
}
func RecorrerHijosCifradoCompleto(NodoPadre string,pag *Pagina,info string,llave []*fernet.Key)string{
	nombre:="Nodo"+strconv.Itoa(contGrafo)
	contGrafo++
	info+=nombre+"[label=<<table><tr><td>Nombre</td>\n"

	for a:=0;a<pag.CantidadElementos;a++{
		tok,_:= fernet.EncryptAndSign([]byte(pag.Datos[a].Nombre), llave[0])
		info+="<td>"+string(tok)+"</td>\n"
	}
	info+="</tr><tr><td>DPI</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{
		tok,_:= fernet.EncryptAndSign([]byte(strconv.Itoa(pag.Datos[a].DPI)), llave[0])
		info+="<td>"+string(tok)+"</td>\n"
	}
	info+="</tr><tr><td>Correo</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{
		tok,_:= fernet.EncryptAndSign([]byte(pag.Datos[a].Correo), llave[0])
		info+="<td>"+string(tok)+"</td>\n"
	}
	info+="</tr><tr><td>Password</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{
		tok,_:= fernet.EncryptAndSign([]byte(pag.Datos[a].Password), llave[0])
		info+="<td>"+string(tok)+"</td>\n"
	}
	info+="</tr><tr><td>Cuenta</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{
		tok,_:= fernet.EncryptAndSign([]byte(pag.Datos[a].Cuenta), llave[0])
		info+="<td>"+string(tok)+"</td>\n"
	}
	info+="</tr></table>>];\n"
	info+=NodoPadre+"->"+nombre+"\n"

	if !pag.CondicionHoja{
		var contarHijos int
		for a:=0;a<5;a++{
			if pag.PaginasHijas[a]!=nil{
				contarHijos++
			}
		}
		for a:=0;a<contarHijos;a++{
			info=RecorrerHijosCifradoCompleto(nombre,pag.PaginasHijas[a],info,llave)
		}
	}


	return info
}






func GenerarGrafoArbolBCifradoSensible(pag *Pagina,llave []*fernet.Key)(info string){
	contGrafo=0
	nombre:="Nodo"+strconv.Itoa(contGrafo)
	contGrafo++
	info+="digraph G{\n"
	info+="node[shape=none];\n"
	info+=nombre+"[label=<<table><tr><td>Nombre</td>\n"

	for a:=0;a<pag.CantidadElementos;a++{

		info+="<td>"+pag.Datos[a].Nombre+"</td>\n"

	}
	info+="</tr><tr><td>DPI</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{
		tok,_:= fernet.EncryptAndSign([]byte(strconv.Itoa(pag.Datos[a].DPI)), llave[0])
		info+="<td>"+string(tok)+"</td>\n"
	}
	info+="</tr><tr><td>Correo</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{
		tok,_:= fernet.EncryptAndSign([]byte(pag.Datos[a].Correo), llave[0])
		info+="<td>"+string(tok)+"</td>\n"
	}
	info+="</tr><tr><td>Password</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{
		tok,_:= fernet.EncryptAndSign([]byte(pag.Datos[a].Password), llave[0])
		info+="<td>"+string(tok)+"</td>\n"
	}
	info+="</tr><tr><td>Cuenta</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{

		info+="<td>"+pag.Datos[a].Cuenta+"</td>\n"
	}
	info+="</tr></table>>];\n"

	if !pag.CondicionHoja{
		var contarHijos int
		for a:=0;a<5;a++{
			if pag.PaginasHijas[a]!=nil{
				contarHijos++
			}
		}
		for a:=0;a<contarHijos;a++{
			info=RecorrerHijosCifradoSensible(nombre,pag.PaginasHijas[a],info,llave)
		}
	}


	info+="\n}"

	return info
}
func RecorrerHijosCifradoSensible(NodoPadre string,pag *Pagina,info string,llave []*fernet.Key)string{
	nombre:="Nodo"+strconv.Itoa(contGrafo)
	contGrafo++
	info+=nombre+"[label=<<table><tr><td>Nombre</td>\n"

	for a:=0;a<pag.CantidadElementos;a++{

		info+="<td>"+pag.Datos[a].Nombre+"</td>\n"
	}
	info+="</tr><tr><td>DPI</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{
		tok,_:= fernet.EncryptAndSign([]byte(strconv.Itoa(pag.Datos[a].DPI)), llave[0])
		info+="<td>"+string(tok)+"</td>\n"
	}
	info+="</tr><tr><td>Correo</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{
		tok,_:= fernet.EncryptAndSign([]byte(pag.Datos[a].Correo), llave[0])
		info+="<td>"+string(tok)+"</td>\n"
	}
	info+="</tr><tr><td>Password</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{
		tok,_:= fernet.EncryptAndSign([]byte(pag.Datos[a].Password), llave[0])
		info+="<td>"+string(tok)+"</td>\n"
	}
	info+="</tr><tr><td>Cuenta</td>\n"
	for a:=0;a<pag.CantidadElementos;a++{

		info+="<td>"+pag.Datos[a].Cuenta+"</td>\n"
	}
	info+="</tr></table>>];\n"
	info+=NodoPadre+"->"+nombre+"\n"

	if !pag.CondicionHoja{
		var contarHijos int
		for a:=0;a<5;a++{
			if pag.PaginasHijas[a]!=nil{
				contarHijos++
			}
		}
		for a:=0;a<contarHijos;a++{
			info=RecorrerHijosCifradoSensible(nombre,pag.PaginasHijas[a],info,llave)
		}
	}


	return info
}
