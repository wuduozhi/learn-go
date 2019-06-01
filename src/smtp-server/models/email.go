package models

import (
	"smtp-server/pkg/logging"
)

type Email struct {
	From      string
	To        string
	Subject   string
	Message   string
	Create_at int
	Is_new    int
}

func (email * Email) Create() error{
	statement := "insert into email (`from`,`to`,message,`subject`,create_at,is_new) values (?, ?, ?, ?, ?, ?)"
	tx, err := Db.Begin()
	
	if err != nil{
		logging.Debug("Create Mail ",err)
		return err
	}

	defer tx.Rollback()
	stmt,err := tx.Prepare(statement)

	if err != nil{
		logging.Debug("Create Mail ",err)
		return err
	}
	_,err = stmt.Exec(email.From,email.To,email.Message,email.Subject,email.Create_at,email.Is_new)

	if err != nil {
		logging.Debug("Create Mail Exec ",err)
		return err
	}

	err = tx.Commit()

	if err != nil{
		logging.Debug("Create Mail Commit ",err)
		return err
	}

	return nil
}

func GetAllEmails(to string) ([]Email,int,error){
	statement := "SELECT `from`,`to`,message,`subject`,create_at,is_new FROM email WHERE `to`=?"

	row,err:= Db.Query(statement,to)
	check_err(err)
	
	emails := make([]Email,0)
	email := Email{}
	sum := 0
	for row.Next(){
		err = row.Scan(&email.From,&email.To,&email.Message,&email.Subject,&email.Create_at,&email.Is_new)

		if err != nil{
			logging.Info(err)
			return emails,sum*8,err
		}
		sum += len(email.Message)
		emails = append(emails,email)
	}


	return emails,sum*8,err
}

