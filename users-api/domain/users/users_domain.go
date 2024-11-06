package users

type User struct {
    ID       int64  `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
    Nombre   string `json:"nombre"`
    Apellido string `json:"apellido"`
    Email    string `json:"email"`
    Admin    bool   `json:"admin"`
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginResponse struct {
    UserID   int64  `json:"user_id"`
    Username string `json:"username"`
    Token    string `json:"token"`
    Admin    bool   `json:"admin"`
}

type Token struct {
    Token   string `json:"token"`
    User_id int    `json:"id_user"`
    Admin   bool   `json:"admin"`
}
