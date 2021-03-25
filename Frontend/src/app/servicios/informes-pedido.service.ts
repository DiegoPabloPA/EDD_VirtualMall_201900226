import { Injectable } from '@angular/core';
import { HttpClient,HttpClientModule,HttpHeaders } from "@angular/common/http";
import { serverGo } from "../Server/serverGo";
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class InformesPedidoService {

  constructor(private http:HttpClient) { }

  getAVLAnios():Observable<any>{ 
    const httpOptions={ headers:new HttpHeaders({ 'Content-Type':'application/json' }), 
  }; return this.http.get<any>(serverGo+'GrafAVLAnios',httpOptions) }
  getAVLMeses():Observable<any>{ 
    const httpOptions={ headers:new HttpHeaders({ 'Content-Type':'application/json' }), 
  }; return this.http.get<any>(serverGo+'GrafAVLMeses',httpOptions) }

  getMenu():Observable<any>{ 
    const httpOptions={ headers:new HttpHeaders({ 'Content-Type':'application/json' }), 
  }; return this.http.get<any>(serverGo+'MenuAnios',httpOptions) }

  postMenu(res:any):Observable<any>{ 
    const httpOptions={ headers:new HttpHeaders({ 'Content-Type':'application/json' }), 
  }; return this.http.post<any>(serverGo+'MenuDiasCategoria',res,httpOptions) }

  postReporteDia(res:any):Observable<any>{ 
    const httpOptions={ headers:new HttpHeaders({ 'Content-Type':'application/json' }), 
  }; return this.http.post<any>(serverGo+'InformeDia',res,httpOptions) }

  
  

  
}
