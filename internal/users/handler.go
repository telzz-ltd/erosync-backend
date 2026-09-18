package users

import (
	"erosync/internal/lib/app"
	"erosync/internal/lib/notification"
	"erosync/internal/service"
	"erosync/internal/store"
	"log"
	"net/http"
)

func SendEmailVerificationCode(store *store.Store, otps *service.OTPService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if otps == nil || store == nil {
			log.Println(store, otps)
			panic("dependencies cannot be nil")
		}

		userId := app.GetValue(r, "userId").(string)
		user, err := store.User.FindByID(userId)
		if err != nil {
			app.JSON(w, 500, app.H{"message": err.Error()})
			return
		}

		if user == nil {
			app.JSON(w, 401, app.H{"message": "unauthorized"})
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
			app.JSON(w, 500, app.H{"message": err.Error()})
			return
		}

		go func() {
			if err := notification.SendEmailVerificationMail(*user, code, expiryMin); err != nil {
				log.Println(err)
			}
		}()

		app.JSON(w, 200, app.H{"message": "verification code sent"})
	}
}

func VerifyEmail(store *store.Store, otps *service.OTPService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.VerifyEmailRequest
		if err := app.ShouldBindJSON(r, &req); err != nil {
			app.JSON(w, 400, app.H{"message": err.Error()})
			return
		}

		userId := app.GetValue(r, "userId").(string)
		user, err := store.User.FindByID(userId)
		if err != nil || user == nil {
			if err != nil {
				log.Println(err)
			}
			app.JSON(w, 401, app.H{"message": "unauthorized"})
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
			app.JSON(w, 401, app.H{"message": "invalid or expired otp"})
			return
		}

		if !user.EmailVerified() {
			user.VerifyEmail()
			if err := store.User.Save(user); err != nil {
				app.JSON(w, 500, app.H{"message": err.Error()})
				return
			}
		}

		app.JSON(w, 200, app.H{"message": "success"})
	}
}
