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

func (p *peopleusecase)GetMan(ctx context.Context,id int64)(projectCRUDapp.PeopleEntity,error)  {
	res,err:= p.peopleRepo.GetMan(ctx,id)
	if err != nil{
		logrus.Println(err)
		return res, err
	}
	return res,err
}

func (p *peopleusecase)DeleteHuman(ctx context.Context,id int64)error{
	err := p.peopleRepo.DeleteHuman(ctx,id)
	if err != nil{
		logrus.Println(err)
		return err
	}
	return err
}

func (p *peopleusecase)UpdateHuman(ctx context.Context,pe *projectCRUDapp.PeopleEntity) error {
	err := p.peopleRepo.UpdateHuman(ctx,pe)
	if err != nil{
		logrus.Println(err)
		return err
	}
	return err
}