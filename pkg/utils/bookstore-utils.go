package utils

import 
(
	"encoding/json"
	
	"net/http"
)
func ParseBody(r *http.Request, dst interface{}) (interface{}, error) {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(dst); err != nil {
		return nil, err
	}
	return dst, nil // Yes, in the success case, we return the parsed body (dst)
}
