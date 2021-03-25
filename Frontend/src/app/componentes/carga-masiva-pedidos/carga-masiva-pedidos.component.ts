import { Component, OnInit } from '@angular/core';
import {  FormControl} from "@angular/forms";
import { PedidosMasivosService } from "../../servicios/pedidos-masivos.service";
@Component({
  selector: 'app-carga-masiva-pedidos',
  templateUrl: './carga-masiva-pedidos.component.html',
  styleUrls: ['./carga-masiva-pedidos.component.css']
})
export class CargaMasivaPedidosComponent implements OnInit {

  constructor(private cargaMInventario:PedidosMasivosService) { }
  informacion=new FormControl('')
  mostrarMensaje=false
  mostrarMensajeError=false




  prueba(){
    this.cargaMInventario.postCargaMasivaPedidos(this.informacion.value).subscribe((res:any)=>{
      this.mostrarMensaje=true
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
