package db

import (
	"database/sql"
	"local_eat/api/model"
)

type DB interface {
	GetProducers() ([]*model.Producers, error)
}

type MySQLDB struct {
	mysql *sql.DB
}

func NewDB(db *sql.DB) DB {
	return MySQLDB{mysql: db}
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