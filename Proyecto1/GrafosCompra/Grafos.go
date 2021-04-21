package GrafosCompra
type EstructuraCargaNodos struct {
	Nodos []struct {
		Nombre  string `json:"Nombre"`
		Enlaces []struct {
			Nombre    string `json:"Nombre"`
			Distancia int    `json:"Distancia"`
		} `json:"Enlaces"`
	} `json:"Nodos"`
	Posicioninicialrobot string `json:"PosicionInicialRobot"`
	Entrega              string `json:"Entrega"`
}
