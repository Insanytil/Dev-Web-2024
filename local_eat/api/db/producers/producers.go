package producers

import (
	"database/sql"
	"local_eat/api/model"
)

func GetProducers(db *sql.DB) ([]*model.Producers, error) {
	rows, err := db.Query("select * from producers")
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