package Inventarios

import (
	"Proyecto1/EstructuraAVL"
	"Proyecto1/ListaDoble"
	"fmt"
	"strings"
)


func Mayor(numa int,numb int)int{
	if numa>numb{
		return numa
	}
	return numb

}
func Altura(evaluar *EstructuraAVL.NodoArbol)int{
	if evaluar==nil{
		return -1
	}
	return evaluar.Altura
}
func RotacionSimpleIzquierda(evaluar*EstructuraAVL.NodoArbol)*EstructuraAVL.NodoArbol{
	aux1:=evaluar.Izquierda
	evaluar.Izquierda=aux1.Derecha
	aux1.Derecha=evaluar
	evaluar.Altura=Mayor(Altura(evaluar.Izquierda),Altura(evaluar.Derecha))+1
	aux1.Altura=Mayor(Altura(evaluar.Izquierda),evaluar.Altura)+1
	evaluar=aux1
	return evaluar
}
func RotacionSimpleDerecha(evaluar *EstructuraAVL.NodoArbol)*EstructuraAVL.NodoArbol{
	aux1:=evaluar.Derecha
	evaluar.Derecha=aux1.Izquierda
	aux1.Izquierda=evaluar
	evaluar.Altura=Mayor(Altura(evaluar.Izquierda),Altura(evaluar.Derecha))+1
	aux1.Altura=Mayor(Altura(evaluar.Derecha),evaluar.Altura)+1
	evaluar=aux1
	return evaluar
}
func RotacionDobleIzquierda(evaluar *EstructuraAVL.NodoArbol)*EstructuraAVL.NodoArbol{
	evaluar=RotacionSimpleDerecha(evaluar)
	evaluar=RotacionSimpleIzquierda(evaluar)
	return evaluar
}
func RotacionDobleDerecha(evaluar *EstructuraAVL.NodoArbol)*EstructuraAVL.NodoArbol{
	evaluar=RotacionSimpleIzquierda(evaluar)
	evaluar=RotacionSimpleDerecha(evaluar)
	return evaluar
}


func  Insertar(nodo *EstructuraAVL.NodoArbol,datos EstructuraAVL.DatosNodoArbol)*EstructuraAVL.NodoArbol{
	nuevo:=&EstructuraAVL.NodoArbol{Izquierda: nil, Derecha:   nil, Altura:    0, Datos:     datos}
	if nodo==nil{
		return nuevo
	}else if  nodo.Datos.Codigo>datos.Codigo{
		nodo.Izquierda=Insertar(nodo.Izquierda,datos)
		if Altura(nodo.Izquierda)-Altura(nodo.Derecha)==-2{
			if datos.Codigo<nodo.Izquierda.Altura{
				nodo=RotacionSimpleIzquierda(nodo)
			}else{
				nodo=RotacionDobleIzquierda(nodo)
			}
		}}else if nodo.Datos.Codigo<datos.Codigo{
		nodo.Derecha=Insertar(nodo.Derecha,datos)
		if Altura(nodo.Derecha)-Altura(nodo.Izquierda)==2{
			if datos.Codigo>nodo.Datos.Codigo{
				nodo=RotacionSimpleDerecha(nodo)

			}else{
				nodo=RotacionDobleDerecha(nodo)
			}
		}
	}else{
		nodo.Datos.Cantidad=datos.Cantidad+nodo.Datos.Cantidad
	}
	nodo.Altura=Mayor(Altura(nodo.Izquierda),Altura(nodo.Derecha))+1
	return nodo
}
func Imprimir(nodo *EstructuraAVL.NodoArbol){
	if nodo!=nil{
		Imprimir(nodo.Izquierda)
		fmt.Println(nodo.Datos.Codigo,"Cantidad:",nodo.Datos.Cantidad)
		Imprimir(nodo.Derecha)
	}
}
func EnviarDatos(nodo *EstructuraAVL.NodoArbol,resultado[]InventarioJson)(resul[]InventarioJson){
	if nodo!=nil{
		resultado=EnviarDatos(nodo.Izquierda,resultado)
		datos:=InventarioJson{
			Nombre:      nodo.Datos.Nombre,
			Codigo:      nodo.Datos.Codigo,
			Descripcion: nodo.Datos.Descripcion,
			Precio:      nodo.Datos.Precio,
			Cantidad:    nodo.Datos.Cantidad,
			Imagen:      nodo.Datos.Imagen,
		}
		resultado=append(resultado,datos)
		resultado=EnviarDatos(nodo.Derecha,resultado)
	}
	return resultado
}


type InventarioJson struct{
	Nombre      string
	Codigo      int
	Descripcion string
	Precio      float64
	Cantidad    int
	Imagen      string
}
type CargadeInventario struct {
	Invetario []struct {
		Tienda       string `json:"Tienda"`
		Departamento string `json:"Departamento"`
		Calificacion int    `json:"Calificacion"`
		Productos    []struct {
			Nombre      string  `json:"Nombre"`
			Codigo      int     `json:"Codigo"`
			Descripcion string  `json:"Descripcion"`
			Precio      float64 `json:"Precio"`
			Cantidad    int     `json:"Cantidad"`
			Imagen      string  `json:"Imagen"`
		} `json:"Productos"`
	} `json:"Invetarios"`
}



