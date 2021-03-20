export class Busqueda {
    Departamento:string
    Nombre:string
    Calificacion:number

    constructor(_dep:string,_nom:string,_cal:number){
        this.Departamento=_dep
        this.Nombre=_nom
        this.Calificacion=_cal
    }
}

export class EnvioCompra{
    Departamento:string
    Nombre:string
    Calificacion:number
    Codigo:number
    Cantidad:number
    constructor(_dep:string,_nombre:string,_cal:number,_cod:number,_cant:number){
        this.Departamento=_dep
        this.Nombre=_nombre
        this.Calificacion=_cal
        this.Codigo=_cod
        this.Cantidad=_cant
    }
}
export class EliminarArticulo{
    Producto:string
    Nombre:string
    Codigo:number
    
    constructor(_prod:string,_nom:string,_cod:number){
        this.Producto=_prod
        this.Nombre=_nom
        this.Codigo=_cod
       
    }
}
