package user

import (
	"fmt"
	"github.com/chat-bot-data/errors"
	"github.com/chat-bot-data/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type database struct {
	db *gorm.DB
}

func New(db *gorm.DB) database {
	return database{db: db}
}

func (d database) Create(c *gin.Context, user models.UserInfo) (models.UserInfo, error) {
	result := d.db.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		return models.UserInfo{}, errors.ErrorResponse{StatusCode: http.StatusInternalServerError, Code: "INTERNAL SERVER ERROR", Reason: "DB ERROR"}
	}

	if result.RowsAffected == 0 {
		return models.UserInfo{}, errors.ErrorResponse{StatusCode: http.StatusBadRequest, Code: "BAD REQUEST", Reason: "Unable to Insert Data"}
	}

	return user, nil
}

func (d database) Get(c *gin.Context, username string, password string) (models.UserInfo, error) {
	var users models.UserInfo
	result := d.db.Find(&users, "email=? and password=?", username, password)
	if result.Error != nil {
		return models.UserInfo{}, errors.ErrorResponse{StatusCode: http.StatusInternalServerError, Code: "INTERNAL SERVER ERROR", Reason: "DB ERROR"}
	}

	return users, nil
}

func (d database) GetByID(c *gin.Context, ID string) (models.UserInfo, error) {
	var data models.UserInfo

	result := d.db.Find(&data, "id=?", ID)
	if result.Error != nil {
		return models.UserInfo{}, errors.ErrorResponse{StatusCode: http.StatusInternalServerError, Code: "INTERNAL SERVER ERROR", Reason: "DB ERROR"}
	}

	if result.RowsAffected == 0 {

		return models.UserInfo{}, errors.ErrorResponse{StatusCode: http.StatusNotFound, Code: "NOT FOUND", Reason: "ID not Found"}

	}

	return data, nil
}

func (d database) PatchByID(c *gin.Context, ID string, user models.UserInfo) (models.UserInfo, error) {
	var data models.UserInfo

	result := d.db.Model(&data).Where("id=?", ID).Update("password", user.Password)
	if result.Error != nil {
		return models.UserInfo{}, errors.ErrorResponse{StatusCode: http.StatusInternalServerError, Code: "INTERNAL SERVER ERROR", Reason: "DB ERROR"}
	}

	if result.RowsAffected == 0 {

		return models.UserInfo{}, errors.ErrorResponse{StatusCode: http.StatusNotFound, Code: "NOT FOUND", Reason: "ID not Found"}

	}

	return data, nil
}
