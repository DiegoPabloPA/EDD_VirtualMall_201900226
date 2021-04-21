import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule, FormsModule } from "@angular/forms";
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { StartComponent } from './componentes/start/start.component';
import { CargartiendasComponent } from './componentes/cargartiendas/cargartiendas.component';
import { HttpClientModule } from "@angular/common/http";
import { CargainventariosComponent } from './componentes/cargainventarios/cargainventarios.component';
import { TiendasdisponiblesComponent } from './componentes/tiendasdisponibles/tiendasdisponibles.component';
import { TiendaComponent } from './componentes/tienda/tienda.component';
import { ConfirmarcarritoComponent } from './componentes/confirmarcarrito/confirmarcarrito.component';
import { InformesComponent } from './componentes/informes/informes.component';
import { CargaMasivaPedidosComponent } from './componentes/carga-masiva-pedidos/carga-masiva-pedidos.component';
import { LoginComponent } from './componentes/login/login.component';
import { AuthGuard } from "./auth.guard";
import { RegistroComponent } from './componentes/registro/registro.component';
import { SubidagrafoComponent } from './componentes/subidagrafo/subidagrafo.component';
import { CargaUsuariosComponent } from './componentes/carga-usuarios/carga-usuarios.component';
import { GestionusuariosComponent } from './componentes/gestionusuarios/gestionusuarios.component';
@NgModule({
  declarations: [
    AppComponent,
    StartComponent,
    CargartiendasComponent,
    CargainventariosComponent,
    TiendasdisponiblesComponent,
    TiendaComponent,
    ConfirmarcarritoComponent,
    InformesComponent,
    CargaMasivaPedidosComponent,
    LoginComponent,
    RegistroComponent,
    SubidagrafoComponent,
    CargaUsuariosComponent,
    GestionusuariosComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    ReactiveFormsModule,
    FormsModule,
    HttpClientModule
  ],
  providers: [AuthGuard],
  bootstrap: [AppComponent]
})
export class AppModule { }
