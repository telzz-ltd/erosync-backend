package handler

import (
	"erosync/internal/dto"
	"erosync/internal/model"
	"erosync/internal/service"
	"erosync/internal/shared/notification"
	"erosync/internal/store"
	"log"

	"github.com/gin-gonic/gin"
)

func Register(s *service.AuthService) gin.HandlerFunc {
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

func Login(s *service.AuthService) gin.HandlerFunc {
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

func SendEmailVerificationCode(store *store.Store, otps *service.OTPService) gin.HandlerFunc {
	return func(c *gin.Context) {
		if otps == nil || store == nil {
			log.Println(store, otps)
			panic("dependencies cannot be nil")
		}

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

		var expiryMin = 10
		code, err := otps.Create(service.CreateOTPParam{
			Purpose:   model.OTPPurposeVerifyEmail,
			Channel:   model.OTPChannelEmail,
			Recipient: user.Email,
			ExpireMin: expiryMin,
		})
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		go func() {
			if err := notification.SendEmailVerificationMail(*user, code, expiryMin); err != nil {
				log.Println(err)
			}
		}()

		c.JSON(200, gin.H{"message": "verification code sent"})
	}
}

func VerifyEmail(store *store.Store, otps *service.OTPService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.VerifyEmailRequest
		if err := c.ShouldBindBodyWithJSON(&req); err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}

		userId := c.MustGet("userId").(string)
		user, err := store.User.FindByID(userId)
		if err != nil || user == nil {
			if err != nil {
				log.Println(err)
			}
			c.JSON(401, gin.H{"message": "unauthorized"})
			return
		}

		err = otps.Validate(service.ValidateOTPParam{
			Code:      req.OtpCode,
			Channel:   model.OTPChannelEmail,
			Purpose:   model.OTPPurposeVerifyEmail,
			Recipient: user.Email,
		})
		if err != nil {
			log.Println(err)
			c.JSON(401, gin.H{"message": "invalid or expired otp"})
			return
		}

		if !user.EmailVerified() {
			user.VerifyEmail()
			if err := store.User.Save(user); err != nil {
				c.JSON(500, gin.H{"message": err.Error()})
				return
			}
		}

		c.JSON(200, gin.H{"message": "success"})
	}
}
