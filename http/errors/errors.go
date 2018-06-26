package errors

import (
	"encoding/json"
	"jgit.me/tools/notify_gate/utils"
	"net/http"
)

type httpError struct {
	Error string `json:"error"`
}

// CheckErrorHTTP is used for check error and write it to http.ResponseWriter, if it not nil.
// Return true if error not nil
func CheckErrorHTTP(err error, w http.ResponseWriter, code int) bool {

	if err != nil {
		var er httpError
		er.Error = err.Error()
		jsn, errr := json.Marshal(er)
		utils.CheckError(errr)

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(code)
		w.Write([]byte(jsn))
		return true
	}
	return false
}
