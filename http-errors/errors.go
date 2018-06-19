package http_errors

import (
	"net/http"
	"encoding/json"
	"log"
)
type Error struct {
	Error string `json:"error"`
}

func CheckErrorHttp(err error, w http.ResponseWriter, code int) bool {

	if err != nil {
		var er Error
		er.Error = err.Error()
		jsn, errr := json.Marshal(er)
		CheckError(errr)

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(code)
		w.Write([]byte(jsn))
		return true
	}
	return false
}


func CheckError(err error) {
	if err != nil {
		log.Println("error:", err)
	}
}
