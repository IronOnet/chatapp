package data 

// test data 
var users = []User{
	{
		Name: "Peter Jones", 
		Email: "peterjones@email.com",
		Password: "password123",
	}, 
	{
		Name: "Conrad Adenauer",
		Email: "conrad9@eu.com", 
		Password: "password123",  

	}, 
}

func setup(){
	ThreadDeleteAll() 
	SessionDeleteAll() 
	UserDeleteAll()
}