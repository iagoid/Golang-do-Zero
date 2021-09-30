package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type DataRequest struct {
	Email string `json:"email" validate:"required"`
	Name  string `json:"name" validate:"required"`
}

type ValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func main() {
	r := gin.Default()

	r.POST("/user", func(c *gin.Context) {
		var q DataRequest
		err := c.BindJSON(&q)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error to bind"})
			return
		}

		err = validator.New().Struct(q)
		if err != nil {
			var verr validator.ValidationErrors

			if errors.As(err, &verr) {
				c.JSON(http.StatusBadRequest, gin.H{"errors": Descriptive(verr)})
				return
			}

			// Outro tipo de erro, como json como json mal formado, ext
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		c.Status(http.StatusAccepted)
	})

	r.Run()
}

// Nomes nos campo
func Simple(verr validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs[f.Field()] = err
	}

	return errs
}

func Descriptive(verr validator.ValidationErrors) []ValidationError {
	errs := []ValidationError{}

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs = append(errs, ValidationError{Field: f.Field(), Reason: err})
	}

	return errs
}
