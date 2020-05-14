package main

import (
	"ProjectCRUD/config"
	"ProjectCRUD/httphand"
	"ProjectCRUD/postgres"
	"ProjectCRUD/projectCRUDlogic"
	"database/sql"
	"flag"
	"fmt"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	//"github.com/gocraft/dbr"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.json", "this configuration")
	flag.Parse()
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		logrus.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DB.ConnectString())
	if err != nil {
		logrus.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		logrus.Fatal(err)
	}

	logrus.Println("DB Connected...")

	e := echo.New()
	people := postgres.NewPostgresRepository(db)
	au := projectCRUDlogic.NewPeopleUsecase(people)
	httphand.NewPeopleHandler(e, au)
	err = e.Start(fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		logrus.Fatal(err)
	}

}
