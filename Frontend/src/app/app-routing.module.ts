import { Component, NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { StartComponent } from "./componentes/start/start.component";
import {  CargartiendasComponent} from "./componentes/cargartiendas/cargartiendas.component";
import {  CargainventariosComponent} from "./componentes/cargainventarios/cargainventarios.component";
import { TiendasdisponiblesComponent } from "./componentes/tiendasdisponibles/tiendasdisponibles.component";
import { TiendaComponent } from "./componentes/tienda/tienda.component";
const routes: Routes = [
  {path:'',
  component:StartComponent,
},{
  path:"cargarTiendas",
  component:CargartiendasComponent,
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
}


];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
