import { Component, OnInit } from '@angular/core';
import { FormControl, FormControlName } from '@angular/forms';
import { DatosgrafoService } from "../../servicios/datosgrafo.service";
@Component({
  selector: 'app-subidagrafo',
  templateUrl: './subidagrafo.component.html',
  styleUrls: ['./subidagrafo.component.css']
})
export class SubidagrafoComponent implements OnInit {

  constructor(private cargarGrafo:DatosgrafoService) { }
  informacion=new FormControl('')
  mostrarMensaje=false
  mostrarMensajeError=false
  ngOnInit(): void {
  }
  prueba(){
    this.cargarGrafo.postCargaGrafo(this.informacion.value).subscribe((res:any)=>{
      this.mostrarMensaje=true
      console.log("Si se envio")
      this.informacion.setValue("")
    },(err)=>{
      this.mostrarMensajeError=true
      console.log(err)
    }
    
    )
  }



  desactivarMensajes(){
    this.mostrarMensaje=false
    this.mostrarMensajeError=false
  }
}
