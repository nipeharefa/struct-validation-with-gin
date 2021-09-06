package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// MyStruct ..
type MyStruct struct {
	Status string `form:"status" validate:"sep-comma"`
	Page   int    `form:"page" validate:"required,gte=1"`
	Name   string `form:"name" validate:"min=3"`
}

type Response struct {
	Message string `json:"message"`
}

// use a single instance of Validate, it caches struct info
// var validate *validator.Validate

func main() {

	r := gin.New()

	r.GET("/car", func(c *gin.Context) {
		var err error

		resp := Response{
			Message: "Bad Requeest",
		}
		validate := NewValidator()
		s := MyStruct{}

		if err := c.Bind(&s); err != nil {
			// c.JSON(http.StatusBadRequest, "bad")
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		err = validate.Struct(s)
		if err != nil {
			a := err.(validator.ValidationErrors)

			msg := GetOneValidatorErrMsg(a)
			resp.Message = msg
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		c.JSON(http.StatusOK, nil)
	})

	r.Run()
}

// ValidateMyVal implements validator.Func
// func ValidateMyVal(fl validator.FieldLevel) bool {
// 	return fl.Field().String() == "awesome"
// }
