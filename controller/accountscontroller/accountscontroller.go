package accountscontroller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/novalsh/finpro-go/entities"
	"github.com/novalsh/finpro-go/models"
)

var accountModel = models.NewAccountsModel()

func Index(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/accounts" {
		http.NotFound(response, request)
		return
	}

	// Logika lainnya untuk mengambil data atau melakukan operasi yang sesuai
	accounts, _ := accountModel.FindAll()

	// Membuat objek map dengan data yang akan diresponse
	data := map[string]interface{}{
		"accounts": accounts,
	}

	// Mengubah data menjadi JSON
	jsonResponse, err := json.Marshal(data)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mengatur header Content-Type sebagai application/json
	response.Header().Set("Content-Type", "application/json")

	// Mengirimkan respons JSON ke client
	response.Write(jsonResponse)
}

func Add(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		// ...
	} else if request.Method == http.MethodPost {
		// ...

		// Membaca body request sebagai JSON
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		// Mengubah body JSON menjadi objek Accounts
		var account entities.Accounts
		err = json.Unmarshal(body, &account)
		if err != nil {
			http.Error(response, err.Error(), http.StatusBadRequest)
			return
		}

		// Memanggil fungsi Create untuk menyimpan data
		accountModel.Create(account)

		// Mengirimkan respons JSON dengan pesan berhasil
		message := map[string]interface{}{
			"message": "Data berhasil disimpan",
		}

		// Mengubah pesan menjadi JSON
		jsonResponse, err := json.Marshal(message)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		// Mengatur header Content-Type sebagai application/json
		response.Header().Set("Content-Type", "application/json")

		// Mengirimkan respons JSON ke client
		response.Write(jsonResponse)
	}
}

func Edit(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		// Logika untuk operasi saat metode HTTP adalah GET
		// Misalnya, mengambil data dari database dan menampilkan form edit
	} else if request.Method == http.MethodPut {
		// Membaca body request sebagai JSON
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		// Mengubah body JSON menjadi objek buku
		var account entities.Accounts
		err = json.Unmarshal(body, &account)
		if err != nil {
			http.Error(response, err.Error(), http.StatusBadRequest)
			return
		}

		// Memanggil fungsi Update untuk mengupdate buku
		err = accountModel.Update(account)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		// Mengirimkan respons JSON dengan pesan berhasil
		message := map[string]interface{}{
			"message": "Data berhasil diupdate",
		}

		// Mengubah pesan menjadi JSON
		jsonResponse, err := json.Marshal(message)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		// Mengatur header Content-Type sebagai application/json
		response.Header().Set("Content-Type", "application/json")

		// Mengirimkan respons JSON ke client
		response.Write(jsonResponse)
	}
}


func Delete(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodDelete {
		// Membaca parameter ID buku yang akan dihapus
		idStr := request.URL.Query().Get("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(response, "Invalid ID", http.StatusBadRequest)
			return
		}

		// Memanggil fungsi Delete untuk menghapus buku
		err = accountModel.Delete(id)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		// Mengirimkan respons JSON dengan pesan berhasil
		message := map[string]interface{}{
			"message": "Data berhasil dihapus",
		}

		// Mengubah pesan menjadi JSON
		jsonResponse, err := json.Marshal(message)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		// Mengatur header Content-Type sebagai application/json
		response.Header().Set("Content-Type", "application/json")

		// Mengirimkan respons JSON ke client
		response.Write(jsonResponse)
	}
}
