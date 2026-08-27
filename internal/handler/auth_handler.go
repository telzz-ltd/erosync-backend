package handler

import (
	"erosync/internal/dto"
	"erosync/internal/service"
	"erosync/internal/shared/notification"
	"erosync/internal/store"
	"log"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(s *service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}

		resp, err := s.Register(req)
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		c.JSON(201, resp)
	}
}

func LoginHandler(s *service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}

		resp, err := s.Login(req)
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		c.JSON(201, resp)
	}
}

func SendEmailVerificationCodeHandler(store *store.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.MustGet("userId").(string)
		user, err := store.User.FindByID(userId)
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		if user == nil {
			c.JSON(401, gin.H{"message": "unauthorized"})
			return
		}

		go func() {
			if err := notification.SendEmailVerificationMail(*user, "189091"); err != nil {
				log.Println(err)
			}
		}()

		c.JSON(200, gin.H{"message": "verification code sent"})
	}
}
