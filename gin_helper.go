package cornerstone

import (
	"github.com/gin-gonic/gin"
)

type response struct {
	ErrorCode    int         `json:"error_code,omitempty"`
	SubErrorCode int         `json:"sub_error_code,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
	TransitID    string      `json:"transit_id,omitempty"`
	Result       interface{} `json:"result,omitempty"`
}

func DoneWithStatus(c *gin.Context, status int, result interface{}) {
	resp := response{
		Result: result,
	}
	c.AbortWithStatusJSON(status, resp)
}

func FromCodeErrorWithStatus(c *gin.Context, status int, err CodeError) {
	resp := response{
		ErrorCode:    err.ErrorCode(),
		SubErrorCode: err.SubErrorCode(),
		ErrorMessage: err.ErrorMessage(),
	}
	c.AbortWithStatusJSON(status, resp)
}
