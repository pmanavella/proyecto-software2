package tokenizers

import (
    "fmt"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

type JWTConfig struct {
    Key      string
    Duration time.Duration
}

type JWT struct {
    config JWTConfig
}

func NewTokenizer(config JWTConfig) JWT {
    return JWT{
        config: config,
    }
}

func (tokenizer JWT) GenerateToken(username string, userID int64) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": username,
        "user_id":  userID,
        "exp":      time.Now().UTC().Add(tokenizer.config.Duration).Unix(),
    })

    tokenString, err := token.SignedString([]byte(tokenizer.config.Key))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func (tokenizer JWT) ValidateToken(tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(tokenizer.config.Key), nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }

    return token, nil
}
