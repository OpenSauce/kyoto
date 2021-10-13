package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	ssc "github.com/yuriizinets/ssceng"
)

func ssatemplate(p ssc.Page) *template.Template {
	return template.Must(template.New("SSA").Funcs(ssc.Funcs()).ParseGlob("*.html"))
}

func pagehandler(p ssc.Page) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ssc.PageHandlerFactory(p, map[string]interface{}{
			"internal:rw": rw,
			"internal:r":  r,
		})(rw, r)
	}
}

func ssahandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ssc.SSAHandlerFactory(ssatemplate, map[string]interface{}{
			"internal:rw": rw,
			"internal:r":  r,
		})(rw, r)
	}
}

func main() {
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/", pagehandler(&PageIndex{}))
	// SSA plugin
	mux.HandleFunc("/SSA/", ssahandler())

	// Run
	if os.Getenv("PORT") == "" {
		log.Println("Listening on localhost:25025")
		http.ListenAndServe("localhost:25025", mux)
	} else {
		log.Println("Listening on 0.0.0.0:" + os.Getenv("PORT"))
		http.ListenAndServe(":"+os.Getenv("PORT"), mux)
	}
}