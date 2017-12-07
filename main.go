package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	b := bytes.NewBufferString("Hello there cowboy")
	_, err := w.Write(b.Bytes())
	if err != nil {
		fmt.Print("herpderp")
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	err := http.ListenAndServe(":80", router)
	if err != nil {
		fmt.Printf("Err: %s \n", err.Error())
	}
}
