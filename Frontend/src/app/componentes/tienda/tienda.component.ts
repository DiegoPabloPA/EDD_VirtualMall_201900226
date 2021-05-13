import { Component, OnInit } from '@angular/core';
import { TiendaService } from "../../servicios/tienda.service";
import { Busqueda, EnvioCompra,ComentarioTienda,SubComentarioTienda, ComentarioArticulo, SubComentarioArticulo } from "../../modelos/busqueda";
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
  textoareacomentarioTienda=new FormControl('')
  textoareacomentarioArticulo=new FormControl('')
  textosubcomentarioTienda=new FormControl('')
  textosubcomentarioArticulo=new FormControl('')
  pInputCant:string[]=[];
  informacion:String[];
  categorias:String[];
  informacion2:String[];
  inventarios:any;
  respuesta:any;
  sub:number;
  indicador1:number;
  indicador2:number;
  almacenarLLaves:string;
  Logo:String="Logo";
  Com:String="Comentarios";
  mostrarMensaje=false
  mostrarMensajeError=false
  escritura:string="";
  codigo:string="";

  ComentarioTienda(){
    const coment:ComentarioTienda={
      DPI:parseInt(localStorage.getItem('DPI')),
      NombreSec:localStorage.getItem('Nombre'),
      Categoria:localStorage.getItem('Categoria'),
      Calificacion:parseInt(localStorage.getItem('calificacion')),
      Nombre:localStorage.getItem('NombreUsuario'),
      Comentario:this.textoareacomentarioTienda.value

    }
    this.tienda.postComentarioPrincipalTienda(coment).subscribe((res:any)=>{
      this.textoareacomentarioTienda.setValue("")
      window.location.reload()
      },(err)=>{
        
      }
      
      )
  }
  ComentarioPrincipalArticulo(codigo:number){
    const coment:ComentarioArticulo={
      Codigo:codigo,
      DPI:parseInt(localStorage.getItem('DPI')),
      NombreSec:localStorage.getItem('Nombre'),
      Categoria:localStorage.getItem('Categoria'),
      Calificacion:parseInt(localStorage.getItem('calificacion')),
      Nombre:localStorage.getItem('NombreUsuario'),
      Comentario:this.textoareacomentarioArticulo.value

    }
    this.tienda.postComentarioPrincipalArticulo(coment).subscribe((res:any)=>{
      this.textoareacomentarioArticulo.setValue("")
      window.location.reload()
      },(err)=>{
        
      }
      
      )
  }



  SubComentarioTienda(clav:number[]){
    
    const coment:SubComentarioTienda={
      Clave:clav.join(';'),
      DPI:parseInt(localStorage.getItem('DPI')),
      NombreSec:localStorage.getItem('Nombre'),
      Categoria:localStorage.getItem('Categoria'),
      Calificacion:parseInt(localStorage.getItem('calificacion')),
      Nombre:localStorage.getItem('NombreUsuario'),
      Comentario:this.textosubcomentarioTienda.value

    }
    console.log(coment)
    this.tienda.postSubComentario(coment).subscribe((res:any)=>{
      this.textosubcomentarioTienda.setValue("")
      window.location.reload()
      },(err)=>{
        
      }
      
      )
  }
  SubComentarioArticulo(cod:number,clav:number[]){
    
    const coment:SubComentarioArticulo={
      Codigo:cod,
      Clave:clav.join(';'),
      DPI:parseInt(localStorage.getItem('DPI')),
      NombreSec:localStorage.getItem('Nombre'),
      Categoria:localStorage.getItem('Categoria'),
      Calificacion:parseInt(localStorage.getItem('calificacion')),
      Nombre:localStorage.getItem('NombreUsuario'),
      Comentario:this.textosubcomentarioArticulo.value

    }
    console.log(coment)
    this.tienda.postSubComentarioArticulo(coment).subscribe((res:any)=>{
      this.textosubcomentarioArticulo.setValue("")
      window.location.reload()
      },(err)=>{
        
      }
      
      )
  }





     botonAnidacion(m:string){
       console.log(m)
     }
 
  


  



  ngOnInit(): void {
    const bus:Busqueda={
      Departamento:localStorage.getItem('Categoria'),
      Nombre:localStorage.getItem('Nombre'),
      Calificacion:parseInt(localStorage.getItem('calificacion'))
      

    }


    

      

    this.tienda.postCargatiendas(bus).subscribe((res:any)=>{
      this.informacion=Object.keys(res)
      this.respuesta=res
      console.log(this.respuesta)
      

      },(err)=>{
        
      }
      
      )


      this.tienda.postCargarInventario(bus).subscribe((res:any)=>{
        this.informacion2=Object.keys(res)
        
        this.inventarios=res
        console.log(res)
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
