package resource

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterProductResource(r *mux.Router) {
	r.HandleFunc("/products", GetProducts).Methods("GET")
}

func GetProducts(w http.ResponseWriter, r *http.Request) {

}
