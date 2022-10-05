package routes

import (
	"fmt"
	"net/http"

	"github.com/irononet/chatapp/data"
	"github.com/irononet/chatapp/utils"
)

func Login(writer http.ResponseWriter, request *http.Request){
	t := utils.ParseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(writer, nil)
}

// GET signup 
// show the signup page 
func Signup(writer http.ResponseWriter, request *http.Request){
	utils.GenerateHTML(writer, nil , "login.layout", "public.navbar", "signup")
}

func SignupAccount(writer http.ResponseWriter, request *http.Request){
	err := request.ParseForm() 
	if err != nil{
		fmt.Errorf("Could not parse form %s", err)
	}

	user := data.User{
		Name: request.PostFormValue("name"), 
		Email: request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}

	if err := user.Create(); err != nil{
		fmt.Errorf("Cannot create user! Error: '%s")
	}
	http.Redirect(writer, request, "/login", 302)
}

// Authenticate user 
func Authenticate(writer http.ResponseWriter, request *http.Request){
	err := request.ParseForm() 
	user, err := data.UserByEmail(request.PostFormValue("email"))
	if err != nil{
		fmt.Errorf("Cannot create Session : %v", err)
	}
	if user.Password == data.Encrypt(request.PostFormValue("password")){
		session, err := user.CreateSession() 
		if err != nil{
			fmt.Errorf("Cannot create session %v", err)
		}
		cookie := http.Cookie{
			Name : "_cookie", 
			Value: session.Uuid, 
			HttpOnly: true,
		}

		http.SetCookie(writer, &cookie) 
		http.Redirect(writer, request, "/", http.StatusFound)
	} else{
		http.Redirect(writer, request, "/login", http.StatusFound)
	}
}

// Get /logout 
// logs the user out  
func Logout(writer http.ResponseWriter, request *http.Request){
	cooke, err := request.Cookie("_cookie") 
	if err != http.ErrNoCookie{
		fmt.Errorf("Failed to get cookie %s", err)
		session := data.Session{Uuid: cooke.Value}
		session.DeleteByUUID()
	}

	http.Redirect(writer, request, "/", 302)
}