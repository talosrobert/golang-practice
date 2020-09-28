package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/talosrobert/golang-practice/pkg/products"
)

// Product handler object
type Product struct {
	l *log.Logger
}

// NewProduct Product handler object constructor
func NewProduct(l *log.Logger) *Product {
	return &Product{
		l: l,
	}
}

// Get returns all the products from our datastore
func (p *Product) Get(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Product Handler GET")

	// fetch product list
	pl := products.GetProducts()
	if err := pl.ToJSON(rw); err != nil {
		http.Error(rw, "unable to encode list of products to json", http.StatusInternalServerError)
	}
}

func (p *Product) Add(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Product Handler POST")

	prod := r.Context().Value(KeyProduct{}).(*products.Product)

	products.AddProduct(prod)
}

func (p *Product) Update(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "failed to convert id to integer", http.StatusBadRequest)
		return
	}

	p.l.Println("Product Handler PUT")

	prod := r.Context().Value(KeyProduct{}).(*products.Product)

	err = products.UpdateProduct(id, prod)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}
}

type KeyProduct struct{}

func (p *Product) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &products.Product{}
		if err := prod.FromJSON(r.Body); err != nil {
			http.Error(rw, "unable to unmarshal json object", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
