package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Port int64 `json:"port"`
	DB   DB    `json:"db"`
}

type DB struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (db *DB) ConnectString() string {
	return fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=disable",
		db.User, db.Password, db.Name, db.Host, db.Port)
}

func NewConfig(path string) (*Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	cfg := new(Config)
	err = json.Unmarshal(file, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
