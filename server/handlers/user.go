package handlers

import (
	"dashboard-api/service/user"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type UserHandler struct {
	get func() (*user.UserType, error)
}

func (handler *UserHandler) Get(c echo.Context) error {
	u, err := handler.get()
	if err != nil {
		return errors.Wrap(err, "Fail to list user")
	}

	return c.JSON(http.StatusOK, u)
}