package models

import (
	"log"
)

type User struct {
	Email      string
	Password  string
	Nickname  string
	Icon      string
	Create_at int
}

func (user * User) Create() error{
	statement := "insert into users (email, password, nickname, icon, create_at ) values (?, ?, ?, ?, ?)"
	tx, err := Db.Begin()
	
	check_err(err)

	defer tx.Rollback()

	stmt,err := tx.Prepare(statement)
	check_err(err)

	_,err = stmt.Exec(user.Email,user.Password,user.Nickname,user.Icon,user.Create_at)

	check_err(err)

	err = tx.Commit()
	check_err(err)

	return nil

}

func Login(email,password string) User{
	user := User{}
	statement := "SELECT email, password, nickname, icon, create_at from users where email=? and password=?"

	rows, err := Db.Query(statement,email,password)
	check_err(err)
	
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&user.Email, &user.Password, &user.Nickname,&user.Icon,&user.Create_at)
	}

	return user
}

func check_err(err error){
	if err != nil {
		log.Fatal(err)
	}
}