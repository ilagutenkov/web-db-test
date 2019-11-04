package utils

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) (map[string]interface{}) {
	return map[string]interface{} {"status" : status, "message" : message}
}

func Respond(w http.ResponseWriter, data map[string] interface{})  {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}


type Container []interface {
}

func (c *Container) Put(elem interface{}) {
	*c = append(*c, elem)
}



func (c Container)  Filer(predicate func(interface{}) bool) Container {

	vsf := make([]interface{}, 0)

	for _, value := range c {
		if(predicate(value)){
			vsf=append(vsf,value)
		}
	}

	return vsf
}