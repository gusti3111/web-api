package accounts

import (
	"fmt"
)

type fileRepo struct {
}

func NewfileRepo() *fileRepo {
	return &fileRepo{}

}

func (r *fileRepo) FindAll() ([]Person, error) {
	var persons []Person

	fmt.Println("FindAll")
	return persons, nil
}

func (r *fileRepo) FindByID(ID int) (Person, error) {
	var person Person

	fmt.Println("FindByID")

	return person, nil

}

func (r *fileRepo) Create(person Person) (Person, error) {
	fmt.Println("Create")
	return person, nil

}
