import { Component, OnInit } from '@angular/core';
import { CompraService } from "../../servicios/compra.service";
import { EliminarArticulo,Nombre } from "../../modelos/busqueda";

@Component({
  selector: 'app-confirmarcarrito',
  templateUrl: './confirmarcarrito.component.html',
  styleUrls: ['./confirmarcarrito.component.css']
})
export class ConfirmarcarritoComponent implements OnInit {

  constructor(private compra:CompraService) { }
  informacion:any[]=[]
  totalfinal:number=0
  ngOnInit(): void {

    this.compra.getInfoCompra().subscribe((res:any)=>{
      
      this.informacion=res
      res.forEach(total => {
        this.totalfinal+=total.Total
      });


      },(err)=>{
        
      }
      
      )
      var Nombre:Nombre={
        Nombre:localStorage.getItem("NombreUsuario")
      }
      this.compra.postNombreCliente(Nombre).subscribe((res:any)=>{
      
        
  
  
        },(err)=>{
          
        }
        
        )



  }
  eliminar(valor1,valor2,valor3){
    var elim:EliminarArticulo={
      Producto:valor1,
      Nombre:valor2,
      Codigo:valor3
    }
   
    this.compra.postEliminarCarrito(elim).subscribe((res:any)=>{
      
      


      },(err)=>{
        
      }
      
      )
      window.location.href="/confirmarPedido"

  }
  redireccionTiendas(){
    window.location.href="/tiendasDisponibles"
  }

  EjecutrarCompra(){
    this.compra.getConfirmarPedido().subscribe((res:any)=>{
      console.log(res)
      
      window.location.href="/confirmarPedido"

    },(err)=>{
      
    }
    
    )

    window.location.href="/confirmarPedido"
  }

}
