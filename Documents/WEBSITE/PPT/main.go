package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", AccountsController.Index)
	http.HandleFunc("/buku", AccountsController.Index)
	http.HandleFunc("/buku/index", AccountsController.Index)
	http.HandleFunc("/buku/add", AccountsController.Add)
	http.HandleFunc("/buku/edit", AccountsController.Edit)
	http.HandleFunc("/buku/delete", AccountsController.Delete)

	http.ListenAndServe(":8080", nil)
}
