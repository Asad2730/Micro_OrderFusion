package common

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	pb "github.com/Asad2730/Micro_OrderFusion/proto"
	"github.com/golang-jwt/jwt/v5"
)

func generateSecretKey() (string, error) {
	key := make([]byte, 32) // 32 bytes = 256 bits
	if _, err := rand.Read(key); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(key), nil
}

func GenerateJWT(user *pb.User) (string, error) {
	secret, err := generateSecretKey()
	if err != nil {
		panic(err)
	}

	jwtSecret := []byte(secret)

	claims := jwt.MapClaims{
		"id":    user.GetId(),
		"email": user.GetEmail(),
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