func CargaInventariosMasivo(arreglo []ListaDoble.ListaDE,informacion CargadeInventario)[]ListaDoble.ListaDE{
	aux:=arreglo
	for a:=0;a<len(informacion.Invetario);a++{
		for b:=0;b<len(arreglo);b++{
			if aux[b].Tamanio>0{
				aux2:=aux[b].Inicio
				for aux2!=nil{
					if strings.EqualFold(informacion.Invetario[a].Tienda,aux2.Datos.Nombre)&&strings.EqualFold(informacion.Invetario[a].Departamento,aux2.Datos.Columna)&&informacion.Invetario[a].Calificacion==aux2.Datos.Calificacion{
						for c:=0;c<len(informacion.Invetario[a].Productos);c++{
							ingreso:=EstructuraAVL.DatosNodoArbol{
								Codigo:      informacion.Invetario[a].Productos[c].Codigo,
								Nombre:      informacion.Invetario[a].Productos[c].Nombre,
								Descripcion: informacion.Invetario[a].Productos[c].Descripcion,
								Imagen:      informacion.Invetario[a].Productos[c].Imagen,
								Precio:      informacion.Invetario[a].Productos[c].Precio,
								Cantidad:    informacion.Invetario[a].Productos[c].Cantidad,
							}
							aux2.Datos.Inventario.Raiz=Insertar(aux2.Datos.Inventario.Raiz,ingreso)

						}
						bandera:=true
						for bandera{
							if strings.EqualFold(aux2.Datos.Nombre,aux[b].Inicio.Datos.Nombre)&&strings.EqualFold(aux2.Datos.Columna,aux[b].Inicio.Datos.Columna)&&aux2.Datos.Calificacion==aux[b].Inicio.Datos.Calificacion{
								bandera=false
							}else{
								aux2=aux2.Anterior
							}
						}
						aux[b].Inicio=aux2
						bandera=true
						for bandera{
							if strings.EqualFold(aux2.Datos.Nombre,aux[b].Final.Datos.Nombre)&&strings.EqualFold(aux2.Datos.Columna,aux[b].Final.Datos.Columna)&&aux2.Datos.Calificacion==aux[b].Final.Datos.Calificacion{
								bandera=false
							}else{
								aux2=aux2.Siguiente
							}
						}
						aux[b].Final=aux2
						aux2=nil

					}else{
					aux2=aux2.Siguiente
					}

				}

			}
		}
	}



	return aux
}

func CargaInventariosIndividual(arreglo []ListaDoble.ListaDE,informacion EstructuraAVL.InventarioIndividual,estructura EstructuraAVL.EstructuraBusqueda)[]ListaDoble.ListaDE{
	aux:=arreglo

		for b:=0;b<len(arreglo);b++{
			if aux[b].Tamanio>0{
				aux2:=aux[b].Inicio
				for aux2!=nil{
					if strings.EqualFold(estructura.Nombre,aux2.Datos.Nombre)&&strings.EqualFold(estructura.Departamento,aux2.Datos.Columna)&&estructura.Calificacion==aux2.Datos.Calificacion{

							ingreso:=EstructuraAVL.DatosNodoArbol{
								Codigo:      informacion.Codigo,
								Nombre:      informacion.Nombre,
								Descripcion: informacion.Descripcion,
								Imagen:      informacion.Imagen,
								Precio:      informacion.Precio,
								Cantidad:    informacion.Cantidad,
							}
							aux2.Datos.Inventario.Raiz=Insertar(aux2.Datos.Inventario.Raiz,ingreso)


						bandera:=true
						for bandera{
							if strings.EqualFold(aux2.Datos.Nombre,aux[b].Inicio.Datos.Nombre)&&strings.EqualFold(aux2.Datos.Columna,aux[b].Inicio.Datos.Columna)&&aux2.Datos.Calificacion==aux[b].Inicio.Datos.Calificacion{
								bandera=false
							}else{
								aux2=aux2.Anterior
							}
						}
						aux[b].Inicio=aux2
						bandera=true
						for bandera{
							if strings.EqualFold(aux2.Datos.Nombre,aux[b].Final.Datos.Nombre)&&strings.EqualFold(aux2.Datos.Columna,aux[b].Final.Datos.Columna)&&aux2.Datos.Calificacion==aux[b].Final.Datos.Calificacion{
								bandera=false
							}else{
								aux2=aux2.Siguiente
							}
						}
						aux[b].Final=aux2
						aux2=nil

					}else{
						aux2=aux2.Siguiente
					}

				}

			}
		}




	return aux
}
