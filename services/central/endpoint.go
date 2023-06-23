package central

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Endpoint struct{}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

func (e *Endpoint) Health(c *gin.Context) {
	response := HealthResponse{
		Code:        200,
		Status:      "up",
		Description: nil,
		Message:     nil,
	}
	c.JSON(http.StatusOK, response)
}

func (e *Endpoint) Add(c *gin.Context) {
	centralService := NewCentralService()
	var requestBody AddNumberRequest
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	if response, err := centralService.AddNumber(requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}
