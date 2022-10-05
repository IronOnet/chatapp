package main

import (
	"net/http"
	"time"

	"github.com/irononet/chatapp/routes"
	"github.com/irononet/chatapp/utils"
	
)

var config utils.Configuration = utils.LoadConfiguration("config.json")

func main(){
	utils.Pstdout("Chatapp", utils.Version(), "Started at", config.Address)

	mux := http.NewServeMux() 
	files := http.FileServer(http.Dir(config.Static)) 
	mux.Handle("/static/", http.StripPrefix("/static/", files))


	// all route patterns mathed here 
	

	// index 
	mux.HandleFunc("/", routes.Index) 
	// error 
	mux.HandleFunc("/err", routes.Err) 

	// Authentication routes 
	mux.HandleFunc("/login", routes.Login) 
	mux.HandleFunc("/logout", routes.Logout) 
	mux.HandleFunc("/signup", routes.Signup) 
	mux.HandleFunc("/signup_account", routes.SignupAccount) 
	mux.HandleFunc("/authenticate", routes.Authenticate)

	// thread routes 
	mux.HandleFunc("/thread/new", routes.NewThread) 
	mux.HandleFunc("/thread/create", routes.CreateThread) 
	mux.HandleFunc("/thread/post", routes.PostThread) 
	mux.HandleFunc("/thread/read", routes.ReadThread)

	// starting the server 
	server := &http.Server{
		Addr: config.Address, 
		Handler: mux, 
		ReadTimeout: time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout: time.Duration(config.WriteTimeout * int64(time.Second)), 
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}