package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/issuerenaming/go/orm"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// RegisterControllers register controllers
func RegisterControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/issuerenaming/go")
	{ // insertion point for registrations
		v1.GET("/v1/bars", GetBars)
		v1.GET("/v1/bars/:id", GetBar)
		v1.POST("/v1/bars", PostBar)
		v1.PATCH("/v1/bars/:id", UpdateBar)
		v1.PUT("/v1/bars/:id", UpdateBar)
		v1.DELETE("/v1/bars/:id", DeleteBar)

		v1.GET("/v1/waldos", GetWaldos)
		v1.GET("/v1/waldos/:id", GetWaldo)
		v1.POST("/v1/waldos", PostWaldo)
		v1.PATCH("/v1/waldos/:id", UpdateWaldo)
		v1.PUT("/v1/waldos/:id", UpdateWaldo)
		v1.DELETE("/v1/waldos/:id", DeleteWaldo)

		v1.GET("/v1/commitfrombacknb", GetLastCommitFromBackNb)
		v1.GET("/v1/pushfromfrontnb", GetLastPushFromFrontNb)
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func GetLastCommitFromBackNb(c *gin.Context) {
	res := orm.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func GetLastPushFromFrontNb(c *gin.Context) {
	res := orm.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
