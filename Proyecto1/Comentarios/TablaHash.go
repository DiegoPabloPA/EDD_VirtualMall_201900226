package Comentarios

import (
	"math"
	"strconv"
	"strings"
)
type DatosComentarios struct {
	NombreUsuario string
	DPI int
	DPIPadre []int
	Comentario string
	SubComentarios []DatosComentarios
}

var I int
var J int
func DefinirTamanioPrimoHash(numero int)int{
	numero+=1
	condicion:=true
	for condicion{
		divisores:=0
		for a:=1;a<=numero;a++{
			if numero%a==0{
				divisores+=1
			}
		}
		if divisores==2{
			condicion=false
		}else{
			numero+=1
		}

	}


	return numero
}
func CalcularOcupacion(analizar[]DatosComentarios) (res int){
	tamanio:=len(analizar)
	elementos:=0
	for a:=0;a<tamanio;a++{
		if !strings.EqualFold(analizar[a].Comentario,""){
			elementos+=1
		}
	}
	res=(100*elementos)/tamanio

	return res
}
func CalcularOcupacionSubComentarios(analizar[]DatosComentarios) (res int){
	tamanio:=len(analizar)
	elementos:=0
	for a:=0;a<tamanio;a++{
		if !strings.EqualFold(analizar[a].Comentario,""){
			elementos+=1
		}
	}
	res=(100*elementos)/tamanio

	return res
}



func CalcularCodigoHash(DPI int,tamanio int)(codigo int){
	A:=0.1563*float64(DPI)
	operador:=math.Mod(A,1)
	resultado:=float64(tamanio)*operador
	codigo=int(math.Floor(resultado))
	return codigo
}
func Rehash(arreglo []DatosComentarios,padres []int,posalterar int)[]DatosComentarios{
	aux:=arreglo
	tamanio:=DefinirTamanioPrimoHash(len(arreglo))
	arreglo=make([]DatosComentarios,tamanio)
	for a:=0;a<len(aux);a++{
		if aux[a].DPI>0{
			m:=CalcularCodigoHash(aux[a].DPI,tamanio)
			if arreglo[m].DPI==0 {
				arreglo[m] = aux[a]
				arreglo[m].DPIPadre=padres
				arreglo[m].DPIPadre=append(arreglo[m].DPIPadre,m)
				arreglo[m].SubComentarios=RealizarCambioArregloPadre(posalterar,m,arreglo[m].SubComentarios)
			}else{
				llaveaux:=Redireccionamiento(arreglo,m)
				arreglo[llaveaux]=aux[a]
				arreglo[llaveaux].DPIPadre=padres
				arreglo[llaveaux].DPIPadre=append(arreglo[llaveaux].DPIPadre,llaveaux)
				arreglo[llaveaux].SubComentarios=RealizarCambioArregloPadre(posalterar,llaveaux,arreglo[llaveaux].SubComentarios)
			}
		}
	}

	return arreglo
}

func RealizarCambioArregloPadre(posicion int,valor int,datos[]DatosComentarios)[]DatosComentarios{
	if len(datos)>0{
		for a:=0;a<len(datos);a++{
			if datos[a].DPI>0{
				datos[a].DPIPadre[posicion]=valor
				datos[a].SubComentarios=RealizarCambioArregloPadre(posicion,valor,datos[a].SubComentarios)
			}
		}
	}
	return datos
}





type EsqueletoComentario struct {
	Dpi        int    `json:"DPI"`
	NombreSec string `json:"NombreSec"`
	Categoria string `json:"Categoria"`
	Calificacion int `json:"Calificacion"`
	Nombre     string `json:"Nombre"`
	Comentario string `json:"Comentario"`
}
type EsqueletoComentarioArticulo struct {
	Dpi        int    `json:"DPI"`
	Codigo 	   int    `json:"Codigo"`
	NombreSec string `json:"NombreSec"`
	Categoria string `json:"Categoria"`
	Calificacion int `json:"Calificacion"`
	Nombre     string `json:"Nombre"`
	Comentario string `json:"Comentario"`
}


type EsqueletoSubComentario struct {
	Clave       string    `json:"Clave"`
	Dpi        int    `json:"DPI"`
	NombreSec string `json:"NombreSec"`
	Categoria string `json:"Categoria"`
	Calificacion int `json:"Calificacion"`
	Nombre     string `json:"Nombre"`
	Comentario string `json:"Comentario"`
}
type EsqueletoSubComentarioArticulo struct {
	Codigo 	   int    `json:"Codigo"`
	Clave       string    `json:"Clave"`
	Dpi        int    `json:"DPI"`
	NombreSec string `json:"NombreSec"`
	Categoria string `json:"Categoria"`
	Calificacion int `json:"Calificacion"`
	Nombre     string `json:"Nombre"`
	Comentario string `json:"Comentario"`
}




