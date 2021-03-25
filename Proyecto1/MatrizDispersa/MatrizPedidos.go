package MatrizDispersa

import (
	"fmt"
	"strings"
)

type DescripcionPedido struct {
	Codigo int
	Cantidad int
}

type Pedido struct{
	Siguiente *Pedido
	Dia int
	Mes string
	Anio int
	Cliente string
	Direccion string
	NombreTienda string
	Departamento string
	Calificacion int
	Descripcion []DescripcionPedido
}



type Nodo struct {
	Arriba,Abajo,Izquierda,Derecha *Nodo
	Informacion *Pedido
}
type MatrizDispersa struct {
	Nodoinicio *Nodo
}
func (m *MatrizDispersa) Init() {
	m.Nodoinicio = &Nodo{
		Arriba:      nil,
		Abajo:       nil,
		Izquierda:   nil,
		Derecha:     nil,
		Informacion: nil,
	}
}


func (m *MatrizDispersa) CrearColumna(dia int){
	aux:=m.Nodoinicio
	if aux.Derecha==nil{
		nuevo:=&Nodo{
			Arriba:      nil,
			Abajo:       nil,
			Izquierda:   aux,
			Derecha:     nil,
			Informacion: &Pedido{
				Siguiente: nil,
				Dia:          dia,
				Mes:          "",
				Anio:         0,
				Cliente:      "",
				Direccion:    "",
				NombreTienda: "",
				Departamento: "",
				Calificacion: 0,
				Descripcion:  nil,
			},
		}
		aux.Derecha=nuevo
	}else{
		bandera:=true
		aux2:=m.Nodoinicio.Derecha
		for aux2!=nil{
			if aux2.Informacion.Dia==dia{
				bandera=false
				aux2=aux2.Derecha
			}else{
				aux2=aux2.Derecha
			}
		}
		if bandera{
			for aux.Derecha!=nil&&aux.Derecha.Informacion.Dia<dia{
				aux=aux.Derecha
			}
			if aux.Derecha==nil{
				nuevo:=&Nodo{
					Arriba:      nil,
					Abajo:       nil,
					Izquierda:   aux,
					Derecha:     nil,
					Informacion: &Pedido{
						Siguiente: nil,
						Dia:          dia,
						Mes:          "",
						Anio:         0,
						Cliente:      "",
						Direccion:    "",
						NombreTienda: "",
						Departamento: "",
						Calificacion: 0,
						Descripcion:  nil,
					},
				}
				aux.Derecha=nuevo
			}else if aux.Derecha!=nil{
				aux3:=aux.Derecha
				nuevo:=&Nodo{
					Arriba:      nil,
					Abajo:       nil,
					Izquierda:   aux,
					Derecha:     aux3,
					Informacion: &Pedido{
						Siguiente: nil,
						Dia:          dia,
						Mes:          "",
						Anio:         0,
						Cliente:      "",
						Direccion:    "",
						NombreTienda: "",
						Departamento: "",
						Calificacion: 0,
						Descripcion:  nil,
					},
				}
				aux.Derecha=nuevo
				aux3.Izquierda=nuevo
			}
		}
	}

}
func (m *MatrizDispersa)GenerarFila(categoria string){
	aux:=m.Nodoinicio
	if aux.Abajo==nil{
		nuevo:=&Nodo{
			Arriba:      aux,
			Abajo:       nil,
			Izquierda:   nil,
			Derecha:     nil,
			Informacion: &Pedido{
				Siguiente: nil,
				Dia:          0,
				Mes:          "",
				Anio:         0,
				Cliente:      "",
				Direccion:    "",
				NombreTienda: "",
				Departamento: categoria,
				Calificacion: 0,
				Descripcion:  nil,
			},
		}
		aux.Abajo=nuevo
	}else{
		bandera:=true
		aux2:=m.Nodoinicio.Abajo
		for aux2!=nil{

			if strings.EqualFold(aux2.Informacion.Departamento,categoria){
				bandera=false
				aux2=aux2.Abajo
			}else{
				aux2=aux2.Abajo
			}
		}
		if bandera{
			for aux.Abajo!=nil{
				aux=aux.Abajo
			}
			if aux.Abajo==nil{
				nuevo:=&Nodo{
					Arriba:      aux,
					Abajo:       nil,
					Izquierda:   nil,
					Derecha:     nil,
					Informacion: &Pedido{
						Siguiente: nil,
						Dia:          0,
						Mes:          "",
						Anio:         0,
						Cliente:      "",
						Direccion:    "",
						NombreTienda: "",
						Departamento: categoria,
						Calificacion: 0,
						Descripcion:  nil,
					},
				}
				aux.Abajo=nuevo
			}
		}
	}
}

