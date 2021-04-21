package ReportesDiaCategoria

import (
	"Proyecto1/AVLAnio"
	"Proyecto1/AVLMeses"
	"strings"
)

type JsonReporteDiaSolicitado struct {
	Anio      int    `json:"Anio"`
	Mes       string `json:"Mes"`
	Dia       int    `json:"Dia"`
	Categoria string `json:"Categoria"`
}
type DescripcionPedido struct {
	Codigo int
	Cantidad int
}
type RespuestaDiaSolicitado struct {
	Categoria string
	Direccion string
	Cliente string
	Descripcion []DescripcionPedido

}
func GenerarReporte(nodo *AVLAnio.NodoAVLAnio,solicitud JsonReporteDiaSolicitado)(respuesta[]RespuestaDiaSolicitado){
	aux1:=nodo
	for aux1!=nil{
		if aux1.Datos.Anio>solicitud.Anio{
			aux1=aux1.Izquierda
		}else if aux1.Datos.Anio<solicitud.Anio{
			aux1=aux1.Derecha
		}else{
			aux2:=aux1.Datos.AVLdeMes.Raiz
			aux1=aux1.Derecha
			for aux2!=nil{
				if aux2.Datos.NoMes>AVLMeses.SeleccionMes(solicitud.Mes){
					aux2=aux2.Izquierda
				}else if aux2.Datos.NoMes<AVLMeses.SeleccionMes(solicitud.Mes){
					aux2=aux2.Derecha
				}else{
					aux3:=aux2.Datos.Matriz.Nodoinicio
					aux2=aux2.Derecha
					for aux3.Abajo!=nil{
						aux3=aux3.Abajo
						if strings.EqualFold(aux3.Informacion.Departamento,solicitud.Categoria){
							aux4:=aux3

							for aux4.Derecha!=nil{
								aux4=aux4.Derecha
								if aux4.Informacion.Dia==solicitud.Dia{
									aux5:=aux4.Informacion

									for aux5!=nil{
										nuevo:=RespuestaDiaSolicitado{
											Categoria:   aux5.Departamento,
											Direccion: aux5.Direccion,
											Cliente: aux5.Cliente,
											Descripcion: nil,
										}
										for a:=0;a<len(aux5.Descripcion);a++{
											nuevo2:=DescripcionPedido{
												Codigo:   aux5.Descripcion[a].Codigo,
												Cantidad: aux5.Descripcion[a].Cantidad,
											}
											nuevo.Descripcion=append(nuevo.Descripcion,nuevo2)
										}
										if len(nuevo.Descripcion)==0{
											aux5=aux5.Siguiente
										}else{
										respuesta=append(respuesta,nuevo)
										aux5=aux5.Siguiente
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}


	return respuesta
}
