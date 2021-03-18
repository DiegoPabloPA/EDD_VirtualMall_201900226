import { Injectable } from '@angular/core';
import { HttpClient,HttpHeaders } from '@angular/common/http';
import { serverGo } from "../Server/serverGo";
import { Observable } from 'rxjs';


@Injectable({
  providedIn: 'root'
})
export class MostrartiendasdisponiblesService {

  constructor(private http:HttpClient) {}
 
  obtenerTiendasDisponibles():Observable<any>{
    const httpOptions={
      headers:new HttpHeaders({
        'Content-Type':'application/json'
      }),
    };
    return this.http.get<any>(serverGo+'JsonFrontEnd',httpOptions)
  }
  
}
