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
import { SubidagrafoComponent } from "./componentes/subidagrafo/subidagrafo.component";
import { CargaMasivaPedidosComponent } from "./componentes/carga-masiva-pedidos/carga-masiva-pedidos.component";
import {  LoginComponent} from "./componentes/login/login.component";
import { CargaUsuariosComponent } from "./componentes/carga-usuarios/carga-usuarios.component";
import { AuthGuard } from "./auth.guard";
import { RegistroComponent } from "./componentes/registro/registro.component";
import { GestionusuariosComponent } from "./componentes/gestionusuarios/gestionusuarios.component";
const routes: Routes = [
  {path:'',
  component:StartComponent,
},
{
  path:"Registro",
  component:RegistroComponent,
},
{
  path:"Login",
  component:LoginComponent,
},
{
  path:"cargarGrafo",
  component:SubidagrafoComponent,
  canActivate:[AuthGuard],
data:{tipoUser:['Admin']}
  
  
 
},
{
  path:"EncriptarUsuarios",
  component:GestionusuariosComponent,
  canActivate:[AuthGuard],
data:{tipoUser:['Admin']}
  
  
 
},
{
  path:"cargarUsuarios",
  component:CargaUsuariosComponent,
  canActivate:[AuthGuard],
data:{tipoUser:['Admin']}
  
  
 
},

{
  path:"cargarTiendas",
  component:CargartiendasComponent,
  canActivate:[AuthGuard],
data:{tipoUser:['Admin']}
  
  
 
},
{
  path:"cargarPedidosMasivo",
  component:CargaMasivaPedidosComponent,
  canActivate:[AuthGuard],
data:{tipoUser:['Admin']}
},

{
  path:"cargarInventarios",
  component:CargainventariosComponent,
  canActivate:[AuthGuard],
data:{tipoUser:['Admin']}
},
{
  path:"tiendasDisponibles",
  component:TiendasdisponiblesComponent,
  canActivate:[AuthGuard],
data:{tipoUser:['Normal']}
},
{
  path:"tienda",
  component:TiendaComponent,
  canActivate:[AuthGuard],
data:{tipoUser:['Normal']}
},
{
  path:"confirmarPedido",
  component:ConfirmarcarritoComponent,
  canActivate:[AuthGuard],
data:{tipoUser:['Normal']}
},
{
  path:"InformePedidos",
  component:InformesComponent,
  canActivate:[AuthGuard],
data:{tipoUser:['Admin']}

}


];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
