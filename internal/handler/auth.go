package handler

import (
	"encoding/json"
	"erosync/internal/domain"
	"erosync/internal/middleware"
	"erosync/internal/schema"
	"erosync/internal/service"
	"erosync/pkg/response"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req schema.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, 400, response.MsgInvalidBody, nil)
		return
	}

	if errs := h.validator.ValidateStruct(req); errs != nil {
		response.Error(w, 400, response.MsgInvalidBody, errs)
		return
	}

	user, err := h.users.Create(r.Context(), req)
	if err != nil {
		response.JSON(w, 500, response.Map{"message": err.Error()})
		return
	}

	accessToken, _ := h.jwt.GenerateToken(user.ID, string(user.Role), 30*time.Minute)
	refreshToken, _ := h.jwt.GenerateToken(user.ID, string(user.Role), 24*time.Hour)

	resp := schema.AuthResponse{
		User:         *user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	response.JSON(w, 201, resp)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req schema.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, 400, response.MsgInvalidBody, nil)
		return
	}

	if errs := h.validator.ValidateStruct(req); errs != nil {
		response.Error(w, 400, response.MsgInvalidBody, errs)
		return
	}

	user, err := h.users.GetByEmail(req.Email)
	if err != nil || user == nil {
		response.JSON(w, 500, response.Map{"message": "invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		response.JSON(w, 500, response.Map{"message": "invalid credentials"})
		return
	}

	accessToken, _ := h.jwt.GenerateToken(user.ID, string(user.Role), 30*time.Minute)
	refreshToken, _ := h.jwt.GenerateToken(user.ID, string(user.Role), 24*time.Hour)

	resp := schema.AuthResponse{
		User:         *user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	response.JSON(w, 200, resp)
}

func (h *Handler) SendEmailVerificationCode(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(string)

	user, err := h.users.GetByID(userID)
	if err != nil {
		log.Println(err)
		response.Error(w, 400, response.MsgUnknown, nil)
		return
	}

	otpCode, err := h.otps.Create(r.Context(), service.CreateOTPParam{
		Purpose:   domain.OTPPurposeVerifyEmail,
		Channel:   domain.OTPChannelEmail,
		Recipient: user.Email,
	})

	if err != nil {
		response.Error(w, 500, err.Error(), nil)
		return
	}

	if err := h.mail.VerifyEmail(*user, otpCode); err != nil {
		response.Error(w, 500, err.Error(), nil)
		return
	}

	response.JSON(w, 200, response.Map{"message": "verification code sent"})
}

type VerifyEmailRequest struct {
	OtpCode string `json:"code" validate:"required"`
}

func (h *Handler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	var req schema.VerifyEmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, 400, response.MsgInvalidBody, nil)
		return
	}

	if err := h.validator.ValidateStruct(req); err != nil {
		response.Error(w, 400, response.MsgInvalidBody, err)
		return
	}

	userID := r.Context().Value(middleware.UserIDKey).(string)

	user, err := h.users.GetByID(userID)
	if err != nil {
		response.Error(w, 500, err.Error(), nil)
		return
	}

	err = h.otps.Validate(r.Context(), service.ValidateOTPParam{
		Code:      req.OtpCode,
		Channel:   domain.OTPChannelEmail,
		Purpose:   domain.OTPPurposeVerifyEmail,
		Recipient: user.Email,
	})
	if err != nil {
		log.Println(err)
		response.Error(w, 400, "invalid otp", nil)
		return
	}

	if !user.EmailVerified() {
		err = h.users.VerifyEmail(r.Context(), *user)
		if err != nil {
			response.Error(w, 500, err.Error(), nil)
			return
		}
	}

	response.JSON(w, 200, response.Map{"message": "success"})
}
