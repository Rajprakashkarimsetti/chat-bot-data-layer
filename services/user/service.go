package user

import (
	"github.com/chat-bot-data/errors"
	"github.com/chat-bot-data/models"
	"github.com/chat-bot-data/store"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
	"net/http"
	"regexp"
)

const (
	EMAIL = `^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$`
	PHONE = `^[6-9]\d{9}$`
)

type service struct {
	User store.User
}

func New(s store.User) service {
	return service{User: s}
}

func (s service) Create(c *gin.Context, user models.UserInfo) (models.UserInfo, error) {

	err := validateUser(user)
	if err != nil {
		return models.UserInfo{}, err
	}

	user.Password = encrypt(user.Password)
	user.ID = uuid.NewString()

	userInfo, err := s.User.Create(c, user)
	if err != nil {
		return models.UserInfo{}, err
	}

	return userInfo, nil
}

func (s service) GetByID(c *gin.Context, ID string) (models.UserInfo, error) {
	userInfo, err := s.User.GetByID(c, ID)
	if err != nil {
		return models.UserInfo{}, err
	}

	return userInfo, nil
}

func (s service) Get(c *gin.Context, username string, password string) (models.UserInfo, error) {
	password = encrypt(password)
	userInfo, err := s.User.Get(c, username, password)
	if err != nil {
		return models.UserInfo{}, err
	}

	return userInfo, nil
}

func (s service) PatchByID(c *gin.Context, ID string, user models.UserInfo) (models.UserInfo, error) {
	user.Password = encrypt(user.Password)
	userInfo, err := s.User.PatchByID(c, ID, user)
	if err != nil {
		return models.UserInfo{}, err
	}

	return userInfo, nil
}

func encrypt(password string) string {

	return password
}

func validateUser(user models.UserInfo) error {
	switch {
	case user.FirstName == "":
		return errors.ErrorResponse{StatusCode: http.StatusBadRequest, Code: "BAD REQUEST", Reason: "Missing field firstName"}
	case user.LastName == "":
		return errors.ErrorResponse{StatusCode: http.StatusBadRequest, Code: "BAD REQUEST", Reason: "Missing field lastName"}
	case !isEmail(user.Email):
		return errors.ErrorResponse{StatusCode: http.StatusBadRequest, Code: "BAD REQUEST", Reason: "Invalid Email"}
	case !isPhone(user.Phone):
		return errors.ErrorResponse{StatusCode: http.StatusBadRequest, Code: "BAD REQUEST", Reason: "Invalid Phone"}
	}
	return nil
}

func isEmail(email string) bool {
	regex := regexp.MustCompile(EMAIL)
	if regex.MatchString(email) {
		return true
	}

	return false
}

func isPhone(phone string) bool {
	regex := regexp.MustCompile(PHONE)
	if regex.MatchString(phone) {
		return true
	}

	return false
}

func validGender(gender string) bool {
	genders := []string{"MALE", "FEMALE"}
	if slices.Contains(genders, gender) {
		return true
	}
	return false
}
