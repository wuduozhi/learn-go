package models

import (
	"time"
	"go-web/pkg/logging"
)


type User struct {
	Id          int
	Uuid        string
	Name        string
	Email       string
	Password    string
	CreatedAt    time.Time
}


func (user *User) Create() (err error){
	statement := "insert into users (uuid, name, email, password, created_at) values (?, ?, ?, ?, ?)"
	tx, err := Db.Begin()
	if err != nil {
		logging.Error(err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(statement)
	if err != nil {
		logging.Error(err)
	}

	_, err = stmt.Exec(createUUID(), user.Name, user.Email, Encrypt(user.Password), time.Now())
	if err != nil {
		logging.Error(err)
	}

	err = tx.Commit()
	if err != nil {
		logging.Error(err)
	}

	return
}


// Get all users in the database and returns it
func Users() (users []User, err error) {
	rows, err := Db.Query("SELECT id, uuid, name, email, password, created_at FROM users")
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			logging.Error("Database get all users")
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

func GetUserByEmail(email string) (user User, err error){
	user = User{}
	statement := "SELECT id, uuid, name, email, password, created_at FROM users WHERE email = ?"
	rows, err := Db.Query(statement,email)
	if err != nil {
		logging.Error(err)
	}
	
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	}

	return
}