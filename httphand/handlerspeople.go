package httphand

import (
	"ProjectCRUD/projectCRUDapp"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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
	e.GET("/peoples/:id",handler.GetMan)
	e.DELETE("/peoples/:id",handler.DeleteHuman)
	e.PUT("/peoples",handler.UpdateHuman)
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

func (p *PeopleHandler)GetMan(c echo.Context)error  {
	ctx := c.Request().Context()
	ids,err := strconv.Atoi(c.Param("id"))
	if err != nil{
		logrus.Println(err)
	}
	id := int64(ids)
	res,err := p.PeopleUsecase.GetMan(ctx,id)
	if err != nil{
		return err
	}
	return c.JSON(http.StatusOK,res)
}

func (p *PeopleHandler)DeleteHuman(c echo.Context)error {
	ctx := c.Request().Context()
	ids,err := strconv.Atoi(c.Param("id"))
	if err != nil{
		logrus.Println(err)
	}
	id := int64(ids)
	err = p.PeopleUsecase.DeleteHuman(ctx,id)
	if err != nil{
		logrus.Println(err)
	}
	return c.String(http.StatusOK,"deleted")
}

func (p *PeopleHandler)UpdateHuman(c echo.Context)error  {
	var human projectCRUDapp.PeopleEntity
	err := c.Bind(&human)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	err = p.PeopleUsecase.UpdateHuman(ctx,&human)
	if err != nil{
		logrus.Println(err)
	}
	return c.JSON(http.StatusOK,"data updated successfully")
}