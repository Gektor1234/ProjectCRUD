package main

import (
	"ProjectCRUD/config"
	"ProjectCRUD/httphand"
	"ProjectCRUD/postgres"
	"ProjectCRUD/projectCRUDlogic"
	"context"
	"flag"
	"fmt"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/gocraft/dbr"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.json", "this configuration")
	flag.Parse()
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		logrus.Fatal(err)
	}

	db, err := dbr.Open("postgres", cfg.DB.ConnectString(),nil)
	if err != nil {
		logrus.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		logrus.Fatal(err)
	}

	logrus.Println("DB Connected...")

	e := echo.New()
	e.Use(DBRSessionMiddleware(db))
	peopleRepo := postgres.NewPostgresRepository(db)
	peopleService := projectCRUDlogic.NewPeopleUsecase(peopleRepo)
	httphand.NewPeopleHandler(e, peopleService)
	err = e.Start(fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		logrus.Fatal(err)
	}

}

func DBRSessionMiddleware(db *dbr.Connection)echo.MiddlewareFunc  {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(eCtx echo.Context)error {
				req:=eCtx.Request()
				ctx := req.Context()
				ctx = Newcontext(ctx,db)
				eCtx.SetRequest(req.WithContext(ctx))
				return next(eCtx)
			}
	}
}


func Newcontext(ctx context.Context, db *dbr.Connection) context.Context {
	return context.WithValue(ctx, "dbrsession", db.NewSession(nil))
}
