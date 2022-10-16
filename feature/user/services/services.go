package services

import (
	"errors"
	"strings"
	"userapi/config"
	"userapi/feature/user/domain"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
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
			return domain.Core{}, errors.New(config.DUPLICATED_DATA)
		}
		return domain.Core{}, errors.New("some problem on database")
	}
	return res, nil
}

// ShowAllUser implements domain.Service
func (rs *repoService) ShowAllUser() ([]domain.Core, error) {
	res, err := rs.qry.GetAll()

	if err == gorm.ErrRecordNotFound {
		log.Error(err.Error())
		return nil, gorm.ErrRecordNotFound
	} else if err != nil {
		log.Error(err.Error())
		return nil, errors.New(config.DATABASE_ERROR)
	}

	if len(res) == 0 {
		log.Info("no data")
		return nil, errors.New(config.DATA_NOTFOUND)
	}
	return res, nil
}

// Profile implements domain.Service
func (rs *repoService) Profile(ID uint) (domain.Core, error) {
	res, err := rs.qry.Get(ID)
	if err != nil {
		log.Error(err.Error())
		if err == gorm.ErrRecordNotFound {
			return domain.Core{}, gorm.ErrRecordNotFound
		} else {
			return domain.Core{}, errors.New(config.DATABASE_ERROR)
		}
	}
	return res, nil
}

// UpdateProfile implements domain.Service
func (rs *repoService) UpdateProfile(updatedData domain.Core, ID uint) (domain.Core, error) {
	res, err := rs.qry.Update(updatedData, ID)
	if err != nil {
		log.Error(err.Error())
		if err == gorm.ErrRecordNotFound {
			return domain.Core{}, gorm.ErrRecordNotFound
		} else {
			return domain.Core{}, errors.New(config.DATABASE_ERROR)
		}
	}
	return res, nil
}

// DeleteUser implements domain.Service
func (rs *repoService) DeleteUser(ID uint) (domain.Core, error) {
	res, err := rs.qry.Delete(ID)
	if err != nil {
		log.Error(err.Error())
		if err == gorm.ErrRecordNotFound {
			return res, gorm.ErrRecordNotFound
		} else {
			return res, errors.New(config.DATABASE_ERROR)
		}
	}
	return domain.Core{}, nil
}
