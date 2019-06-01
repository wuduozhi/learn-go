package pop3

import (
	"smtp-server/pkg/logging"
	"smtp-server/models"
	"strconv"
	"fmt"
)

//pop3 协议，参考https://kewl.lu/articles/pop3/

func userProcessor(req *Request) error{
	if len(req.Line) < 2 {
		logging.Info(req.Line,"Not enough arguments")
		return req.TextProto.PrintfLine("%d %s", 501, "Not enough arguments")
	}

	user := req.Line[1]

	if req.Authable == true{
		err := models.UserExists(user)
		if err != nil{
			logging.Info(err)
			return req.TextProto.PrintfLine("-ERR The user %s doesn't belong here!",user)
		}
		req.AuthUser = user
		logging.Info(user," auth begin")
		return req.TextProto.PrintfLine("+OK %s selected",user)
	}else{
		return req.TextProto.PrintfLine("+OK User signed in")
	}
}


func passProcessor(req *Request) error{
	if len(req.Line) < 2 {
		logging.Info(req.Line,"Not enough arguments")
		return req.TextProto.PrintfLine("%d %s", 501, "Not enough arguments")
	}

	if req.Authable == true{
		user := req.AuthUser
		pass := req.Line[1]
		err := models.AuthUser(user,pass)
		if err != nil{
			logging.Info(err)
			return req.TextProto.PrintfLine("-ERR Password incorrect!")
		}

		req.Authable = false
		logging.Info(user," auth success")
		return req.TextProto.PrintfLine("+OK User signed in")
	}else{
		return req.TextProto.PrintfLine("+OK User signed in")
	}
}

func statProcessor(req *Request) error{
	if len(req.Line) < 1 {
		logging.Info(req.Line,"Not enough arguments")
		return req.TextProto.PrintfLine("%d %s", 501, "Not enough arguments")
	}
	emails,sum,_ := models.GetAllEmails(req.AuthUser)

	logging.Info(emails)
	return req.TextProto.PrintfLine("+OK %d %d",len(emails),sum)
}

func listProcessor(req *Request) error {
	if len(req.Line) < 1 {
		logging.Info(req.Line,"Not enough arguments")
		return req.TextProto.PrintfLine("%d %s", 501, "Not enough arguments")
	}
	emails,sum,_ := models.GetAllEmails(req.AuthUser)

	_ = req.TextProto.PrintfLine("+OK %d  messages (%d octets)",len(emails),sum)

	for i:=0;i<len(emails);i++{
		_ = req.TextProto.PrintfLine("%d %d",i+1,len(emails[i].Message)*8)
	}

	return nil
}

func retrProcessor(req *Request) error {
	if len(req.Line) < 2 {
		logging.Info(req.Line,"Not enough arguments")
		return req.TextProto.PrintfLine("%d %s", 501, "Not enough arguments")
	}
	emails,_,_ := models.GetAllEmails(req.AuthUser)

	index,_ := strconv.Atoi(req.Line[1])
	if index > len(emails){
		return req.TextProto.PrintfLine("%d %s", index, "Not such email")
	}

	email := emails[index-1]
	_ = req.TextProto.PrintfLine("+OK %d octets",len(email.Message)*8)
	// _ = req.TextProto.PrintfLine(email.Message)
	fmt.Fprintf(req.Conn,email.Message)
	logging.Info(email.Message)
	return nil
}