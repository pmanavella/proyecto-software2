package usersapi

// msotrar usuario

type UserResponse struct {
	Firstname string `json:"firstname"` 
	Lastname  string `json:"lastname"` 
	Email     string `json:"email"` 
	Type      string `json:"type"`
}


// login

type UserLoginRequest struct {
	Email     string `json:"email"` 
	Password  string `json:"password"` 
}

type UserLoginResponse struct {
	Token string `json:"token"` 
}
