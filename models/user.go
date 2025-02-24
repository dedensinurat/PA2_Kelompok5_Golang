package models

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Sample in-memory database
var Users = []User{
	{ID: 1, Name: "John Doe", Email: "john@example.com"},
	{ID: 2, Name: "Alice Smith", Email: "alice@example.com"},
}
