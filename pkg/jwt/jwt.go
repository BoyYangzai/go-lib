package jwt

import (
	"errors"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/BoyYangZai/go-server-lib/pkg/config_reader"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	userTokenMap = make(map[string]string)
	tokenMutex   sync.Mutex
)

type User struct {
	Username string `json:"username"`
	ID       uint64 `json:"id"`
	jwt.RegisteredClaims
}

var CurrentAuthUserId uint64

func GenerateToken(username string, id uint64) (string, error) {
	secretKey := getSecretKey()
	claims := User{
		Username: username,
		ID:       id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ParseJwt(tokenString string) (*User, error) {
	secretKey := getSecretKey()
	t, err := jwt.ParseWithClaims(tokenString, &User{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if t.Valid {
		if claims, ok := t.Claims.(*User); ok {
			return claims, nil
		}
	}
	return nil, errors.New("Failed to parse JWT claims")
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET_KEY")

	if secretKey == "" {
		secretKey = config_reader.GetConfigByKey("jwt.secret_key")
	}

	return secretKey
}

func Auth(c *gin.Context, isMatchedSuccess bool, username string, id uint64) {
	// login auth storage
	isLoginAuthStorage, token := LoginAuthStorage(c)
	if isLoginAuthStorage {
		c.JSON(http.StatusOK, gin.H{
			"msg":   "login success",
			"token": token,
		})
		return
	}

	// normal login auth
	if isMatchedSuccess {
		token, err := GenerateToken(username, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "generate token error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg":   "login success",
			"token": token,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "email and password not match",
		})
	}
}
