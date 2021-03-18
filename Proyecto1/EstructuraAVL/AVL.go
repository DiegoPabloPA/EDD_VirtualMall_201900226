package EstructuraAVL

type AVL struct {
	Raiz *NodoArbol
}
type DatosNodoArbol struct{
	Codigo int
	Nombre,Descripcion,Imagen string
	Precio float64
	Cantidad int
}



type NodoArbol struct{
	Izquierda,Derecha *NodoArbol
	Altura int
	Datos DatosNodoArbol
}
type InventarioIndividual struct {
	Nombre      string  `json:"Nombre"`
	Codigo      int     `json:"Codigo"`
	Descripcion string  `json:"Descripcion"`
	Precio      float64 `json:"Precio"`
	Cantidad    int     `json:"Cantidad"`
	Imagen      string  `json:"Imagen"`
}
type EstructuraBusqueda struct {
	Departamento string `json:"Departamento"`
	Nombre string `json:"Nombre"`
	Calificacion int `json:"Calificacion"`

}
