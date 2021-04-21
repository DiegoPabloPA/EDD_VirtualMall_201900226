import { Injectable } from '@angular/core';
import { HttpClient,HttpClientModule,HttpHeaders } from "@angular/common/http";
import { serverGo } from "../Server/serverGo";
import { Observable } from 'rxjs';
@Injectable({
  providedIn: 'root'
})
export class CompraService {

  constructor(private http:HttpClient) { }

  getInfoCompra():Observable<any>{
    const httpOptions={
      headers:new HttpHeaders({
        'Content-Type':'application/json'
      }),
    };
    return this.http.get<any>(serverGo+'InfoCompra',httpOptions)

  }
  postEliminarCarrito(infor):Observable<any>{
    const httpOptions={
      headers:new HttpHeaders({
        'Content-Type':'application/json'
      }),
    };
    return this.http.post<any>(serverGo+'EliminarCarrito',infor,httpOptions)

  }
  getConfirmarPedido():Observable<any>{
    const httpOptions={
      headers:new HttpHeaders({
        'Content-Type':'application/json'
      }),
    };
    return this.http.post<any>(serverGo+'EjecutarCompra',httpOptions)

  }
  postNombreCliente(infor):Observable<any>{
    const httpOptions={
      headers:new HttpHeaders({
        'Content-Type':'application/json'
      }),
    };
    return this.http.post<any>(serverGo+'NombreCliente',infor,httpOptions)

  }
 


}
