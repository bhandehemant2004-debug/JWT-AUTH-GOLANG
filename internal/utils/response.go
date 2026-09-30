package utils

import (
	"encoding/json"
	"net/http"
)


func WriteJsonResponse(statuscode int , resWriter http.ResponseWriter , data map[string]any){
	resWriter.Header().Set("Content-Type","application/json")
	resWriter.WriteHeader(statuscode)
	json.NewEncoder(resWriter).Encode(data)
}