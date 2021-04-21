import { Component, OnInit } from '@angular/core';
import {  FormControl} from "@angular/forms";
import {NuevoUsuario} from "../../modelos/busqueda";
import { NuevoUsuarioService } from "../../servicios/nuevo-usuario.service";
@Component({
  selector: 'app-registro',
  templateUrl: './registro.component.html',
  styleUrls: ['./registro.component.css']
})
export class RegistroComponent implements OnInit {

  constructor(private usuario:NuevoUsuarioService) { }
  nom=new FormControl('')
  corr=new FormControl('')
  dpi_=new FormControl('')
  contra=new FormControl('')

  agregarUsuario(){
    const nuevo:NuevoUsuario={
      Dpi:this.dpi_.value,
      Nombre: this.nom.value,
      Correo:this.corr.value,
      Password:this.contra.value,
      Cuenta:"Usuario"
    }

    this.usuario.postGenerarUsuario(nuevo).subscribe((res:any)=>{
      window.location.href="/Login"
      


      },(err)=>{
        window.location.href="/Registro"
        console.log(nuevo)
      }
      
      )
    
  }



  ngOnInit(): void {
  }

}
