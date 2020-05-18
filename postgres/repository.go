package postgres

import (
	"ProjectCRUD/projectCRUDapp"
	"context"
	"errors"
	"github.com/gocraft/dbr"
	"github.com/sirupsen/logrus"
)

type postgresRepository struct {
}

func NewPostgresRepository(*dbr.Connection) projectCRUDapp.PeopleRepository {
	return &postgresRepository{}
}
func (p *postgresRepository) fetch(ctx context.Context, query string) (result []projectCRUDapp.PeopleEntity,err error) {
	session := ctx.Value("dbrsession").(*dbr.Session)
	_,err = session.SelectBySql(query).LoadContext(ctx,&result)
	if err != nil {
		logrus.Error(err)
		return nil,err
	}
	return result,nil

}

func (p *postgresRepository) Fetch(ctx context.Context) (res []projectCRUDapp.PeopleEntity, err error) {
	query := "SELECT * FROM PeopleEntity order by id"
	res,err = p.fetch(ctx, query)
	if err != nil{
		logrus.Println(err)
	}
	return
}

func (p *postgresRepository) AddHuman(ctx context.Context, a *projectCRUDapp.PeopleEntity) error {
	session := ctx.Value("dbrsession").(*dbr.Session)
	query := "INSERT INTO PeopleEntity (id, firstname, lastname, age)VALUES($1, $2, $3, $4)"
	res, err := session.QueryContext(ctx, query, a.Id, a.Firstname, a.Lastname, a.Age)
	if err != nil {
		logrus.Println(err)
	} else {
		logrus.Println(res)
	}
	return err
}

func (p *postgresRepository)getMan(ctx context.Context,id int64)( result projectCRUDapp.PeopleEntity,err error)  {
	query := "SELECT * FROM PeopleEntity WHERE id = ?"
	result = projectCRUDapp.PeopleEntity{}
	session := ctx.Value("dbrsession").(*dbr.Session)
	err = session.SelectBySql(query,id).LoadOneContext(ctx,&result)
	if err != nil {
		logrus.Error(err)
		return result,err
	}
	return result, nil
}


func (p *postgresRepository)GetMan(ctx context.Context,id int64) (res projectCRUDapp.PeopleEntity,err error)  {
	res,err = p.getMan(ctx, id)
	if err != nil{
		logrus.Println(err)
		return res, err
	}
	return
}

func (p *postgresRepository)DeleteHuman(ctx context.Context,id int64)error  {
	session,ok := ctx.Value("dbrsession").(*dbr.Session)
	if !ok{
		return errors.New("is not dbr.session")
	}
	query := "DELETE  FROM PeopleEntity WHERE id = $1"
	_,err := session.QueryContext(ctx,query,id)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return err
}

func (p *postgresRepository)UpdateHuman(ctx context.Context, pe *projectCRUDapp.PeopleEntity)error  {
	query := "UPDATE PeopleEntity SET firstname =$2,lastname=$3,age=$4 WHERE id=$1"
	session := ctx.Value("dbrsession").(*dbr.Session)
	_,err := session.QueryContext(ctx,query,pe.Id,pe.Firstname,pe.Lastname,pe.Age)
	if err != nil{
		logrus.Println(err)
		return err
	}
	return err
}