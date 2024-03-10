package query

import (
	"github.com/chat-bot-data/errors"
	"github.com/chat-bot-data/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type database struct {
	db *gorm.DB
}

func New(db *gorm.DB) database {
	return database{db: db}
}

func (d database) Create(c *gin.Context, input models.QueryInfo) (models.QueryInfo, error) {

	result := d.db.Create(&input)
	if result.Error != nil {
		log.Print(result.Error)
		return models.QueryInfo{}, errors.ErrorResponse{StatusCode: http.StatusInternalServerError, Code: "INTERNAL SERVER ERROR", Reason: "DB ERROR"}
	}

	if result.RowsAffected == 0 {
		return models.QueryInfo{}, errors.ErrorResponse{StatusCode: http.StatusBadRequest, Code: "BAD REQUEST", Reason: "Unable to Insert Data"}
	}

	return input, nil
}

func (d database) Get(c *gin.Context) ([]models.QueryInfo, error) {
	var data []models.QueryInfo
	result := d.db.Find(&data)
	if result.Error != nil {
		return []models.QueryInfo{}, errors.ErrorResponse{StatusCode: http.StatusInternalServerError, Code: "INTERNAL SERVER ERROR", Reason: "DB ERROR"}
	}

	return data, nil
}

func (d database) GetByQuestion(c *gin.Context, question string) (models.QueryInfo, error) {
	var data models.QueryInfo

	result := d.db.Find(&data, "question=?", question)
	if result.Error != nil {
		return models.QueryInfo{}, errors.ErrorResponse{StatusCode: http.StatusInternalServerError, Code: "INTERNAL SERVER ERROR", Reason: "DB ERROR"}
	}

	if result.RowsAffected == 0 {

		return models.QueryInfo{}, errors.ErrorResponse{StatusCode: http.StatusNotFound, Code: "NOT FOUND", Reason: "question Not Found"}

	}

	return data, nil
}

func (d database) PatchByQuestion(c *gin.Context, count int64, question string) (models.QueryInfo, error) {
	var data models.QueryInfo

	result := d.db.Model(&data).Where("question=?", question).Update("count", count)
	if result.Error != nil {
		return models.QueryInfo{}, errors.ErrorResponse{StatusCode: http.StatusInternalServerError, Code: "INTERNAL SERVER ERROR", Reason: "DB ERROR"}
	}

	if result.RowsAffected == 0 {

		return models.QueryInfo{}, errors.ErrorResponse{StatusCode: http.StatusNotFound, Code: "NOT FOUND", Reason: "question Not Found"}

	}

	return data, nil
}

func (d database) GetFrequentQuestions(c *gin.Context) ([]models.QueryInfo, error) {
	var data []models.QueryInfo
	result := d.db.Order("count DESC").Find(&data)
	if result.Error != nil {
		return []models.QueryInfo{}, errors.ErrorResponse{StatusCode: http.StatusInternalServerError, Code: "INTERNAL SERVER ERROR", Reason: "DB ERROR"}
	}

	return data, nil
}
