import { Component, OnInit } from '@angular/core';
import {  FormControl} from "@angular/forms";
import {  InventarioscargaService} from "../../servicios/inventarioscarga.service";

@Component({
  selector: 'app-cargainventarios',
  templateUrl: './cargainventarios.component.html',
  styleUrls: ['./cargainventarios.component.css']
})
export class CargainventariosComponent implements OnInit {

  constructor(private cargaMInventario:InventarioscargaService) { }
  informacion=new FormControl('')
  mostrarMensaje=false
  mostrarMensajeError=false

  prueba(){
    this.cargaMInventario.postCargaMasivaInventarios(this.informacion.value).subscribe((res:any)=>{
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
