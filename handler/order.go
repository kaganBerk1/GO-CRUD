package handler

import (
	"fmt"
	"net/http"
)

type Order struct {
	name string
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create ORDER")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List All Orders")
}
func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Order by ID")
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Order by ID")
}
func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Order by ID")
}
