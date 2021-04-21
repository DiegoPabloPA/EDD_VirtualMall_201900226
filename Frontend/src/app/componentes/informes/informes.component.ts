import { Component, OnInit } from '@angular/core';
import { InformesPedidoService } from 'src/app/servicios/informes-pedido.service';
import { DiaCategoria,DiaInforme,TrackNombre } from "../../modelos/busqueda";
@Component({
  selector: 'app-informes',
  templateUrl: './informes.component.html',
  styleUrls: ['./informes.component.css']
})
export class InformesComponent implements OnInit {

  constructor(private informe: InformesPedidoService) { }
  CategoriaDia:any[]=[]
  anio:any[]=[]
  Track:any[]=[]
  ReporteDia:any[]=[]
  condicion:boolean
  condicion2:boolean
  titulodia:string
  titulomes:string
  tituloanio:string
  titulocate:string
  ngOnInit(): void {
    this.condicion=false
    this.informe.getTracking().subscribe((res: any) => {
     
      this.Track=res


    }, (err) => {

    }

    )



    this.informe.getAVLAnios().subscribe((res: any) => {
      



    }, (err) => {

    }

    )
    this.informe.getAVLMeses().subscribe((res: any) => {
      



    }, (err) => {

    }

    )


    this.informe.getMenu().subscribe((res: any) => {
      
      this.anio=res
      

    }, (err) => {

    }

    )

  }
  GirarOrden(anio,mes){
    this.tituloanio=anio
    this.titulomes=mes
    this.condicion2=false
    var bus:DiaCategoria={
      Anio:anio,
      Mes:mes
    }
    this.informe.postMenu(bus).subscribe((res: any) => {
      
      this.CategoriaDia=res
      this.condicion=true
      localStorage.setItem("Anio",anio)
      localStorage.setItem("Mes",mes)
    }, (err) => {

    }

    )

  }
  GenerarReporteDia(categoria,dia){
    this.condicion2=true
    this.titulodia=dia
    this.titulocate=categoria
    var bus:DiaInforme={
      
      Anio:parseInt(localStorage.getItem("Anio")),
      Mes:localStorage.getItem("Mes"),
      Dia:dia,
      Categoria:categoria
    }
    
    this.informe.postReporteDia(bus).subscribe((res: any) => {
      this.ReporteDia=res
      
      
    }, (err) => {

    }

    )

  }
  DescargarTrack(info){
    

    window.location.href="http://localhost:3000/DescargaTrack/"+info
  }
  




}
