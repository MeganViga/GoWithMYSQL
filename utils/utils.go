package utils

import (
	"encoding/json"

	"io/ioutil"

	"net/http"
)

func ParseBody(r *http.Request, x interface{}){
	// body, _ := ioutil.ReadAll(r.Body)
	// fmt.Println(r.Body,[]byte(body))
	// _ = json.Unmarshal([]byte(body),x)
	// fmt.Println(x)
	if body, err := ioutil.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal([]byte(body),x); err != nil{
			return
		}
	}
}