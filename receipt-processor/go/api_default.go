/*
 * api_default.go - API Endpoints for Receipt Processor
 *
 * This file defines two endpoints:
 * 1. **POST /receipts/process**:
 *    - Validates and processes a receipt.
 *    - Calculates points and stores them with a unique ID.
 *    - Returns the receipt ID or a `400 Bad Request` on validation failure.
 * 2. **GET /receipts/:id/points**:
 *    - Retrieves points for a receipt by its ID.
 *    - Returns `404 Not Found` if the ID does not exist.
 *
 * Includes custom field validation and uses a map to store receipt points.
 */
package openapi

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type DefaultAPI struct{}

var validate *validator.Validate

func init() {
	validate = validator.New()

	validate.RegisterValidation("retailer", func(fl validator.FieldLevel) bool {
		match, _ := regexp.MatchString(`^[\w\s\-&]+$`, fl.Field().String())
		return match
	})

	validate.RegisterValidation("total", func(fl validator.FieldLevel) bool {
		match, _ := regexp.MatchString(`^\d+\.\d{2}$`, fl.Field().String())
		return match
	})

	validate.RegisterValidation("price", func(fl validator.FieldLevel) bool {
		match, _ := regexp.MatchString(`^\d+\.\d{2}$`, fl.Field().String())
		return match
	})
}

var receiptStore = make(map[string]int)

func (api *DefaultAPI) ReceiptsProcessPost(c *gin.Context) {
	var receipt Receipt

	if err := c.ShouldBindJSON(&receipt); err != nil || validate.Struct(receipt) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "The receipt is invalid.",
		})
		return
	}

	points := calculatePoints(receipt)

	id := uuid.New().String()
	receiptStore[id] = points

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (api *DefaultAPI) ReceiptsIdPointsGet(c *gin.Context) {
	id := c.Param("id")

	points, exists := receiptStore[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No receipt found for that ID.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
}
