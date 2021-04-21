import { Injectable } from '@angular/core';
import { HttpClient,HttpClientModule,HttpHeaders } from "@angular/common/http";
import { serverGo } from "../Server/serverGo";
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AutentificacionService {

  constructor(private http:HttpClient) { }

logueado(){
  return localStorage.getItem('tipoUsuario')
}

postBuscarUsuario(archivo):Observable<any>{
  const httpOptions={
    headers:new HttpHeaders({
      'Content-Type':'application/json'
    }),
  };
  return this.http.post<any>(serverGo+'BuscarUsuario',archivo,httpOptions)

}



}
