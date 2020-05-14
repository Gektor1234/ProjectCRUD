package projectCRUDlogic

import (
	"ProjectCRUD/projectCRUDapp"
	"context"
	"github.com/sirupsen/logrus"
)

type peopleusecase struct {
	peopleRepo projectCRUDapp.PeopleRepository
}

func NewPeopleUsecase(p projectCRUDapp.PeopleRepository) projectCRUDapp.PeopleUsecase {
	return &peopleusecase{
		peopleRepo: p,
	}
}

func (p *peopleusecase) Fetch(c context.Context) (res []projectCRUDapp.PeopleEntity, err error) {

	res, err = p.peopleRepo.Fetch(c)
	if err != nil {
		return nil, err
	}
	return
}

func (p *peopleusecase) AddHuman(ctx context.Context, a *projectCRUDapp.PeopleEntity) (err error) {
	err = p.peopleRepo.AddHuman(ctx, a)
	if err != nil {
		logrus.Println(err)
	}
	return
}
