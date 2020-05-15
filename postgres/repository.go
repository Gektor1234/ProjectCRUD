package postgres

import (
	"ProjectCRUD/projectCRUDapp"
	"context"
	"github.com/sirupsen/logrus"
	"github.com/gocraft/dbr"
)

type postgresRepository struct {
}

func NewPostgresRepository(db *dbr.Connection) projectCRUDapp.PeopleRepository {
	return &postgresRepository{}
}
func (p *postgresRepository) fetch(ctx context.Context, query string) (result []projectCRUDapp.PeopleEntity) {
	session := ctx.Value("dbrsession").(*dbr.Session)
	_,err := session.SelectBySql(query).LoadContext(ctx,&result)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return result
}

func (p *postgresRepository) Fetch(ctx context.Context) (res []projectCRUDapp.PeopleEntity, err error) {
	query := "SELECT * FROM PeopleEntity order by id"
	res = p.fetch(ctx, query)
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
