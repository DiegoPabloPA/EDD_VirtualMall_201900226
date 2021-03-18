import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { MostrartiendasdisponiblesService } from "../../servicios/mostrartiendasdisponibles.service";


@Component({
  selector: 'app-tiendasdisponibles',
  templateUrl: './tiendasdisponibles.component.html',
  styleUrls: ['./tiendasdisponibles.component.css']
})
export class TiendasdisponiblesComponent implements OnInit {

  constructor(private tiendasdisponibles:MostrartiendasdisponiblesService) { }
  
  respuesta:any
  ngOnInit(): void {
    this.tiendasdisponibles.obtenerTiendasDisponibles().subscribe((res:any)=>{
    this.respuesta=res
    },(err)=>{
      
    }
    
    )
  }
  nombre=new FormControl('')
  categoria=new FormControl('')
  calif=new FormControl('')
 
  impresion(valor1:string,valor2:string,valor3:number){
   localStorage.setItem('Categoria',valor1);
   localStorage.setItem('Nombre',valor2);
   localStorage.setItem('calificacion',valor3.toString());
   
    window.location.href="/tienda"
    
  }
  

}
