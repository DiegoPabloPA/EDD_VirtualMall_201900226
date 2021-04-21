import { Component, OnInit } from '@angular/core';
import { GestionusersService } from "../../servicios/gestionusers.service";
import { FormControl } from "@angular/forms";
@Component({
  selector: 'app-gestionusuarios',
  templateUrl: './gestionusuarios.component.html',
  styleUrls: ['./gestionusuarios.component.css']
})
export class GestionusuariosComponent implements OnInit {
  condicion:boolean=false
  constructor(private Gestion:GestionusersService) { }
  codigo=new FormControl('')
  ngOnInit(): void {
    this.Gestion.getArbolNormalB().subscribe((res:any)=>{
     
    },(err)=>{
     
    }
    
    )


  }
  enviarClave(){
    this.Gestion.getArbolesCifradosB(this.codigo.value).subscribe((res:any)=>{
      this.codigo.setValue("")
     this.condicion=true
    },(err)=>{
      this.codigo.setValue("")
    }
    
    )
  }
  Descarga1(){
    window.location.href="http://localhost:3000/ArbolNormalB"
  }
  Descarga2(){
    window.location.href="http://localhost:3000/Cifrado2B"
  }
  Descarga3(){
    window.location.href="http://localhost:3000/Cifrado3B"
  }


}
