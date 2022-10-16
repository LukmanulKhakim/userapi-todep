package services

import (
	"errors"
	"testing"
	"userapi/config"
	"userapi/feature/user/domain"
	"userapi/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAddUser(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sucses Add ", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{ID: uint(1), Name: "Lukman", HP: "123", Addres: "lmg"}, nil).Once()
		srv := New(repo)
		input := domain.Core{Name: "Lukman", HP: "123", Addres: "lmg"}
		res, err := srv.AddUser(input)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.NotEmpty(t, res.ID, "harusnya ada id yang terbuat")
		assert.Equal(t, input.Name, res.Name, "seharusnya nama sama")
		repo.AssertExpectations(t)
	})

	t.Run("Duplicate data", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{}, errors.New(config.DUPLICATED_DATA)).Once()
		srv := New(repo)
		input := domain.Core{ID: uint(1), Name: "Lukman", HP: "123", Addres: "lmg"}
		res, err := srv.AddUser(input)
		assert.NotNil(t, err)
		assert.EqualError(t, err, config.DUPLICATED_DATA, "pesan error tidak sesuai")
		assert.Empty(t, res, "karena object gagal dibuat")
		assert.Equal(t, uint(0), res.ID, "id harusnya 0 karena tidak ada data")
		repo.AssertExpectations(t)
	})
}

func TestShowAllUser(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sucses show all user", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return([]domain.Core{{ID: uint(1), Name: "Lukman", HP: "123", Addres: "lmg"}}, nil).Once()
		srv := New(repo)
		res, err := srv.ShowAllUser()
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Greater(t, res[0].ID, uint(0)) //lebih besar
		assert.GreaterOrEqual(t, len(res), 1) //lebih besar atau sama
		repo.AssertExpectations(t)
	})
	t.Run("Cant Retrive on database", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return(nil, errors.New(config.DATABASE_ERROR)).Once()
		srv := New(repo)
		res, err := srv.ShowAllUser()
		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, config.DATABASE_ERROR, "pesan error tidak sesuai")
		assert.Equal(t, len(res), 0, "len harusnya 0 karena tidak ada data")
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return(nil, gorm.ErrRecordNotFound).Once()
		srv := New(repo)
		res, err := srv.ShowAllUser()
		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, gorm.ErrRecordNotFound.Error(), "pesan error tidak sesuai")
		assert.Equal(t, len(res), 0, "len harusnya 0 karena tidak ada data")
		repo.AssertExpectations(t)
	})

}

func TestProfile(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("sucses get profile", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(domain.Core{ID: uint(1), Name: "Lukman", HP: "123", Addres: "lmg"}, nil).Once()
		srv := New(repo)
		var input = uint(1)
		res, err := srv.Profile(input)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, uint(1), res.ID, "nilai id harus sama dengan nilai input")
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(domain.Core{}, gorm.ErrRecordNotFound).Once()
		srv := New(repo)
		var input = uint(1)
		res, err := srv.Profile(input)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.EqualError(t, err, gorm.ErrRecordNotFound.Error(), "pesan error tidak sesuai")
		assert.NotEqual(t, uint(1), res.ID, "harusnya tidak sama dengan nilai id")
		repo.AssertExpectations(t)

	})

	t.Run("cant retrive data on database", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(domain.Core{}, errors.New(config.DATABASE_ERROR)).Once()
		srv := New(repo)
		var input = uint(1)
		res, err := srv.Profile(input)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.EqualError(t, err, config.DATABASE_ERROR, "pesan error tidak sesuai")
		assert.NotEqual(t, uint(1), res.ID, "harusnya tidak sama dengan nilai id")
		repo.AssertExpectations(t)
	})

}

func TestUpdateProfil(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sucses update profil", func(t *testing.T) {
		repo.On("Update", mock.Anything, mock.Anything).Return(domain.Core{ID: uint(1), Name: "Lukman", HP: "123", Addres: "lmg"}, nil).Once()
		srv := New(repo)
		var input = domain.Core{Name: "Lukman", HP: "123", Addres: "lmg"}
		var id = uint(1)
		res, err := srv.UpdateProfile(input, id)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, input.Name, res.Name, "Nama input dan hasil harus sama")
		assert.Equal(t, id, res.ID, "nilai id harus sama dengan nilai input")
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("Update", mock.Anything, mock.Anything).Return(domain.Core{}, gorm.ErrRecordNotFound).Once()
		srv := New(repo)
		var input = domain.Core{Name: "Lukman", HP: "123", Addres: "lmg"}
		var id = uint(1)
		res, err := srv.UpdateProfile(input, id)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.EqualError(t, err, gorm.ErrRecordNotFound.Error(), "pesan error tidak sesuai")
		assert.NotEqual(t, id, res.ID, "harusnya tidak sama dengan nilai id")
		repo.AssertExpectations(t)
	})

	t.Run("cant retrive data on database", func(t *testing.T) {
		repo.On("Update", mock.Anything, mock.Anything).Return(domain.Core{}, errors.New(config.DATABASE_ERROR)).Once()
		srv := New(repo)
		var input = domain.Core{Name: "Lukman", HP: "123", Addres: "lmg"}
		var id = uint(1)
		res, err := srv.UpdateProfile(input, id)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.EqualError(t, err, config.DATABASE_ERROR, "pesan error tidak sesuai")
		assert.NotEqual(t, id, res.ID, "harusnya tidak sama dengan nilai id")
		repo.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sucses delete profil", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.Core{}, nil).Once()
		srv := New(repo)
		var id = uint(1)
		_, err := srv.DeleteUser(id)
		assert.Nil(t, err)
		//assert.Empty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Data no found", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.Core{}, gorm.ErrRecordNotFound).Once()
		srv := New(repo)
		var id = uint(1)
		_, err := srv.DeleteUser(id)
		assert.NotNil(t, err)
		assert.EqualError(t, err, gorm.ErrRecordNotFound.Error(), "pesan error tidak sesuai")
		repo.AssertExpectations(t)
	})

	t.Run("error data on database", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.Core{}, errors.New(config.DATABASE_ERROR)).Once()
		srv := New(repo)
		var id = uint(1)
		_, err := srv.DeleteUser(id)
		assert.NotNil(t, err)
		assert.EqualError(t, err, config.DATABASE_ERROR, "pesan error tidak sesuai")
		repo.AssertExpectations(t)
	})

}