func (m*MatrizDispersa)InsertarNodo(FilaCategoria string, ColumnaDia int,Informacion Pedido){
	auxFila:=m.Nodoinicio
	banderaUniversal:=true
	nuevo:=&Nodo{
		Arriba:      nil,
		Abajo:       nil,
		Izquierda:   nil,
		Derecha:     nil,
		Informacion:&Pedido{
			Siguiente:    nil,
			Dia:          Informacion.Dia,
			Mes:          Informacion.Mes,
			Anio:         Informacion.Anio,
			Cliente:      Informacion.Cliente,
			Direccion:    Informacion.Direccion,
			NombreTienda: Informacion.NombreTienda,
			Departamento: Informacion.Departamento,
			Calificacion: Informacion.Calificacion,
			Descripcion:  Informacion.Descripcion,
		},
	}

	//Insertar el nodo en la Fila
	bandera1Fila:=true
	for auxFila.Abajo!=nil{
		auxFila=auxFila.Abajo
		if strings.EqualFold(FilaCategoria,auxFila.Informacion.Departamento){
			if auxFila.Derecha==nil{

				auxFila.Derecha=nuevo
				nuevo.Izquierda=auxFila
			}else if auxFila.Derecha!=nil{
				for auxFila.Derecha!=nil {
					auxFila = auxFila.Derecha
					aux1:=auxFila
					for aux1.Arriba!=nil{
						aux1=aux1.Arriba
					}

					if ColumnaDia<aux1.Informacion.Dia{
						bandera1Fila=false
						aux2:=auxFila.Izquierda
						nuevo.Derecha=auxFila
						nuevo.Izquierda=aux2
						auxFila.Izquierda=nuevo
						aux2.Derecha=nuevo
					}else if ColumnaDia>aux1.Informacion.Dia&&auxFila.Derecha==nil&&bandera1Fila{
						bandera1Fila=false
						nuevo.Izquierda=auxFila
						auxFila.Derecha=nuevo
					}else if ColumnaDia==aux1.Informacion.Dia&&bandera1Fila{
						banderaUniversal=false
						aux2:=auxFila.Informacion
						for aux2.Siguiente!=nil{
							aux2=aux2.Siguiente
						}


						aux2.Siguiente=&Pedido{
							Siguiente:    nil,
							Dia:          Informacion.Dia,
							Mes:          Informacion.Mes,
							Anio:         Informacion.Anio,
							Cliente:      Informacion.Cliente,
							Direccion:    Informacion.Direccion,
							NombreTienda: Informacion.NombreTienda,
							Departamento: Informacion.Departamento,
							Calificacion: Informacion.Calificacion,
							Descripcion:  Informacion.Descripcion,
						}
					}
				}

			}
		}
	}

	//<---------------------------------------------------------------------------------------------------------------->




	auxColumna:=m.Nodoinicio
	banderaColumna:=true
	for auxColumna.Derecha!=nil{
		auxColumna=auxColumna.Derecha
		if auxColumna.Informacion.Dia==ColumnaDia{
			if auxColumna.Abajo==nil{
				auxColumna.Abajo=nuevo
				nuevo.Arriba=auxColumna
			}else if auxColumna.Abajo!=nil{
				for auxColumna.Abajo!=nil {
					auxColumna = auxColumna.Abajo
					//ahora voy a encontrar la categoria
					comparadorFila:=auxColumna
					for comparadorFila.Izquierda!=nil{
						comparadorFila=comparadorFila.Izquierda
					}
					for comparadorFila.Arriba!=nil{
						if strings.EqualFold(comparadorFila.Informacion.Departamento,FilaCategoria)&&banderaUniversal{
							banderaColumna=false
							aux3:=auxColumna.Arriba
							nuevo.Arriba=aux3
							nuevo.Abajo=auxColumna
							auxColumna.Arriba=nuevo
							aux3.Abajo=nuevo
						}
						comparadorFila=comparadorFila.Arriba
					}


				}



			}
			if banderaColumna&&banderaUniversal{
				auxColumna.Abajo=nuevo
				nuevo.Arriba=auxColumna
			}
		}
	}

}

func (m MatrizDispersa)InsertarPedido(datos Pedido){
	m.CrearColumna(datos.Dia)
	m.GenerarFila(datos.Departamento)
	m.InsertarNodo(datos.Departamento,datos.Dia,datos)

}
func (m MatrizDispersa)Imprimir(){
	auxColumna:=m.Nodoinicio
	auxFila:=m.Nodoinicio

	for auxColumna.Derecha!=nil{
		auxColumna=auxColumna.Derecha
		fmt.Print("       	   ",auxColumna.Informacion.Dia," ")
	}
	fmt.Println(" ")
	for auxFila.Abajo!=nil{

		fmt.Print(auxFila.Abajo.Informacion.Departamento," ")
		auxFila=auxFila.Abajo
		aux2:=auxFila
		for aux2.Derecha!=nil{

			aux2=aux2.Derecha

			aux3:=aux2


			for aux3.Arriba!=nil{

				aux3=aux3.Arriba
			}
			fmt.Print("(",auxFila.Informacion.Departamento,",",aux3.Informacion.Dia,")",aux2.Informacion.Cliente)
			desaux:=aux2.Informacion
			for desaux.Siguiente!=nil{
				desaux=desaux.Siguiente

				fmt.Print(" (",auxFila.Informacion.Departamento,",",desaux.Dia,")",desaux.Cliente)
			}
		}
		fmt.Println(" ")
	}
}


