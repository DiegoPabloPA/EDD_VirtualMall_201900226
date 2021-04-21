import { Injectable } from '@angular/core';
import { HttpClient,HttpClientModule,HttpHeaders } from "@angular/common/http";
import { serverGo } from "../Server/serverGo";
import { Observable } from 'rxjs';
@Injectable({
  providedIn: 'root'
})
export class NuevoUsuarioService {

  constructor(private http:HttpClient) { }


  postGenerarUsuario(res:any):Observable<any>{ 
    const httpOptions={ headers:new HttpHeaders({ 'Content-Type':'application/json' }), 
  }; return this.http.post<any>(serverGo+'RegistrarUsuario',res,httpOptions) }
}
