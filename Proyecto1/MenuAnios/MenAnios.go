package MenuAnios

import (
	"Proyecto1/AVLAnio"
	"Proyecto1/AVLMeses"
	"Proyecto1/MatrizDispersa"
)

type Meses struct {
	Mes string
	Dias []Dia
}
type Dia struct {
	Dia int
}

type Anio struct{
	Anio int
	Meses []Meses
}
func RecepAnios(nodo*AVLAnio.NodoAVLAnio)[]Anio{
	nuevo:=Anio{
		Anio:  nodo.Datos.Anio,
		Meses: nil,
	}
	nuevo.Meses=RecepMeses(nodo.Datos.AVLdeMes.Raiz)
	var infor []Anio
	infor=append(infor,nuevo)
	infor=DevolverAnios(nodo,infor)
	return infor
}



func DevolverAnios(nodo *AVLAnio.NodoAVLAnio,infor[]Anio)[]Anio{
	if nodo.Izquierda!=nil{
		nuevo:=Anio{
			Anio:  nodo.Izquierda.Datos.Anio,
			Meses: nil,
		}
		nuevo.Meses=RecepMeses(nodo.Izquierda.Datos.AVLdeMes.Raiz)
		infor=append(infor,nuevo)
		infor=DevolverAnios(nodo.Izquierda,infor)
	}
	if nodo.Derecha!=nil{
		nuevo:=Anio{
			Anio:  nodo.Derecha.Datos.Anio,
			Meses: nil,
		}
		nuevo.Meses=RecepMeses(nodo.Derecha.Datos.AVLdeMes.Raiz)
		infor=append(infor,nuevo)
		infor=DevolverAnios(nodo.Derecha,infor)
	}
	return infor
}



func RecepMeses(nodo*AVLMeses.NodoAVLMes)[]Meses{

		nuevo:=Meses{
			Mes:  nodo.Datos.Mes,
			Dias: nil,
		}
		nuevo.Dias=DevolverDias(nodo.Datos.Matriz.Nodoinicio)
		var infor []Meses
		infor=append(infor,nuevo)
		infor=DevolverMeses(nodo,infor)
		return infor

}



func DevolverMeses(nodo *AVLMeses.NodoAVLMes,info[]Meses)[]Meses{
	if nodo.Izquierda!=nil{
		nuevo:=Meses{
			Mes:  nodo.Izquierda.Datos.Mes,
			Dias: nil,
		}
		nuevo.Dias=DevolverDias(nodo.Izquierda.Datos.Matriz.Nodoinicio)
		info=append(info,nuevo)
		info=DevolverMeses(nodo.Izquierda,info)
	}
	if nodo.Derecha!=nil{
		nuevo:=Meses{
			Mes:  nodo.Derecha.Datos.Mes,
			Dias: nil,
		}
		nuevo.Dias=DevolverDias(nodo.Derecha.Datos.Matriz.Nodoinicio)
		info=append(info,nuevo)
		info=DevolverMeses(nodo.Derecha,info)
	}
	return info
}


func DevolverDias(nodo*MatrizDispersa.Nodo)(info []Dia){
	for nodo.Derecha!=nil{
		nodo=nodo.Derecha
		nuevo:=Dia{Dia: nodo.Informacion.Dia}
		info=append(info,nuevo)
	}
	return info
}
func DevolverDiasCategoria(nodo*MatrizDispersa.Nodo)(info[]DiasConCateogira) {
	for nodo.Abajo!=nil{
		nodo=nodo.Abajo
		aux1:=nodo
		nuevo:=DiasConCateogira{
			Categoria: aux1.Informacion.Departamento,
			Dias:      nil,
		}
		for aux1.Derecha!=nil{
		aux1=aux1.Derecha

		nuevo2:=Dia{Dia: aux1.Informacion.Dia}
		nuevo.Dias=append(nuevo.Dias,nuevo2)
		}
		info=append(info,nuevo)
	}

	return info
}
type DiasConCateogira struct {
	Categoria string
	Dias []Dia
}
type RespInfo struct {
	Anio int    `json:"Anio"`
	Mes  string `json:"Mes"`
}