package postgres

import (
	"ProjectCRUD/projectCRUDapp"
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) projectCRUDapp.PeopleRepository {
	return &postgresRepository{db}
}
func (p *postgresRepository) fetch(ctx context.Context, query string) (result []projectCRUDapp.PeopleEntity) {
	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()
	result = make([]projectCRUDapp.PeopleEntity, 0)
	for rows.Next() {
		p := projectCRUDapp.PeopleEntity{}
		err = rows.Scan(
			&p.Id,
			&p.Firstname,
			&p.Lastname,
			&p.Age,
		)
		if err != nil {
			logrus.Error(err)
			return nil
		}
		result = append(result, p)

	}
	return result
}

func (p *postgresRepository) Fetch(ctx context.Context) (res []projectCRUDapp.PeopleEntity, err error) {
	query := "SELECT * FROM PeopleEntity order by id"
	res = p.fetch(ctx, query)
	return
}

func (p *postgresRepository) AddHuman(ctx context.Context, a *projectCRUDapp.PeopleEntity) error {
	query := "INSERT INTO PeopleEntity (id, firstname, lastname, age)VALUES($1, $2, $3, $4)"
	res, err := p.db.QueryContext(ctx, query, a.Id, a.Firstname, a.Lastname, a.Age)
	if err != nil {
		logrus.Println(err)
	} else {
		logrus.Println(res)
	}
	return err
}
