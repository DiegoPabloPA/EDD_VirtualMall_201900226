package GenerarGraficoReportes

import (
	"Proyecto1/ReportesDiaCategoria"
	"strconv"
)

func GenerarReporteDia(info []ReportesDiaCategoria.RespuestaDiaSolicitado)string{
	archie:="digraph dot{\n"
	archie+="n[shape=none label=<<table><tr>\n<td>Categoria</td>\n<td>Direccion</td>\n<td>Cantidad</td>\n<td>Codigo</td>\n</tr>"
	for a:=0;a<len(info);a++{
		archie+="\n<tr>\n"
		archie+="<td>\n"
		archie+=info[a].Categoria
		archie+="</td>\n"
		archie+="<td>\n"
		archie+=info[a].Direccion
			archie+="</td>\n"
		archie+="<td>\n"

			for b:=0;b<len(info[a].Descripcion);b++{
				archie+=strconv.Itoa(info[a].Descripcion[b].Cantidad)+";"
			}

			archie+="</td>\n"
		archie+="<td>\n"

		for b:=0;b<len(info[a].Descripcion);b++{
			archie+=strconv.Itoa(info[a].Descripcion[b].Codigo)+";"
		}

		archie+="</td>\n"
		archie+="\n</tr>\n"
	}
	archie+="</table>>];\n"
	archie+="}"
	return archie
}