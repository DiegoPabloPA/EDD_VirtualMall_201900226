import { Component, OnInit } from '@angular/core';
import {  FormControl} from "@angular/forms";
import {  UsuariosmasivoService} from "../../servicios/usuariosmasivo.service";
@Component({
  selector: 'app-carga-usuarios',
  templateUrl: './carga-usuarios.component.html',
  styleUrls: ['./carga-usuarios.component.css']
})
export class CargaUsuariosComponent implements OnInit {

  constructor(private cargaUsuario:UsuariosmasivoService) { }
  informacion=new FormControl('')
  mostrarMensaje=false
  mostrarMensajeError=false
  ngOnInit(): void {
  }
  prueba(){
    this.cargaUsuario.postCargaUsuarios(this.informacion.value).subscribe((res:any)=>{
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
}
