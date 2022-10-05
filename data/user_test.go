package data 

import (
	"database/sql"
	"testing"
)

func TestUserCreate(t *testing.T){
	setup() 
	if err := users[0].Create(); err != nil{
		t.Error(err, "Cannot create user.")
	}
	if users[0].Id == 0{
		t.Errorf("No id or created_at in user")
	}
	u, err := UserByEmail(users[0].Email) 
	if err != nil{
		t.Error(err, "User not created!")
	}
	if users[0].Email != u.Email{
		t.Error("user retrieved is not the same as the one created.")
	}
}

func TestUserDelete(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	if err := users[0].Delete(); err != nil {
		t.Error(err, "- Cannot delete user")
	}
	_, err := UserByEmail(users[0].Email)
	if err != sql.ErrNoRows {
		t.Error(err, "- User not deleted.")
	}
}

func TestUserUpdate(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	users[0].Name = "Random User"
	if err := users[0].Update(); err != nil {
		t.Error(err, "- Cannot update user")
	}
	u, err := UserByEmail(users[0].Email)
	if err != nil {
		t.Error(err, "- Cannot get user")
	}
	if u.Name != "Random User" {
		t.Error(err, "- User not updated")
	}
}

func TestUserByUUID(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	u, err := UserByUUID(users[0].Uuid)
	if err != nil {
		t.Error(err, "User not created.")
	}
	if users[0].Email != u.Email {
		t.Errorf("User retrieved is not the same as the one created.")
	}
}

func TestUsers(t *testing.T) {
	setup()
	for _, user := range users {
		if err := user.Create(); err != nil {
			t.Error(err, "Cannot create user.")
		}
	}
	u, err := Users()
	if err != nil {
		t.Error(err, "Cannot retrieve users.")
	}
	if len(u) != 2 {
		t.Error(err, "Wrong number of users retrieved")
	}
	if u[0].Email != users[0].Email {
		t.Error(u[0], users[0], "Wrong user retrieved")
	}
}

func TestCreateSession(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	session, err := users[0].CreateSession()
	if err != nil {
		t.Error(err, "Cannot create session")
	}
	if session.UserId != users[0].Id {
		t.Error("User not linked with session")
	}
}

func TestGetSession(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	session, err := users[0].CreateSession()
	if err != nil {
		t.Error(err, "Cannot create session")
	}

	s, err := users[0].Session()
	if err != nil {
		t.Error(err, "Cannot get session")
	}
	if s.Id == 0 {
		t.Error("No session retrieved")
	}
	if s.Id != session.Id {
		t.Error("Different session retrieved")
	}
}

func TestcheckValidSession(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	session, err := users[0].CreateSession()
	if err != nil {
		t.Error(err, "Cannot create session")
	}

	uuid := session.Uuid

	s := Session{Uuid: uuid}
	valid, err := s.Check()
	if err != nil {
		t.Error(err, "Cannot check session")
	}
	if valid != true {
		t.Error(err, "Session is not valid")
	}

}

func TestcheckInvalidSession(t *testing.T) {
	setup()
	s := Session{Uuid: "123"}
	valid, err := s.Check()
	if err == nil {
		t.Error(err, "Session is not valid but is validated")
	}
	if valid == true {
		t.Error(err, "Session is valid")
	}

}

func TestDeleteSession(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	session, err := users[0].CreateSession()
	if err != nil {
		t.Error(err, "Cannot create session")
	}

	err = session.DeleteByUUID()
	if err != nil {
		t.Error(err, "Cannot delete session")
	}
	s := Session{Uuid: session.Uuid}
	valid, err := s.Check()
	if err == nil {
		t.Error(err, "Session is valid even though deleted")
	}
	if valid == true {
		t.Error(err, "Session is not deleted")
	}
}
