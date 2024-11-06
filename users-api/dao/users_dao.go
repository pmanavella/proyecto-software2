package users

type User struct {
    UserID   int64  `gorm:"primaryKey;autoIncrement"`
    Username string `gorm:"size:100;not null;unique" binding:"required"`
    Password string `gorm:"size:255;not null" binding:"required"`
    Nombre   string `gorm:"size:100;not null" binding:"required"`
    Apellido string `gorm:"size:100;not null" binding:"required"`
    Email    string `gorm:"size:100;not null;unique" binding:"required"`
    Admin    bool   `gorm:"not null"`
}

type Users []User
