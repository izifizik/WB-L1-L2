package cmd

import "fmt"

type Repository struct {
}

func NewRepo() Repository {
	return Repository{}
}

func (r Repository) Insert() {
	fmt.Println("Insert")
}

func (r Repository) Update() {
	fmt.Println("Update")
}

func (r Repository) Delete() {
	fmt.Println("Delete")
}

func (r Repository) Find() {
	fmt.Println("Find")
}
