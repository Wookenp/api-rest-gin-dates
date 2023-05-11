package dates

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/dates/create", CreateDate)
	r.GET("/dates", ListDates)
	r.GET("/dates/:id", RetrieveDate)
	r.PUT("/dates/:id", UpdateDate)
	r.DELETE("/dates/:id", DeleteDate)
}
