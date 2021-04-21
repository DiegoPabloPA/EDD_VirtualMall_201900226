import { Component, OnInit } from '@angular/core';
import {BuscarLogin} from "../../modelos/busqueda";
import { FormControl } from "@angular/forms";
import { AutentificacionService } from "../../servicios/autentificacion.service";
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  constructor(private usuario:AutentificacionService) { }

  dpi_=new FormControl('')
  contra=new FormControl('')
  ngOnInit(): void {
  }
  consolidar(){
    localStorage.setItem("tipoUsuario","Normal")
    localStorage.setItem("Logueado","true")
  }
  Loguear(){
    const requisa:BuscarLogin={
      Dpi:this.dpi_.value,
      Password:this.contra.value
    }

    this.usuario.postBuscarUsuario(requisa).subscribe((res:any)=>{
      console.log(res)
      if (res.DPI===-1){
        window.location.href="/Login"
      }else if (res.Cuenta==="Admin"){
        localStorage.setItem("tipoUsuario","Admin")
        localStorage.setItem("Logueado","true")
        localStorage.setItem("NombreUsuario",res.Nombre)
        localStorage.setItem("Correo",res.Correo)
        localStorage.setItem("DPI",res.DPI)
        window.location.href="/"
      }else if(res.Cuenta==="Usuario"){
        localStorage.setItem("tipoUsuario","Normal")
        localStorage.setItem("Logueado","true")
        localStorage.setItem("NombreUsuario",res.Nombre)
        localStorage.setItem("Correo",res.Correo)
        localStorage.setItem("DPI",res.DPI)
        window.location.href="/"
      }


     
      


      },(err)=>{
        
      }
      
      )
    

  }




}
