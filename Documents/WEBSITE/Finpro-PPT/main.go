package finproppt

import (
	"net/http"

	"github.com/novalsh/go-finpro/controller/accountscontroller"
)

func main() {
	http.HandleFunc("/", accountscontroller.Index)
	http.HandleFunc("/buku", accountscontroller.Index)
	http.HandleFunc("/buku/index", accountscontroller.Index)
	http.HandleFunc("/buku/add", accountscontroller.Add)
	http.HandleFunc("/buku/edit", accountscontroller.Edit)
	http.HandleFunc("/buku/delete", accountscontroller.Delete)

	http.ListenAndServe(":9000", nil)
}