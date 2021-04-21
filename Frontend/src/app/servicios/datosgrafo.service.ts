import { Injectable } from '@angular/core';
import { HttpClient,HttpClientModule,HttpHeaders } from "@angular/common/http";
import { serverGo } from "../Server/serverGo";
import { Observable } from 'rxjs';
@Injectable({
  providedIn: 'root'
})
export class DatosgrafoService {

  constructor(private http:HttpClient) { }

  postCargaGrafo(archivo):Observable<any>{
    const httpOptions={
      headers:new HttpHeaders({
        'Content-Type':'application/json'
      }),
    };
    return this.http.post<any>(serverGo+'CargaNodosGrafo',archivo,httpOptions)

  }
}
