package userHandler

import (
	"net/http"

	"github.com/labstack/echo"
	"app/models"
)

type (
    resultLists struct {
	    Users []user.User `json:"users"`
    }

    handler struct {
	    UserModel user.UserModelImpl
    }
)

func NewHandler(u user.UserModelImpl) *handler {
	return &handler{u}
}

func (h *handler) GetIndex(c echo.Context) error {
	lists := h.UserModel.FindAll()
	u := &resultLists{
		Users: lists,
	}
	return c.JSON(http.StatusOK, u)
}

func (h *handler) GetDetail(c echo.Context) error {
	id := c.Param("id")
	u := h.UserModel.FindByID(id)
	return c.JSON(http.StatusOK, u)
}
