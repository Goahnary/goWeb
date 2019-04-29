package main

import (
	"html/template"
	//"os"
	"strconv"
	//"text/template"
	"net/http"
	"strings"
	"getWkuTemp"
);

type TemperatureLocation struct {
	Location	string
	Temp		float64
}

func sayHello(w http.ResponseWriter, r *http.Request){
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")

	tmp, err := strconv.ParseFloat(getWkuTemp.GetTemp(), 32)
	if err != nil { panic(err) }
	wkuTemp := TemperatureLocation{"WKU", tmp}

	tmpl, err := template.New("test").Parse("<html><head><title>Noah's Go-Lang Website</title></head><body><p>Hello " + message + ". The tempurature at WKU is {{.Temp}}°F</p></body></html>")
	if err != nil { panic(err) }
	err = tmpl.Execute(w, wkuTemp)
	if err != nil { panic(err) }
	//w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
