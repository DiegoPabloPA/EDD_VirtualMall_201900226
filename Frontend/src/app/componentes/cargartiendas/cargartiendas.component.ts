import { Component, OnInit } from '@angular/core';
import { FormControl, FormControlName } from '@angular/forms';
import { CargartiendasService } from "../../servicios/cargartiendas.service";
@Component({
  selector: 'app-cargartiendas',
  templateUrl: './cargartiendas.component.html',
  styleUrls: ['./cargartiendas.component.css']
})
export class CargartiendasComponent implements OnInit {

  constructor(private cargatienda:CargartiendasService) { }
  
  informacion=new FormControl('')
  mostrarMensaje=false
  mostrarMensajeError=false

prueba(){
  this.cargatienda.postCargatiendas(this.informacion.value).subscribe((res:any)=>{
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

  ngOnInit(): void {
  }

}
