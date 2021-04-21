import { Injectable } from '@angular/core';
import { HttpClient,HttpClientModule,HttpHeaders } from "@angular/common/http";
import { serverGo } from "../Server/serverGo";
import { Observable } from 'rxjs';
@Injectable({
  providedIn: 'root'
})
export class GestionusersService {

  constructor(private http:HttpClient) { }
  getArbolNormalB():Observable<any>{
    const httpOptions={
      headers:new HttpHeaders({
        'Content-Type':'application/json'
      }),
    };
    return this.http.get<any>(serverGo+'GenerarUsuariosNormal',httpOptions)

  }

  getArbolesCifradosB(clave:string):Observable<any>{
    const httpOptions={
      headers:new HttpHeaders({
        'Content-Type':'application/json'
      }),
    };
    return this.http.get<any>(serverGo+'Encriptacion/'+clave,httpOptions)

  }

}
