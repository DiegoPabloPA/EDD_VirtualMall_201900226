package MatrizDispersa

import "strings"

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
	informacion *Pedido
}
type MatrizDispersa struct {
	Nodoinicio *Nodo
}
func (m *MatrizDispersa) CrearColumna(dia int){
	aux:=m.Nodoinicio
	if aux.Derecha==nil{
		nuevo:=&Nodo{
			Arriba:      nil,
			Abajo:       nil,
			Izquierda:   aux,
			Derecha:     nil,
			informacion: &Pedido{
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
		aux2:=m.Nodoinicio
		for aux2!=nil{
			if aux2.informacion.Dia==dia{
				bandera=false
			}else{
				aux2=aux2.Derecha
			}
		}
		if bandera{
			for aux.Derecha!=nil&&aux.informacion.Dia<dia{
				aux=aux.Derecha
			}
			if aux.Derecha==nil{
				nuevo:=&Nodo{
					Arriba:      nil,
					Abajo:       nil,
					Izquierda:   aux,
					Derecha:     nil,
					informacion: &Pedido{
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
					informacion: &Pedido{
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
			informacion: &Pedido{
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
		aux2:=m.Nodoinicio
		for aux2!=nil{
			if strings.EqualFold(aux2.informacion.Departamento,categoria){
				bandera=false
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
					informacion: &Pedido{
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

	nuevo:=&Nodo{
		Arriba:      nil,
		Abajo:       nil,
		Izquierda:   nil,
		Derecha:     nil,
		informacion:&Pedido{
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
	for auxFila.Abajo!=nil{
		auxFila=auxFila.Abajo
		if strings.EqualFold(auxFila.informacion.Departamento,FilaCategoria){
			if auxFila.Derecha==nil{
				auxFila.Derecha=nuevo
				nuevo.Izquierda=auxFila
			}else{
				bandera:=true
				for auxFila.Derecha!=nil{
					auxFila=auxFila.Derecha
					aux3:=auxFila
					for aux3.Arriba!=nil{
						aux3=aux3.Arriba
					}
					 if ColumnaDia<aux3.informacion.Dia{
							bandera=false
							aux4:=auxFila.Izquierda
							aux4.Derecha=nuevo
							nuevo.Izquierda=aux4
							nuevo.Derecha=auxFila
							auxFila.Izquierda=nuevo

					 }else if ColumnaDia==aux3.informacion.Dia{
					 		bandera=false
					 		aux5:=auxFila.informacion
							for aux5.Siguiente!=nil{
								aux5=aux5.Siguiente
							}
							aux5.Siguiente=&Pedido{
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
				if bandera{
					auxFila.Derecha=nuevo
					nuevo.Izquierda=auxFila
				}
			}

		}
	}
	//-----------------------------------------------------------------------------------------------
	auxColumna:=m.Nodoinicio
	for auxColumna.Derecha!=nil{
		auxColumna=auxColumna.Derecha
		if auxColumna.informacion.Dia==ColumnaDia{
			if auxColumna.Abajo==nil{
				auxColumna.Abajo=nuevo
				nuevo.Arriba=auxColumna
			}else{
				bandera:=true
				for auxColumna.Abajo!=nil{
					auxColumna=auxColumna.Abajo
					aux3:=auxColumna
					for aux3.Izquierda!=nil{
						aux3=aux3.Izquierda
					}
					aux7:=aux3
					bandera2:=true
					for aux7.Arriba!=nil{
					if strings.EqualFold(aux7.Arriba.informacion.Departamento,FilaCategoria){
						bandera2=false
						bandera=false
						aux4:=auxColumna.Arriba
						aux4.Abajo=nuevo
						nuevo.Arriba=aux4
						nuevo.Abajo=auxFila
						auxFila.Arriba=nuevo

					}
					aux7=aux7.Arriba
					}
					if strings.EqualFold(FilaCategoria,aux3.informacion.Departamento)&&bandera2{
						bandera=false
						aux5:=auxColumna.informacion
						for aux5.Siguiente!=nil{
							aux5=aux5.Siguiente
						}
						aux5.Siguiente=&Pedido{
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
				if bandera{
					auxColumna.Abajo=nuevo
					nuevo.Arriba=auxColumna
				}
			}
		}
	}
}


