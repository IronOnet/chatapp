package routes 

import (
	"net/http" 
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetLogin(t *testing.T){
	mux := http.NewServeMux() 
	mux.HandleFunc("/login", Login) 

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/login", nil) 
	mux.ServeHTTP(writer, request) 

	if writer.Code != 200{
		t.Errorf("Response code is %v", writer.Code)
	}

	body := writer.Body.String() 
	if strings.Contains(body, "Sign in") == false{
		t.Errorf("Body doesn't contain Sign in")
	}
	
}