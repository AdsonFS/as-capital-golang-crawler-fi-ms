package routes

import (
	fisolver "as-capital-crawler-fi-ms/internal/fi_solver"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	v1 := router.Group("api/v1")
	{
		v1.GET("/:fi", func(ctx *gin.Context) {
			fiName := ctx.Params.ByName("fi")
			data, err := fisolver.GetData(fiName)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"err": err})
				return
			}
			ctx.JSON(http.StatusOK, data)
		})
	}
	return v1
}
