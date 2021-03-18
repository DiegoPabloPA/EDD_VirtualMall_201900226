import { Injectable } from '@angular/core';
import { HttpClient,HttpClientModule,HttpHeaders } from "@angular/common/http";
import { serverGo } from "../Server/serverGo";
import { Observable } from 'rxjs';
@Injectable({
  providedIn: 'root'
})
export class CargartiendasService {

  constructor(private http:HttpClient) { 
    
  }
  postCargatiendas(archivo):Observable<any>{
    const httpOptions={
      headers:new HttpHeaders({
        'Content-Type':'application/json'
      }),
    };
    return this.http.post<any>(serverGo+'cargartienda',archivo,httpOptions)

  }
}
