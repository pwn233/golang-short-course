package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Endpoint struct{}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

// AddUser is a endpoint(controller) for taking action of add user to the system.
func (e *Endpoint) AddUser(c *gin.Context) {
	centralService := NewUserService()
	// Prepare Model(struct)
	var requestBody AddUserRequest
	// Mapping requeest body to Model
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	// Call function (from logic) to deep step inside such as validation and insert to database
	if err := centralService.AddUser(requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
	} else {
		c.JSON(http.StatusCreated, nil)
	}
}
