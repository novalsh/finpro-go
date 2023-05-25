package main

import (
	"net/http"

	"github.com/novalsh/finpro-go/controller/accountscontroller"
)

func main() {
	http.HandleFunc("/", accountscontroller.Index)
	http.HandleFunc("/accounts", accountscontroller.Index)
	http.HandleFunc("/accounts/index", accountscontroller.Index)
	http.HandleFunc("/accounts/add", accountscontroller.Add)
	http.HandleFunc("/accounts/edit", accountscontroller.Edit)
	http.HandleFunc("/accounts/delete", accountscontroller.Delete)

	http.ListenAndServe(":3000", nil)
}