func Redireccionamiento (datos []DatosComentarios,llaveocupada int)(llave int){
	I+=1
	if I>len(datos)-1{
		I=0
	}

	llave=int(math.Mod(float64(llaveocupada+(I*I)),float64(len(datos))))

	if datos[llave].DPI!=0{
		llave=Redireccionamiento(datos,llave)

	}


	return llave
}
func IngresarComentarioPrincipalInventario(comentario EsqueletoComentarioArticulo,datos []DatosComentarios)[]DatosComentarios{
	posicion:=CalcularCodigoHash(comentario.Dpi,len(datos))
	posArreglos=0
	var padrePrincipal []int

	if datos[posicion].DPI==0{
		datos[posicion].DPI=comentario.Dpi
		datos[posicion].Comentario=comentario.Comentario
		datos[posicion].NombreUsuario=comentario.Nombre
		datos[posicion].SubComentarios=make([]DatosComentarios,7)
		datos[posicion].DPIPadre=append(datos[posicion].DPIPadre,posicion)
		ocupacion:=CalcularOcupacion(datos)

		if ocupacion>=60{
			datos=Rehash(datos,padrePrincipal,0)
		}

	}else{
		posicionRedireccionada:=Redireccionamiento(datos,posicion)
		datos[posicionRedireccionada].DPI=comentario.Dpi
		datos[posicionRedireccionada].Comentario=comentario.Comentario
		datos[posicionRedireccionada].NombreUsuario=comentario.Nombre
		datos[posicionRedireccionada].SubComentarios=make([]DatosComentarios,7)
		datos[posicionRedireccionada].DPIPadre=append(datos[posicionRedireccionada].DPIPadre,posicionRedireccionada)
		ocupacion:=CalcularOcupacion(datos)
		if ocupacion>=60{
			datos=Rehash(datos,padrePrincipal,0)
		}
	}


	return datos
}






var posArreglos int

func IngresarComentarioPrincipal(comentario EsqueletoComentario,datos []DatosComentarios)[]DatosComentarios{
	posicion:=CalcularCodigoHash(comentario.Dpi,len(datos))
	posArreglos=0
	var padrePrincipal []int

	if datos[posicion].DPI==0{
		datos[posicion].DPI=comentario.Dpi
		datos[posicion].Comentario=comentario.Comentario
		datos[posicion].NombreUsuario=comentario.Nombre
		datos[posicion].SubComentarios=make([]DatosComentarios,7)
		datos[posicion].DPIPadre=append(datos[posicion].DPIPadre,posicion)
		ocupacion:=CalcularOcupacion(datos)

		if ocupacion>=60{
			datos=Rehash(datos,padrePrincipal,0)
		}

	}else{
		posicionRedireccionada:=Redireccionamiento(datos,posicion)
		datos[posicionRedireccionada].DPI=comentario.Dpi
		datos[posicionRedireccionada].Comentario=comentario.Comentario
		datos[posicionRedireccionada].NombreUsuario=comentario.Nombre
		datos[posicionRedireccionada].SubComentarios=make([]DatosComentarios,7)
		datos[posicionRedireccionada].DPIPadre=append(datos[posicionRedireccionada].DPIPadre,posicionRedireccionada)
		ocupacion:=CalcularOcupacion(datos)
		if ocupacion>=60{
			datos=Rehash(datos,padrePrincipal,0)
		}
	}


	return datos
}


func IngresarSubComentarios(comentario EsqueletoSubComentario,datos []DatosComentarios)[]DatosComentarios{
	avanza:=0
	auxiliar:=strings.Split(comentario.Clave,";")
	var claves []int
	claves=make([]int,len(auxiliar))
	for a:=0;a<len(auxiliar);a++{
		claves[a],_=strconv.Atoi(auxiliar[a])
	}
	datos[claves[avanza]].SubComentarios=IngresarSubComentariosRecursivo(avanza,comentario,datos[claves[avanza]].SubComentarios,claves,datos[claves[avanza]].DPIPadre)

	return datos
}

