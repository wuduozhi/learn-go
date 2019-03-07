package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"go-web/models"
	"fmt"
	"go-web/pkg/logging"
	"go-web/pkg/utils"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	users,_ := models.Users()

	for _, user := range users {
		fmt.Println(user)
	}
	logging.Info("Router GetAllUsers")
	fmt.Fprintf(w,"All users")

}

func LoginTemplate(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	t := utils.ParseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(w, nil)
}


func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	err := r.ParseForm()
	if err != nil {
		logging.Error("Router Login ",err)
	}

	email := r.PostFormValue("email")
	user,err := models.GetUserByEmail(email)

	if err != nil {
		logging.Error("Router Login ,can't find user")
	}

	if user.Password == models.Encrypt(r.PostFormValue("password")){
		logging.Info("Login Success")
	}else{
		logging.Info("Login Fail")
	}

	fmt.Fprintf(w,"Login Success")

}

func Signup(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	utils.GenerateHTML(w, nil, "login.layout", "public.navbar", "signup")
}

func DoSignup(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	err := r.ParseForm()
	if err != nil {
		logging.Error("Router DoSignup ",err)
	}

	user := models.User{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}

	if err := user.Create(); err != nil {
		logging.Error("Can't create user")
	}

	http.Redirect(w,r,"/login",302)

}