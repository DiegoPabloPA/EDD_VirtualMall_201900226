import { Component, OnInit } from '@angular/core';
import { TiendaService } from "../../servicios/tienda.service";
import { Busqueda, EnvioCompra } from "../../modelos/busqueda";
import { FormControl } from '@angular/forms';
import { flushMicrotasks } from '@angular/core/testing';

@Component({
  selector: 'app-tienda',
  templateUrl: './tienda.component.html',
  styleUrls: ['./tienda.component.css']
})
export class TiendaComponent implements OnInit {

  constructor(private tienda:TiendaService) { 
   
  }
  textoarea=new FormControl('')

  pInputCant:string[]=[];
  informacion:String[];
  categorias:String[];
  informacion2:String[];
  inventarios:any;
  respuesta:any;
  Logo:String="Logo";
  mostrarMensaje=false
  mostrarMensajeError=false



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
        this.AsignarInfo()
        },(err)=>{
          
        }
        
        )

      

       
    
  }
  AsignarInfo(){
    this.inventarios.forEach(element => {
      this.pInputCant.push(element.Codigo)
    });
    
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
    cantidad:number

    desactivarMensajes(){
this.mostrarMensaje=false
this.mostrarMensajeError=false
    }
  imprimirInventarios(valor){
    var reco= (<HTMLInputElement>document.getElementById(valor)).value;
    var compra:EnvioCompra={
      Departamento:localStorage.getItem('Categoria'),
      Nombre:localStorage.getItem('Nombre'),
      Calificacion:parseInt(localStorage.getItem('calificacion')),
      Codigo:parseInt(valor),
      Cantidad:parseInt(reco)
    }
    this.tienda.postAgregarCarrito(compra).subscribe((res:any)=>{
      if (res.Res==="Si"){
        this.mostrarMensaje=true;
        
      }else if (res.Res==="No"){
        this.mostrarMensajeError=true
      }
    },(err)=>{
      
    }
    
    )
    
    

  }

}
