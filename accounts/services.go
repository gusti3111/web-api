package accounts

import "encoding/json"

type Services interface {
	FindAll() ([]Person, error)
	FindByID(ID int) (Person, error)
	Create(user User) (Person, error)
	Update(ID int, user User) (Person, error)
}

type services struct {
	repo Repo
}

func NewServices(repo Repo) *services {
	return &services{repo}

}

func (s *services) FindAll() ([]Person, error) {
	return s.repo.FindAll()

}

func (s *services) FindByID(ID int) (Person, error) {
	person, err := s.repo.FindByID(ID)
	return person, err

}

func (s *services) Create(user User) (Person, error) {
	pass, _ := json.Marshal(user.Password)

	person := Person{
		Name:     user.Name,
		Password: string(pass),
	}
	newperson, err := s.repo.Create(person)
	return newperson, err

}
func (s *services) Update(ID int,user User) (Person, error) {
	user,err := s.repo.FindByID(ID)
	pass, _ := json.Marshal(user.Password)

	user.Name=     user.Name
	user.Password = string(pass)
	newperson, err := s.repo.Update(user)
	return newperson, err

}