func IngresarSubComentariosArticulo(comentario EsqueletoSubComentarioArticulo,datos []DatosComentarios)[]DatosComentarios{
	avanza:=0
	auxiliar:=strings.Split(comentario.Clave,";")
	var claves []int
	claves=make([]int,len(auxiliar))
	for a:=0;a<len(auxiliar);a++{
		claves[a],_=strconv.Atoi(auxiliar[a])
	}
	datos[claves[avanza]].SubComentarios=IngresarSubComentariosRecursivoArticulo(avanza,comentario,datos[claves[avanza]].SubComentarios,claves,datos[claves[avanza]].DPIPadre)

	return datos
}


func IngresarSubComentariosRecursivoArticulo(avanza int,comentarios EsqueletoSubComentarioArticulo,datos[]DatosComentarios,llaves []int,padres[]int)[]DatosComentarios{

	if len(llaves)-1==avanza{
		posicion:=CalcularCodigoHash(comentarios.Dpi,len(datos))


		if datos[posicion].DPI==0{
			datos[posicion].DPI=comentarios.Dpi
			datos[posicion].Comentario=comentarios.Comentario
			datos[posicion].NombreUsuario=comentarios.Nombre
			datos[posicion].SubComentarios=make([]DatosComentarios,7)
			datos[posicion].DPIPadre=padres
			datos[posicion].DPIPadre=append(datos[posicion].DPIPadre,posicion)
			ocupacion:=CalcularOcupacion(datos)

			if ocupacion>=60{
				datos=Rehash(datos,padres,avanza)
			}

		}else{
			posicionRedireccionada:=Redireccionamiento(datos,posicion)
			datos[posicionRedireccionada].DPI=comentarios.Dpi
			datos[posicionRedireccionada].Comentario=comentarios.Comentario
			datos[posicionRedireccionada].NombreUsuario=comentarios.Nombre
			datos[posicionRedireccionada].SubComentarios=make([]DatosComentarios,7)
			datos[posicionRedireccionada].DPIPadre=padres
			datos[posicionRedireccionada].DPIPadre=append(datos[posicionRedireccionada].DPIPadre,posicionRedireccionada)
			ocupacion:=CalcularOcupacion(datos)
			if ocupacion>=60{
				datos=Rehash(datos,padres,avanza)
			}
		}

	}else{
		avanza += 1
		datos[llaves[avanza]].SubComentarios = IngresarSubComentariosRecursivoArticulo(avanza,comentarios, datos[llaves[avanza]].SubComentarios, llaves, datos[llaves[avanza]].DPIPadre)

	}
	return datos
}


func IngresarSubComentariosRecursivo(avanza int,comentarios EsqueletoSubComentario,datos[]DatosComentarios,llaves []int,padres[]int)[]DatosComentarios{

	if len(llaves)-1==avanza{
		posicion:=CalcularCodigoHash(comentarios.Dpi,len(datos))


		if datos[posicion].DPI==0{
			datos[posicion].DPI=comentarios.Dpi
			datos[posicion].Comentario=comentarios.Comentario
			datos[posicion].NombreUsuario=comentarios.Nombre
			datos[posicion].SubComentarios=make([]DatosComentarios,7)
			datos[posicion].DPIPadre=padres
			datos[posicion].DPIPadre=append(datos[posicion].DPIPadre,posicion)
			ocupacion:=CalcularOcupacion(datos)

			if ocupacion>=60{
				datos=Rehash(datos,padres,avanza)
			}

		}else{
			posicionRedireccionada:=Redireccionamiento(datos,posicion)
			datos[posicionRedireccionada].DPI=comentarios.Dpi
			datos[posicionRedireccionada].Comentario=comentarios.Comentario
			datos[posicionRedireccionada].NombreUsuario=comentarios.Nombre
			datos[posicionRedireccionada].SubComentarios=make([]DatosComentarios,7)
			datos[posicionRedireccionada].DPIPadre=padres
			datos[posicionRedireccionada].DPIPadre=append(datos[posicionRedireccionada].DPIPadre,posicionRedireccionada)
			ocupacion:=CalcularOcupacion(datos)
			if ocupacion>=60{
				datos=Rehash(datos,padres,avanza)
			}
		}

	}else{
			avanza += 1
			datos[llaves[avanza]].SubComentarios = IngresarSubComentariosRecursivo(avanza,comentarios, datos[llaves[avanza]].SubComentarios, llaves, datos[llaves[avanza]].DPIPadre)

		}
	return datos
}




