package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type ViewModel struct {
	ResultOfEquation string
	Equation         string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func server() {
	http.HandleFunc("/", BodServer)
	http.HandleFunc("/processbodmas", processbodmas)
	http.ListenAndServe(":8080", nil)

}

func BodServer(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "bodcal.gohtml", &ViewModel{})
}
func processbodmas(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Printf("Request equation: %v\n", r.PostForm.Get("equation"))

	store := r.PostForm.Get("equation")

	temp := strings.Split(store, "")
	cleanedParams := removeSpaces(temp)
	cleanedParams = parser(cleanedParams)
	fmt.Println(cleanedParams)
	Result := (BodmasFunc(cleanedParams))
	ResultConversion := (Result[0])
	tpl.ExecuteTemplate(w, "bodcal.gohtml", &ViewModel{
		ResultOfEquation: ResultConversion,
		Equation:         store,
	})

}
