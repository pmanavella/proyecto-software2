package users

import (
    "users-api/domain/users" // Importar el paquete de domain

    "gorm.io/gorm"
)

type User struct {
    ID       int64  `gorm:"primaryKey;autoIncrement"`                    // Auto-increment primary key
    Username string `gorm:"size:100;not null;unique" binding:"required"` // Unique username, required
    Password string `gorm:"size:255;not null" binding:"required"`        // Password field, required
}

func CreateUser(db *gorm.DB, user users.User) (int64, error) {
    // Convertir `users.User` a `UserDAO`
    userDAO := User{
        Username: user.Username,
        Password: user.Password,
    }

    if err := db.Create(&userDAO).Error; err != nil {
        return 0, err
    }

    return userDAO.ID, nil
}

func GetUserByID(db *gorm.DB, id int64) (users.User, error) {
    var userDAO User
    if err := db.First(&userDAO, id).Error; err != nil {
        return users.User{}, err
    }

    // Convertir `UserDAO` a `users.User`
    user := users.User{
        ID:       userDAO.ID,
        Username: userDAO.Username,
        Password: userDAO.Password,
    }

    return user, nil
}

func GetAllUsers(db *gorm.DB) ([]users.User, error) {
    var userDAOs []User
    if err := db.Find(&userDAOs).Error; err != nil {
        return nil, err
    }

    var usersList []users.User
    for _, userDAO := range userDAOs {
        usersList = append(usersList, users.User{
            ID:       userDAO.ID,
            Username: userDAO.Username,
            Password: userDAO.Password,
        })
    }

    return usersList, nil
}

func UpdateUser(db *gorm.DB, user users.User) error {
    // Convertir `users.User` a `UserDAO`
    userDAO := User{
        ID:       user.ID,
        Username: user.Username,
        Password: user.Password,
    }

    if err := db.Save(&userDAO).Error; err != nil {
        return err
    }

    return nil
}

func DeleteUser(db *gorm.DB, id int64) error {
    if err := db.Delete(&User{}, id).Error; err != nil {
        return err
    }

    return nil
}

func GetUserByUsername(db *gorm.DB, username string) (users.User, error) {
    var userDAO User
    if err := db.Where("username = ?", username).First(&userDAO).Error; err != nil {
        return users.User{}, err
    }

    // Convertir `UserDAO` a `users.User`
    user := users.User{
        ID:       userDAO.ID,
        Username: userDAO.Username,
        Password: userDAO.Password,
    }

    return user, nil
}
