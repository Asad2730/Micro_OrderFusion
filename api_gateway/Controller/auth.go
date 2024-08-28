package controller

import (
	user "github.com/Asad2730/Micro_OrderFusion/proto/user"
	"github.com/gin-gonic/gin"
)

type AuthClient struct {
	gRPCClient user.UserServiceClient
}

func NewAuthClient(client user.UserServiceClient) *AuthClient {
	return &AuthClient{gRPCClient: client}
}

func (client *AuthClient) SignUp(c *gin.Context) {
	var auth *user.User
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(500, gin.H{"Error binding data": err.Error()})
		return
	}
	req := &user.RequestSignup{
		Name:     auth.Name,
		Email:    auth.Email,
		Password: auth.Email,
	}
	res, err := client.gRPCClient.Signup(c, req)
	if err != nil {
		c.JSON(402, gin.H{"Error Creating": err.Error()})
		return
	}

	c.JSON(201, res.Message)
}

func (client *AuthClient) Login(c *gin.Context) {
	var auth *user.User
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(500, gin.H{"Error binding data": err.Error()})
		return
	}
	req := &user.RequestLogin{
		Email:    auth.Email,
		Password: auth.Email,
	}
	res, err := client.gRPCClient.Login(c, req)
	if err != nil {
		c.JSON(402, gin.H{"Error Creating": err.Error()})
		return
	}

	c.JSON(200, res)
}
