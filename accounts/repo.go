package accounts

import "gorm.io/gorm"

type Repo interface {
	FindAll() ([]Person, error)
	FindByID(ID int) (Person, error)
	Create(person Person) (Person, error)
	Update(person Person) (Person, error)
	Delete(person Person)(Person,error)
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo {
	return &repo{db}

}

func (r *repo) FindAll() ([]Person, error) {
	var persons []Person

	err := r.db.Find(&persons).Error

	return persons, err
}

func (r *repo) FindByID(ID int) (Person, error) {
	var person Person

	err := r.db.Find(&person, ID).Error

	return person, err

}

func (r *repo) Create(person Person) (Person, error) {
	err := r.db.Create(&person).Error
	return person, err

}
func (r *repo) Update(person Person) (Person, error) {
	err := r.db.Save(&person).Error
	return person, err
}
func (r *repo) Delete(person Person) (Person, error) {
	err := r.db.Delete(&person).Error
	return person, err
}
