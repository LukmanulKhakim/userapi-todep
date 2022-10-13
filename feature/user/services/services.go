package services

import (
	"errors"
	"strings"
	"userapi/feature/user/domain"

	"github.com/labstack/gommon/log"
)

type repoService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &repoService{
		qry: repo,
	}
}

// AddUser implements domain.Service
func (rs *repoService) AddUser(newUser domain.Core) (domain.Core, error) {
	res, err := rs.qry.Insert(newUser)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Core{}, errors.New("rejected from database")
		}
		return domain.Core{}, errors.New("some problem on database")
	}
	return res, nil
}

// ShowAllUser implements domain.Service
func (rs *repoService) ShowAllUser() ([]domain.Core, error) {
	res, err := rs.qry.GetAll()
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return nil, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return nil, errors.New("no data")
		}
	}
	if len(res) == 0 {
		log.Info("no data")
		return nil, errors.New("no data")
	}
	return res, nil
}

// Profile implements domain.Service
func (rs *repoService) Profile(ID uint) (domain.Core, error) {
	res, err := rs.qry.Get(ID)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return domain.Core{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.Core{}, errors.New("no data")
		}
	}
	return res, nil
}

// UpdateProfile implements domain.Service
func (rs *repoService) UpdateProfile(updatedData domain.Core, ID uint) (domain.Core, error) {
	res, err := rs.qry.Update(updatedData, ID)
	if err != nil {
		if strings.Contains(err.Error(), "column") {
			return domain.Core{}, errors.New("rejected from database")
		}
	}
	return res, nil
}

// DeleteUser implements domain.Service
func (rs *repoService) DeleteUser(ID uint) (domain.Core, error) {
	res, err := rs.qry.Delete(ID)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "column") {
			return domain.Core{}, errors.New("rejected from database")
		}
	}
	return res, nil
}
