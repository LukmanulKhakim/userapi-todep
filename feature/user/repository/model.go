package repository

import (
	"userapi/feature/user/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string
	HP     string
	Addres string
}

func FromDomain(du domain.Core) User {
	return User{
		Model:  gorm.Model{ID: du.ID},
		Name:   du.Name,
		HP:     du.HP,
		Addres: du.Addres,
	}
}

func ToDomain(u User) domain.Core {
	return domain.Core{
		ID:     u.ID,
		Name:   u.Name,
		HP:     u.HP,
		Addres: u.Addres,
	}
}

func ToDomainArray(au []User) []domain.Core {
	var res []domain.Core
	for _, val := range au {
		res = append(res, domain.Core{ID: val.ID, Name: val.Name, HP: val.HP, Addres: val.Addres})
	}

	return res
}
