import { Injectable } from '@angular/core';
import { serverGo } from "../Server/serverGo";
import { HttpClient,HttpHeaders } from "@angular/common/http";
import { Observable } from 'rxjs';
@Injectable({
  providedIn: 'root'
})
export class InventarioscargaService {

  constructor(private http:HttpClient) { }

  postCargaMasivaInventarios(archivo):Observable<any>{
    const httpOptions={
      headers:new HttpHeaders({
        'Content-Type':'application/json'
      }),
    };
    return this.http.post<any>(serverGo+'inventario',archivo,httpOptions)
  }
}
