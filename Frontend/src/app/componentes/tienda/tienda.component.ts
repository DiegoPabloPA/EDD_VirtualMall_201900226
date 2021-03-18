import { Component, OnInit } from '@angular/core';
import { TiendaService } from "../../servicios/tienda.service";
import { Busqueda } from "../../modelos/busqueda";
import { FormControl } from '@angular/forms';

@Component({
  selector: 'app-tienda',
  templateUrl: './tienda.component.html',
  styleUrls: ['./tienda.component.css']
})
export class TiendaComponent implements OnInit {

  constructor(private tienda:TiendaService) { 
   
  }
  textoarea=new FormControl('')
  informacion:String[];
  informacion2:String[];
  inventarios:any;
  respuesta:any;
  Logo:String="Logo";
  ngOnInit(): void {
    const bus:Busqueda={
      Departamento:localStorage.getItem('Categoria'),
      Nombre:localStorage.getItem('Nombre'),
      Calificacion:parseInt(localStorage.getItem('calificacion'))
    }



    this.tienda.postCargatiendas(bus).subscribe((res:any)=>{
      this.informacion=Object.keys(res)
      this.respuesta=res

      },(err)=>{
        
      }
      
      )


      this.tienda.postCargarInventario(bus).subscribe((res:any)=>{
        this.informacion2=Object.keys(res)
        
        this.inventarios=res
        console.log(this.inventarios)
        },(err)=>{
          
        }
        
        )
    
  }
 
  imprimir(){
    console.log(localStorage.getItem('Nombre'))
  }

  enviarInventario(){
    const bus:Busqueda={
      Departamento:localStorage.getItem('Categoria'),
      Nombre:localStorage.getItem('Nombre'),
      Calificacion:parseInt(localStorage.getItem('calificacion'))


    }
    this.tienda.postUbicarTienda(bus).subscribe((res:any)=>{
      
      },(err)=>{
        
      }
      
      )

      this.tienda.postCargarInventarioIndividual(this.textoarea.value).subscribe((res:any)=>{
      
      },(err)=>{
        
      }
      
      )
      
      window.location.href="/tienda"
  }

}
