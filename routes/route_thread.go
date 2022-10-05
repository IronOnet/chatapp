package routes 

import (
	"fmt" 
	"net/http" 

	"github.com/irononet/chatapp/data"
	"github.com/irononet/chatapp/utils"
)

func NewThread(writer http.ResponseWriter, request *http.Request){
	_, err := utils.Session(writer, request)
	if err != nil{
		http.Redirect(writer, request, "/login", 302)
	} else{
		utils.GenerateHTML(writer, nil, "layout", "private.navbar", "new.thread")
	}
}

// Create the user account 
func CreateThread(writer http.ResponseWriter, request *http.Request){
	sess, err := utils.Session(writer, request) 
	if err != nil{
		http.Redirect(writer, request, "/login", 302)
	} else{
		err = request.ParseForm() 
		if err != nil{
			fmt.Errorf("cannot parse form %v", err) 
			
		}
		user, err := sess.User() 
		if err != nil{
			fmt.Errorf("cannot get user from session %v", err)
		}
		topic := request.PostFormValue("topic") 
		if _, err := user.CreateThread(topic); err != nil{
			fmt.Errorf("cannot create thread %v", err)
		}
		http.Redirect(writer, request, "/", http.StatusFound)
	}
}

// Show the details of the thread, including the posts and the form 
// to write a post 
func ReadThread(writer http.ResponseWriter, request *http.Request){
	vals := request.URL.Query() 
	uuid := vals.Get("id") 
	thread, err := data.ThreadByUUID(uuid) 
	if err != nil{
		utils.ErrorMessage(writer, request, "Cannot read thread")
	} else{
		_, err := utils.Session(writer, request) 
		if err != nil{
			utils.GenerateHTML(writer, &thread, "layout", "public.navbar", "public.thread")
		} else{
			utils.GenerateHTML(writer, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}

// Create the Post 
func PostThread(writer http.ResponseWriter, request *http.Request){
	sess, err := utils.Session(writer, request) 
	if err != nil{
		http.Redirect(writer, request, "/login", http.StatusFound)
	}else{
		err = request.ParseForm() 
		if err != nil{
			fmt.Errorf("cannot parse form %v", err)
		}
		user, err := sess.User() 
		if err != nil{
			fmt.Errorf("cannot get user from session %v", err)
		}
		body := request.PostFormValue("body") 
		uuid := request.PostFormValue("uuid") 
		thread, err := data.ThreadByUUID(uuid)
		if err != nil{
			utils.ErrorMessage(writer, request, "Cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil{
			fmt.Errorf("cannot create post %v", err)
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}