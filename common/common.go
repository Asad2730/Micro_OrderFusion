package common

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"

	pb "github.com/Asad2730/Micro_OrderFusion/proto/user"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret string
var mu sync.Mutex

func generateSecretKey() (string, error) {
	key := make([]byte, 32) // 32 bytes = 256 bits
	if _, err := rand.Read(key); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(key), nil
}

func GenerateJWT(user *pb.User) (string, error) {
	mu.Lock()
	defer mu.Unlock()

	if jwtSecret == "" {
		var err error
		jwtSecret, err = generateSecretKey()
		if err != nil {
			panic(err)
		}
	}

	claims := jwt.MapClaims{
		"id":    user.GetId(),
		"email": user.GetEmail(),
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func GetJWTSecret() string {
	mu.Lock()
	defer mu.Unlock()
	return jwtSecret
}
