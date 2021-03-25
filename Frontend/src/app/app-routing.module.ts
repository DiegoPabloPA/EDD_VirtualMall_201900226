import { Component, NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { StartComponent } from "./componentes/start/start.component";
import {  CargartiendasComponent} from "./componentes/cargartiendas/cargartiendas.component";
import {  CargainventariosComponent} from "./componentes/cargainventarios/cargainventarios.component";
import { TiendasdisponiblesComponent } from "./componentes/tiendasdisponibles/tiendasdisponibles.component";
import { TiendaComponent } from "./componentes/tienda/tienda.component";
import { ConfirmarcarritoComponent } from './componentes/confirmarcarrito/confirmarcarrito.component';
import { InformesComponent } from './componentes/informes/informes.component';
import { InformesPedidoService } from "./servicios/informes-pedido.service";
import { CargaMasivaPedidosComponent } from "./componentes/carga-masiva-pedidos/carga-masiva-pedidos.component";
const routes: Routes = [
  {path:'',
  component:StartComponent,
},{
  path:"cargarTiendas",
  component:CargartiendasComponent,
},
{
  path:"cargarPedidosMasivo",
  component:CargaMasivaPedidosComponent
},

{
  path:"cargarInventarios",
  component:CargainventariosComponent
},
{
  path:"tiendasDisponibles",
  component:TiendasdisponiblesComponent
},
{
  path:"tienda",
  component:TiendaComponent
},
{
  path:"confirmarPedido",
  component:ConfirmarcarritoComponent
},
{
  path:"InformePedidos",
  component:InformesComponent

}


];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
