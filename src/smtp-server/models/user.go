package models

import (
	"smtp-server/pkg/logging"
	"errors"
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


func UserExists(email string) error{
	var count int
	statement := "select count(*) from users where email=?"
	row,err:= Db.Query(statement,email)
	check_err(err)

	for row.Next(){
		err = row.Scan(&count)

		if err != nil{
			logging.Info(err)
			return err
		}
	}

	if count > 0{
		return nil
	}else{
		return errors.New("User don't exist")
	}
}

func AuthUser(email,password string) error{
	statement := "SELECT count(*) from users where email=? and password=?"

	rows, err := Db.Query(statement,email,password)
	check_err(err)
	
	defer rows.Close()
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}

	if count > 0{
		return nil
	}else{
		return errors.New("User don't exist")
	}
}

func check_err(err error){
	if err != nil {
		logging.Info(err)
	}
}