package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func router() http.Handler {
	r := mux.NewRouter()

	r.PathPrefix("/__/auth/").HandlerFunc(firebaseAuthProxy)

	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/", handleIndex)
	return r
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("error parsing template: %s \n", err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("error executing template: %s", err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
}

func firebaseAuthProxy(w http.ResponseWriter, r *http.Request) {
	firebaseProject := "test-project-7aeb1"
	target, _ := url.Parse("https://" + firebaseProject + ".firebaseapp.com")
	proxy := httputil.NewSingleHostReverseProxy(target)

	r.URL.Host = target.Host
	r.URL.Scheme = target.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = target.Host

	proxy.ServeHTTP(w, r)
}
