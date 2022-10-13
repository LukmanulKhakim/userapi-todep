package delivery

import (
	"net/http"
	"strconv"
	"userapi/feature/user/domain"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type userHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := userHandler{srv: srv}

	e.POST("/users", handler.AddUser())
	e.GET("/users", handler.ShowAllUser())
	e.GET("/users/:id", handler.Profile())
	e.PUT("/users/:id", handler.EditProfile())
	e.DELETE("/users/:id", handler.DeleteUser())
}

func (us *userHandler) AddUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input addUserFormat
		if err := c.Bind(&input); err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusBadRequest, FailResponses("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := us.srv.AddUser(cnv)
		if err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusInternalServerError, FailResponses(err.Error()))
		}
		return c.JSON(http.StatusCreated, SuccessResponses("sucses input user", ToResponse(res, "insert")))
	}
}

func (us *userHandler) ShowAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := us.srv.ShowAllUser()
		if err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusInternalServerError, FailResponses(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponses("success get all user", ToResponse(res, "all")))
	}
}

func (us *userHandler) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, err := strconv.Atoi(c.Param("id"))
		res, err := us.srv.Profile(uint(ID))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponses(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponses("sucses get userByid", res))
	}
}

func (us *userHandler) EditProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		//var update domain.Basic
		ID, err := strconv.Atoi(c.Param("id"))

		var new editUserFormat
		if err := c.Bind(&new); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponses("cannot bind input"))
		}

		update := domain.Core{Name: new.Name, HP: new.HP, Addres: new.Addres}
		res, err := us.srv.UpdateProfile(update, uint(ID))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponses(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponses("sucses edit profile", ToResponse(res, "edit")))
	}
}

func (us *userHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, err := strconv.Atoi(c.Param("id"))
		if _, err = us.srv.DeleteUser(uint(ID)); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponses(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponses("sucses delete user", " "))
	}
}
