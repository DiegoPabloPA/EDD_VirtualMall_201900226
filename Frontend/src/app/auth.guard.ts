import { Injectable } from '@angular/core';
import { CanActivate } from '@angular/router';
import { Observable } from 'rxjs';
import { AutentificacionService } from "./servicios/autentificacion.service";
import {  Router,ActivatedRouteSnapshot,RouterStateSnapshot} from "@angular/router";
@Injectable({
  providedIn: 'root'
})
export class AuthGuard implements CanActivate {
  constructor(private authService:AutentificacionService,private router:Router){}


  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot):boolean{
    
    const condicion=route.data['tipoUser']

    if (condicion[0]===localStorage.getItem('tipoUsuario')&&localStorage.getItem('Logueado')==="true"){

      return true
    }else{
    
    window.location.href="/Login"
    localStorage.clear()
    return false
    }
  }
  
  
  
   
  }
  

