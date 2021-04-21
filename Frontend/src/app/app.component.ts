import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Frontend';
condicion1():Boolean{
  if(localStorage.getItem("Logueado")==="true"){
    return true
  }
  return false
}
condicion2():Boolean{
  if(localStorage.getItem("tipoUsuario")==="Admin"){
   return true
  }
  return false
}

cerrar(){
  localStorage.clear()
}

}
