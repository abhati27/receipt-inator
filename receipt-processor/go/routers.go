package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

func NewRouter(handleFunctions ApiHandleFunctions) *gin.Engine {
	return NewRouterWithGinEngine(gin.Default(), handleFunctions)
}

func NewRouterWithGinEngine(router *gin.Engine, handleFunctions ApiHandleFunctions) *gin.Engine {
	for _, route := range getRoutes(handleFunctions) {
		if route.HandlerFunc == nil {
			route.HandlerFunc = DefaultHandleFunc
		}
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		}
	}
	return router
}

func DefaultHandleFunc(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 Not Implemented")
}

type ApiHandleFunctions struct {
	DefaultAPI DefaultAPI
}

func getRoutes(handleFunctions ApiHandleFunctions) []Route {
	return []Route{
		{
			"ReceiptsIdPointsGet",
			http.MethodGet,
			"/receipts/:id/points",
			handleFunctions.DefaultAPI.ReceiptsIdPointsGet,
		},
		{
			"ReceiptsProcessPost",
			http.MethodPost,
			"/receipts/process",
			handleFunctions.DefaultAPI.ReceiptsProcessPost,
		},
	}
}
