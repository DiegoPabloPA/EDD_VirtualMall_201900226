import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient,HttpClientModule,HttpHeaders } from "@angular/common/http";
import { serverGo } from "../Server/serverGo";

@Injectable({
  providedIn: 'root'
})
export class TiendaService {

  constructor(private http:HttpClient) { }

  postCargatiendas(archivo):Observable<any>{
    const httpOptions={
      headers:new HttpHeaders({
        'Content-Type':'application/json'
      }),
    };
    return this.http.post<any>(serverGo+'TiendaEspecifica',archivo,httpOptions)

  }
  postComentarioPrincipalTienda(archivo):Observable<any>{
    const httpOptions={
      headers:new HttpHeaders({
        'Content-Type':'application/json'
      }),
    };
    return this.http.post<any>(serverGo+'ComentarioTienda',archivo,httpOptions)

  }
  postSubComentario(archivo):Observable<any>{
    const httpOptions={
      headers:new HttpHeaders({
        'Content-Type':'application/json'
      }),
    };
    return this.http.post<any>(serverGo+'agregarSubComentario',archivo,httpOptions)

  }
  postComentarioPrincipalArticulo(archivo):Observable<any>{
    const httpOptions={
      headers:new HttpHeaders({
        'Content-Type':'application/json'
      }),
    };
    return this.http.post<any>(serverGo+'ComentarioArticulo',archivo,httpOptions)

  }
  postSubComentarioArticulo(archivo):Observable<any>{
    const httpOptions={
      headers:new HttpHeaders({
        'Content-Type':'application/json'
      }),
    };
    return this.http.post<any>(serverGo+'agregarSubComentarioArticulo',archivo,httpOptions)

  }




  postCargarInventario(archivo):Observable<any>{
    const httpOptions={
      headers:new HttpHeaders({
        'Content-Type':'application/json'
      }),
    };
    return this.http.post<any>(serverGo+'EnvioInventario',archivo,httpOptions)

  }
  postUbicarTienda(archivo):Observable<any>{
    const httpOptions={
      headers:new HttpHeaders({
        'Content-Type':'application/json'
      }),
    };
    return this.http.post<any>(serverGo+'ubicaTienda',archivo,httpOptions)

  }
  postCargarInventarioIndividual(archivo):Observable<any>{
    const httpOptions={
      headers:new HttpHeaders({
        'Content-Type':'application/json'
      }),
    };
    return this.http.post<any>(serverGo+'subirInventarioIndividual',archivo,httpOptions)

  }
  postAgregarCarrito(archivo):Observable<any>{
    const httpOptions={
      headers:new HttpHeaders({
        'Content-Type':'application/json'
      }),
    };
    return this.http.post<any>(serverGo+'AgregarCarrito',archivo,httpOptions)

  }

 



}
