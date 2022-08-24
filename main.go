package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Alisher struct {
	ID     int    `json:"id"`
	URL    string `json:"url"`
	Method string `json:"method"`
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", Kemal)

	if err := http.ListenAndServe(":8888", mux); err != nil {
		log.Fatalln(err)
	}
}
func Kemal(w http.ResponseWriter, r *http.Request) {
	alish := Alisher{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	if err = json.Unmarshal(body, &alish); err != nil {
		fmt.Println(err)
	}

	buf := new(bytes.Buffer)

	req, err := http.NewRequest(alish.Method, alish.URL, buf)
	if err != nil {
		fmt.Println(err)
	}
	// req.Header.Add()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range resp.Header {
		fmt.Println(k, v)
	}
	// resp.ContentLength
	anon := struct {
		Status string `json:"status"`
	}{
		Status: resp.Status,
	}

	if err = json.NewEncoder(w).Encode(&anon); err != nil {
		fmt.Println(err)
	}
}
