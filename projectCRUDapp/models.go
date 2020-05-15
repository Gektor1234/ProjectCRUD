package projectCRUDapp

import (
	"context"
	"github.com/gocraft/dbr"
)

type PeopleEntity struct {
	Id        dbr.NullInt64  `json:"id" db:"id"`
	Firstname dbr.NullString `json:"firstname" db:"firstname"`
	Lastname  dbr.NullString `json:"lastname" db:"lastname"`
	Age       dbr.NullInt64  `json:"age" db:"age"`
}

type PeopleRepository interface {
	Fetch(ctx context.Context) (res []PeopleEntity, err error)
	AddHuman(ctx context.Context, a *PeopleEntity) error
}

type PeopleUsecase interface {
	Fetch(ctx context.Context) ([]PeopleEntity, error)
	AddHuman(ctx context.Context, a *PeopleEntity) (err error)
}
