package main

import (
	"crawler/frontend/controller"
	"log"
	"net/http"
)

func main() {
	// windows 设置为绝对路径
	http.Handle("/", http.FileServer(
		http.Dir("crawler/frontend/view")))
	http.Handle("/search",
		controller.CreateSearchResultHandler(
			"crawler/frontend/view/template.html"))

	log.Println("[INFO] start server...")

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}
