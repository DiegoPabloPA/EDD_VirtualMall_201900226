package Inform

import (
	vector "Proyecto1/ListaDVec"
	"encoding/json"
)

func TomarInfo(mensaje []byte)vector.InfoVector{
	var deco vector.InfoVector

	if err := json.Unmarshal(mensaje, &deco); err != nil {
		panic(err)
	}
	return deco
}
