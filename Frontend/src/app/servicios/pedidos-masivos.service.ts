import { Injectable } from '@angular/core';
import { HttpClient,HttpHeaders } from '@angular/common/http';
import { serverGo } from "../Server/serverGo";
import { Observable } from 'rxjs';


@Injectable({
  providedIn: 'root'
})
export class PedidosMasivosService {

  constructor(private http:HttpClient) { }


  postCargaMasivaPedidos(infor):Observable<any>{
    const httpOptions={
      headers:new HttpHeaders({
        'Content-Type':'application/json'
      }),
    };
    return this.http.post<any>(serverGo+'PedidosMasivo',infor,httpOptions)
  }
}
