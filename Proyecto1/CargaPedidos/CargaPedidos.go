package CargaPedidos

import (
	"Proyecto1/AVLAnio"
	"Proyecto1/ListaDoble"
	"Proyecto1/MatrizDispersa"
	"fmt"
	"strconv"
	"strings"
)


type JsonPedidosMasivo struct {
	Pedidos []struct {
		Fecha        string `json:"Fecha"`
		Tienda       string `json:"Tienda"`
		Departamento string `json:"Departamento"`
		Calificacion int    `json:"Calificacion"`
		Productos    []struct {
			Codigo int `json:"Codigo"`
		} `json:"Productos"`
	} `json:"Pedidos"`
}

func CargaMasiva(datos JsonPedidosMasivo,Arreglo []ListaDoble.ListaDE,pedidos *AVLAnio.NodoAVLAnio)*AVLAnio.NodoAVLAnio{
bandera:=false

for a:=0;a<len(datos.Pedidos);a++{
	fmt.Println(a)
	for b:=0;b<len(Arreglo);b++{

		if Arreglo[b].Tamanio>0{
			aux:=Arreglo[b].Inicio
			for aux!=nil{

				if strings.EqualFold(datos.Pedidos[a].Tienda,aux.Datos.Nombre)&&strings.EqualFold(datos.Pedidos[a].Departamento,aux.Datos.Columna)&&datos.Pedidos[a].Calificacion==aux.Datos.Calificacion{


					var confirmar []int
					for c:=0;c<len(datos.Pedidos[a].Productos);c++{
						aux2:=aux.Datos.Inventario.Raiz
						for aux2!=nil{

							if aux2.Datos.Codigo>datos.Pedidos[a].Productos[c].Codigo{
								aux2=aux2.Izquierda
							}else if aux2.Datos.Codigo<datos.Pedidos[a].Productos[c].Codigo{
								aux2=aux2.Derecha
							}else{
								confirmar=append(confirmar,1)
								aux2=aux2.Derecha
							}
						}
					}
						if len(confirmar)==len(datos.Pedidos[a].Productos){


							bandera=true

						}
				confirmar=nil
				}
				aux=aux.Siguiente

			}
		}
	}
if bandera{


	fecha:=strings.Split(datos.Pedidos[a].Fecha,"-")
	dia,_:=strconv.Atoi(fecha[0])
	mes,_:=strconv.Atoi(fecha[1])
	anio,_:=strconv.Atoi(fecha[2])
nuevo:=MatrizDispersa.Pedido{
	Siguiente:    nil,
	Dia:          dia,
	Mes:          SeleccionMesIngles(mes),
	Anio:         anio,
	Cliente:      "",
	Direccion:    "Guatemala",
	NombreTienda: datos.Pedidos[a].Tienda,
	Departamento: datos.Pedidos[a].Departamento,
	Calificacion: datos.Pedidos[a].Calificacion,
	Descripcion:  nil,
}


var registro []Producto
for m:=0;m<len(datos.Pedidos[a].Productos);m++{
	if len(registro)<0{
		nuevo2:=Producto{
			Codigo:   datos.Pedidos[a].Productos[m].Codigo,
			Cantidad: 1,
		}
		registro=append(registro,nuevo2)
	}else{
		bandera2:=true
		for t:=0;t<len(registro);t++{
			if registro[t].Codigo==datos.Pedidos[a].Productos[m].Codigo{
				registro[t].Cantidad+=1
				bandera2=false
			}
		}
		if bandera2{
			nuevo3:=Producto{
				Codigo:  datos.Pedidos[a].Productos[m].Codigo ,
				Cantidad: 1,
			}
			registro=append(registro,nuevo3)
		}
	}
}
for v:=0;v<len(registro);v++{

	recupera:=MatrizDispersa.DescripcionPedido{
		Codigo:   registro[v].Codigo,
		Cantidad: registro[v].Cantidad,
	}
	nuevo.Descripcion=append(nuevo.Descripcion,recupera)
}
fmt.Println(nuevo)
pedidos=AVLAnio.Insertar(pedidos,nuevo)

}

bandera=false

}


return pedidos
}

type Producto struct {
	Codigo int
	Cantidad int
}

func SeleccionMesIngles(mes int)string{
if mes==01{
	return "January"
}else if mes==2{
	return"February"
}else if mes==3{
	return"March"
}else if mes==4{
	return"April"
}else if mes==5{
	return"May"
}else if mes==6{
	return"June"
}else if mes==7{
	return"July"
}else if mes==8{
	return"August"
}else if mes==9{
	return"September"
}else if mes==10{
	return"October"
}else if mes==11{
	return"November"
}else if mes==12{
	return"December"
}else{
	return "nil"
}
}