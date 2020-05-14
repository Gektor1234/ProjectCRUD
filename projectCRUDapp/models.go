package projectCRUDapp

import (
	"context"
	//"github.com/gocraft/dbr"
)

type PeopleEntity struct {
	Id        int64  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int64  `json:"age"`
}

type PeopleRepository interface {
	Fetch(ctx context.Context) (res []PeopleEntity, err error)
	AddHuman(ctx context.Context, a *PeopleEntity) error
}

type PeopleUsecase interface {
	Fetch(ctx context.Context) ([]PeopleEntity, error)
	AddHuman(ctx context.Context, a *PeopleEntity) (err error)
}
