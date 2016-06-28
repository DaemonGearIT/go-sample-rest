package webcontrollers

import (
	"github.com/DaemonGearIT/go-sample-rest/usecases"
	"github.com/fatih/color"
	"github.com/labstack/echo"
	"net/http"
)

type (
	UserInteractor interface {
		SaveUser(user usecases.User) error
		FindAllUsers() error
	}

	UserController struct {
		UserInteractor UserInteractor
	}
)

func (ctrl *UserController) SaveUser(c echo.Context) error {
	color.Cyan("Register User called")
	user := new(usecases.User)

	if err := c.Bind(user); err != nil {
		return err
	}

	err = ctrl.UserInteractor.SaveUser(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}
