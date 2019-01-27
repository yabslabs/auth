package model

type Logins struct {
	Logins []Login `json:"Logins"`
}

type Login struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Roles    []Role `json:"roles"`
}

type Role string
