package controllers

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thanhpk/randstr"
	"github.com/wpcodevo/golang-gorm-postgres/initializers"
	"github.com/wpcodevo/golang-gorm-postgres/models"
	"github.com/wpcodevo/golang-gorm-postgres/utils"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

// [...] SignUp User
func (ac *AuthController) SignUpUser(ctx *gin.Context) {
	var payload *models.SignUpInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if payload.Password != payload.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	now := time.Now()
	newUser := models.User{
		Name:      payload.Name,
		Email:     strings.ToLower(payload.Email),
		Password:  hashedPassword,
		Role:      "user",
		Verified:  false,
		Photo:     payload.Photo,
		Provider:  "local",
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := ac.DB.Create(&newUser)

	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
		return
	}

	config, _ := initializers.LoadConfig(".")

	// Generate Verification Code
	code := randstr.String(20)

	verification_code := utils.Encode(code)

	// Update User in Database
	newUser.VerificationCode = verification_code
	ac.DB.Save(newUser)

	var firstName = newUser.Name

	if strings.Contains(firstName, " ") {
		firstName = strings.Split(firstName, " ")[1]
	}

	// 👇 Send Email
	emailData := utils.EmailData{
		URL:       config.ClientOrigin + "/verifyemail/" + code,
		FirstName: firstName,
		Subject:   "Your account verification code",
	}

	utils.SendEmail(&newUser, &emailData, "verificationCode.html")

	message := "We sent an email with a verification code to " + newUser.Email
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": message})
}

// [...] SignIn User
func (ac *AuthController) SignInUser(ctx *gin.Context) {
	var payload *models.SignInInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var user models.User
	result := ac.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}

	if !user.Verified {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "Please verify your email"})
		return
	}

	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}

	config, _ := initializers.LoadConfig(".")

	// Generate Token
	token, err := utils.GenerateToken(config.TokenExpiresIn, user.ID, config.TokenSecret)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("token", token, config.TokenMaxAge*60, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "token": token})
}

// [...] SignOut User
func (ac *AuthController) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

// [...] Verify Email
func (ac *AuthController) VerifyEmail(ctx *gin.Context) {

	code := ctx.Params.ByName("verificationCode")
	verification_code := utils.Encode(code)

	var updatedUser models.User
	result := ac.DB.First(&updatedUser, "verification_code = ?", verification_code)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid verification code or user doesn't exists"})
		return
	}

	if updatedUser.Verified {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User already verified"})
		return
	}

	updatedUser.VerificationCode = ""
	updatedUser.Verified = true
	ac.DB.Save(&updatedUser)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Email verified successfully"})
}

func (ac *AuthController) ForgotPassword(ctx *gin.Context) {
	var payload *models.ForgotPasswordInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	message := "You will receive a reset email if user with that email exist"

	var user models.User
	result := ac.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}

	if !user.Verified {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Account not verified"})
		return
	}

	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load config", err)
	}

	// Generate Verification Code
	resetToken := randstr.String(20)

	passwordResetToken := utils.Encode(resetToken)
	user.PasswordResetToken = passwordResetToken
	user.PasswordResetAt = time.Now().Add(time.Minute * 15)
	ac.DB.Save(&user)

	var firstName = user.Name

	if strings.Contains(firstName, " ") {
		firstName = strings.Split(firstName, " ")[1]
	}

	// 👇 Send Email
	emailData := utils.EmailData{
		URL:       config.ClientOrigin + "/forgotPassword/" + resetToken,
		FirstName: firstName,
		Subject:   "Your password reset token (valid for 10min)",
	}

	utils.SendEmail(&user, &emailData, "resetPassword.html")

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
}

func (ac *AuthController) ResetPassword(ctx *gin.Context) {
	var payload *models.ResetPasswordInput
	resetToken := ctx.Params.ByName("resetToken")

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if payload.Password != payload.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
		return
	}

	hashedPassword, _ := utils.HashPassword(payload.Password)

	passwordResetToken := utils.Encode(resetToken)

	var updatedUser models.User
	result := ac.DB.First(&updatedUser, "password_reset_token = ? AND password_reset_at > ?", passwordResetToken, time.Now())
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "The reset token is invalid or has expired"})
		return
	}

	updatedUser.Password = hashedPassword
	updatedUser.PasswordResetToken = ""
	ac.DB.Save(&updatedUser)

	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Password data updated successfully"})
}
