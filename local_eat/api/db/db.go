package db

import (
	"database/sql"
	"local_eat/api/model"
)

type DB interface {
	GetTechnologies() ([]*model.Technology, error)
	GetProducers() ([]*model.Producers, error)
}

type MySQLDB struct {
	mysql *sql.DB
}

func NewDB(db *sql.DB) DB {
	return MySQLDB{mysql: db}
}

func (db MySQLDB) GetTechnologies() ([]*model.Technology, error) {
	rows, err := db.mysql.Query("select name, details from technologies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var techs []*model.Technology
	for rows.Next() {
		tech := new(model.Technology)
		err = rows.Scan(&tech.Name, &tech.Details)
		if err != nil {
			return nil, err
		}
		techs = append(techs, tech)
	}
	return techs, nil
}

func (db MySQLDB) GetProducers() ([]*model.Producers, error) {
	rows, err := db.mysql.Query("select * from producers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var prods []*model.Producers
	for rows.Next() {
		prod := new(model.Producers)
		err = rows.Scan(&prod.Id, &prod.Name, &prod.Picture, &prod.Created)
		if err != nil {
			return nil, err
		}
		prods = append(prods, prod)
	}
	return prods, nil
}