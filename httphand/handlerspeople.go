package httphand

import (
	"ProjectCRUD/projectCRUDapp"
	"github.com/labstack/echo"
	"net/http"
)

type PeopleHandler struct {
	PeopleUsecase projectCRUDapp.PeopleUsecase
}

func NewPeopleHandler(e *echo.Echo, us projectCRUDapp.PeopleUsecase) {
	handler := &PeopleHandler{
		PeopleUsecase: us,
	}
	e.GET("/peoples", handler.Fetch)
	e.POST("/peoples", handler.AddHuman)
}

func (p *PeopleHandler) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	listPe, err := p.PeopleUsecase.Fetch(ctx)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, listPe)
}

func (p *PeopleHandler) AddHuman(c echo.Context) error {
	var human projectCRUDapp.PeopleEntity
	err := c.Bind(&human)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	err = p.PeopleUsecase.AddHuman(ctx, &human)
	if err != nil {
		return err
	}
	res := "OK"
	return c.String(http.StatusOK, res)
}
