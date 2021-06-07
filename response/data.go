package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type ApiDataResponse struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta,omitempty"`
}

func DataResponse(c *gin.Context, status int, data interface{}, meta interface{}) {

	if meta == nil || (reflect.ValueOf(meta).Kind() == reflect.Ptr && reflect.ValueOf(meta).IsNil()) {
		c.JSON(status, gin.H{"data": data})
	} else {
		c.JSON(status, gin.H{"meta": meta, "data": data})
	}
}

func OkDataResponse(c *gin.Context, response *ApiDataResponse) {
	c.JSON(http.StatusOK, response)
}

func CreatedDataResponse(c *gin.Context, response *ApiDataResponse) {
	c.JSON(http.StatusCreated, response)
}

func OkResponse(c *gin.Context) {
	c.Status(http.StatusOK)
}

func CreatedResponse(c *gin.Context) {
	c.Status(http.StatusCreated)
}
