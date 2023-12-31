package handler

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pwn233/golang-short-course/middleware"
	"github.com/pwn233/golang-short-course/services/central"
	"github.com/pwn233/golang-short-course/services/user"
)

type route struct {
	Name        string
	Description string
	Method      string
	Pattern     string
	Endpoint    gin.HandlerFunc
	Validation  gin.HandlerFunc
}

type Routes struct {
	router         *gin.Engine
	centralService []route
	userService    []route
}

func (r Routes) InitRouter() http.Handler {
	centralEndpoint := central.NewEndpoint()
	userEndpoint := user.NewEndpoint()
	r.centralService = []route{
		{
			Name:        "Get : Health",
			Description: "get health status from server",
			Method:      http.MethodGet,
			Pattern:     "health",
			Endpoint:    centralEndpoint.Health,
			Validation:  middleware.GeneralValidation,
		},
		{
			Name:        "Post : Add Number",
			Description: "add number with 2 variables from server",
			Method:      http.MethodPost,
			Pattern:     "add",
			Endpoint:    centralEndpoint.Add,
			Validation:  middleware.GeneralValidation,
		},
	}
	r.userService = []route{
		{
			Name:        "Post : Add User",
			Description: "add user to the database",
			Method:      http.MethodPost,
			Pattern:     "/user/add",
			Endpoint:    userEndpoint.AddUser,
			Validation:  middleware.GeneralValidation,
		},
	}
	ro := gin.Default()
	ro.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "HEAD", "OPTIONS", "DELETE", "CONNECT", "TRACE"},
		AllowHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Origin", "X-Correlation-Id"},
		MaxAge:       12 * time.Hour,
	}))
	mainRoute := ro.Group("golang-short-course")
	for _, e := range r.centralService {
		mainRoute.Handle(e.Method, e.Pattern, e.Validation, e.Endpoint)
	}
	for _, e := range r.userService {
		mainRoute.Handle(e.Method, e.Pattern, e.Validation, e.Endpoint)
	}
	return ro
}
