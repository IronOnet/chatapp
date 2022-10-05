package data 

import (
	"time"
)

type User struct{
	Id			int		`json:"id"`
	Uuid		string 	`json:"uuid"`
	Name		string	`json:"name"`
	Email		string	`json:"email"`
	Password	string	`json:"password"`
	CreatedAt	time.Time	`json:"created_at"`
}

type Session struct{
	Id			int 		`json:"id"`
	Uuid		string		`json:"uuid"`
	Email		string		`json:"email"`
	UserId		int			`json:"user_id"`
	CreatedAt	time.Time	`json:"created_at"`
}

// create a new session for an existing user 
func (user *User) CreateSession() (session Session, err error){
	statement := "insert into sessions (uuid, email, user_id, created_at) value ($1, $2, $3, $4) returning id, uuid, email, user_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil{
		return 
	}
	defer stmt.Close()
	// use QueryRow to return a rown and scan the returned id into  the Sesson struct 
	err = stmt.QueryRow(createUUID(), user.Email, user.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return 
}

// get the session for the existing user 
func (user *User) Session() (session Session, err error){
	session = Session{} 
	err  = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = $1", user.Id). 
	Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return 
}

// Check if the session is valid in the database
func (session *Session) Check() (valid bool, err error){
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1", session.Uuid). 
	Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt) 

	if err != nil{
		valid = false
		return 
	}
	if session.Id != 0{
		valid = true 
	}
	return 
}

// Delete Session from database 
func (session *Session) DeleteByUUID() (err error){
	statement := "DELETE FROM sessions WHERE uuid = $1"
	stmt, err := Db.Prepare(statement) 
	if err != nil{
		return 
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.Uuid)
	return 
}

// Get the user from the session 
func (session *Session) User() (user User, err error){
	user = User{} 
	err = Db.QueryRow("SELECT id, uuid, name, email, created_at FROM sessions WHERE id=$1", session.UserId). 
	Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return 
}

// Delete all sessions from database 
func SessionDeleteAll() (err error){
	statement := "DELETE FROM sessions" 
	_, err = Db.Exec(statement) 
	return 
}

// create a new user and save the user info in the database 
func (user *User) Create() (err error){
	statement := "INSERT INTO users (uuid, name, email, password, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, uuid, created_at"
	stmt, err := Db.Prepare(statement) 
	if err != nil{
		return 
	}
	defer stmt.Close()

	// use QueryRow to return a row and scan the returned id into user struct 
	err = stmt.QueryRow(createUUID(), user.Name, user.Email, Encrypt(user.Password), time.Now()). 
	Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	return 
}

// Delete from database 
func (user *User) Delete() (err error){
	statement := "DELETE FROM users WHERE id =$1" 
	stmt, err := Db.Prepare(statement) 
	if err != nil{
		return 
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)	
	return 
}

// Update user information in the database 
func (user *User) Update() (err error){
	statement := "UPDATE users SET name = $2, email = $3 WHERE id= $1"
	stmt, err := Db.Prepare(statement) 
	if err != nil{
		return 
	}
	defer stmt.Close() 

	_, err =  stmt.Exec(user.Id, user.Name, user.Email)
	return 
}

// Delete all users from database 
func UserDeleteAll() (err error){
	statement := "DELETE FROM users" 
	_, err = Db.Exec(statement)
	return 
}

// Get all users 
func Users() (users []User, err error){
	rows, err := Db.Query("SELECT id, uuid, name, email, password, created_at FROM users")
	if err != nil{
		return 
	}
	for rows.Next(){
		user := User{} 
		if err = rows.Scan(&user.Id, &user.Uuid ,&user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil{
			return 
		}
		users = append(users, user)
	}
	rows.Close()
	return 
}

// get a single user by email 
func UserByEmail(email string) (user User, err error){
	user = User{} 
	err = Db.QueryRow("SELEC id, uuid, name, email, password, creted_at FROM users WHERE email = $1", email). 
	Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return 
}

// get a single user using the id 
func UserByUUID(uuid string) (user User, err error){
	user = User{} 
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE id = $1", uuid). 
	Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return 
}