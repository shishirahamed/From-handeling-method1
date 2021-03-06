package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/request", request)
	http.HandleFunc("/features", features)
	http.HandleFunc("/docs", docs)

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8888", nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

	//fmt.Fprintln(w, `Welcome`)
}
func features(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/features.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

	//fmt.Fprintln(w, `Welcome`)
}
func docs(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/docs.gohtml")
	if err != nil {
		fmt.Println(err.Error())

	}

	ptmp.Execute(w, nil)

	//fmt.Fprintln(w, `Welcome`)
}

func request(w http.ResponseWriter, r *http.Request) {

	//r.Parsefrom()
	name := r.FormValue("name")
	company := r.FormValue("company")
	email := r.FormValue("email")

	fmt.Println(name, company, email)
	fmt.Fprintf(w, "Well-Come %s %s %s", name, company, email)

}
