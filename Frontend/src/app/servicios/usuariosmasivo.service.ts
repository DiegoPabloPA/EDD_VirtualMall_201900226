import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient,HttpClientModule,HttpHeaders } from "@angular/common/http";
import { serverGo } from "../Server/serverGo";
@Injectable({
  providedIn: 'root'
})
export class UsuariosmasivoService {

  constructor(private http:HttpClient) { }

  postCargaUsuarios(archivo):Observable<any>{
    const httpOptions={
      headers:new HttpHeaders({
        'Content-Type':'application/json'
      }),
    };
    return this.http.post<any>(serverGo+'UsuarioMasivo',archivo,httpOptions)

  }

}
