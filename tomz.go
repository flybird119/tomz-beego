package main

import (
	// "github.com/astaxie/beego"
	// "tomz/controllers"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// beego.Router("/", &controllers.MainController{})
	// beego.Router("/edit", &controllers.EditController{})
	// beego.Run()

	http.HandleFunc("/", indexHandle)
	err := http.ListenAndServe(8080, handler)
	log.Fatal(err)
}

func indexHandle(rw http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("index.html", "footer.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(rw, nil)
}
