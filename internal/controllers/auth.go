package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"simple_auth/internal/errorhandling"
	"simple_auth/internal/lib"
	"simple_auth/internal/lib/db"
	"simple_auth/internal/lib/schema"
	"simple_auth/internal/types"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func HandleSignup(w http.ResponseWriter, r *http.Request) {
	var user schema.UserSignUpSchema
	w.Header().Set("Content-Type", "application/json")

	// Read and decode the JSON body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Error(err)
		errorhandling.RequestErrorHandler(w, "Input payload seems to be malformed, kindly check the input body.", http.StatusBadRequest)
		return
	}

	validationErrors := lib.ValidateUserInput(user)
	if validationErrors != nil {
		log.Error(errors.New("validation Failed"))
		errorhandling.ValidationErrorHandler(w, validationErrors)
		return
	}

	pgDb, getDbErr := db.GetDb()
	if getDbErr != nil {
		log.Error(getDbErr)
		errorhandling.InternalErrorHandler(w)
		return
	}
	var resultUser db.User
	queryResult := pgDb.Where("email = ?", user.Email).First(&resultUser)
	if queryResult.Error == nil {
		response := types.ApiResponse[any]{
			StatusCode: http.StatusBadRequest,
			Message:    "Email already exists, try another email.",
			Payload:    nil,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	password := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		log.Error(err)
		errorhandling.InternalErrorHandler(w)
		return
	}
	hashedPwd := string(hashedPassword)
	pgDb.Create(&db.User{
		Email:    user.Email,
		Password: &hashedPwd,
		FullName: user.FullName,
		// Role:     "Admin",
	})

	json.NewEncoder(w).Encode(types.ApiResponse[any]{
		StatusCode: http.StatusBadRequest,
		Message:    "User has been signed up successfully.",
		Payload:    nil,
	})
}
func HandleSignin(w http.ResponseWriter, r *http.Request) {

	var user schema.UserSignInSchema
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Error(r)
		errorhandling.RequestErrorHandler(w, "Input payload seems to be malformed, kindly check the input body.", http.StatusBadRequest)
		return
	}

	validationErrors := lib.ValidateUserInput(user)
	if validationErrors != nil {
		log.Error(errors.New("validation Failed"))
		errorhandling.ValidationErrorHandler(w, validationErrors)
		return
	}

	pgDb, getDbErr := db.GetDb()
	if getDbErr != nil {
		log.Error(getDbErr)
		errorhandling.InternalErrorHandler(w)
	}
	var resultUser db.User
	queryResult := pgDb.Where("email = ?", user.Email).First(&resultUser)
	if queryResult.Error != nil {
		response := types.MessageResponse{
			StatusCode: http.StatusNotFound,
			Message:    "User does not exist",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	compareErr := bcrypt.CompareHashAndPassword([]byte(*resultUser.Password), []byte(user.Password))
	if compareErr != nil {
		json.NewEncoder(w).Encode(types.MessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Password incorrect",
		})
		return
	}

	jwtPayload := schema.UserJwtPayloadSchema{
		Email:           resultUser.Email,
		FullName:        resultUser.FullName,
		Role:            resultUser.Role,
		IsEmailVerified: resultUser.EmailVerified,
	}
	refreshToken := lib.GenerateJWT(&jwtPayload, 7*24*time.Hour)
	accessToken := lib.GenerateJWT(&jwtPayload, 10*time.Minute)
	if refreshToken.Error != nil || accessToken.Error != nil {
		log.Error(refreshToken.Error)
		log.Error(accessToken.Error)
		errorhandling.InternalErrorHandler(w)
		return
	}

	createQueryResult := pgDb.Create(&db.Session{
		UserID:       resultUser.ID,
		ExpiresAt:    *refreshToken.ExpiresAt,
		RefreshToken: *refreshToken.Token,
	})

	if createQueryResult.Error != nil {
		log.Error(createQueryResult.Error)
		errorhandling.InternalErrorHandler(w)
		return
	}

	json.NewEncoder(w).Encode(types.UserSignInResponse{
		StatusCode:   http.StatusOK,
		Message:      "User signed in successfully",
		AccessToken:  *accessToken.Token,
		RefreshToken: *refreshToken.Token,
	})
}
