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
export class DiaCategoria{
    
    Anio:number
    Mes:string
    
    constructor(_an:number,_me:string){
        this.Anio=_an
        this.Mes=_me
       
    }
}
export class DiaInforme{
    
    Anio:number
    Mes:string
    Dia:number
    Categoria:string
    
    constructor(_an:number,_me:string,_di:number,_ca:string){
        this.Anio=_an
        this.Mes=_me
        this.Dia=_di
        this.Categoria=_ca
       
    }
}

export class NuevoUsuario{
    
    Dpi:number
    Nombre:string
    Correo:string
    Password:string
    Cuenta:string

    constructor(dpi_:number,nom:string,cor:string,pas:string,cue:string){
        this.Dpi=dpi_
        this.Nombre=nom
        this.Correo=cor
        this.Password=pas
        this.Cuenta=cue
    }
}

export class BuscarLogin{
    Dpi:number
    Password:string

    constructor(dpi:number,pass:string){
        this.Dpi=dpi
        this.Password=pass

    }
}

export class Nombre{
    Nombre:string
    constructor(nom:string){
        this.Nombre=nom
    }
}
export class TrackNombre{
    Imagen:string
    constructor(nom:string){
        this.Imagen=nom
    }
}