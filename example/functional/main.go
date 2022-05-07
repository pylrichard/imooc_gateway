package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func main() {
	resp := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", bytes.NewBuffer([]byte("test")))

	hf := HandlerFunc(HelloHandler)
	hf.ServeHTTP(resp, req)

	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}

func HelloHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello world"))
}
