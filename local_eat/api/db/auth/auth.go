package auth

import (
	"database/sql"
	"local_eat/api/model"
	"log"
)

func CreateUser(db *sql.DB, user *model.UsersSignup) (error){
	query := "INSERT INTO `users` (`username`, `password`, `email`, `age`, `gender`, `address`, `locality`, `cellphone`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(query, user.Username, user.Password, user.Email, user.Age, user.Gender, user.Address, user.Locality, user.Cellphone)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetUser(db *sql.DB, username string) (model.UsersLogin, error) {
	var user model.UsersLogin
	query := "SELECT username, password, email FROM `users` WHERE `username` = ?"
	err := db.QueryRow(query, username).Scan(&user.Username, &user.Password, &user.Email)
	if err != nil {
		log.Println(err)
		return user, err
	}
	return user, nil
}