package data 


import (
	"testing"
)

func ThreadDeleteAll() (err error){
	db := Db() 
	defer db.Close() 
	statement := "delete from threads"
	_, err = db.Exec(statement)
	if err != nil{
		return 
	}
	return 
}

func TestCreateThread(t *testing.T){
	setup() 
	if err := users[0].Create(); err != nil{
		t.Error(err, "Cannot create user.") 

	}
	conv, err := users[0].CreateThread("My First Thread Ever") 
	if err != nil{
		t.Error(err, "cannot create thread")
	}
	if conv.UserId	!= users[0].Id{
		t.Errorf("User not linked with this thread")
	}
}